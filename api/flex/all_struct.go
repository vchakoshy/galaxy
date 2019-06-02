package flex

type ActionInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Tabs struct {
	Title  string `json:"title"`
	PageID string `json:"pageId"`
	Active bool   `json:"active"`
	Action Action `json:"action"`
}

type Setting struct {
	Tabs []Tabs `json:"tabs"`
}

type InputAction struct {
	Key        string      `json:"key"`
	ArrayValue []string    `json:"arrayValue,omitempty"`
	Value      interface{} `json:"value,omitempty"`
}

type Action struct {
	Type      string            `json:"type"`
	Input     []InputAction `json:"input"`
	ExtraData interface{}       `json:"extraData,omitempty"`
	Method    string            `json:"method"`
}

type OutputComponent struct {
	Title string `json:"title,omitempty"`
	Icon  string `json:"icon,omitempty"`
	Data  struct {
		Items struct {
			Generic []Generic     `json:"generic"`
			Model   []interface{} `json:"model"`
		} `json:"items"`
	} `json:"data,omitempty"`
	Type         string  `json:"type,omitempty"`
	ResourceType string  `json:"resource_type,omitempty"`
	Action       *Action `json:"action,omitempty"`
	ActionTitle  string  `json:"actionTitle,omitempty"`
}

type Output struct {
	Components []OutputComponent `json:"components"`
	FlexErrors []string          `json:"flexErrors"`
	IsLastPage bool              `json:"isLastPage"`
	Setting    Setting           `json:"setting"`
	Title      string            `json:"title"`
	Result     bool              `json:"result"`
}

type Response struct {
	Output     Output `json:"output"`
	Error      string `json:"error"`
	Message    string `json:"message"`
	IsPlusMode bool   `json:"isPlusMode"`
}
