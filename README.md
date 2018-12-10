# gorm-paginator

Custom modifications based on the project [paginator](https://github.com/Prabandham/paginator)

### Install 
`go get -u github.com/william-tu/gorm-paginator`

### Usage

```
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
    paginator := Paginator{DB: &db, OrderBy: order_by, Page: "1", PerPage: "10"}
    data := p.Paginate(&users)
    totalRecords := data.TotalRecords // total records return int
    currentPage := data.CurrentPage //cureent page return int
    totalPages := data.TotalPages //total pages return int
    hasPrev := data.HasPrev // true if has prev
    hasNext := data.HasNext //true if has next
    var records []User
    //type conversion
    for _, obj := range data.Records {
    	records = append(records, obj.(User))
    }
}

```

        


