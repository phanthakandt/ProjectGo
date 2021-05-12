package main
import "fmt"
func main() {
  fmt.Print("Input number: ")
  var input float64
  fmt.Scanf("%f",&input)

  //condition:=input>5

  if input>5 {
    fmt.Println("condition is true")
  }else{
    fmt.Println("condition is false")
  }

  switch input {
  case 0: fmt.Println("zero")
  case 1:fmt.Println("one")
  case 2:fmt.Println("Two")
  default:fmt.Println("Unknow")
  }

  var i float64 = 0
  for i < input {
    fmt.Println("hello")
    i++
  }

  for i := 0; i < 5; i++ {
    fmt.Println("hell")
  }

}
