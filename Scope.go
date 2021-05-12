package main
import "fmt"

var global1 int= 50

func main() {

  fmt.Println(global1)
  anotherFunction()
}

func anotherFunction(){
  fmt.Println(global1)
}
