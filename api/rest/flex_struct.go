package rest

type flexComponentSettings struct {
	Settings struct {
		Type         string `json:"type"`
		Setup        Setup  `json:"setup"`
		DataProvider string `json:"dataProvider"`
	} `json:"settings"`
	Elements struct {
		Title struct {
			Value struct {
				SettingType string `json:"settingType"`
			} `json:"value"`
			Setup []struct {
				Type  string `json:"type"`
				Setup struct {
				} `json:"setup"`
			} `json:"setup"`
			Action struct {
			} `json:"action"`
		} `json:"title"`
		SubTitle struct {
			Value struct {
				SettingType string `json:"settingType"`
			} `json:"value"`
			Setup []struct {
				Type  string `json:"type"`
				Setup struct {
				} `json:"setup"`
			} `json:"setup"`
			Action struct {
			} `json:"action"`
		} `json:"subTitle"`
	} `json:"elements"`
}

type flexActionInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type flexAction struct {
	Method string            `json:"method"`
	Type   string            `json:"type"`
	Input  []flexActionInput `json:"input"`
}

type flexTabs struct {
	Title  string     `json:"title"`
	PageID string     `json:"pageId"`
	Active bool       `json:"active"`
	Action flexAction `json:"action"`
}

type flexSetting struct {
	Tabs []flexTabs `json:"tabs"`
}

type flexComponent struct {
	Title string `json:"title,omitempty"`
	Icon  string `json:"icon,omitempty"`
	Data  struct {
		Items struct {
			Generic []interface{} `json:"generic"`
		} `json:"items"`
	} `json:"data"`
	Type         string `json:"type"`
	ResourceType string `json:"resource_type"`
	Action       struct {
		Type  string `json:"type"`
		Input []struct {
			Key        string   `json:"key"`
			ArrayValue []string `json:"arrayValue,omitempty"`
			Value      bool     `json:"value,omitempty"`
		} `json:"input"`
		ExtraData interface{} `json:"extraData"`
		Method    string      `json:"method"`
	} `json:"action,omitempty"`
	ActionTitle string `json:"actionTitle,omitempty"`
}

type flexOutput struct {
	Components []flexComponent `json:"components"`
	FlexErrors []string        `json:"flexErrors"`
	IsLastPage bool            `json:"isLastPage"`
	Setting    flexSetting     `json:"setting"`
	Title      string          `json:"title"`
	Result     bool            `json:"result"`
}

type flexStruct struct {
	Output     flexOutput `json:"output"`
	Error      string     `json:"error"`
	Message    string     `json:"message"`
	IsPlusMode bool       `json:"isPlusMode"`
}
