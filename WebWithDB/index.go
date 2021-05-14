package main

import(
  "fmt"
  //"net/http"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func main() {
  //connect DB
  db, err := sql.Open("mysql","root:@/test")
  if err!=nil{
    fmt.Println(err)
  }
  defer db.Close()

  //insert
  stmt,err:=db.Prepare("insert into customer (customerID,cfirstname,clastname,address) values(?,?,?,?)")
  stmt.Exec("c004","Nanthicha","Tunvichai","Unknown")
  if(err!=nil){
    fmt.Println(err)
  }

  //query
  rows,err:=db.Query("select cfirstname,clastname,address from customer")
  if err!=nil{
    fmt.Println(err)
  }

  //loop
  for rows.Next(){
      var firstname string
      var lastname string
      var address string
      err=rows.Scan(&firstname,&lastname,&address)
      fmt.Printf("firstname : %s lastname: %s address: %s\n",firstname,lastname,address)
  }
  //http.HandleFunc("/",index)
  //http.ListenAndServe(":8080",nil)
}
