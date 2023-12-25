package tr064model

import "encoding/xml"

// GetFilelinkListPathResponse AUTO-GENERATED (do not edit) model from [x_filelinksSCPD],
// based on SOAP action 'GetFilelinkListPath', Fritz!Box-System-Version 141.07.57
//
// [x_filelinksSCPD]: http://fritz.box:49000/x_filelinksSCPD.xml
type GetFilelinkListPathResponse struct {
	XMLName          xml.Name // rather for debug purpose
	FilelinkListPath string   `xml:"NewFilelinkListPath"` //
}
