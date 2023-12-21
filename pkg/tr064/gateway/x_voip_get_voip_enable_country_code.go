package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetVoIPEnableCountryCode AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'GetVoIPEnableCountryCode', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func GetVoIPEnableCountryCode(session *soap.SoapSession) (tr064model.GetVoIPEnableCountryCodeResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("GetVoIPEnableCountryCode").
		Do().Body.Data
	result := tr064model.GetVoIPEnableCountryCodeResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
