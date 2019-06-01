package flex

type Action struct {
	PanelType  string
	ClientType string
	Method     string
}

// Action types definition
const (
	ActionContentListPanelType      = "CONTENT_LIST"
	ActionBookPanelType             = "BOOK_PAGE"
	ActionProposedListPagePanelType = "PROPOSED_LIST_PAGE"
	ActionPublisherListPanelType    = "PUBLISHER_LIST"
	ActionNewsListPanelType         = "NEWS_LIST"
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

func PublisherListAction() Action {
	return Action{
		PanelType:  "PUBLISHER_LIST",
		ClientType: "publisher_list",
		Method:     "/v2/general/list/publisher",
	}
}

func ProposedListPageAction() Action {
	return Action{
		PanelType:  "PROPOSED_LIST_PAGE",
		ClientType: "proposed_list_page",
		Method:     "/general/proposed-list/get",
	}
}

func NewsListAction() Action {
	return Action{
		PanelType:  "NEWS_LIST",
		ClientType: "news_list",
		Method:     "/v2/general/list/news",
	}
}

func GetActionByPanelType(t string) (action Action) {
	switch t {
	case ActionContentListPanelType:
		action = ContentListAction()
	case ActionBookPanelType:
		action = BookAction()
	case ActionProposedListPagePanelType:
		action = ProposedListPageAction()
	case ActionPublisherListPanelType:
		action = PublisherListAction()
	default:
		action = Action{}
	}
	return
}
