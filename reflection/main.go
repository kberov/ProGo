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

var stringPtrType = reflect.TypeOf((*string)(nil))

func transformString(val any) {
	elemValue := reflect.ValueOf(val)
	if elemValue.Type() == stringPtrType {
		upperStr := strings.ToUpper(elemValue.Elem().String())
		if elemValue.Elem().CanSet() {
			elemValue.Elem().SetString(upperStr)
		}
	}
}

func checkElemType(val any, arrOrSlice any) bool {
	elemType := reflect.TypeOf(val)
	arrOrSliceType := reflect.TypeOf(arrOrSlice)
	return (arrOrSliceType.Kind() == reflect.Array ||
		arrOrSliceType.Kind() == reflect.Slice) &&
		arrOrSliceType.Elem() == elemType
}

/*
The setValue function changes the value of an element in a slice or array, but
each kind of type has to be handled differently. Slices are the easiest to work
with and can be passed as values, like this:
...
setValue(slice, 1, newCity)
...
As I explained in Chapter 7, slices are references and are not copied when they
are used as function arguments. In Listing 28-6, the setValue method uses the
Kind method to detect the slice, uses the Index (761 Chapter 28 ■ Using
Reflection, Part 2) method to get the Value for the element at the specified
location, and uses the Set method to change the value of the element. Arrays
must be passed as pointers, like this:
...
setValue(&array, 1, newCity)
...
If you don’t use a pointer, then you won’t be able to set new values, and the
CanSet method will return false. The Kind method is used to detect the pointer,
and the Elem method is used to confirm it points to an array: ...
*/
func setValue(arrayOrSlice any, index int, replacement any) {
	arrayOrSliceVal := reflect.ValueOf(arrayOrSlice)
	replacementVal := reflect.ValueOf(replacement)
	if arrayOrSliceVal.Kind() == reflect.Slice {
		elemVal := arrayOrSliceVal.Index(index)
		if elemVal.CanSet() {
			elemVal.Set(replacementVal)
		}
	} else if arrayOrSliceVal.Kind() == reflect.Ptr &&
		arrayOrSliceVal.Elem().Kind() == reflect.Array &&
		arrayOrSliceVal.Elem().CanSet() {
		arrayOrSliceVal.Elem().Index(index).Set(replacementVal)
	}
}

func enumerateStrings(arrayOrSlice any) {
	arrayOrSliceVal := reflect.ValueOf(arrayOrSlice)
	// 762 Chapter 28 ■ Using Reflection, Part 2
	if (arrayOrSliceVal.Kind() == reflect.Array ||
		arrayOrSliceVal.Kind() == reflect.Slice) &&
		arrayOrSliceVal.Type().Elem().Kind() == reflect.String {
		for i := 0; i < arrayOrSliceVal.Len(); i++ {
			Printfln("Element: %v, Value: %v", i, arrayOrSliceVal.Index(i).String())
		}
	}
}

func findAndSplit(slice interface{}, target interface{}) interface{} {
	sliceVal := reflect.ValueOf(slice)
	targetType := reflect.TypeOf(target)
	if sliceVal.Kind() == reflect.Slice && sliceVal.Type().Elem() == targetType {
		for i := 0; i < sliceVal.Len(); i++ {
			if sliceVal.Index(i).Interface() == target {
				return sliceVal.Slice(0, i+1)
			}
		}
	}
	return slice
}

// The pickValues function creates a new slice using the Type reflected from an
// existing slice and uses the Append function to add values to the new slice.
func pickValues(slice interface{}, indices ...int) interface{} {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() == reflect.Slice {
		// 765 Chapter 28 ■ Using Reflection, Part 2
		newSlice := reflect.MakeSlice(sliceVal.Type(), 0, 10)
		for _, index := range indices {
			newSlice = reflect.Append(newSlice, sliceVal.Index(index))
		}
		return newSlice
	}
	return nil
}

// The Kind method is used to confirm that the describeMap function has
// received a map and the Key and Elem methods are used to write out the key
// and value types.
func describeMap(m interface{}) {
	mapType := reflect.TypeOf(m)
	/* The Kind method is used to confirm that the describeMap function has
	* received a map and the Key and Elem methods are used to write out the key
	* and value types. */
	if mapType.Kind() == reflect.Map {
		Printfln("Key type: %v, Val type: %v", mapType.Key(), mapType.Elem())
	} else {
		Printfln("Not a map")
	}
}

/* The reflect package provides two different ways to enumerate the contents of
* a map. The first is to use the MapKeys method to get a slice containing the
* reflected key values and obtain each reflected map value using the MapIndex
* method, as shown in Listing 28-11. */
func printMapContents(m any) {
	mapValue := reflect.ValueOf(m)
	if mapValue.Kind() == reflect.Map {
		for _, keyVal := range mapValue.MapKeys() {
			reflectedVal := mapValue.MapIndex(keyVal)
			Printfln("Map Key: %v, Value: %v", keyVal, reflectedVal)
		}
	} else {
		Printfln("Not a map")
	}
}

/*
	The same effect can be achieved using the MapRange method, which returns a

* pointer to a MapIter value, which defines the methods described in Table
* 28-11.
The MapIter struct provides a cursor-based approach to enumerating maps, where
the Next method advances through the map contents, and the Key and Value
methods provide access to the key and value at the current position. The result
of the Next method indicates whether there are remaining values to be read,
which makes it convenient to use with a for loop, as shown in Listing 28-12.
*/
func printMapContentsWithMapRange(m any) {
	mapValue := reflect.ValueOf(m)
	if mapValue.Kind() == reflect.Map {
		iter := mapValue.MapRange()
		for iter.Next() {
			/* It is important to call the Next method before calling the Key
			* and Value methods and to avoid calling those methods when the
			* Next method returns false */
			Printfln("Map Key: %v, Value: %v", iter.Key(), iter.Value())
		}
	} else {
		Printfln("Not a map")
	}
}

/* The SetMapIndex method is used to add, modify, or remove key-value pairs in
* a map. Listing 28-13 defines functions for modifying a map. */
func setMap(m interface{}, key interface{}, val interface{}) {
	mapValue := reflect.ValueOf(m)
	keyValue := reflect.ValueOf(key)
	valValue := reflect.ValueOf(val)
	if mapValue.Kind() == reflect.Map &&
		mapValue.Type().Key() == keyValue.Type() &&
		mapValue.Type().Elem() == valValue.Type() {
		mapValue.SetMapIndex(keyValue, valValue)
	} else {
		Printfln("Not a map or mismatched types")
	}
}

func createMap(slice any, op func(any) any) any {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() == reflect.Slice {
		mapType := reflect.MapOf(sliceVal.Type().Elem(), sliceVal.Type().Elem())
		mapVal := reflect.MakeMap(mapType)
		for i := 0; i < sliceVal.Len(); i++ {
			elemVal := sliceVal.Index(i)
			mapVal.SetMapIndex(elemVal, reflect.ValueOf(op(elemVal.Interface())))
		}
		return mapVal.Interface()
	}
	return nil
}

func removeFromMap(m interface{}, key interface{}) {
	mapValue := reflect.ValueOf(m)
	keyValue := reflect.ValueOf(key)
	if mapValue.Kind() == reflect.Map &&
		mapValue.Type().Key() == keyValue.Type() {
		// This is a handy trick that ensures that the (float64) value is
		// removed from the map.
		mapValue.SetMapIndex(keyValue, reflect.Value{})
	}
}

func inspectStructs(structs ...interface{}) {
	for _, s := range structs {
		structType := reflect.TypeOf(s)
		if structType.Kind() == reflect.Struct {
			inspectStructType([]int{}, structType)
		}
	}
}

func inspectStructType(baseIndex []int, structType reflect.Type) {
	Printfln("--- Struct Type: %v", structType)
	for i := 0; i < structType.NumField(); i++ {
		fieldIndex := append(baseIndex, i)
		field := structType.Field(i)
		Printfln("Field %v: Name: %v, Type: %v, Exported: %v",
			fieldIndex, field.Name, field.Type, field.PkgPath == "")
		if field.Type.Kind() == reflect.Struct {
			field := structType.FieldByIndex(fieldIndex)
			inspectStructType(fieldIndex, field.Type)
		}
	}
	Printfln("--- End Struct Type: %v", structType)
}

func describeField(s interface{}, fieldName string) {
	structType := reflect.TypeOf(s)
	field, found := structType.FieldByName(fieldName)
	if found {
		Printfln("Found: %v, Type: %v, Index: %v",
			field.Name, field.Type, field.Index)
		index := field.Index
		for len(index) > 1 {
			index = index[0 : len(index)-1]
			field = structType.FieldByIndex(index)
			Printfln("Parent : %v, Type: %v, Index: %v",
				field.Name, field.Type, field.Index)
		}
		Printfln("Top-Level Type: %v", structType)
	} else {
		Printfln("Field %v not found", fieldName)
	}
}

func inspectTags(s interface{}, tagName string) {
	structType := reflect.TypeOf(s)
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		tag := field.Tag
		valGet := tag.Get(tagName)
		valLookup, ok := tag.Lookup(tagName)
		Printfln("Field: %v, Tag %v: %v", field.Name, tagName, valGet)
		Printfln("Field: %v, Tag %v: %v, Set: %v",
			field.Name, tagName, valLookup, ok)
	}
}

type Person struct {
	// 778 Chapter 28 ■ Using Reflection, Part 2
	Name    string `alias:"id"`
	City    string `alias:""`
	Country string
}

func getFieldValues(s any) {
	structValue := reflect.ValueOf(s)
	if structValue.Kind() == reflect.Struct {
		for i := 0; i < structValue.NumField(); i++ {
			fieldType := structValue.Type().Field(i)
			fieldVal := structValue.Field(i)
			Printfln("Name: %v, Type: %v, Value: %v",
				fieldType.Name, fieldType.Type, fieldVal)
		}
	} else {
		Printfln("Not a struct")
	}
}

/*
As with other data types, reflection can only be used to change values via a
pointer to the struct. The Elem method is used to follow the pointer so that
the Value that reflects the field can be obtained using one of the methods
described in Table 28-17. The CanSet method is used to determine if a field can
be set.  An additional step is required for fields that are not nested structs,
which is to create a pointer to the field value using the Addr method
*/
func setFieldValue(s interface{}, newVals map[string]interface{}) {
	structValue := reflect.ValueOf(s)
	if structValue.Kind() == reflect.Ptr &&
		structValue.Elem().Kind() == reflect.Struct {
		for name, newValue := range newVals {
			fieldVal := structValue.Elem().FieldByName(name)
			if fieldVal.CanSet() {
				fieldVal.Set(reflect.ValueOf(newValue))
			} else if fieldVal.CanAddr() {
				ptr := fieldVal.Addr()
				if ptr.CanSet() {
					ptr.Set(reflect.ValueOf(newValue))
				} else {
					Printfln("Cannot set field via pointer")
				}
			} else {
				Printfln("Cannot set field")
			}
		}
	} else {
		panic("Not a pointer to a struct! Cannot change a struct value")
	}
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

	Printfln("%s%s%s%s",
		"\n---------\n",
		"Using Reflection, Part 2",
		"\n Preparing for This Chapter",
		"\n Working with Pointers")
	Printfln("")

	name = "Alice"
	t := reflect.TypeOf(name)
	Printfln("Original Type: %v for value %v", t, name)
	pt := createPointerType(t)
	Printfln("Pointer Type: %v", pt)
	Printfln("Follow pointer type: %v", followPointerType(pt))

	transformString(&name)
	Printfln("Follow pointer value: %v", name)

	Printfln("\n Working with Array and Slice Types")

	name = "Alice"
	city = "London"
	hobby := "Running"

	sliceOfStr := []string{name, city, hobby}
	array := [3]string{name, city, hobby}
	Printfln("Slice (string): %v", checkElemType("testString", sliceOfStr))
	Printfln("Array (string): %v", checkElemType("testString", array))
	Printfln("Array (int): %v", checkElemType(10, array))
	Printfln("Array (string): %v", checkElemType("10", array))

	Printfln("\n Working with Array and Slice Values")

	Printfln("Original slice: %v", sliceOfStr)
	newCity := "Paris"
	setValue(sliceOfStr, 1, newCity)
	Printfln("Modified slice: %v", sliceOfStr)
	Printfln("Original slice: %v", array)
	newCity = "Rome"
	setValue(&array, 1, newCity)
	Printfln("Modified slice: %v", array)

	Printfln("\n  Enumerating Slices and Arrays")
	/* The Len method can be used to set the limit in a for loop to enumerate
	* the elements in an array or slice, as shown in Listing 28-7. */
	enumerateStrings(sliceOfStr)
	enumerateStrings(array)

	Printfln("\n  Creating New Slices from Existing Slices")

	Printfln("Strings: %v", findAndSplit(sliceOfStr, "Paris"))
	numbers := []int{1, 3, 4, 5, 7}
	Printfln("Numbers: %v", findAndSplit(numbers, 4))

	Printfln("\n  Creating, Copying, and Appending Elements to Slices")
	sliceOfStr = append(sliceOfStr, "Bob", "Paris", "Soccer")
	Printfln("sliceOfStr: %v", sliceOfStr)
	picked := pickValues(sliceOfStr, 0, 3, 5)
	Printfln("Picked values: %v", picked)

	Printfln("\n Working with Map Types")
	pricesMap := map[string]float64{
		"Kayak": 279, "Lifejacket": 48.95, "Soccer Ball": 19.50,
	}
	describeMap(pricesMap)

	Printfln("\n Working with Map Values")
	printMapContents(pricesMap)
	printMapContentsWithMapRange(pricesMap)

	Printfln("\n  Setting and Removing Map Values")
	/* As noted in Chapter 7, maps are not copied when they are used as
	* arguments and so a pointer isn’t required to modify the contents of a
	* map. */
	setMap(pricesMap, "Kayak", 100.00)
	setMap(pricesMap, "Hat", 10.00)
	removeFromMap(pricesMap, "Lifejacket")
	for k, v := range pricesMap {
		// 770 Chapter 28 ■ Using Reflection, Part 2
		Printfln("Key: %v, Value: %v", k, v)
	}

	Printfln("\n  Creating New Maps")
	names = []string{"Alice", "Bob", "Charlie"}
	reverse := func(val any) any {
		if str, ok := val.(string); ok {
			return strings.ToUpper(str)
		}
		return val
	}
	namesMap := createMap(names, reverse).(map[string]string)
	for k, v := range namesMap {
		Printfln("Key: %v, Value:%v", k, v)
	}

	Printfln("%s%s", "\n Working with Struct Types", "\n  Processing Nested Fields")
	inspectStructs(Purchase{})

	Printfln("\n  Locating a Field by Name")

	describeField(Purchase{}, "Price")
	describeField(Product{}, "Name")
	describeField(Purchase{}, "AlaBala")

	Printfln("\n  Inspecting Struct Tags")
	inspectTags(Person{}, "alias")

	Printfln("\n  Creating Struct Types")

	stringType := reflect.TypeOf("this is a string")
	Printfln("The type of a created by reflect.TypeOf(\"\") as a string: %s", stringType)
	structType := reflect.StructOf([]reflect.StructField{
		{Name: "Name", Type: stringType, Tag: `alias:"id"`},
		{Name: "City", Type: stringType, Tag: `alias:""`},
		{Name: "Country", Type: stringType},
	})
	inspectTags(reflect.New(structType), "alias")

	Printfln("\n Working with Struct Values")

	purchase := Purchase{
		Customer: Customer{Name: "Acme", City: "Chicago"},
		Product:  Product{Name: "Kayak", Category: "Watersports", Price: 279},
		Total:    279,
		taxRate:  10,
	}

	setFieldValue(&purchase, map[string]any{
		"City": "London", "Category": "Boats", "Total": 100.50,
	})
	getFieldValues(purchase)

	// Using Reflection, Part 3
	Printfln("\n\nUsing Reflection, Part 3\n")

}
