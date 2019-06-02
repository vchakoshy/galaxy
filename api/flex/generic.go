package flex

// Generic struct
type Generic struct {
	ChildAction  *Action `json:"childAction,omitempty"`
	Title        string  `json:"title,omitempty"`
	SubTitle     string  `json:"subTitle,omitempty"`
	Icon         string  `json:"icon,omitempty"`
	IconSubtitle string  `json:"iconSubtitle,omitempty"`
	IconTitle    string  `json:"iconTitle,omitempty"`
	FooterText   string  `json:"footerText,omitempty"`
	Image        string  `json:"image,omitempty"`
	Ratio        string  `json:"ratio,omitempty"`
	Format       string  `json:"format,omitempty"`
	ContentType  string  `json:"content_type,omitempty"`
	Badge        *Badge  `json:"badge,omitempty"`
	Action       *Action `json:"action,omitempty"`
	BookID       string  `json:"bookId,omitempty"`
	Rate         string  `json:"rate,omitempty"`
}
