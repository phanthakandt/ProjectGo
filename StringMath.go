package main
import "fmt"
func main() {
  //math operation
  fn := 5
  sn := 3
  fmt.Println(fn+sn)

  //string operation
  p1 := "text1"
  p2 := "text2"
  fmt.Println(p1+p2)
  //start at 1 end at before 3(2)
  fmt.Println(p1[1:3])
  //start at 1
  fmt.Println(p1[1:])

  //askii of p1[1]
  fmt.Println(p1[1])
}
