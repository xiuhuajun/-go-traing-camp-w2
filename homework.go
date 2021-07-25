import (
	"database/sql"
	"github.com/pkg/errors"
	"fmt"
)

var DB *sql.DB

func OpenDB() error {
	dsn := "root:root@tcp(127.0.0.2:3306)/test?charset=utf8"
	DB, _ = sql.Open("mysql", dsn)

	if err := DB.Ping(); err != nil {
		// 数据库链接错误是一个严重的错误，应该Wrap这个error，抛给上层
		return errors.Wrap(err, "Connection DB error")
	}

	return nil
}

type Customer struct {
	ID	string
	Name	string
}

func getCustomer(ID string) (*Customer, error) {
	user := &Customer{}
	err := DB.QueryRow("SELECT id,name FROM user WHERE id = ?", ID).Scan(&user.ID, &user.Name)
	if err != nil {
		//提示未查询到数据，无需Wrap
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		//若为其他错误，抛出
		return nil, err
	}

	return user, nil
}

func main() {
	err := OpenDB()
	if err != nil {
		fmt.Printf("query customer err : %+v", err)
  		return nil
	}

	user, err := getCustomer("123")
	if err != nil {
		fmt.Printf("query customer err : %+v", err)
	} else {
		if user != nil {
			fmt.Printf("customer: %+v", user)
		} else {
			fmt.Println("No Rows!")
		}
	}
}
