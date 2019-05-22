package flex

type Field struct {
	Page   string
	Panel  string
	Method string
	Action string
}

func CategoryField() Field {
	return Field{
		Page:   "categoryId",
		Panel:  "category",
		Method: "category",
		Action: "categoryId",
	}
}
