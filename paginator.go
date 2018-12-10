package paginator

import (
	"github.com/jinzhu/gorm"
	"math"
	"reflect"
)

type Data struct {
	TotalRecords int           `json:"total_records"`
	Records      []interface{} `json:"records"`
	CurrentPage  int           `json:"current_page"`
	TotalPages   int           `json:"total_pages"`
	HasPrev      bool          `json:"has_prev"`
	HasNext      bool          `json:"has_next"`
}

type Paginator struct {
	DB      *gorm.DB
	OrderBy []string
	Page    int
	PerPage int
}

func (p *Paginator) Paginate(modelSource interface{}) *Data {
	db := p.DB

	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			db = db.Order(o)
		}
	}

	done := make(chan bool, 1)
	var output Data
	var count int
	var offset int

	go p.countRecords(modelSource, done, &count)

	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Page == 1 {
		offset = 0
	} else {

		offset = (p.Page - 1) * p.PerPage
	}
	db.Limit(p.PerPage).Offset(offset).Find(modelSource)
	<-done

	output.TotalRecords = count
	output.Records = toSlice(modelSource)
	output.CurrentPage = p.Page
	output.TotalPages = int(math.Ceil(float64(count) / float64(p.PerPage)))
	output.HasPrev = output.CurrentPage > 1
	output.HasNext = output.CurrentPage < output.TotalPages

	return &output
}

func (p *Paginator) countRecords(countDataSource interface{}, done chan bool, count *int) {
	p.DB.Model(countDataSource).Count(count)
	done <- true
}

func toSlice(arr interface{}) []interface{} {
	v1 := reflect.ValueOf(arr)
	v := reflect.Indirect(v1)
	if v.Kind() != reflect.Slice {
		panic(" func toslice:arr is not slice")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}
