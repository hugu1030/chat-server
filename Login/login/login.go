package login

import (
	"net/http"
	"database/sql"
	_"github.com/lib/pq"
	"encoding/json"
	"log"
)

type User struct {
	Id   int
	Name string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=postgres dbname=postgres password=hugu1030 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func ReturnCanLogin(w http.ResponseWriter,r *http.Request){
	  e := r.ParseForm()
	  log.Println(e)
	  
	  u := User{}
	  err := Db.QueryRow("SELECT id,name FROM chatlogin WHERE mail = ? && password = ? ",r.FormValue("mail"),r.FormValue("password")).Scan(&u.Id,&u.Name)

      if err != nil{
		  panic(err)
	  }else{
		 log.Println(u)
		 log.Println(json.NewEncoder(w).Encode(u))
		 json.NewEncoder(w).Encode(u)
	  }
}

func ReturnMadeAccount(w http.ResponseWriter, r *http.Request){    
	e := r.ParseForm()
	log.Println(e)
	// category := r.Form.Get("category")

	stmt,err := Db.Prepare("INSERT INTO postgres (name,mail,password) VALUES(? ,? ,?)")
	if err != nil {
		return
	}
	defer stmt.Close()
    u := User{}
	ret,err1 := stmt.Exec(r.FormValue("name"),r.FormValue("mail"),r.FormValue("password"))
	if(err1 != nil){
		return
	}
	log.Println(ret)

	err2 := Db.QueryRow("SELECT id,name FROM postgres WHERE mail = ? && name = ? && password = ? ", r.FormValue("mail"),r.FormValue("name"),r.FormValue("password")).Scan(&u.Id,&u.Name)
	if err2 != nil{
		return 
	}else{
		json.NewEncoder(w).Encode(u)
	}
}
func ReturnResetAccount(w http.ResponseWriter, r *http.Request){
	e := r.ParseForm()
	log.Println(e)
	var password string
	err := Db.QueryRow("SELECT password FROM postgres WHERE mail = ? && name = ? ",r.FormValue("mail"),r.FormValue("name")).Scan(&password)
	if err != nil{
		return 
	}else{
		json.NewEncoder(w).Encode(password)
	}
}