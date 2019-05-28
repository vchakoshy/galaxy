package flex

// DataProvider interface
type DataProvider interface {
	getGeneric(ComponentSettings, string) []interface{}
}
