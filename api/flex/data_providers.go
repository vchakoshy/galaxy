package flex

// DataProvider interface
type DataProvider interface {
	getGeneric(flexComponentSettings, string) []interface{}
}
