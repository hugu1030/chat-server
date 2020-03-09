package login

import (
	"net/http"
	"database/sql"
	_"github.com/lib/pq"
	"encoding/json"
	"log"
	"io/ioutil"
)

type User struct {
	Id   int      `json:"id"`
	Name string   `json:"name"`
}

type Pass struct {
	Mail string
	Password string    
}

type AllInfo struct {
	Name   string   
	Mail   string
	Password  string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=postgres dbname=postgres password=hugu1030 sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = Db.Ping()
	if err != nil {
		panic(err)
	}
}

func ReturnCanLogin(w http.ResponseWriter,r *http.Request){

	  body,_ := ioutil.ReadAll(r.Body)
	  u := User{}
	  p := Pass{}
	
	  json.Unmarshal(body,&p)     
	  
	  err := Db.QueryRow("SELECT id,name FROM chatlogin WHERE mail = $1 AND password = $2 ", p.Mail, p.Password).Scan(&u.Id,&u.Name)

      if err != nil{
		  w.WriteHeader(http.StatusInternalServerError)
		//   w.Write([]byte("No rows in result"))
  		 return
	  }
		w.Header().Set("Content-Type","application/json")
		log.Println("u:",u)
		json.NewEncoder(w).Encode(u)
}


func ReturnMadeAccount(w http.ResponseWriter, r *http.Request){
	
	body,_ := ioutil.ReadAll(r.Body)
    allinfo := AllInfo{}
	json.Unmarshal(body,&allinfo)
    u := User{}
	
	isPassword := Db.QueryRow("SELECT id FROM chatlogin WHERE mail = $1", allinfo.Mail).Scan(&u.Id)
	log.Println(u.Id)
	log.Println(allinfo.Mail)
	log.Println(isPassword)
	if isPassword == nil {
		log.Println("このメールアドレスは既に登録されています")
		w.WriteHeader(501)
       return 
	}else{
	stmt,err := Db.Prepare("INSERT INTO chatlogin (name,mail,password) VALUES($1,$2,$3)")
	if err != nil {
		log.Println("データベース作成で失敗しました")
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}

	// defer stmt.Close()

	_,err1 := stmt.Exec(allinfo.Name,allinfo.Mail,allinfo.Password) 
	if(err1 != nil){
	   log.Println("データベース挿入に失敗しました")
	   w.WriteHeader(http.StatusInternalServerError)
	   return
	}

	err2 := Db.QueryRow("SELECT id,name FROM chatlogin WHERE mail = $1 AND name = $2 AND password = $3 ", allinfo.Mail,allinfo.Name,allinfo.Password).Scan(&u.Id,&u.Name)
	if err2 != nil {
		log.Println("データベース取り出しに失敗しました")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
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