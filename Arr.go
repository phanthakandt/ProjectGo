package main
import "fmt"
func main() {
  var x[5] int
  fmt.Println(x)

  y:=[5]float64{1,2,5,4}
  //start at 0 end at 2
  fmt.Println(y[0:3])

  var total float64 = 0
  for _,value:=range y{
    total += value
  }
  fmt.Println(total/float64(len(y)))
}
