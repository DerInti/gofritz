package main

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/nitram509/gofitz/pkg/scpd"
	"os"
	"strings"
	"text/template"
	"unicode"
)

const HOST = "fritz.box"
const URL = "http://" + HOST + ":49000"

const debug = false
const spaces = "                                                                                                     "

type structData struct {
	Name           string
	SCDPShortName  string
	SCDPUrlPath    string
	SystemVersion  string
	SoapActionName string
	Variables      []string
}

//go:embed struct.go.tmpl
var templateStructGo string

func findScdpUrlPath(rootSpec scpd.ServiceControlledProtocolDescriptions, serviceId string) string {
	for _, service := range rootSpec.Device.ServiceList {
		if serviceId == service.ServiceId {
			return service.SCPDURL
		}
	}
	for _, device := range rootSpec.Device.DeviceList {
		for _, service := range device.ServiceList {
			if serviceId == service.ServiceId {
				return service.SCPDURL
			}
		}
	}
	panic(errors.New(fmt.Sprintf("serviceId '%s' not found", serviceId)))
}

func deriveScdpShortName(scpdUrl string) string {
	result := strings.TrimLeft(scpdUrl, "/")
	if strings.HasSuffix(result, ".xml") {
		result = result[:len(result)-4]
	}
	return result
}

func filterArguments(arguments []scpd.Argument, direction string) (result []scpd.Argument) {
	for _, arg := range arguments {
		if arg.Direction == direction {
			result = append(result, arg)
		}
	}
	return result
}

func determineVariableName(argumentName string) string {
	if strings.HasPrefix(argumentName, "New") {
		argumentName = argumentName[3:]
	}
	return argumentName
}

func determineTypeName(spec scpd.ServiceControlledProtocolDescriptions, relatedStateVariable string) string {
	for _, serviceState := range spec.ServiceStateTable {
		if relatedStateVariable == serviceState.Name {
			switch strings.ToLower(serviceState.DataType) {
			// return type with trailing spaces for formatting
			case "string":
				return "string"
			case "datetime":
				return "string"
			case "uuid":
				return "string"
			case "ui4":
				return "int   "
			case "i4":
				return "int   "
			case "ui2":
				return "int   "
			case "ui1":
				return "int   "
			case "boolean":
				return "bool  "
			default:
				panic(errors.New(fmt.Sprintf("relatedStateVariable '%s' unknown type '%s'",
					relatedStateVariable,
					strings.ToLower(serviceState.DataType))))
			}
		}
	}
	panic(errors.New(fmt.Sprintf("relatedStateVariable '%s' not found", relatedStateVariable)))
}

func determineVariableComment(spec scpd.ServiceControlledProtocolDescriptions, relatedStateVariable string) string {
	result := "//"
	for _, serviceState := range spec.ServiceStateTable {
		if relatedStateVariable == serviceState.Name {
			if len(serviceState.DefaultValue) > 0 {
				result = result + " default=" + serviceState.DefaultValue
			}
			if len(serviceState.AllowedValueList) > 0 {
				if strings.Contains(result, "default") {
					result = result + ";"
				}
				for i, allowedValue := range serviceState.AllowedValueList {
					if i == 0 {
						result = result + " oneOf=[" + allowedValue
					} else {
						result = result + "," + allowedValue
					}
				}
				result = result + "]"
			}
			if strings.ToLower(serviceState.DataType) == "datetime" {
				result = result + "; type=Datetime"
			}
			if strings.ToLower(serviceState.DataType) == "uuid" {
				result = result + "; type=UUID"
			}
		}
	}
	return result
}

func serviceId2SnakeCase(serviceId string) string {
	serviceId = strings.ReplaceAll(serviceId, "X_AVM-DE_", "Avm")
	serviceId = strings.ReplaceAll(serviceId, "WebDAV", "Webdav")
	serviceId = strings.ReplaceAll(serviceId, "FRITZ", "Fritz")
	serviceId = strings.ReplaceAll(serviceId, "VoIP", "Voip")
	serviceId = strings.ReplaceAll(serviceId, "WLAN", "Wlan")
	serviceId = strings.ReplaceAll(serviceId, "UUID", "Uuid")
	serviceId = strings.ReplaceAll(serviceId, "SSID", "Ssid")
	serviceId = strings.ReplaceAll(serviceId, "DECT", "Dect")
	serviceId = strings.ReplaceAll(serviceId, "UPnP", "Upnp")
	serviceId = strings.ReplaceAll(serviceId, "DNS", "Dns")
	serviceId = strings.ReplaceAll(serviceId, "FTP", "Ftp")
	serviceId = strings.ReplaceAll(serviceId, "NTP", "Ntp")
	serviceId = strings.ReplaceAll(serviceId, "SMB", "Smb")
	serviceId = strings.ReplaceAll(serviceId, "SID", "Sid")
	serviceId = strings.ReplaceAll(serviceId, "UID", "Uid")
	serviceId = strings.ReplaceAll(serviceId, "URL", "Url")
	serviceId = strings.ReplaceAll(serviceId, "WAN", "Wan")
	serviceId = strings.ReplaceAll(serviceId, "VPN", "Vpn")
	serviceId = strings.ReplaceAll(serviceId, "PFS", "Pfs")
	serviceId = strings.ReplaceAll(serviceId, "DSL", "Dsl")
	serviceId = strings.ReplaceAll(serviceId, "LAN", "Lan")
	serviceId = strings.ReplaceAll(serviceId, "USP", "Usp")
	serviceId = strings.ReplaceAll(serviceId, "TAM", "Tam")
	serviceId = strings.ReplaceAll(serviceId, "MAC", "Mac")
	serviceId = strings.ReplaceAll(serviceId, "CGI", "CGI")
	serviceId = strings.ReplaceAll(serviceId, "IP", "Ip")
	serviceId = strings.ReplaceAll(serviceId, "ID", "Id")
	serviceId = strings.ReplaceAll(serviceId, "TR", "Tr")
	serviceId = strings.ReplaceAll(serviceId, "TV", "Tv")
	serviceId = strings.ReplaceAll(serviceId, "X_", "X")
	sb := strings.Builder{}
	for i := 0; i < len(serviceId); i++ {
		s := serviceId[i]
		if unicode.IsUpper(rune(s)) && (sb.Len() > 0) {
			sb.WriteString("_")
		}
		sb.WriteString(strings.ToLower(string(s)))
	}
	return sb.String()
}

func determineFileName(deviceType string, serviceId string, actionName string) (result string) {
	var packageName string
	switch deviceType {
	case "urn:dslforum-org:device:InternetGatewayDevice:1":
		packageName = "gateway"
	case "urn:dslforum-org:device:LANDevice:1":
		packageName = "lan"
	case "urn:dslforum-org:device:WANDevice:1":
		packageName = "wan"
	default:
		panic(errors.New("missing mapping for deviceType '" + deviceType + "'"))
	}

	serviceGroup := strings.Split(serviceId, ":")[1]
	if !strings.HasSuffix(serviceGroup, "-com") {
		panic(errors.New("unsupported serviceId mapping for '" + serviceId + "'"))
	}
	serviceGroup = serviceGroup[:len(serviceGroup)-4]
	serviceGroup = serviceId2SnakeCase(serviceGroup)

	fileName := serviceId2SnakeCase(actionName) + ".go"
	return fmt.Sprintf("pkg/tr064model/%s_%s_%s", packageName, serviceGroup, fileName)
}

func generateResponseStruct(deviceType string, serviceId string, rootSpec scpd.ServiceControlledProtocolDescriptions, serviceSpec scpd.ServiceControlledProtocolDescriptions) {

	for _, action := range serviceSpec.ActionList {
		tmpl, err := template.New("struct.go.tmpl").Parse(templateStructGo)
		if err != nil {
			panic(err)
		}
		variables := []string{}
		varNameMaxLen := 0
		outArguments := filterArguments(action.ArgumentList, "out")
		// create variable declaration
		for _, argument := range outArguments {
			varName := determineVariableName(argument.Name)
			varDeclaration := fmt.Sprintf("%s\t%s `xml:\"%s\"`",
				varName,
				determineTypeName(serviceSpec, argument.RelatedStateVariable),
				argument.Name,
			)
			variables = append(variables, varDeclaration)
			if len(varName) > varNameMaxLen {
				varNameMaxLen = len(varName)
			}
		}
		// fix block spacing for alignment
		var varDeclMaxLen = 0
		for i, _ := range variables {
			varDecl := variables[i]
			idx := strings.Index(varDecl, "\t")
			varDecl = strings.Replace(varDecl, "\t", spaces[:varNameMaxLen-idx+1], 1)
			if len(varDecl) > varDeclMaxLen {
				varDeclMaxLen = len(varDecl)
			}
			variables[i] = varDecl
		}
		// append commends with block ident
		for i, _ := range variables {
			varDecl := variables[i]
			comment := determineVariableComment(serviceSpec, outArguments[i].RelatedStateVariable)
			if len(comment) > 0 {
				varDecl = varDecl + spaces[:varDeclMaxLen-len(varDecl)+1] + comment
			}
			variables[i] = varDecl
		}
		// render the template
		sd := structData{
			Name:           action.Name + "Response",
			SCDPShortName:  deriveScdpShortName(findScdpUrlPath(rootSpec, serviceId)),
			SCDPUrlPath:    findScdpUrlPath(rootSpec, serviceId),
			SystemVersion:  rootSpec.SystemVersion.Display,
			SoapActionName: action.Name,
			Variables:      variables,
		}
		sb := strings.Builder{}
		err = tmpl.Execute(&sb, sd)
		if err != nil {
			panic(err)
		}
		if !debug {
			err := os.WriteFile(determineFileName(deviceType, serviceId, action.Name), []byte(sb.String()), 0644)
			if err != nil {
				panic(err)
			}
		} else {
			println("===== " + determineFileName(deviceType, serviceId, action.Name) + "=====")
			println(sb.String())
		}
	}

}

func generateModels(description tr64Desc) {
	for _, service := range description.services {
		generateResponseStruct(service.deviceType, service.serviceId, description.root, service.spec)
	}
}

func main() {
	tr64SDescription := fetchAllTr64SDescription()
	println("URL..............: " + tr64SDescription.root.Device.PresentationURL)
	println("Model............: " + tr64SDescription.root.Device.ModelDescription)
	println("System-Version...: " + tr64SDescription.root.SystemVersion.Display)
	println("No# services.....: " + fmt.Sprintf("%d", len(tr64SDescription.services)))
	for _, service := range tr64SDescription.services {
		println(fmt.Sprintf("- (%s) - %s", service.deviceType, service.serviceId))
	}

	generateModels(tr64SDescription)
}
