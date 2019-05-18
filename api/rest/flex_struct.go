package rest

type flexComponentSettings struct {
	Settings struct {
		Type  string `json:"type"`
		Setup struct {
			Book struct {
				Type    string   `json:"type"`
				Mapping string   `json:"mapping"`
				Ids     []string `json:"ids"`
				Setup   struct {
				} `json:"setup"`
				DataProvider string `json:"dataProvider"`
			} `json:"book"`
			Author struct {
				Type    string        `json:"type"`
				Mapping string        `json:"mapping"`
				Ids     []interface{} `json:"ids"`
				Setup   struct {
				} `json:"setup"`
				DataProvider string `json:"dataProvider"`
			} `json:"author"`
			Category struct {
				Type    string        `json:"type"`
				Mapping string        `json:"mapping"`
				Ids     []interface{} `json:"ids"`
				Setup   struct {
				} `json:"setup"`
				DataProvider string `json:"dataProvider"`
			} `json:"category"`
			Channel struct {
				Type    string        `json:"type"`
				Mapping string        `json:"mapping"`
				Ids     []interface{} `json:"ids"`
				Setup   struct {
				} `json:"setup"`
				DataProvider string `json:"dataProvider"`
			} `json:"channel"`
			Publisher struct {
				Type    string        `json:"type"`
				Mapping string        `json:"mapping"`
				Ids     []interface{} `json:"ids"`
				Setup   struct {
				} `json:"setup"`
				DataProvider string `json:"dataProvider"`
			} `json:"publisher"`
			ProposedList struct {
				Type    string        `json:"type"`
				Mapping string        `json:"mapping"`
				Ids     []interface{} `json:"ids"`
				Setup   struct {
				} `json:"setup"`
				DataProvider string `json:"dataProvider"`
			} `json:"proposedList"`
			Tag struct {
				Type    string        `json:"type"`
				Mapping string        `json:"mapping"`
				Ids     []interface{} `json:"ids"`
				Setup   struct {
				} `json:"setup"`
				DataProvider string `json:"dataProvider"`
			} `json:"tag"`
			Sort struct {
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
		} `json:"setup"`
		DataProvider string `json:"dataProvider"`
	} `json:"settings"`
	Elements struct {
		Title struct {
			Value struct {
				SettingType string `json:"settingType"`
			} `json:"value"`
			Setup []struct {
				Type  string `json:"type"`
				Setup struct {
				} `json:"setup"`
			} `json:"setup"`
			Action struct {
			} `json:"action"`
		} `json:"title"`
		SubTitle struct {
			Value struct {
				SettingType string `json:"settingType"`
			} `json:"value"`
			Setup []struct {
				Type  string `json:"type"`
				Setup struct {
				} `json:"setup"`
			} `json:"setup"`
			Action struct {
			} `json:"action"`
		} `json:"subTitle"`
	} `json:"elements"`
}

type flexGenericChildAction struct {
	Type      string            `json:"type"`
	Input     []flexActionInput `json:"input"`
	ExtraData interface{}       `json:"extraData"`
	Method    string            `json:"method"`
}

type flexGenericBook struct {
	ChildAction flexGenericChildAction `json:"childAction"`
	Title       string                 `json:"title"`
	SubTitle    string                 `json:"subTitle"`
	Icon        string                 `json:"icon"`
	Image       string                 `json:"image"`
	Format      string                 `json:"format"`
	ContentType string                 `json:"content_type"`
	Badge       struct {
		Text     string `json:"text"`
		BgColor  string `json:"bg_color"`
		TxtColor string `json:"txt_color"`
	} `json:"badge"`
	Action flexGenericChildAction `json:"action"`
	BookID string                 `json:"bookId"`
}

type flexActionInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type flexAction struct {
	Method string            `json:"method"`
	Type   string            `json:"type"`
	Input  []flexActionInput `json:"input"`
}

type flexTabs struct {
	Title  string     `json:"title"`
	PageID string     `json:"pageId"`
	Active bool       `json:"active"`
	Action flexAction `json:"action"`
}

type flexSetting struct {
	Tabs []flexTabs `json:"tabs"`
}

type flexComponent struct {
	Title string `json:"title,omitempty"`
	Icon  string `json:"icon,omitempty"`
	Data  struct {
		Items struct {
			Generic []interface{} `json:"generic"`
		} `json:"items"`
	} `json:"data"`
	Type         string `json:"type"`
	ResourceType string `json:"resource_type"`
	Action       struct {
		Type  string `json:"type"`
		Input []struct {
			Key        string   `json:"key"`
			ArrayValue []string `json:"arrayValue,omitempty"`
			Value      bool     `json:"value,omitempty"`
		} `json:"input"`
		ExtraData interface{} `json:"extraData"`
		Method    string      `json:"method"`
	} `json:"action,omitempty"`
	ActionTitle string `json:"actionTitle,omitempty"`
}

type flexOutput struct {
	Components []flexComponent `json:"components"`
	FlexErrors []string        `json:"flexErrors"`
	IsLastPage bool            `json:"isLastPage"`
	Setting    flexSetting     `json:"setting"`
	Title      string          `json:"title"`
	Result     bool            `json:"result"`
}

type flexStruct struct {
	Output     flexOutput `json:"output"`
	Error      string     `json:"error"`
	Message    string     `json:"message"`
	IsPlusMode bool       `json:"isPlusMode"`
}
