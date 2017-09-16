package contract

type Rig struct {
	Register struct {
		Hostname string `json:"hostname"`
		JWT      string `json:"jwt"`
	}

	StreamInit struct {
		Name string `json:"name"`
	}

	StreamDrop struct {
		Name string `json:"name"`
	}
}

type Agent struct {
	Register struct {
		Hostname string `json:"hostname"`
		Secret   string `json:"secret"`
	}

	StreamInit struct {
		Name  string `json:"name"`
		Token string `json:"token"`
		JWT   string `json:"jwt"`
	}
}
