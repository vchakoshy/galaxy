package flex

type ProviderSimpleSetup struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

func (ps ProviderSimpleSetup) getInputAction(key string) (action flexComponentAction) {
	if ps.Type == "" {
		return
	}

	switch ps.Value.(type) {
	case string:
		if ps.Value == "" {
			return
		}
	}

	action.Key = key
	action.Value = ps.Value

	return
}
