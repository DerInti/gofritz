package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan3XavmGetWLANDeviceListPath AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetWLANDeviceListPath', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan3XavmGetWLANDeviceListPath(session *soap.SoapSession) (tr064model.XavmGetWLANDeviceListPathResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig3").
		Uri("urn:dslforum-org:service:WLANConfiguration:3").
		Action("X_AVM-DE_GetWLANDeviceListPath").
		Do().Body.Data
	result := tr064model.XavmGetWLANDeviceListPathResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
