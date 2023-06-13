package domain

import (
	"strings"
	"time"
)

type Book struct {
	Id       int64
	Title    string
	Author   string
	Year     int32
	Genres   []string
	CreateAt time.Time
}

type Filter struct {
	Title  string
	Author string
	Genres []string
	Size   int64
	Sort   string
}

func (f Filter) SortColumn() string {
	if f.Sort == "" {
		f.Sort = "id"
	}
	for _, safeValue := range []string{"title", "-title", "author", "-author", "year", "-year", "id", "-id"} {
		if f.Sort == safeValue {
			return strings.TrimPrefix(f.Sort, "-")
		}
	}
	panic("unsafe sort parameter: " + f.Sort)
}

func (f Filter) SortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}
	return "ASC"
}
