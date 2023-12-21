package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetTicketIDStatus AUTO-GENERATED (do not edit) code from [x_hostfilterSCPD],
// based on SOAP action 'GetTicketIDStatus', Fritz!Box-System-Version 164.07.57
//
// [x_hostfilterSCPD]: http://fritz.box:49000/x_hostfilterSCPD.xml
func GetTicketIDStatus(session *soap.SoapSession) (tr064model.GetTicketIDStatusResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_hostfilter").
		Uri("urn:dslforum-org:service:X_AVM-DE_HostFilter:1").
		Action("GetTicketIDStatus").
		Do().Body.Data
	result := tr064model.GetTicketIDStatusResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
