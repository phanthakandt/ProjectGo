package main
import "fmt"
func main() {
  x:=10
  var p *int
  p=&x // p -> address of x

  fmt.Println("Address of x is ",p)

  *p = 20 //change value of holding address

  fmt.Println(x)
}
