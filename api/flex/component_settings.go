package flex

type flexComponentSettings struct {
	Settings struct {
		Type         string `json:"type"`
		Setup        Setup  `json:"setup"`
		DataProvider string `json:"dataProvider"`
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
		MoreTitle struct {
			Value  string            `json:"value"`
			Setup  []moreTtitleSetup `json:"setup"`
			Action struct {
				Type  string `json:"type"`
				Setup Setup  `json:"setup"`
			} `json:"action"`
		} `json:"moreTitle"`
	} `json:"elements"`
}

type moreTtitleSetup struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	Setup struct {
		Value string `json:"value"`
	} `json:"setup"`
}
