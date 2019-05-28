package flex

type DataProvidersProposedList struct{}

func (b DataProvidersProposedList) getGeneric(cs ComponentSettings, t string) OutputComponent {
	com := OutputComponent{
		Type:         t,
		ResourceType: "PROPOSED_LIST",
		Title:        cs.Elements.Title.Value.Static,
	}

	a := getAction(cs)
	if a.Type != "" {
		com.Action = a
	}

	if cs.Elements.MoreTitle.Value != "" {
		com.ActionTitle = cs.Elements.MoreTitle.Value
	}

	return com
}
