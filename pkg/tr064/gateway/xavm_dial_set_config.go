package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// XavmDialSetConfig AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_DialSetConfig', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func XavmDialSetConfig(session *soap.SoapSession, avmPhoneName string) (tr064model.XavmDialSetConfigResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("X_AVM-DE_DialSetConfig").
		AddStringParam("NewX_AVM-DE_PhoneName", avmPhoneName).
		Do().Body.Data
	result := tr064model.XavmDialSetConfigResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
