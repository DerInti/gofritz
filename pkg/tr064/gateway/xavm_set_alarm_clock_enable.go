package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// XavmSetAlarmClockEnable AUTO-GENERATED (do not edit) code from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_SetAlarmClockEnable', Fritz!Box-System-Version 141.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
func XavmSetAlarmClockEnable(session *soap.SoapSession, index int, avmAlarmClockEnable bool) (tr064model.XavmSetAlarmClockEnableResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_voip").
		Uri("urn:dslforum-org:service:X_VoIP:1").
		Action("X_AVM-DE_SetAlarmClockEnable").
		AddIntParam("NewIndex", index).
		AddBoolParam("NewX_AVM-DE_AlarmClockEnable", avmAlarmClockEnable).
		Do().Body.Data
	result := tr064model.XavmSetAlarmClockEnableResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}