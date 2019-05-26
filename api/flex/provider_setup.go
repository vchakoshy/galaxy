package flex

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

func (ps ProviderSetup) GetKeyID() string {
	switch ps.DataProvider {
	case "BOOK":
		return "bookId"
	case "CATEGORY":
		return "categoryId"
	}

	return ""
}

func (ps ProviderSetup) getInputAction(key string) flexComponentAction {
	action := flexComponentAction{}

	switch ps.Type {
	case "STATIC":
		action.Key = ps.GetKeyID()
		action.ArrayValue = ps.Ids
	}

	return action
}

func (ps ProviderSetup) getQuery() []qm.QueryMod {

	queries := make([]qm.QueryMod, 0)

	qidis := make([]interface{}, 0)
	for _, id := range ps.Ids {
		qidis = append(qidis, id)
	}

	if len(qidis) > 0 {
		queries = append(queries, qm.WhereIn("id in ?", qidis...))
	}

	return queries
}