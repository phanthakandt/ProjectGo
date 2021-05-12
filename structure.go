package main
import "fmt"

type Book struct{
  title string
  author string
  subtitle string
  price float64
}

func main() {
  var Book1 Book
  Book1.title = "Go Programming"
  Book1.author = "Phanthakan"
  Book1.subtitle = "PanthakanEiEi"
  Book1.price = 1990.50

  Golang:=Book{title:"golang"}
  fmt.Println(Golang)
}
