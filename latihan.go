package main
import "fmt"

func cariDigit(n, m int)bool{
	var b bool
	var num int
	b = false
	for m > 0{
		num = m%10
		if n == num {
			b = true
		}
		m = m/10
	}
	return b
}
func main(){
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Print(cariDigit(n, m))
}