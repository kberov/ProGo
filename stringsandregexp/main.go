package main

import (
	"bytes"
	. "fmt"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	Println("String Processing and Regular Expressions")
	product := "Kayak"
	Println("Product: ", product)
	Println("\nProcessing Strings")
	Println("Comparing Strings")
	Println("Браво е брАво: ", strings.EqualFold("Браво", "брАво"))

	Println("Contains:", strings.Contains(product, "yak"))
	Println("ContainsAny:", strings.ContainsAny(product, "abc"))
	Println("ContainsRune:", strings.ContainsRune(product, 'K'))
	Println("EqualFold:", strings.EqualFold(product, "KAYAK"))
	Println("HasPrefix:", strings.HasPrefix(product, "Ka"))
	Println("HasSuffix:", strings.HasSuffix(product, "yak"))

	/* For all the functions in the strings package, which operate on
	* characters, there is a corresponding function in the bytes package that
	* operates on a byte slice */
	price := "€100"
	Println("Strings Prefix:", strings.HasPrefix(price, "€"))
	Println("Bytes Prefix:", bytes.HasPrefix([]byte(price), []byte{226, 130}))

	Println("\nConverting String Case")
	description := "Кораб за плаване."
	Println("Original:", description)
	Println("Титле:", strings.Title(description))
	description = "Boat for traveling"
	Println("Original:", description)
	Println("Титле:", strings.Title(description))

	specialChar := "\u01c9"
	Println("Original:", specialChar, []byte(specialChar))
	upperChar := strings.ToUpper(specialChar)
	Println("Upper:", upperChar, []byte(upperChar))
	titleChar := strings.ToTitle(specialChar)
	Println("Title:", titleChar, []byte(titleChar))
	product = "Каяк"
	for _, char := range product {
		Printf("%s Upper case: %t\n", string(char), unicode.IsUpper(char))
	}

	Println("\nInspecting Strings")
	Println("How many 'ая' in "+product+":", strings.Count(product, "ая"))
	product = "Kayak"
	Println("How many 'ay' in "+product+":", strings.Count(product, "ay"))

	description = "A boat for one person"
	Println("description is now:", description)
	Println("Count o:", strings.Count(description, "o"))
	Println("Index o:", strings.Index(description, "o"))
	Println("LastIndex o:", strings.LastIndex(description, "o"))
	Println("IndexAny abcd:", strings.IndexAny(description, "abcd"))
	Println("LastIndex o:", strings.LastIndex(description, "o"))
	Println("LastIndexAny abcd:", strings.LastIndexAny(description, "abcd"))

	Println("\nInspecting Strings with Custom Functions")

	isLetterB := func(r rune) bool {
		return r == 'B' || r == 'b'
	}
	isLetterБ := func(r rune) bool {
		return r == 'Б' || r == 'б'
	}

	Println("IndexFunc B||b:", strings.IndexFunc(description, isLetterB))
	Println("IndexFunc ^Б||б:", strings.IndexFunc(description, isLetterБ))

	Println("\nManipulating Strings")
	splits := strings.Split(description, " ")
	for _, x := range splits {
		Println("Split >>" + x + "<<")
	}
	splitsAfter := strings.SplitAfter(description, " ")
	for _, x := range splitsAfter {
		Println("SplitAfter >>" + x + "<<")
	}
	Printf("Fields: %#v\n", strings.Fields(description))
	description = "Лодка за един човек на ден до пладне."
	Printf("Fields: %#v\n", strings.Fields(description))
	Println("\nRestricting the Number of Results")
	Printf("SplitN: %#v\n", strings.SplitN(description, "д", 3))

	Println("\nSplitting on Whitespace Characters")
	description = "С  някои  двойни пространства  " + description
	// the SplitN function splits only on the first space
	splits = strings.SplitN(description, " ", 4)
	for _, x := range splits {
		Println("SplitN(s, \" \", 4) >>" + x + "<<")
	}
	// To deal with repeated whitespace characters, the Fields function breaks
	// strings on any whitespace character.
	Printf("Fields: %#v\n", strings.Fields(description))

	Println("\nSplitting Using a Custom Function to Split String")
	Printf("FieldsFunc: %#v\n", strings.FieldsFunc(description,
		func(c rune) bool { return c == 'д' }))
	Printf("FieldsFunc: %#v\n", strings.FieldsFunc(description,
		func(c rune) bool { return unicode.IsSpace(c) }))

	Println("\n\nTrimming Strings")
	/* The process of trimming removes leading and trailing characters from a
	 * string and is most often used to remove whitespace characters. */
	Println("\nTrimming Whitespace")
	username := " Alice"
	Printf("username: [%s]\n", username)
	trimmed := strings.TrimSpace(username)
	Println("Trimmed:", ">>"+trimmed+"<<")

	Println("\nTrimming Character Sets")
	trimmed = strings.Trim(description, "Ск. ")
	Printf("Trimmed:[%s]\n", trimmed)

	Println("\nTrimming Substrings")
	description = "A boat for one person"
	// The start or end of the target string must exactly match the specified prefix or suffix.
	prefixTrimmed := strings.TrimPrefix(description, "A boat ")
	wrongPrefix := strings.TrimPrefix(description, "A hat ")
	Println("Trimmed:", prefixTrimmed)
	Println("Not trimmed:", wrongPrefix)

	Println("\nTrimming with Custom Functions")
	trimmer := func(r rune) bool {
		return r == 'A' || r == 'n'
	}
	trimmed = strings.TrimFunc(description, trimmer)
	Println("Trimmed:", trimmed)

	Println("\nAltering Strings")

	text := "It was a boat. A small boat."
	replace := strings.Replace(text, "boat", "canoe", 1)
	replaceAll := strings.ReplaceAll(text, "boat", "truck")
	Println("Replace:", replace)
	Println("Replace All:", replaceAll)

	Println("\nAltering Strings with a Map Function")
	mapper := func(r rune) rune {
		if r == 's' {
			return 'Z'
		}
		return r
	}
	mapped := strings.Map(mapper, text)
	Println("Mapped:", mapped)

	Println("\nUsing a String Replacer")
	replacer := strings.NewReplacer("boat", "kayak", "small", "huge")
	replaced := replacer.Replace(text)
	Println("Replaced:", replaced)

	Println("\nBuilding and Generating Strings")
	elements := strings.Fields(text)
	joined := strings.Join(elements, "--")
	Println("Joined:", joined)

	var builder strings.Builder
	for _, sub := range strings.Fields(text) {
		if sub == "small" {
			builder.WriteString("very ")
		}
		builder.WriteString(sub)
		builder.WriteRune(' ')
	}
	Println("String:", builder.String())

	Println("\nUsing Regular Expressions")
	// The regexp package provides support for regular expressions, which allow
	// complex patterns to be found in strings.
	re := "[A-z]oat"
	match, err := regexp.MatchString(re, description)
	if err == nil {
		Printf("Match '%s' with /%s/: %t\n", description, re, match)
	} else {
		Println("Error:", err)
	}

	Println("\nCompiling and Reusing Patterns")
	pattern, compileErr := regexp.Compile("[A-z]oat")
	question := "Is that a goat?"
	preference := "I like oats"
	if compileErr == nil {
		Println("Description:", pattern.MatchString(description))
		Println("Question:", pattern.MatchString(question))
		Println("Preference:", pattern.MatchString(preference))
	} else {
		Println("Error:", compileErr)
	}

	pattern = regexp.MustCompile("K[a-z]{4}|[A-z]oat")
	description = "Kayak. A boat for one person."
	firstIndex := pattern.FindStringIndex(description)
	allIndices := pattern.FindAllStringIndex(description, -1)
	Println("First index", firstIndex[0], "-", firstIndex[1], "=", getSubstring(description, firstIndex))
	for i, idx := range allIndices {
		Println("Index", i, "=", idx[0], "-",
			idx[1], "=", getSubstring(description, idx))
	}

	// index, substring
	firstMatch := pattern.FindString(description)
	allMatches := pattern.FindAllString(description, -1)
	Println("First match:", firstMatch)
	for i, m := range allMatches {
		Println("Match", i, "=", m)
	}

	Println("\nSplitting Strings Using a Regular Expression")
	pattern = regexp.MustCompile(" |boat|one")
	split := pattern.Split(description, -1)
	for _, s := range split {
		if s != "" {
			Println("Substring:", s)
		}
	}

	Println("\nUsing Subexpressions")
	/* Subexpressions allow parts of a regular expression to be accessed, which can
	 * make it easier to extract substrings from within a matched region. */

	pattern = regexp.MustCompile(`A\s([A-z]*)\sfor\s([A-z]*?)\sperson`)
	str := pattern.FindString(description)
	Printf("Match: %s\n", str)
	subs := pattern.FindStringSubmatch(description)
	for _, str := range subs {
		Printf("Match: %s\n", str)
	}

	Println("\nUsing Named Subexpressions")
	pattern = regexp.MustCompile("A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")
	subs = pattern.FindStringSubmatch(description)
	for _, name := range []string{"type", "capacity"} {
		// The SubexpIndex method returns the position of a named subexpression
		// in the result.
		Println(name, "=", subs[pattern.SubexpIndex(name)])
	}

	Println("\nReplacing Substrings Using a Regular Expression")
	template := "(type: ${type}, capacity: ${capacity})"
	replaced = pattern.ReplaceAllString(description, template)
	Println("Replaced:", replaced)

	Println("\nReplacing Matched Content with a Function")
	replaced = pattern.ReplaceAllStringFunc(description, func(s string) string {
		println("The matched portion of the string:", s)
		return "This is the replacement content"
	})
	Println("Replaced:", replaced)

}

func getSubstring(s string, indices []int) string {
	return string(s[indices[0]:indices[1]])
}
