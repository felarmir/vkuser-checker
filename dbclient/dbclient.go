package dbclient

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/felarmir/vkuser-checker/vkclient"
)

type DBConnection struct {
	Host     string
	Port     int
	User     string
	Password string
	DBname   string
}

func (self *DBConnection) PGConnect() (*sql.DB, error) {
	credential := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		self.Host, self.Port, self.User, self.Password, self.DBname)

	db, err := sql.Open("postgres", credential)
	if err != nil {
		return nil, err
	}
	//defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InsertRow(db *sql.DB, user vkclient.User) error {
	sqlQuery := `
	INSERT INTO user_status (path, first_name, last_name, isOnline)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	path := "https://vk.com/id" + strconv.Itoa(user.ID)
	id := 0
	err := db.QueryRow(sqlQuery, path, user.FirstName, user.LastName, user.Online).Scan(&id)
	if err != nil {
		return err
	}
	fmt.Println("New record ID is:", id)
	return nil
}
