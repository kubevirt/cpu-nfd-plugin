package collector

//HostDomCapabilities represents structure for parsing output of virsh capabilities
type HostDomCapabilities struct {
	CPU CPU `xml:"cpu"`
}

//CPU represents slice of cpu modes
type CPU struct {
	Mode []Mode `xml:"mode"`
}

//Mode represents slice of cpu models
type Mode struct {
	Model []Model `xml:"model"`
}

//Model represents cpu model
type Model struct {
	Name   string `xml:",chardata"`
	Usable string `xml:"usable,attr"`
}
