package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// XavmGetChangeCounter AUTO-GENERATED (do not edit) code from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_GetChangeCounter', Fritz!Box-System-Version 141.07.57
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
func XavmGetChangeCounter(session *soap.SoapSession) (tr064model.XavmGetChangeCounterResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_GetChangeCounter").
		Do().Body.Data
	result := tr064model.XavmGetChangeCounterResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}