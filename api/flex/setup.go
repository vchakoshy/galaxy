package flex

import (
	"strings"
)

type Setup struct {
	Book         ProviderSetup       `json:"book"`
	Author       ProviderSetup       `json:"author"`
	Category     ProviderSetup       `json:"category"`
	Channel      ProviderSetup       `json:"channel"`
	Publisher    ProviderSetup       `json:"publisher"`
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
		Value        QueryIdis   `json:"value"`
	} `json:"contentType"`
	Format struct {
		DataProvider interface{} `json:"dataProvider"`
		Type         string      `json:"type"`
		Value        QueryIdis   `json:"value"`
	} `json:"format"`
}

func (s Setup) getInputActions() []ComponentAction {
	q := make([]ComponentAction, 0)

	f := []ComponentAction{
		s.Book.getInputAction("book"),
		s.Free.getInputAction("free"),
		s.Size.getInputAction("size"),
		s.Subscription.getInputAction("subscription"),
		s.Category.getInputAction("category"),
	}

	for _, ca := range f {
		if ca.Key != "" {
			q = append(q, ca)
		}
	}

	return q
}

type QueryIdis []interface{}

func (q QueryIdis) getInterfaceList() []interface{} {
	qidis := make([]interface{}, 0)
	for _, id := range q {
		qidis = append(qidis, id)
	}

	return qidis
}

func (q QueryIdis) getInterfaceStringLowerList() []interface{} {
	contentTypeList := q.getInterfaceList()

	for index, ct := range contentTypeList {
		contentTypeList[index] = strings.ToLower(ct.(string))
	}

	return contentTypeList
}
