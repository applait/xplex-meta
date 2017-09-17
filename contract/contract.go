package contract

type RigRegister struct {
	Hostname string `json:"hostname"`
	JWT      string `json:"jwt"`
}

type RigStreamInit struct {
	Name string `json:"name"`
}

type RigStreamDrop struct {
	Name string `json:"name"`
}

type AgentRegister struct {
	Hostname string `json:"hostname"`
	Secret   string `json:"secret"`
}

type AgentStreamInit struct {
	Name  string `json:"name"`
	Token string `json:"token"`
	JWT   string `json:"jwt"`
}
