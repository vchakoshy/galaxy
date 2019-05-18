package rest

type ProviderSetup struct {
	Type    string   `json:"type"`
	Mapping string   `json:"mapping"`
	Ids     []string `json:"ids"`
	Setup   struct {
	} `json:"setup"`
	DataProvider string `json:"dataProvider"`
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
