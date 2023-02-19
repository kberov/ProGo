package main

import (
	"fmt"
	. "fmt"
	"os"
	"syscall"
	"time"

	// EMBEDDING THE TIME ZONE DATABASE
	// To ensure that time zone data will always be available, declare a
	// dependency on the package like this:
	_ "time/tzdata"
)

func writeToChannel(nameChannel chan<- string) {
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	Printfln("Waiting for initial duration...")
	_ = <-time.After(time.Second * 2)
	Printfln("Initial duration elapsed.")
	index := 0
	for {
		Println("sending", names[index], "to the channel from a goroutine",
			"and sleeping for a second")
		nameChannel <- names[index]
		index++
		time.Sleep(time.Second * 1)
		if index == len(names) {
			close(nameChannel)
			break
		}
	}
}

func _writeToChannel(channel chan<- string) {
	timer := time.NewTimer(time.Minute * 10)
	go func() {
		time.Sleep(time.Second * 2)
		Printfln("Resetting timer")
		timer.Reset(time.Second)
	}()
	Printfln("Waiting for initial duration...")
	<-timer.C
	Printfln("Initial duration elapsed.")
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
		// time.Sleep(time.Second * 3)
	}
	close(channel)
}

func __writeToChannel(nameChannel chan<- string) {
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	ticker := time.NewTicker(time.Second / 10)
	index := 0
	for {
		<-ticker.C
		nameChannel <- names[index]
		index++
		if index == len(names) {
			ticker.Stop()
			close(nameChannel)
			break
		}
	}
}

func PrintTime(label string, t *time.Time) {
	// Printfln("%s: Day: %v: Month: %v Year: %v",
	// label, t.Day(), t.Month(), t.Year())
	// Printf("%s: %s\n", label, t.Format("Day: 02 Month: Jan Year: 06"))
	// Printf("%s: %s\n", label, t.Format("2006-01-02"))
	Printf("%s: %s\n", label, t.Format(time.RFC822Z))
}
func main() {
	Println("\nDates, Times, and Durations")
	current := time.Now()
	spec := time.Date(1995, time.April, 9, 0, 0, 0, 0, time.Local)
	unix := time.Unix(3600*24*31, 0)
	PrintTime("Current", &current)
	PrintTime("Spec", &spec)
	PrintTime("Unix", &unix)
	Println("\nFormatting Times as Strings")
	home, _ := syscall.Getenv("HOME")
	Printf("HOME: %v\n", home)

	Println("\nParsing Time Values from Strings")
	Println("time.RFC3339:", time.RFC3339)

	layout := "2006-Jan-02"
	dates := []string{
		"1995-Jun-09",
		"2015-Jun-02",
	}
	for _, d := range dates {
		time, err := time.Parse(layout, d)
		if err == nil {
			PrintTime("Parsed", &time)
		} else {
			Printfln("Error: %s", err.Error())
		}
	}
	// Specifying a Parsing Location

	layout = "02 Jan 06 15:04"
	date := "09 Jun 95 19:30"
	london, lonerr := time.LoadLocation("Europe/London")
	newyork, nycerr := time.LoadLocation("America/New_York")
	Printfln("Londontime: %s %s", london, lonerr)
	Printfln("NewYorkTime: %s %s", newyork, nycerr)

	Println("Using the Local Location")
	local, _ := time.LoadLocation("Local")
	Println("Specifying Time Zones Directly")
	london = time.FixedZone("BST", 1*60*60)
	newyork = time.FixedZone("EDT", -4*60*60)
	local = time.FixedZone("Local", +3*3600)
	// if lonerr == nil && nycerr == nil {
	nolocation, _ := time.Parse(layout, date)
	londonTime, _ := time.ParseInLocation(layout, date, london)
	newyorkTime, _ := time.ParseInLocation(layout, date, newyork)
	localTime, _ := time.ParseInLocation(layout, date, local)
	PrintTime("No location:", &nolocation)
	PrintTime("London:", &londonTime)
	PrintTime("New York:", &newyorkTime)
	PrintTime("Local:", &localTime)
	//} else {
	//fmt.Println(lonerr.Error(), nycerr.Error())
	//}

	Println("\nManipulating Time Values")

	t, err := time.Parse(time.RFC822, "09 Jun 95 04:59 BST")
	if err == nil {
		Printfln("After: %v", t.After(time.Now()))
		Printfln("Round: %v", t.Round(time.Hour))
		Printfln("Truncate: %v", t.Truncate(time.Hour))
	} else {
		fmt.Println(err.Error())
	}
	t1, _ := time.Parse(time.RFC822Z, "09 Jun 95 04:59 +0100")
	t2, _ := time.Parse(time.RFC822Z, "08 Jun 95 23:59 -0400")
	Printfln("Equal Method: %v", t1.Equal(t2))
	Printfln("Equality Operator: %v", t1 == t2)
	Printfln("Equality Operator: t1 and t2 as strings: %s, %s", t1.String(), t2.String())
	Println("\nRepresenting Durations")
	/*
	   The Duration type is an alias to the int64 type and is used to represent a
	   specific number of milliseconds.  Custom Duration values are composed from
	   constant Duration values defined in the time package, described in Table
	   19-11.
	*/
	var d time.Duration = time.Hour + (30 * time.Minute)
	Printfln("Hours: %v", d.Hours())
	Printfln("Mins: %v", d.Minutes())
	Printfln("Seconds: %v", d.Seconds())
	Printfln("Millseconds: %v", d.Milliseconds())
	rounded := d.Round(time.Hour)
	Printfln("Rounded Hours: %v", rounded.Hours())
	Printfln("Rounded Mins: %v", rounded.Minutes())
	trunc := d.Truncate(time.Hour)
	Printfln("Truncated  Hours: %v", trunc.Hours())
	Printfln("Rounded Mins: %v", trunc.Minutes())

	Println("\nCreating Durations Relative to a Time")
	toYears := func(d time.Duration) int {
		return int(d.Hours() / (24 * 365))
	}

	future := time.Date(2051, 0, 0, 0, 0, 0, 0, time.Local)
	past := time.Date(1965, 0, 0, 0, 0, 0, 0, time.Local)

	Printfln("Future: %v", toYears(time.Until(future)))
	Printfln("Past: %v", toYears(time.Since(past)))

	Println("\nCreating Durations from Strings")
	d, err = time.ParseDuration("1h30m")
	if err == nil {
		Printfln("Hours: %v", d.Hours())
		Printfln("Mins: %v", d.Minutes())
		Printfln("Seconds: %v", d.Seconds())
		Printfln("Millseconds: %v", d.Milliseconds())
	} else {
		Println("Error parsing duration:", err.Error())
	}

	Println("\nUsing the Time Features for Goroutines and Channels")

	Printfln("Putting a Goroutine to Sleep in two seconds...\n")
	Printf("HOME: %v\n", os.Getenv("HOME"))
	time.Sleep(time.Second * 2)

	nameChannel := make(chan string)

	go writeToChannel(nameChannel)

	for name := range nameChannel {
		Printfln("Read name: %v in a range from the channel in the main goroutine", name)
	}

	Println("\nDeferring Execution of a Function (for 5 seconds)")
	// create the channel again to write to it and read from it.
	nameChannel = make(chan string)

	// AfterFunc(duration, func) -- This function executes the specified function in its own goroutine after the specified duration
	time.AfterFunc(time.Second*5, func() { writeToChannel(nameChannel) })

	for name := range nameChannel {
		Printfln("Read name: %v in a range from the channel in the main goroutine", name)
	}

	Println("\nUsing Notifications as Timeouts in Select Statements")
	// create the channel again to write to it and read from it.
	nameChannel = make(chan string)
	go writeToChannel(nameChannel)
	channelOpen := true
	for channelOpen {
		Printfln("Starting channel read")
		select {
		case name, ok := <-nameChannel:
			if !ok {
				channelOpen = false
				break
			} else {
				Printfln("Read name: %v", name)
			}
		case <-time.After(time.Second * 2):
			Printfln("Timeout")
		}
	}
	Println("\nStopping and Resetting Timers")
	nameChannel = make(chan string)
	go _writeToChannel(nameChannel)
	for name := range nameChannel {
		Printfln("Read name: %v", name)
	}

	Println("\nReceiving Recurring Notifications")
	nameChannel = make(chan string)
	go __writeToChannel(nameChannel)
	for name := range nameChannel {
		Printfln("Read name: %v", name)
	}
}
