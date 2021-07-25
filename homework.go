import (
   "database/sql"
   "github.com/pkg/errors"
)

var DB *sql.DB

func OpenDB() error {
	dsn := "root:root@tcp(127.0.0.2:3306)/test?charset=utf8"
	DB, _ = sql.Open("mysql", dsn)

	if err := DB.Ping(); err != nil {
		// 数据库链接错误是一个致命错误，应Wrap，并抛出
		return errors.Wrap(err, "Connection DB error")
	}

	return nil
}

type Customer struct {
   CustomerId string
   Name       string
}

func getUser(ID int) (*User, error) {
	user := &User{}
	err := global.DB.QueryRow("SELECT id,username,birthday FROM blog_user WHERE id = ?", ID).Scan(&user.ID, &user.Username, &user.Birthday)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func main() {
	err := global.OpenDB()
	if err != nil {
		log.Fatalf("FATAL: %+v\n", err)
	}

	user, err := getUser(1024)
	if err != nil {
		log.Println(err)
	} else {
		if user != nil {
			fmt.Printf("User: %+v", user)
		} else {
			fmt.Println("No Rows!")
		}
	}
}
