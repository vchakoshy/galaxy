package flex

import (
	"strings"

	"github.com/volatiletech/sqlboiler/queries/qm"
)

type Setup struct {
	Book         ProviderSetup       `json:"book"`
	Author       ProviderSetup       `json:"author"`
	Category     ProviderSetup       `json:"category"`
	Channel      ProviderSetup       `json:"channel"`
	Publisher    ProviderSetup       `json:"publisher"`
	ProposedList ProviderSetup       `json:"proposedList"`
	Tag          ProviderSetup       `json:"tag"`
	Free         ProviderSimpleSetup `json:"free"`
	Subscription ProviderSimpleSetup `json:"subscription"`
	Size         ProviderSimpleSetup `json:"size"`
	Filter       ProviderSimpleSetup `json:"filter"`
	Sort         ProviderSimpleSetup `json:"sort"`
	Query        struct {
		DataProvider interface{}   `json:"dataProvider"`
		Type         string        `json:"type"`
		Fields       []interface{} `json:"fields"`
	} `json:"query"`
	ContentType struct {
		DataProvider interface{} `json:"dataProvider"`
		Type         string      `json:"type"`
		Value        QueryIdis   `json:"value"`
	} `json:"contentType"`
	Format struct {
		DataProvider interface{} `json:"dataProvider"`
		Type         string      `json:"type"`
		Value        QueryIdis   `json:"value"`
	} `json:"format"`
}

func (s Setup) getQueries() []qm.QueryMod {
	q := make([]qm.QueryMod, 0)

	q = append(q, s.Book.getQuery()...)
	q = append(q, s.Author.getQuery()...)
	q = append(q, s.Category.getQuery()...)
	q = append(q, s.Channel.getQuery()...)
	q = append(q, s.Publisher.getQuery()...)
	q = append(q, s.ProposedList.getQuery()...)
	q = append(q, s.Tag.getQuery()...)

	return q
}

func (s Setup) getInputActions() []ComponentAction {
	q := make([]ComponentAction, 0)

	f := []ComponentAction{
		s.Book.getInputAction("book"),
		s.Free.getInputAction("free"),
		s.Size.getInputAction("size"),
		s.Subscription.getInputAction("subscription"),
		s.Category.getInputAction("category"),
	}

	for _, ca := range f {
		if ca.Key != "" {
			q = append(q, ca)
		}
	}

	return q
}

func (s Setup) getSort() []qm.QueryMod {
	q := make([]qm.QueryMod, 0)

	switch s.Sort.Value {
	case "RECENT":
		q = append(q, qm.OrderBy("id DESC"))
	case "BESTSELLER":
		q = append(q, qm.InnerJoin("book_stats ON book_stats.book_id=book.id"))
		q = append(q, qm.OrderBy("book_stats.all_sales_count DESC"))
	case "WEEK_BESTSELLER":
		q = append(q, qm.InnerJoin("book_stats ON book_stats.book_id=book.id"))
		q = append(q, qm.OrderBy("book_stats.week_sales_count DESC"))
	case "RECENT_PUBLISH":
		q = append(q, qm.OrderBy("book.publish_date DESC"))
	case "POPULAR":
		q = append(q, qm.InnerJoin("book_stats ON book_stats.book_id=book.id"))
		q = append(q, qm.OrderBy("book_stats.month_download_count DESC"))
	case "LOWEST_PRICE":
		q = append(q, qm.OrderBy("book.price ASC"))
	case "MOST_COMMENTED":
		q = append(q, qm.InnerJoin("comment ON comment.book_id = book.id"))
		q = append(q, qm.Where("comment.publish = 1"))
		q = append(q, qm.OrderBy("COUNT(comment.id) DESC"))
	}

	return q
}

type QueryIdis []interface{}

func (q QueryIdis) getInterfaceList() []interface{} {
	qidis := make([]interface{}, 0)
	for _, id := range q {
		qidis = append(qidis, id)
	}

	return qidis
}

func (q QueryIdis) getInterfaceStringLowerList() []interface{} {
	contentTypeList := q.getInterfaceList()

	for index, ct := range contentTypeList {
		contentTypeList[index] = strings.ToLower(ct.(string))
	}

	return contentTypeList
}
