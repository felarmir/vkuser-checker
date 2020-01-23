package apiservice

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/felarmir/vkuser-checker/dbclient"
	"github.com/gorilla/mux"
)

var db *sql.DB

func UserStatistic(w http.ResponseWriter, r *http.Request) {
	rowID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		rowID = 0
	}

	rows, err := dbclient.UserStatisticsList(db, rowID)
	if err != nil {
		fmt.Fprintln(w, "{status: false}")
	}
	json.NewEncoder(w).Encode(rows)
}

func StartServer(dbPointer *sql.DB) error {
	db = dbPointer
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/userStatistics/pivot/{id}", UserStatistic)
	http.ListenAndServe(":8080", router)
	fmt.Println("IP: 127.0.0.1 Listen Port: 8080")
	return nil
}
