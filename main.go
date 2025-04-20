package main

import (
	"fmt"
	"math"
)


func main(){
	
	fmt.Println("**Калькулятор индекса массы тела**")
	userKg, userHight := getUserInput()
	IMT := calculateIMT(userHight, userKg)
	outputResult(IMT)
	
	
}

func outputResult(imt float64){
	res := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Println(res)
}

func calculateIMT(userHight, userKg float64) float64{
	const powerIMT = 2
	IMT := userKg/math.Pow(userHight/100, powerIMT)

	return IMT
}

func getUserInput()(float64, float64){
	var userHight float64
	var userKg float64 
	fmt.Println("Введите свой рост в см: ")
	fmt.Scan(&userHight)
	fmt.Println("Введите свой вес в кг: ")
	fmt.Scan(&userKg)

	return userKg, userHight

}