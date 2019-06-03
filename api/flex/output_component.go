package flex

// for more information please look at:
// https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
// https://www.reddit.com/r/golang/comments/6mtbx9/useful_constructors_in_go/

// OutputComponent struct
type OutputComponent struct {
	Title        string  `json:"title,omitempty"`
	Icon         string  `json:"icon,omitempty"`
	SubTitle     string  `json:"subTitle,omitempty"`
	Type         string  `json:"type,omitempty"`
	ResourceType string  `json:"resource_type,omitempty"`
	Action       *Action `json:"action,omitempty"`
	ActionTitle  string  `json:"actionTitle,omitempty"`
	Data         struct {
		Items struct {
			Generic []Generic     `json:"generic"`
			Model   []interface{} `json:"model"`
		} `json:"items"`
	} `json:"data,omitempty"`
}

// OutputComponentOptionFunc function
type OutputComponentOptionFunc func(*OutputComponent) error

// NewOutputComponent return new output comment
func NewOutputComponent(options ...OutputComponentOptionFunc) (*OutputComponent, error) {
	o := &OutputComponent{}
	for _, option := range options {
		if err := option(o); err != nil {
			return o, err
		}
	}

	o.Data.Items.Generic = make([]Generic, 0)
	o.Data.Items.Model = make([]interface{}, 0)

	return o, nil
}

// OutputComponentSetTitle set title of struct
func OutputComponentSetTitle(title string) OutputComponentOptionFunc {
	return func(o *OutputComponent) error {
		o.Title = title
		return nil
	}
}

// OutputComponentSetSubTitle set subtitle of struct
func OutputComponentSetSubTitle(title string) OutputComponentOptionFunc {
	return func(o *OutputComponent) error {
		o.SubTitle = title
		return nil
	}
}

// OutputComponentSetResourceType set resource type of struct
func OutputComponentSetResourceType(t string) OutputComponentOptionFunc {
	return func(o *OutputComponent) error {
		o.ResourceType = t
		return nil
	}
}

// OutputComponentSetType set  type of struct
func OutputComponentSetType(t string) OutputComponentOptionFunc {
	return func(o *OutputComponent) error {
		o.Type = t
		return nil
	}
}
