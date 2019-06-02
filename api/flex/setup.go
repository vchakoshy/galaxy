package flex

import (
	"github.com/olivere/elastic"
)

// Setup struct
type Setup struct {
	Book         ProviderSetup       `json:"book"`
	Author       ProviderSetup       `json:"author"`
	Category     ProviderSetup       `json:"category"`
	Channel      ProviderSetup       `json:"channel"`
	Publisher    ProviderSetup       `json:"publisher"`
	News         ProviderSetup       `json:"news"`
	ProposedList ProviderSetup       `json:"proposedList"`
	Tag          ProviderSetup       `json:"tag"`
	Free         ProviderSimpleSetup `json:"free"`
	Subscription ProviderSimpleSetup `json:"subscription"`
	Size         ProviderSimpleSetup `json:"size"`
	Filter       ProviderSimpleSetup `json:"filter"`
	Sort         ProviderSimpleSetup `json:"sort"`
	Query        struct {
		DataProvider interface{}   `json:"dataProvider"`
		Type         string        `json:"type"`
		Fields       []interface{} `json:"fields"`
	} `json:"query"`
	ContentType struct {
		DataProvider interface{} `json:"dataProvider"`
		Type         string      `json:"type"`
		Value        queryIdis   `json:"value"`
	} `json:"contentType"`
	Format Format `json:"format"`
}

type Format struct {
	DataProvider interface{} `json:"dataProvider"`
	Type         string      `json:"type"`
	Value        queryIdis   `json:"value"`
}

func (ps Format) getInputAction(key string) (action InputAction) {
	if ps.Type == "" {
		return
	}

	action.Key = key
	action.Value = ps.Value

	return
}

// GetQueries of setup
func (s Setup) GetQueries() *elastic.BoolQuery {
	q := elastic.NewBoolQuery()

	bookIdis := s.Book.GetIdis()
	if len(bookIdis) > 0 {
		q.Must(elastic.NewTermsQuery("id", bookIdis...))
	}

	catIds := s.Category.GetIdis()
	if len(catIds) > 0 {
		q.Must(elastic.NewBoolQuery().
			Should(elastic.NewTermsQuery("categories.id", catIds...)))
	}

	formatList := s.Format.Value.getInterfaceList()
	if len(formatList) > 0 {
		q.Must(elastic.NewTermsQuery("format.keyword", formatList...))
	}

	contentTypeList := s.ContentType.Value.getInterfaceStringLowerList()
	if len(contentTypeList) > 0 {
		q.Must(elastic.NewTermsQuery("content_type.keyword", contentTypeList...))
	}

	return q
}

func (s Setup) getInputActions() []InputAction {
	q := make([]InputAction, 0)

	// TODO: complete input actions
	f := []InputAction{
		s.Book.getInputAction("book"),
		s.Free.getInputAction("free"),
		s.Size.getInputAction("size"),
		s.Subscription.getInputAction("subscription"),
		s.Category.getInputAction("category"),
		s.Filter.getInputAction("filter"),
		s.Format.getInputAction("format"),

		// TODO: @ali this line not working
		s.Sort.getInputAction("sort"),
	}

	for _, ca := range f {
		if ca.Key != "" {
			q = append(q, ca)
		}
	}

	return q
}
