package paginator

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"math"
	"math/rand"
	"testing"
)

type User struct {
	ID       int    `gorm:"column:id;"`
	UserName string `gorm:"not null;size:100;unique"`
}

var (
	db *gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	db.DropTable(&User{})
	db.AutoMigrate(&User{})
	db.Create(User{ID: 1, UserName: "william"})
	db.Create(User{ID: 2, UserName: "jack"})
	db.Create(User{ID: 3, UserName: "james"})
	db.Create(User{ID: 4, UserName: "alice"})
	db.Create(User{ID: 5, UserName: "bob"})
	db.Create(User{ID: 6, UserName: "tom"})
	db.Create(User{ID: 7, UserName: "lucy"})
	db.Create(User{ID: 8, UserName: "echo"})
}
func assertEqual(t *testing.T, caseInfo interface{}, got, expected interface{}, variableName string) {
	if got != expected {
		t.Errorf("[case %v][variable name:%v] expected `%v`, but got `%v`", caseInfo, variableName, expected, got)
	}
	t.Logf("%v:%v ", variableName, got)

}

func TestPaginator_Paginate(t *testing.T) {
	loop := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, i := range loop {
		t.Logf("\n%v loop \n", i)
		page := rand.Intn(10)
		perPage := rand.Intn(10)
		if page == 0 {
			page = 1
		}
		if perPage == 0 {
			perPage = 1
		}
		var users []User
		orderBy := []string{"id desc"}
		p := Paginator{DB: db, Page: page, PerPage: perPage, OrderBy: orderBy}
		data := p.Paginate(&users)
		caseInfo := fmt.Sprintf("Page:%v PerPage:%v OrderBy:%v", page, perPage, orderBy)
		t.Log(caseInfo)
		assertEqual(t, caseInfo, data.TotalRecords, 8, "TotalRecords")
		assertEqual(t, caseInfo, data.CurrentPage, page, "CurrentPage")
		assertEqual(t, caseInfo, data.TotalPages, int(math.Ceil(float64(data.TotalRecords)/float64(perPage))), "TotalPages")
		assertEqual(t, caseInfo, data.HasNext, page < data.TotalPages, "HasNext")
		assertEqual(t, caseInfo, data.HasPrev, page > 1, "HasPrev")
		t.Log(data.Records)
	}
}
