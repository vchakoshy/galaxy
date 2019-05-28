package flex

import "gitlab.fidibo.com/backend/galaxy/api/models"

// Book struct
type Book struct {
	BookID          int         `json:"book_id"`
	BookTitle       string      `json:"book_title"`
	Price           int         `json:"price"`
	PaperPrice      float32     `json:"paper_price"`
	BookImage       string      `json:"book_image"`
	BookImageSquare interface{} `json:"book_image_square"`
	Author          string      `json:"author"`
	AuthorID        int         `json:"author_id"`
	Translator      interface{} `json:"translator"`
	TranslatorID    interface{} `json:"translator_id"`
	Narrator        interface{} `json:"narrator"`
	NarratorID      interface{} `json:"narrator_id"`
	PublisherID     int         `json:"publisher_id"`
	PublisherTitle  string      `json:"publisher_title"`
	Rtl             bool        `json:"rtl"`
	Format          string      `json:"format"`
	ContentType     string      `json:"content_type"`
	Rate            float64     `json:"rate"`
	Free            int8        `json:"free"`
	Action          struct {
		Type  string `json:"type"`
		Input []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"input"`
		ExtraData interface{} `json:"extraData"`
		Method    string      `json:"method"`
	} `json:"action"`
	Path  string `json:"path"`
	Badge struct {
		Text     string `json:"text"`
		BgColor  string `json:"bg_color"`
		TxtColor string `json:"txt_color"`
	} `json:"badge"`
	PublishDate   string      `json:"publishDate"`
	RateCount     string      `json:"rate_count"`
	Duration      string      `json:"duration"`
	FileSize      string      `json:"file_size"`
	CategoryID    string      `json:"category_id"`
	CategoryName  string      `json:"category_name"`
	PageCount     string      `json:"page_count"`
	Price2        string      `json:"price2"`
	Type          string      `json:"type"`
	SampleCrc     string      `json:"sample_crc"`
	Pass          string      `json:"pass"`
	Crc           interface{} `json:"crc"`
	LinkTitle     string      `json:"link_title"`
	LanguageTitle string      `json:"languageTitle"`
	PriceTitle    string      `json:"priceTitle"`
	PriceTitle2   string      `json:"priceTitle2"`
	Description   string      `json:"description"`
	LinkURL       string      `json:"link_url"`
	Bought        bool        `json:"bought"`
	New           bool        `json:"new"`
	Featured      bool        `json:"featured"`
	Favorite      bool        `json:"favorite"`
	MyRate        int         `json:"my_rate"`
	AuthorLogo    string      `json:"authorLogo"`
}

func newBookByModel(b models.Book) {

}
