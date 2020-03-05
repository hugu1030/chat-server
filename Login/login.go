package login

import (
	"fmt"
	"net/http"
	"database/sql"
	-"github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=postgres dbname=postgres password=hugu1030 sslmode=disable")
	if err != nil {
			panic(err)
	}
}


func ReturnCanLogin(w http.ResponseWriter,r *http.Request){
      

}

func ReturnMadeAccount(w http.ResponseWriter, r *http.Request){

}
func ReturnResetAccount(w http.ResponseWriter, r *http.Request){

}