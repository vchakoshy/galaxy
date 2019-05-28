package flex

func handleListComponent(cs flexComponentSettings, t string) flexComponent {

	switch cs.Settings.DataProvider {
	case "BOOK":
		com := DataProvidersBook{}
		return com.getGeneric(cs, t)

	case "PROPOSED_LIST":
		com := DataProvidersProposedList{}
		return com.getGeneric(cs, t)
	}

	return flexComponent{}
}

func handleBooksLibraryComponent(cs flexComponentSettings) (com flexComponent) {
	m := make(map[string]string)
	m["format"] = cs.Settings.Format
	m["content_type"] = cs.Settings.ContentType
	com.Data.Items.Generic = append(com.Data.Items.Generic, m)
	com.Title = cs.Elements.Title.Value.Static
	com.Icon = cs.Elements.Icon.Value
	com.Type = "HL_BOOKS_LIBRARY"
	com.ResourceType = "CUSTOM"
	return
}
