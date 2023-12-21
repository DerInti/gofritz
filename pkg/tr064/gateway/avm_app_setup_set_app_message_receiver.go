package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetAppMessageReceiver AUTO-GENERATED (do not edit) code from [x_appsetupSCPD],
// based on SOAP action 'SetAppMessageReceiver', Fritz!Box-System-Version 164.07.57
//
// [x_appsetupSCPD]: http://fritz.box:49000/x_appsetupSCPD.xml
func SetAppMessageReceiver(session *soap.SoapSession) (tr064model.SetAppMessageReceiverResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_appsetup").
		Uri("urn:dslforum-org:service:X_AVM-DE_AppSetup:1").
		Action("SetAppMessageReceiver").
		Do().Body.Data
	result := tr064model.SetAppMessageReceiverResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
