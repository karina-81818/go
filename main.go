package main
import(
"fmt"
	"math"
)
 	


func main(){
	var userHight float64
	var userKg float64 
	fmt.Println("**Калькулятор индекса массы тела**")
	fmt.Println("Введите свой рост в см: ")
	fmt.Scan(&userHight)
	fmt.Println("Введите свой вес в кг: ")
	fmt.Scan(&userKg)
	var IMT = userKg/math.Pow(userHight/100, 2)

	fmt.Printf("Ваш индекс массы тела: %.0f", IMT)
}