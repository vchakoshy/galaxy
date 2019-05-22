package flex 


type FlexActionInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}



type flexTabs struct {
	Title  string     `json:"title"`
	PageID string     `json:"pageId"`
	Active bool       `json:"active"`
	Action FlexAction `json:"action"`
}

type Setting struct {
	Tabs []flexTabs `json:"tabs"`
}

type flexComponentAction struct {
	Key        string   `json:"key"`
	ArrayValue []string `json:"arrayValue,omitempty"`
	Value      bool     `json:"value,omitempty"`
}

type flexBaseAction struct {
	Type      string                `json:"type"`
	Input     []flexComponentAction `json:"input"`
	ExtraData interface{}           `json:"extraData"`
	Method    string                `json:"method"`
}

type flexComponent struct {
	Title string `json:"title,omitempty"`
	Icon  string `json:"icon,omitempty"`
	Data  struct {
		Items struct {
			Generic []interface{} `json:"generic"`
		} `json:"items"`
	} `json:"data"`
	Type         string         `json:"type"`
	ResourceType string         `json:"resource_type"`
	Action       flexBaseAction `json:"action,omitempty"`
	ActionTitle  string         `json:"actionTitle,omitempty"`
}

type FlexOutput struct {
	Components []flexComponent `json:"components"`
	FlexErrors []string        `json:"flexErrors"`
	IsLastPage bool            `json:"isLastPage"`
	Setting    Setting     `json:"setting"`
	Title      string          `json:"title"`
	Result     bool            `json:"result"`
}

type FlexStruct struct {
	Output     FlexOutput `json:"output"`
	Error      string     `json:"error"`
	Message    string     `json:"message"`
	IsPlusMode bool       `json:"isPlusMode"`
}
