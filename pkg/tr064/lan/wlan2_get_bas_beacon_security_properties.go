package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// Wlan2GetBasBeaconSecurityProperties AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'GetBasBeaconSecurityProperties', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func Wlan2GetBasBeaconSecurityProperties(session *soap.SoapSession) (tr064model.GetBasBeaconSecurityPropertiesResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig2").
		Uri("urn:dslforum-org:service:WLANConfiguration:2").
		Action("GetBasBeaconSecurityProperties").
		Do().Body.Data
	result := tr064model.GetBasBeaconSecurityPropertiesResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
