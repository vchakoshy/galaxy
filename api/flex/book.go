package flex

import (
	"strings"
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
	BookImageSquare interface{} `json:"book_image_square,omitempty"`
	Author          string      `json:"author"`
	AuthorID        int         `json:"author_id"`
	Translator      interface{} `json:"translator,omitempty"`
	TranslatorID    interface{} `json:"translator_id,omitempty"`
	Narrator        interface{} `json:"narrator,omitempty"`
	NarratorID      interface{} `json:"narrator_id,omitempty"`
	PublisherID     int         `json:"publisher_id"`
	PublisherTitle  string      `json:"publisher_title"`
	Rtl             bool        `json:"rtl"`
	Format          string      `json:"format"`
	ContentType     string      `json:"content_type"`
	Rate            float64     `json:"rate"`
	Free            bool        `json:"free"`
	Action          *Action     `json:"action,omitempty"`
	Path            string      `json:"path"`
	Badge           *Badge      `json:"badge,omitempty"`
	PublishDate     time.Time   `json:"publishDate"`
	RateCount       uint        `json:"rate_count"`
	Duration        int         `json:"duration"`
	FileSize        int64       `json:"file_size,omitempty"`
	CategoryID      string      `json:"category_id"`
	CategoryName    string      `json:"category_name"`
	PageCount       string      `json:"page_count"`
	Price2          string      `json:"price2,omitempty"`
	Type            string      `json:"type,omitempty"`
	SampleCrc       uint        `json:"sample_crc,omitempty"`
	Pass            string      `json:"pass,omitempty"`
	Crc             interface{} `json:"crc,omitempty"`
	LinkTitle       string      `json:"link_title,omitempty"`
	LanguageTitle   string      `json:"languageTitle,omitempty"`
	PriceTitle      string      `json:"priceTitle,omitempty"`
	PriceTitle2     string      `json:"priceTitle2,omitempty"`
	Description     string      `json:"description,omitempty"`
	LinkURL         string      `json:"link_url,omitempty"`
	Bought          bool        `json:"bought"`
	New             bool        `json:"new"`
	Featured        bool        `json:"featured"`
	Favorite        bool        `json:"favorite"`
	MyRate          int         `json:"my_rate"`
	AuthorLogo      string      `json:"authorLogo,omitempty"`
}

func newBookByModel(b *models.Book) Book {
	// TODO Complete the incomplete fields
	rs := Book{
		BookID:          b.ID,
		BookTitle:       b.Title,
		PaperPrice:      b.PaperPrice.Float32,
		BookImage:       modext.GetBookNormalImage(b),
		BookImageSquare: b.ImageSquare,
		Format:          strings.ToLower(b.Format),
		ContentType:     strings.ToLower(b.ContentType),
		Free:            modext.IsFree(b),
		PublisherID:     b.PublisherID.Int,
		AuthorID:        b.AuthorID.Int,
		// Price:           b.Price,
		// Translator:      nil,
		// TranslatorID:    b.TranslatorID,
		// Narrator:        nil,
		// NarratorID:      b.NarratorID,
		// Rtl:             modext.IsRtl(b),
		// Rate:            b.Rate,
		// Path:            "",
		// PublishDate:     b.PublishDate.Time,
		// RateCount:       b.RateCount,
		// Duration:        b.Duration.Int,
		// FileSize:        b.Filesize.Int64,
		// CategoryID:      "",
		// CategoryName:    "",
		// PageCount:       "",
		// Price2:          "",
		// Type:            "",
		// SampleCrc:       b.SampleCRC.Uint,
		// Pass:            b.Password.String,
		// Crc:             b.CRC,
		// LinkTitle:       "",
		// LanguageTitle:   "",
		// PriceTitle:      "",
		// PriceTitle2:     "",
		// Description:     "",
		// LinkURL:         "",
		// Bought:          false,
		// New:             false,
		// Featured:        false,
		// Favorite:        false,
		// MyRate:          0,
	}

	if b.R.Publisher != nil {
		rs.PublisherTitle = b.R.Publisher.Title
	}

	if b.R.Author != nil {
		rs.Author = b.R.Author.Name
		rs.AuthorLogo = b.R.Author.Logo.String
	}

	return rs
}
