package rest

import (
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type ProviderSetup struct {
	Type         string   `json:"type"`
	Mapping      string   `json:"mapping"`
	Ids          []string `json:"ids"`
	Setup        struct{} `json:"setup"`
	DataProvider string   `json:"dataProvider"`
}

func (ps ProviderSetup) getQuery() []qm.QueryMod {
	queries := make([]qm.QueryMod, 0)

	qidis := make([]interface{}, len(ps.Ids))
	for _, id := range ps.Ids {
		qidis = append(qidis, id)
	}

	if len(qidis) > 0 {
		queries = append(queries, qm.WhereIn("id in ?", qidis...))
	}

	return queries
}

func (s Setup) getQueries() []qm.QueryMod {
	q := make([]qm.QueryMod, 0)

	q = append(q, s.Book.getQuery()...)
	q = append(q, s.Author.getQuery()...)
	q = append(q, s.Category.getQuery()...)
	q = append(q, s.Channel.getQuery()...)
	q = append(q, s.Publisher.getQuery()...)
	q = append(q, s.ProposedList.getQuery()...)
	q = append(q, s.Tag.getQuery()...)

	return q
}

func (s Setup) getSort() []qm.QueryMod {
	q := make([]qm.QueryMod, 0)

	switch s.Sort.Value {
	case "RECENT":
		q = append(q, qm.OrderBy("id DESC"))
	}

	return q
}

type Setup struct {
	Book         ProviderSetup `json:"book"`
	Author       ProviderSetup `json:"author"`
	Category     ProviderSetup `json:"category"`
	Channel      ProviderSetup `json:"channel"`
	Publisher    ProviderSetup `json:"publisher"`
	ProposedList ProviderSetup `json:"proposedList"`
	Tag          ProviderSetup `json:"tag"`
	Sort         struct {
		DataProvider interface{} `json:"dataProvider"`
		Type         string      `json:"type"`
		Value        string      `json:"value"`
	} `json:"sort"`
	Query struct {
		DataProvider interface{}   `json:"dataProvider"`
		Type         string        `json:"type"`
		Fields       []interface{} `json:"fields"`
	} `json:"query"`
	ContentType struct {
		DataProvider interface{}   `json:"dataProvider"`
		Type         string        `json:"type"`
		Value        []interface{} `json:"value"`
	} `json:"contentType"`
	Format struct {
		DataProvider interface{} `json:"dataProvider"`
		Type         string      `json:"type"`
		Value        string      `json:"value"`
	} `json:"format"`
	Free struct {
		DataProvider interface{} `json:"dataProvider"`
		Type         string      `json:"type"`
		Value        string      `json:"value"`
	} `json:"free"`
	Subscription struct {
		DataProvider interface{} `json:"dataProvider"`
		Type         string      `json:"type"`
		Value        string      `json:"value"`
	} `json:"subscription"`
	Size struct {
		DataProvider interface{} `json:"dataProvider"`
		Type         string      `json:"type"`
		Value        string      `json:"value"`
	} `json:"size"`
	Filter struct {
		DataProvider interface{} `json:"dataProvider"`
		Type         string      `json:"type"`
		Value        string      `json:"value"`
	} `json:"filter"`
}
