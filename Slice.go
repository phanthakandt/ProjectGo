package main
import "fmt"
func main() {
  x:=make([]int, 5)
  y:=[]int{0,2,3}
  z:=append(y,1,2,3,4)
  copy(x,z)

  fmt.Println(x)
  fmt.Println(y)
  fmt.Println(z)
}
