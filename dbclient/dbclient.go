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

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func UserStatisticsList(db *sql.DB, lastID int) ([]vkclient.UserStatisticRow, error) {
	sqlQuery := `
	SELECT * FROM user_status
	WHERE id > $1
	`
	rows, err := db.Query(sqlQuery, lastID)
	if err != nil {
		return nil, err
	}
	var tmpUsers []vkclient.UserStatisticRow
	for rows.Next() {
		var u vkclient.UserStatisticRow
		if err := rows.Scan(&u.ID, &u.Path, &u.FirstName, &u.LastName, &u.Isonline, &u.LastRequest); err != nil {
			fmt.Println(err)
		}
		tmpUsers = append(tmpUsers, u)
	}
	return tmpUsers, nil
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
	return nil
}
