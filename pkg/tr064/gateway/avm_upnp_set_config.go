package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetAvmUpnpConfig AUTO-GENERATED (do not edit) code from [x_upnpSCPD],
// based on SOAP action 'SetConfig', Fritz!Box-System-Version 164.07.57
//
// [x_upnpSCPD]: http://fritz.box:49000/x_upnpSCPD.xml
func SetAvmUpnpConfig(session *soap.SoapSession) (tr064model.SetAvmUpnpConfigResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_upnp").
		Uri("urn:dslforum-org:service:X_AVM-DE_UPnP:1").
		Action("SetConfig").
		Do().Body.Data
	result := tr064model.SetAvmUpnpConfigResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}