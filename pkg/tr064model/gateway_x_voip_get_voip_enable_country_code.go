package tr064model

// GetVoIPEnableCountryCodeResponse auto generated model from [x_voipSCPD],
// based on SOAP action 'GetVoIPEnableCountryCode', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type GetVoIPEnableCountryCodeResponse struct {
	VoIPEnableCountryCode bool `xml:"NewVoIPEnableCountryCode"` // default=0
}