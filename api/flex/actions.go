package flex

type Action struct {
	PanelType  string
	ClientType string
	Method     string
}

func ContentListAction() Action {
	return Action{
		PanelType:  "CONTENT_LIST",
		ClientType: "content_list",
		Method:     "/v2/general/list/book",
	}
}

func BookAction() Action {
	return Action{
		PanelType:  "BOOK_PAGE",
		ClientType: "book",
		Method:     "/",
	}
}

func GetActionByPanelType(t string) (action Action) {
	switch t {
	case ContentListAction().PanelType:
		action = ContentListAction()
	case BookAction().PanelType:
		action = BookAction()
	default:
		action = Action{}
	}
	return
}
