package main
import(
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
  "os"
  "io"
  "time"
)

type Cookie struct{
  Name string
  Value string
  Expires time.Time
}

func main() {

  productDB:=map[string]int{
    "iphone":20,
    "table":50,
    "ipad":150,
  }

  //use gorilla mux router
  router:=mux.NewRouter()
  router.HandleFunc("/",func (w http.ResponseWriter, r *http.Request)  {
    //http.ServeFile(w,r,"upload.html")
    expiration:=time.Now().Add(time.Hour*24*365)
    cookie:=http.Cookie{Name:"user",Value:"phanthakan",Expires:expiration}
    http.SetCookie(w,&cookie)
    fmt.Fprintf(w,"Create Cookie")
  })
  router.HandleFunc("/signup",func (w http.ResponseWriter, r *http.Request)  {
    http.ServeFile(w,r,"signup.html")
  })
  router.HandleFunc("/upload",func (w http.ResponseWriter, r *http.Request)  {
    //file = file that uploaded
    //handle = description of file that uploaded
    //err = error when upload fail
    file,handle,err:=r.FormFile("file")
    defer file.Close()
    if err!=nil{
      fmt.Println(err)
      return
    }
    //fmt.Fprintf(w,"%v",handle.Header)
    //f pick target's address, err pick error status
    f,err:=os.OpenFile("./test/" + handle.Filename,os.O_CREATE,0666)
    if err!=nil{
      fmt.Println(err)
      return
    }
    defer f.Close()
    io.Copy(f,file)
    fmt.Fprintf(w,"Upload Complete")
  })

  router.HandleFunc("/login",login)

  router.HandleFunc("/product/{name}",func (w http.ResponseWriter, r *http.Request)  {
    vars:=mux.Vars(r)
    name:=vars["name"]
    price:=productDB[name]
    fmt.Fprintf(w,"%s : %d baht",name,price)
  }).Methods("GET")

  //use http router
  /*http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request)  {
      http.ServeFile(w,r,"login.html")
  })

  /*http.HandleFunc("/product/",func (w http.ResponseWriter, r *http.Request)  {
    name:=r.URL.Path[len("/product/"):]
    price:=productDB[name]
    fmt.Fprintf(w,"%s : %d baht",name,price)

  })*/
  //http.HandleFunc("/login",login)

  //http.ListenAndServe(":8080",nil)
  http.ListenAndServe(":8080",router)
}

/*func product (w http.ResponseWriter, r *http.Request)  {
      fmt.Fprintf(w,"product request")
}*/

func login(w http.ResponseWriter, r *http.Request){

  fmt.Println("Method:",r.Method)
  r.ParseForm()
  fmt.Println("Username:",r.Form["username"])
  fmt.Println("Password:",r.Form["password"])
  http.ServeFile(w,r,"login.html")
}
