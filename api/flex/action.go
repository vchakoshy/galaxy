package flex

type FlexAction struct {
	Method string        `json:"method"`
	Type   string        `json:"type"`
	Input  []ActionInput `json:"input"`
}
