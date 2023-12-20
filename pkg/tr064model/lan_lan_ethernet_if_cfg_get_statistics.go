package tr064model

// GetLanEthernetIfCfgStatisticsResponse auto generated model from [ethifconfigSCPD],
// based on SOAP action 'GetStatistics', Fritz!Box-System-Version 164.07.57
//
// [ethifconfigSCPD]: http://fritz.box:49000/ethifconfigSCPD.xml
type GetLanEthernetIfCfgStatisticsResponse struct {
	BytesSent       int `xml:"NewBytesSent"`       // default=0
	BytesReceived   int `xml:"NewBytesReceived"`   // default=0
	PacketsSent     int `xml:"NewPacketsSent"`     // default=0
	PacketsReceived int `xml:"NewPacketsReceived"` // default=0
}