package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func printDetails(values ...any) {
	for _, elem := range values {
		switch val := elem.(type) {
		case Product:
			Printfln("Product: Name: %v, Category: %v, Price: %v",
				val.Name, val.Category, val.Price)
		case Customer:
			Printfln("Customer: Name: %v, City: %v", val.Name, val.City)
		}
	}
}

func printReflectedDetails(values ...any) {
	for _, elem := range values {
		details := []string{}
		eTyp := reflect.TypeOf(elem)
		eVal := reflect.ValueOf(elem)
		//Printfln("\nType:%v; Value:v%", eTyp, eVal)
		Printfln("Name: %v, PkgPath: %v, Kind: %v",
			eTyp.Name(), getTypePath(eTyp), eTyp.Kind())
		//  if this is a struct got into it and enumerate and discover fields types
		if eTyp.Kind() == reflect.Struct {
			// NumField(): number of fields in the structure
			for i := 0; i < eTyp.NumField(); i++ {
				fieldName := eTyp.Field(i).Name
				fieldVal := eVal.Field(i)
				details = append(details, fmt.Sprintf("%v: %v", fieldName, fieldVal))
			}
			Printfln("\n%v: %v", eTyp.Name(), strings.Join(details, ", "))
		} else if eTyp == byteSliceType {
			Printfln("Byte slice: %v", eVal.Bytes())
		} else if eTyp == intPtrType {
			Printfln("Pointer to Int: %v", eVal.Elem().Int())
		} else {
			Printfln("\n%v: %v", eTyp.Name(), eVal)
		}
	}
}

type Payment struct {
	Currency string
	Amount   float64
}

func getTypePath(t reflect.Type) (path string) {
	path = t.PkgPath()
	if path == "" {
		path = "(built-in)"
	}
	return
}

func printDetailsReflectValues(values ...any) {
	for _, elem := range values {
		elemValue := reflect.ValueOf(elem)
		switch elemValue.Kind() {
		case reflect.Bool:
			var val bool = elemValue.Bool()
			Printfln("Bool: %v", val)
		case reflect.Int:
			var val int64 = elemValue.Int()
			Printfln("Int: %v", val)
		case reflect.Float32, reflect.Float64:
			var val float64 = elemValue.Float()
			Printfln("Float: %v", val)
		case reflect.String:
			var val string = elemValue.String()
			Printfln("String: %v", val)
		case reflect.Ptr:
			var val reflect.Value = elemValue.Elem()
			if val.Kind() == reflect.Int {
				Printfln("Pointer to Int: %v", val.Int())
			}
		default:
			Printfln("Other: %v", elemValue.String())
		}
	}
}

var intPtrType = reflect.TypeOf((*int)(nil))
var byteSliceType = reflect.TypeOf([]byte(nil))

func selectValue(data any, index int) (result any) {
	dataVal := reflect.ValueOf(data)
	if dataVal.Kind() == reflect.Slice {
		result = dataVal.Index(index).Interface()
	}
	return
}

func incrementOrUpper(values ...any) {
	for _, elem := range values {
		elemValue := reflect.ValueOf(elem)
		if elemValue.Kind() == reflect.Ptr {
			elemValue = elemValue.Elem()
		}
		//returns true if `values` are pointers
		if elemValue.CanSet() {
			switch elemValue.Kind() {
			case reflect.Int:
				elemValue.SetInt(elemValue.Int() + 1)
			case reflect.String:
				elemValue.SetString(strings.ToUpper(elemValue.String()))
			}
			// 741 Chapter 27 ■ Using Reflection
			Printfln("Modified Value: %v", elemValue)
		} else {
			Printfln("Cannot set %v: %v", elemValue.Kind(), elemValue)
		}
	}
}

func setAll(src any, targets ...any) {
	srcVal := reflect.ValueOf(src)
	for _, target := range targets {
		targetVal := reflect.ValueOf(target)
		if (targetVal.Kind() == reflect.Ptr) &&
			(targetVal.Elem().Type() == srcVal.Type()) &&
			targetVal.Elem().CanSet() {
			targetVal.Elem().Set(srcVal)
		}
	}
}

func contains(slice any, target any) (found bool) {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() == reflect.Slice {
		for i := 0; i < sliceVal.Len(); i++ {
			if reflect.DeepEqual(sliceVal.Index(i).Interface(), target) {
				found = true
				break
			}
		}
	}
	return
}

func convert(src, target interface{}) (result interface{}, assigned bool) {
	srcVal := reflect.ValueOf(src)
	targetVal := reflect.ValueOf(target)
	if srcVal.Type().ConvertibleTo(targetVal.Type()) {
		if (IsInt(targetVal) && IsInt(srcVal)) &&
			targetVal.OverflowInt(srcVal.Int()) {
			// 750 Chapter 27 ■ Using Reflection
			Printfln("Int overflow")
			return src, false
		} else if IsFloat(targetVal) && IsFloat(srcVal) &&
			targetVal.OverflowFloat(srcVal.Float()) {
			Printfln("Float overflow")
			return src, false
		}
		result = srcVal.Convert(targetVal.Type()).Interface()
		assigned = true
	} else {
		result = src
	}
	return
}
func IsInt(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	}
	return false
}
func IsFloat(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

func swap(first interface{}, second interface{}) {
	firstValue, secondValue := reflect.ValueOf(first), reflect.ValueOf(second)
	if firstValue.Type() == secondValue.Type() &&
		firstValue.Kind() == reflect.Ptr &&
		firstValue.Elem().CanSet() && secondValue.Elem().CanSet() {
		temp := reflect.New(firstValue.Elem().Type())
		temp.Elem().Set(firstValue.Elem())
		firstValue.Elem().Set(secondValue.Elem())
		secondValue.Elem().Set(temp.Elem())
	}
}

func createPointerType(t reflect.Type) reflect.Type {
    return reflect.PtrTo(t)
}
func followPointerType(t reflect.Type) reflect.Type {
    if t.Kind() == reflect.Ptr {
        return t.Elem()
    }
    return t
}

func main() {
	product := Product{
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	printDetails(product)

	Printfln("\n Understanding the Need for Reflection")
	customer := Customer{Name: "Alice", City: "New York"}
	printDetails(product, customer)

	a, err := strconv.ParseBool("true")
	b, err := strconv.ParseBool("0")
	c, err := strconv.ParseBool("1")
	Printfln("a: %v; b: %v; c: %v; err: %v", a, b, c, err)
	payment := Payment{Currency: "USD", Amount: 100.50}

	Printfln("\n Using Reflection")
	printReflectedDetails(product, customer, a, b, c, 12, "aaaa", 123.123, payment)

	Printfln("%s%s%s",
		"\nUsing the Basic Type Features",
		"\n Using the Basic Value Features", "\nIdentifying Types")

	number := 100
	printDetailsReflectValues(true, 10, 23.30, "Alice", &number, product)

	printReflectedDetails(product, customer, a, b, c, 1, &number)

	Printfln("\n Identifying Byte Slices")
	slice := []byte("Alice")

	printReflectedDetails(true, 10, 23.30, "Alice", &number, product, slice)

	Printfln("\n Obtaining Underlying Values")
	// The Value struct defines the methods described in Table 27-7 for
	// obtaining an underlying value.
	names := []string{"Alice", "Bob", "Charlie"}
	val := selectValue(names, 1).(string)
	Printfln("Selected: %v; Type: %T", val, val)

	Printfln("\n Setting a Value Using Reflection")
	// The Value struct defines methods that allow values to be set using
	// reflection, as described in Table 27-8.
	name := "Alice"
	price := 279
	city := "London"
	incrementOrUpper(&name, &price, &city)
	for _, val := range []any{name, price, city} {
		Printfln("Value: %v", val)
	}

	Printfln("\n Setting One Value Using Another")
	setAll("New String", &name, &price, &city)
	setAll(10, &name, &price, &city)

	for _, val := range []any{name, price, city} {
		Printfln("Value: %v", val)
	}

	Printfln("\n Comparing Values%s", "\n Using the Comparison Convenience Function")
	city = "London"
	citiesSlice := []string{"Paris", "Rome", "London"}
	Printfln("Found #1: %v", contains(citiesSlice, city))
	sliceOfSlices := [][]string{
		citiesSlice, {"First", "Second", "Third"}}
	Printfln("Found #2: %v", contains(sliceOfSlices, citiesSlice))

	Printfln("\n Converting Values")
	name = "Alice"
	price = 279
	city = "London"
	newVal, ok := convert(price, 100.00)
	Printfln("Converted %v: %v, %T", ok, newVal, newVal)
	newVal, ok = convert(name, 100.00)
	Printfln("Converted %v: %v, %T", ok, newVal, newVal)

	Printfln("\n Converting Numeric Types")
	newVal, ok = convert(5000, int8(100))
	Printfln("Converted %v: %v, %T", ok, newVal, newVal)

	Printfln("\n Creating new values")

	swap(&name, &city)
	for _, val := range []interface{}{name, price, city} {
		Printfln("Value: %v", val)
	}
}
