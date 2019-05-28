package flex

type ComponentSettings struct {
	Settings struct {
		Type          string `json:"type"`
		Setup         Setup  `json:"setup"`
		DataProvider  string `json:"dataProvider"`
		Format        string `json:"format"`
		ContentType   string `json:"contentType"`
		ChildElements struct {
			Image struct {
				Value struct {
					SettingType string `json:"settingType"`
					Static      string `json:"static"`
				} `json:"value"`
				Setup []struct {
					Value string `json:"value"`
					Type  string `json:"type"`
					Setup struct {
						Value string `json:"value"`
					} `json:"setup"`
				} `json:"setup"`
				Action ActionCS `json:"action"`
			} `json:"image"`
			Ratio struct {
				Value string `json:"value"`
				Setup []struct {
					Type  string `json:"type"`
					Value string `json:"value"`
					Setup struct {
						Value string `json:"value"`
					} `json:"setup"`
				} `json:"setup"`
				Action ActionCS `json:"action"`
			} `json:"ratio"`
			Action struct {
				Value string `json:"value"`
				Setup []struct {
					Type  string `json:"type"`
					Value string `json:"value"`
					Setup struct {
						Value string `json:"value"`
					} `json:"setup"`
				} `json:"setup"`
				Action ActionCS `json:"action"`
			} `json:"action"`
		} `json:"childElements"`
	} `json:"settings"`
	Elements struct {
		Title struct {
			Value struct {
				SettingType string `json:"settingType"`
				Static      string `json:"static"`
			} `json:"value"`
			Setup []struct {
				Type  string `json:"type"`
				Setup struct {
				} `json:"setup"`
			} `json:"setup"`
			Action ActionCS `json:"action"`
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
			Action ActionCS `json:"action"`
		} `json:"subTitle"`
		Icon struct {
			Value string `json:"value"`
			Setup []struct {
				Type  string `json:"type"`
				Value string `json:"value"`
				Setup struct {
					Value string `json:"value"`
				} `json:"setup"`
			} `json:"setup"`
			Action ActionCS `json:"action"`
		} `json:"icon"`
		MoreTitle struct {
			Value  string           `json:"value"`
			Setup  []moreTitleSetup `json:"setup"`
			Action ActionCS         `json:"action"`
		} `json:"moreTitle"`
	} `json:"elements"`
}

type moreTitleSetup struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	Setup struct {
		Value string `json:"value"`
	} `json:"setup"`
}

type ActionCS struct {
	Type  string `json:"type"`
	Setup Setup  `json:"setup"`
}
