package main
import(
  "net/http"
  "github.com/gorilla/mux"
  "html/template"
)

type Product struct{
  Name string
  Price int
}

func main() {
  myProduct:=Product{"นมสด",5}

  var templates = template.Must(template.ParseFiles("index.html"))
  router:=mux.NewRouter()
  router.HandleFunc("/",func (w http.ResponseWriter, r *http.Request)  {
    templates.ExecuteTemplate(w,"index.html",myProduct)
  })
  http.ListenAndServe(":8080",router)
}
