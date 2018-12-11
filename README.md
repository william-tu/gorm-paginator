# gorm-paginator

Custom modifications based on the project [paginator](https://github.com/Prabandham/paginator)

### Install 
`go get -u github.com/william-tu/gorm-paginator`

### Usage

```go
import (
	pg "github.com/william-tu/gorm-paginator"
)

type User struct {
	ID       int    `gorm:"column:id;"`
	UserName string `gorm:"not null;size:100;unique"`
	.
	.
	.
}

func main(){
    var users  []User
    db, err := gorm.Open("postgres", ....)
    order_by := []string{"id asc"}
    p := pg.Paginator{DB: db, OrderBy: order_by, Page: "1", PerPage: "10"}
    data := p.Paginate(&users)
    totalRecords := data.TotalRecords 
    currentPage := data.CurrentPage
    totalPages := data.TotalPages 
    hasPrev := data.HasPrev 
    hasNext := data.HasNext
    var records []User
    //type conversion
    for _, obj := range data.Records {
    	records = append(records, obj.(User))
    }
}

```

        


