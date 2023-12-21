package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// X_SetTR069Enable AUTO-GENERATED (do not edit) code from [mgmsrvSCPD],
// based on SOAP action 'X_SetTR069Enable', Fritz!Box-System-Version 164.07.57
//
// [mgmsrvSCPD]: http://fritz.box:49000/mgmsrvSCPD.xml
func X_SetTR069Enable(session *soap.SoapSession) (tr064model.X_SetTR069EnableResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/mgmsrv").
		Uri("urn:dslforum-org:service:ManagementServer:1").
		Action("X_SetTR069Enable").
		Do().Body.Data
	result := tr064model.X_SetTR069EnableResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
