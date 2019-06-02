package flex

import (
	"strconv"
	"strings"
)

type queryIdis []interface{}

func (q queryIdis) getInts() []int {
	r := make([]int, 0)
	for _, id := range q {
		switch id.(type) {
		case string:
			intval, err := strconv.Atoi(id.(string))
			if err == nil {
				r = append(r, intval)
			}

		case int:
			r = append(r, id.(int))
		}

	}

	return r
}

func (q queryIdis) getInterfaceList() []interface{} {
	qidis := make([]interface{}, 0)
	for _, id := range q {
		qidis = append(qidis, id)
	}

	return qidis
}

func (q queryIdis) getInterfaceStringLowerList() []interface{} {
	contentTypeList := q.getInterfaceList()

	for index, ct := range contentTypeList {
		contentTypeList[index] = strings.ToLower(ct.(string))
	}

	return contentTypeList
}
