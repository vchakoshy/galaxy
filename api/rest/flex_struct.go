package rest

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

type flexOutput struct {
	Components []string    `json:"components"`
	FlexErrors []string    `json:"flexErrors"`
	IsLastPage bool        `json:"isLastPage"`
	Setting    flexSetting `json:"setting"`
	Title      string      `json:"title"`
	Result     bool        `json:"result"`
}

type flexStruct struct {
	Output     flexOutput `json:"output"`
	Error      string     `json:"error"`
	Message    string     `json:"message"`
	IsPlusMode bool       `json:"isPlusMode"`
}
