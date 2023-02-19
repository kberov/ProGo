package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("Hello, Operations", 22%2)

	price, tax := 275.00, 27.40
	sum := price + tax
	difference := price - tax
	product := price * tax
	quotient := price / tax
	fmt.Println(sum)
	fmt.Println(difference)
	fmt.Println(product)
	fmt.Println(quotient)

	fmt.Println("Integer overflow")
	var intVal = math.MaxInt64
	var floatVal = math.MaxFloat64
	fmt.Println(intVal * 2)
	fmt.Println(floatVal * 2)
	fmt.Println(math.IsInf((floatVal * 2), 0))

	fmt.Println("Using the Remainder Operator")
	posResult := 3 % 2
	negResult := -3 % 2
	absResult := math.Abs(float64(negResult))
	fmt.Println(posResult)
	fmt.Println(negResult)
	fmt.Println(absResult)

	fmt.Println("Using the Increment and Decrement Operators")
	value := 10.2
	value++
	fmt.Println(value)
	value += 2
	fmt.Println(value)
	value -= 2
	fmt.Println(value)
	value--
	fmt.Println(value)

	fmt.Println("Concatenating Strings")
	greeting := "Hello"
	language := "Go"
	combinedString := greeting + ", " + language
	fmt.Println(combinedString)

	fmt.Println("Understanding the Comparison Operators")
	first := 100
	const second = 200.00
	equal := first == second
	notEqual := first != second
	lessThan := first < second
	lessThanOrEqual := first <= second
	greaterThan := first > second
	greaterThanOrEqual := first >= second
	fmt.Println(equal)
	fmt.Println(notEqual)
	fmt.Println(lessThan)
	fmt.Println(lessThanOrEqual)
	fmt.Println(greaterThan)
	fmt.Println(greaterThanOrEqual)
	fmt.Println('A' > 'B')      //false
	fmt.Println("ala" > "bala") //false
	max := 0
	if first > second {
		max = first
	} else {
		max = second
	}
	fmt.Println(max)
	fmt.Println("Comparing Pointers")
	_second := &first
	third := &first
	alpha := 100
	beta := &alpha
	fmt.Println(_second == third)
	fmt.Println(_second == beta)
	// comparing the values to which the pointers point
	fmt.Println(*_second == *third)
	fmt.Println(*_second == *beta)

	fmt.Println("Understanding the Logical Operators")
	maxMph := 50
	passengerCapacity := 4
	airbags := true
	familyCar := passengerCapacity > 2 && airbags
	sportsCar := maxMph > 100 || passengerCapacity == 2
	canCategorize := !familyCar && !sportsCar
	fmt.Println(familyCar)
	fmt.Println(sportsCar)
	fmt.Println(canCategorize)

	fmt.Println("Converting, Parsing, and Formatting Values")
	kayak := 275
	soccerBall := 19.50
	//total := float64(kayak) + soccerBall
	total := kayak + int(math.Round(soccerBall))
	fmt.Println(total)

	fmt.Println("Parsing from Strings")
	val1 := "true"
	val2 := "false"
	val3 := "not true"
	bool1, b1err := strconv.ParseBool(val1)
	bool2, b2err := strconv.ParseBool(val2)
	bool3, b3err := strconv.ParseBool(val3)
	fmt.Println("Bool 1", bool1, b1err)
	fmt.Println("Bool 2", bool2, b2err)
	fmt.Println("Bool 3", bool3, b3err)
	val1_ := "0"
	bool1_, b1err_ := strconv.ParseBool(val1_)
	if b1err_ == nil {
		fmt.Println("Parsed value:", bool1_)
	} else {
		fmt.Println("Cannot parse", val1_)
	}

	_val1 := "0b1100100"
	int1, int1err := strconv.ParseInt(_val1, 0, 8)

	if int1err == nil {
		smallInt := int8(int1)
		fmt.Println("Parsed value:", smallInt)
	} else {
		fmt.Println("Cannot parse", _val1, int1err)
	}
}
