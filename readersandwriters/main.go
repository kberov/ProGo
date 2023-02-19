package main

import (
	"bufio"
	"encoding/json"
	. "fmt"

	"io"
	"os"
	"strings"
)

func processData(reader io.Reader, writer io.Writer) {
	buf := make([]byte, 2)
	for {
		count, err := reader.Read(buf)
		if count > 0 {
			c, _ := writer.Write(buf[0:count])
			Printfln("Read %d bytes: %s", count, string(buf[0:count]))
			Printfln("Wrote %d bytes: %s", c, string(buf[0:count]))
		}
		if err == io.EOF {
			break
		}

	}
	println(string(buf))
}

func _processData(reader io.Reader, writer io.Writer) {
	count, err := io.Copy(writer, reader)
	if err == nil {
		Printfln("Read %v bytes", count)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func GenerateData(writer io.Writer) {
	data := []byte("Kayak, Lifejacket")
	writeSize := 4
	for i := 0; i < len(data); i += writeSize {
		end := i + writeSize
		if end > len(data) {
			end = len(data)
		}
		count, err := writer.Write(data[i:end])
		Printfln("Wrote %v byte(s): %v", count, string(data[i:end]))
		if err != nil {
			Printfln("Error: %v", err.Error())
		}
	}

	if closer, ok := writer.(io.Closer); ok {
		_ = closer.Close()
	}
}

func ConsumeData(reader io.Reader) {
	data := make([]byte, 0, 10)
	slice := make([]byte, 2)
	for {
		count, err := reader.Read(slice)
		if count > 0 {
			Printfln("Read data: %v", string(slice[0:count]))
			data = append(data, slice[0:count]...)
		}
		if err == io.EOF {
			break
		}
	}
	Printfln("All data: %v", string(data))
}

func main() {
	Printfln("\nReading and Writing Data")
	Printfln("Product: %v, Price : %v", Kayak.Name, Kayak.Price)
	Printfln("\nUnderstanding Readers")
	//r := strings.NewReader("Kayak")
	//processData(r)

	Printfln("\nUnderstanding Writers")
	r := strings.NewReader("Kayakaa")
	var builder strings.Builder
	processData(r, &builder)
	Printfln("String builder contents: %s", builder.String())

	Printfln("\nUtility Functions for Readers and Writers")
	r = strings.NewReader("A string wihich will be written to standart output\n\n")
	_processData(r, os.Stdout)
	Printfln("\nSpecialized Readers and Writers")
	Printfln("Pipes")
	pipeReader, pipeWriter := io.Pipe()
	go GenerateData(pipeWriter)
	ConsumeData(pipeReader)

	Printfln("\nConcatenating Multiple Readers")

	r1 := strings.NewReader("Kayak")
	r2 := strings.NewReader("Lifejacket")
	r3 := strings.NewReader("Canoe")
	concatReader := io.MultiReader(r1, r2, r3)
	ConsumeData(concatReader)

	Printfln("\nCombining Multiple Writers")
	var w1 strings.Builder
	var w2 strings.Builder
	var w3 strings.Builder
	combinedWriter := io.MultiWriter(&w1, &w2, &w3)
	GenerateData(combinedWriter)
	Printfln("Writer #1: %v", w1.String())
	Printfln("Writer #2: %v", w2.String())
	Printfln("Writer #3: %v", w3.String())

	Printfln("\nEchoing Reads to a Writer")
	// The TeeReader function returns a Reader that echoes the data that it
	// receives to a Writer, as shown in Listing 20-15.
	r1 = strings.NewReader("Kayak")
	r2 = strings.NewReader("Lifejacket")
	r3 = strings.NewReader("Canoe")
	concatReader = io.MultiReader(r1, r2, r3)
	var writer strings.Builder
	// The TeeReader function is used to create a Reader that will echo data to
	// a strings.Builder
	teeReader := io.TeeReader(concatReader, &writer)
	ConsumeData(teeReader)
	Printfln("Echo data: %v", writer.String())

	Printfln("\nLimiting Read Data")
	r1 = strings.NewReader("Kayak")
	r2 = strings.NewReader("Lifejacket")
	r3 = strings.NewReader("Canoe")
	concatReader = io.MultiReader(r1, r2, r3)
	// The LimitReader function is used to restrict the amount of data that can be
	// obtained from a Reader, as shown in Listing 20-16.
	// The second argument is the maximum number of bytes that can be read
	limited := io.LimitReader(concatReader, 8)
	ConsumeData(limited)

	Printfln("\nBuffering Data",
		"\nUsing the Additional Buffered Reader Methods")
	// The bufio package provides support for adding buffers to readers and writers.
	text := "It was a boat. A small boat."
	var reader io.Reader = NewCustomReader(strings.NewReader(text))
	//var writer strings.Builder - ALREADY DECLARED SOMEWHERE ABOVE

	// It is the size of the byte slice passed to the Read function that
	// determines how data is consumed. In this case, the size of the slice is
	// five, which means that a maximum of five bytes is read for each call to
	// the Read function.
	slice := make([]byte, 5)
	// using bufio Reader instad of our custom reader
	// NewReader creates a Reader with the default buffer size(4096 bytes)
	buffered := bufio.NewReader(reader)
	for {
		count, err := buffered.Read(slice)
		if count > 0 {
			Printfln("Buffer size: %v, buffered: %v",
				buffered.Size(), buffered.Buffered())
			writer.Write(slice[0:count])
		}
		// nothing to read, exit the loop
		if err != nil {
			break
		}
	}
	Printfln("Read buffered data: %v", writer.String())

	Printfln("")
	/*
		The NewReader and NewReaderSize functions return bufio.Reader values, which
		implement the io.  Reader interface and which can be used as drop-in wrappers
		for other types of Reader methods, seamlessly introducing a read buffer.  The
		bufio.Reader struct defines additional methods that make direct use of the
		buffer, as described in Table 20-9.
	*/
	Printfln("Performing Buffered Writes")
	// The bufio package also provides support for creating writers that use a
	// buffer, using the functions described in Table 20-10.

	//var cwriter = NewCustomWriter(&builder)
	var cwriter = bufio.NewWriterSize(NewCustomWriter(&builder), 20)
	for i := 0; true; {
		end := i + 5
		if end >= len(text) {
			_, _ = cwriter.Write([]byte(text[i:]))
			// The transition to a buffered Writer isn’t entirely seamless
			// because it is important to call the Flush method to ensure that
			// all the data is written out.
			_ = cwriter.Flush()
			break
		}
		_, _ = cwriter.Write([]byte(text[i:end]))
		i = end
	}
	Printfln("Written data: %v", builder.String())

	Printfln("Formatting and Scanning with Readers and Writers")

	scannedReader := strings.NewReader("Kayak Watersports $279.00")
	var name, category string
	var price float64
	scanTemplate := "%s %s $%f"
	_, err := scanFromReader(scannedReader, scanTemplate, &name, &category, &price)
	if err != nil {
		Printfln("Error: %v", err.Error())
	} else {
		Printfln("Name: %v", name)
		Printfln("Category: %v", category)
		Printfln("Price: %.2f", price)
	}

	reader = strings.NewReader("Kayak Watersports $279.00")
	for {
		var str string
		scannedItem, err := scanSingle(reader, &str)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		Printfln("Scanned Item: %d Value: %v", scannedItem, str)
	}

	Printfln("\nWriting Formatted Strings to a Writer")
	template := "Name: %s, Category: %s, Price: $%.2f"
	writeFormatted(&writer, template, "Kayak", "Watersports", float64(279))
	Printfln(writer.String())

	Printfln("\nUsing a Replacer with a Writer")
	text = "It was a boat. A small boat."
	subs := []string{"boat", "kayak", "small", "huge"}
	var mockwriter strings.Builder
	replacer := strings.NewReplacer(subs...)

	lenwritten, _ := replacer.WriteString(&mockwriter, text)
	Printfln("Source: '%s' Replaced '%s'.written text: %d", text, mockwriter.String(), lenwritten)

	Printfln("Reading and Writing JSON Data")
	encoder := json.NewEncoder(os.Stdout)
	for _, val := range []any{20, 18.5, "нещо", Products} {
		err := encoder.Encode(val)
		if err != nil {
			println(err)
		}

	}
	names := []string{"Kayak", "Lifejacket", "Soccer Ball"}
	numbers := [3]int{10, 20, 30}
	var byteArray [5]byte
	copy(byteArray[0:], []byte(names[0]))
	byteSlice := []byte(names[0])
	// formatted printing to stdout
	Fprintf(os.Stdout, "%v", "ArrOfStrs: ")
	_ = encoder.Encode(names)
	Fprintf(os.Stdout, "%v", "ArrOfInts: ")
	_ = encoder.Encode(numbers)
	Fprintf(os.Stdout, "%v", "ArrOfBytes: ")
	_ = encoder.Encode(byteArray)
	Fprintf(os.Stdout, "%v", "SliceOfBytes: ")
	_ = encoder.Encode(byteSlice)

	m := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 49.95,
	}
	Fprintf(os.Stdout, "%v", "Map: ")
	_ = encoder.Encode(m)
	dp := DiscountedProduct{
		Product:  &Products[0],
		Discount: 10.50,
	}
	Fprintf(os.Stdout, "%v", "Struct with embedded struct: ")
	_ = encoder.Encode(dp)

	Println("\nCustomizing the JSON Encoding of Structs")
	/*
		How a struct is encoded can be customized using struct tags, which are string
		literals that follow fields. Struct tags are part of the Go support for
		reflection, which I describe in Chapter 28
	*/
	dp2 := DiscountedProduct{Discount: 10.50}
	Fprintf(os.Stdout, "%v", "Struct with omitted null embedded struct: ")
	_ = encoder.Encode(&dp2)
	dp2 = DiscountedProduct{Discount: 10.50, Product: &Products[1]}
	Fprintf(os.Stdout, "%v", "Struct with formatted float64 as string:\n")
	_ = encoder.Encode(&dp2)

	Println("\nEncoding Interfaces")
	namedItems := []Named{&dp2, &Person{PersonName: "Alice"}}
	/*
	   No aspect of the interface is used to adapt the JSON, and all the exported
	   fields of each value in the slice are included in the JSON. This can be a
	   useful feature, but care must be taken when decoding this kind of JSON, because
	   each value can have a different set of fields, as I explain in the “Decoding
	   Arrays” section.
	*/
	Fprintf(os.Stdout, "%v", "List of structs implementing the `Named` interface:\n")
	_ = encoder.Encode(namedItems)

	Printfln("\nCreating Completely Custom JSON Encodings\n....\nDecoding JSON Data")

	/*
		//go run . < Names.json
		namesDecoder := json.NewDecoder(os.Stdin)
		var anything []Named
		if err := namesDecoder.Decode(&anything); err != nil {
			println("Error decoding json: ", err)
		}
		Printfln("JSON from STDIN:\n %v", anything)
	*/

	Printfln("\nDecoding Number Values")
	reader = strings.NewReader(`true "Hello" 99.99 200`)
	vals := []any{}
	decoder := json.NewDecoder(reader)
	decoder.UseNumber()
	for {
		var decodedVal any
		err := decoder.Decode(&decodedVal)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		vals = append(vals, decodedVal)
	}
	for _, val := range vals {
		if num, ok := val.(json.Number); ok {
			if ival, err := num.Int64(); err == nil {
				Printfln("Decoded Integer: %v", ival)
			} else {

				Printfln("Error decoding Integer: %s, %v", err, ival)
			}
			if fpval, err := num.Float64(); err == nil {
				Printfln("Decoded Floating Point: %v", fpval)
			} else {
				Printfln("Decoded String: %v", num.String())
			}
		} else {
			Printfln("Decoded (%T): %v", val, val)
		}
	}

	Printfln("\nSpecifying Types for Decoding")
	reader = strings.NewReader(`true "Hello" 99.99 200`)
	var bval bool
	var sval string
	var fpval float64
	var ival int
	vals = []any{&bval, &sval, &fpval, &ival}
	decoder = json.NewDecoder(reader)
	for i := 0; i < len(vals); i++ {
		err := decoder.Decode(vals[i])
		if err != nil {
			Printfln("Error: %v", err.Error())
			break
		}
	}
	Printfln("Decoded (%T): %v", bval, bval)
	Printfln("Decoded (%T): %v", sval, sval)
	Printfln("Decoded (%T): %v", fpval, fpval)
	Printfln("Decoded (%T): %v", ival, ival)

	Printfln("Decoding Arrays...\nDexoding Maps")
	reader = strings.NewReader(`{"Kayak" : 279, "Lifejacket" : 49.95}`)
	m = map[string]float64{}
	decoder = json.NewDecoder(reader)
	err = decoder.Decode(&m)
	if err != nil {
		Printfln("Error: %v", err.Error())
	} else {
		Printfln("Map: %T, %#v", m, m)
		for k, v := range m {
			Printfln("Key: %v, Value: %v", k, v)
		}
	}

	Println("\nDecoding Structs\n    Disallowing Unused Keys")
	// See DECODING TO INTERFACE TYPES
	reader = strings.NewReader(`
{"Name":"Kayak","Category":"Watersports","Price":279}
{"Name":"Lifejacket","Category":"Watersports" }
{"name":"Canoe","category":"Watersports", "price": 100, "inStock": true }
`)
	decoder = json.NewDecoder(reader)
	decoder.DisallowUnknownFields()
	for {
		// We know the type of the incomming struct and we decode directly to it.
		var val Product
		err := decoder.Decode(&val)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		} else {
			Printfln("Name: %v, Category: %v, Price: %v",
				val.Name, val.Category, val.Price)
		}
	}

	Println("\nUsing Struct Tags to Control Decoding")
	reader = strings.NewReader(`
{"Name":"Kayak","Category":"Watersports","Price":279, "Offer": "10"}`)
	decoder = json.NewDecoder(reader)
	for {
		var val DiscountedProduct
		err := decoder.Decode(&val)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		} else {
			Printfln("Name: %v, Category: %v, Price: %v, Discount: %v",
				val.Name, val.Category, val.Price, val.Discount)
		}
	}

	Println("Creating Completely Custom JSON Decoders")
	/*
	   The Decoder checks to see whether a struct implements the Unmarshaler
	   interface, which denotes a type that has a custom encoding, and which defines
	   the method described in Table 21-10.
	*/
}

func writeFormatted(writer io.Writer, template string, vals ...interface{}) {
	Fprintf(writer, template, vals...)
}
func scanSingle(reader io.Reader, val interface{}) (int, error) {
	return Fscan(reader, val)
}

func scanFromReader(reader io.Reader, template string,
	vals ...interface{}) (int, error) {
	return Fscanf(reader, template, vals...)
}
