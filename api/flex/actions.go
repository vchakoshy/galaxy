package flex

type Action struct {
	PanelType  string
	ClientType string
	Method     string
}

// Action types definition
const (
	ActionContentListPanelType = "CONTENT_LIST"
	ActionBookPanelType        = "BOOK_PAGE"
	ActionFieldStatic          = "STATIC"
)

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
	case ActionContentListPanelType:
		action = ContentListAction()
	case ActionBookPanelType:
		action = BookAction()
	default:
		action = Action{}
	}
	return
}
