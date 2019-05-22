package flex

import (
	"strconv"
)

type ProviderSimpleSetup struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (ps ProviderSimpleSetup) getInputAction(key string) flexComponentAction {
	action := flexComponentAction{}

	if ps.Value == "" {
		return action
	}

	switch ps.Type {
	case "CUSTOM":
		action.Key = key
		action.Value, _ = strconv.ParseBool(ps.Value)
	}

	return action
}
