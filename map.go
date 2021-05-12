package main
import "fmt"
func main() {
  //x:[5]int{1,2,3,4,5}
  y:=make(map[string] string)
  y["TH"]="Thailand"
  y["JP"]="Japan"
  fmt.Println(y)

  x:=map[string] string{
    "H":"Hidrogen",
    "Li":"Lithium",
  }
  fmt.Println(x)
}
