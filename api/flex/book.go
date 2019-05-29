package flex

import (
	"time"

	"gitlab.fidibo.com/backend/galaxy/api/models"
	"gitlab.fidibo.com/backend/galaxy/api/modext"
)

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
	Free            bool        `json:"free"`
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
	PublishDate   time.Time   `json:"publishDate"`
	RateCount     uint        `json:"rate_count"`
	Duration      int         `json:"duration"`
	FileSize      int64       `json:"file_size"`
	CategoryID    string      `json:"category_id"`
	CategoryName  string      `json:"category_name"`
	PageCount     string      `json:"page_count"`
	Price2        string      `json:"price2"`
	Type          string      `json:"type"`
	SampleCrc     uint        `json:"sample_crc"`
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

func newBookByModel(b *models.Book) Book {
	rs := Book{
		BookID:          b.ID,
		BookTitle:       b.Title,
		Price:           b.Price,
		PaperPrice:      b.PaperPrice.Float32,
		BookImage:       modext.GetBookNormalImage(b),
		BookImageSquare: nil,
		Author:          b.R.Author.Name,
		AuthorID:        b.R.Author.ID,
		Translator:      nil,
		TranslatorID:    nil,
		Narrator:        nil,
		NarratorID:      nil,
		PublisherID:     b.R.Publisher.ID,
		PublisherTitle:  b.R.Publisher.Title,
		Rtl:             modext.IsRtl(b),
		Format:          b.Format,
		ContentType:     b.ContentType,
		Rate:            b.Rate,
		Free:            modext.IsFree(b),
		Path:            "",
		PublishDate:     b.PublishDate.Time,
		RateCount:       b.RateCount,
		Duration:        b.Duration.Int,
		FileSize:        b.Filesize.Int64,
		CategoryID:      "",
		CategoryName:    "",
		PageCount:       "",
		Price2:          "",
		Type:            "",
		SampleCrc:       b.SampleCRC.Uint,
		Pass:            b.Password.String,
		Crc:             b.CRC,
		LinkTitle:       "",
		LanguageTitle:   "",
		PriceTitle:      "",
		PriceTitle2:     "",
		Description:     "",
		LinkURL:         "",
		Bought:          false,
		New:             false,
		Featured:        false,
		Favorite:        false,
		MyRate:          0,
		AuthorLogo:      "",
	}

	return rs
}
