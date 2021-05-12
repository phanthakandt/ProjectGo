package main
import "fmt"
func main() {
  f1("phanthakan")
  fmt.Println(f2(1,2))
  sumation(1,2,3,4,5,6,7,8,9,10)
}

func f1(str string){
  fmt.Println(str)
}

func f2(a int,b int) int{
  return a+b
}

//variadic function
func sumation(nums...int){
  total:=0
  for _,n:=range nums{
    total += n
  }

  fmt.Println(total)
}
