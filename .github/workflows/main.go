package main
import "fmt"
func main() {
	m := map[string]float64{
		"pi": 3.1416,
		"e":  2.71828,
	}
	fmt.Println(m) // "map[e:2.71828 pi:3.1416]"
	
	for key, value := range m { // Order not specified 
		fmt.Println(key, value)
	}
}