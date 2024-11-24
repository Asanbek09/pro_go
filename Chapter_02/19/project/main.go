package main

import (
	//"fmt"
	"time"
)

/*
func PrintTime(label string, t *time.Time) {
	//layout := "Day: 02 Month: Jan Year: 2006"
	fmt.Println(label, t.Format(time.RFC822Z))
	//Printfln("%s: Day: %v: Month: %v Year: %v", label, t.Day(), t.Month(), t.Year())
}
*/

func writeChannel(channel chan <- string) {
	names := []string {"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
		time.Sleep(time.Second * 1)
	}
	close(channel)
}

func main() {
	
	nameChannel := make (chan string)

	go writeChannel(nameChannel)
	for name := range nameChannel {
		Printfln("Read name: %v", name)
	}

	/*
	d, err := time.ParseDuration("1h30m")
	if( err == nil) {
		Printfln("Hours: %v", d.Hours())
		Printfln("Minutes: %v", d.Minutes())
		Printfln("Seconds: %v", d.Seconds())
		Printfln("MilliSeconds: %v", d.Milliseconds())
	} else {
		fmt.Println(err.Error())
	}
	*/

	/*
	toYears := func(d time.Duration) int {
		return int(d.Hours() / (24 * 365))
	}

	future := time.Date(2051, 0, 0, 0, 0, 0, 0, time.Local)
	past := time.Date(1965, 0, 0, 0, 0, 0, 0, time.Local)

	Printfln("Future: %v", toYears(time.Until(future)))
	Printfln("Past: %v", toYears(time.Since(past)))
	*/

	/*
	var d time.Duration = time.Hour + (30 * time.Minute) 

	Printfln("Hours: %v", d.Hours())
	Printfln("Mins: %v", d.Minutes())
	Printfln("Seconds: %v", d.Seconds())
	Printfln("Milliseconds: %v", d.Milliseconds())

	rounded := d.Round(time.Hour)
	Printfln("Rounded Hours: %v", rounded.Hours())
	Printfln("Rounded Mins: %v", rounded.Minutes())

	trunc := d.Truncate(time.Hour)

	Printfln("Truncated Hours: %v", trunc.Hours())
	Printfln("Rounded Mins: %v", trunc.Minutes())
	*/

	/*
	t1, _ := time.Parse(time.RFC822Z, "09 Jun 95 04:59 +0100")
	t2, _ := time.Parse(time.RFC822Z, "08 Jun 95 23:59 -0400")

	Printfln("Equal Method: %v", t1.Equal(t2))
	Printfln("Equality Operator: %v", t1 == t2)
	*/

	/*
	t, err := time.Parse(time.RFC822, "09 Jun 95 04:59 BST")
	if (err == nil) {
		Printfln("After: %v", t.After(time.Now()))
		Printfln("Round: %v", t.Round(time.Hour))
		Printfln("Truncate: %v", t.Truncate(time.Hour))
	} else {
		fmt.Println(err.Error())
	}
	*/

	/*
	layout := "02 Jan 06 15:04"
	date := "09 Jun 95 16:30"

	london := time.FixedZone("BST", 1 * 60 * 60)
	newyork := time.FixedZone("EDT", -4 * 60 * 60)
	local := time.FixedZone("Local", 0)

	nolocation, _ := time.Parse(layout, date)
	londonTime, _ := time.ParseInLocation(layout, date, london)
	newyorkTime, _ := time.ParseInLocation(layout, date, newyork)
	localTime, _ := time.ParseInLocation(layout, date, local)

	PrintTime("No location:", &nolocation)
	PrintTime("London:", &londonTime)
	PrintTime("NewYork:", &newyorkTime)
	PrintTime("Local:", &localTime)
	*/

	/*
	layout := "02 Jan 06 15:04"
	date := "09 Jun 95 19:30"

	london, lonerr := time.LoadLocation("Europe/London")
	newyork, nycerr := time.LoadLocation("America/New_York")
	local, _ := time.LoadLocation("Local")

	if (lonerr == nil && nycerr == nil) {
		nolocation, _ := time.Parse(layout, date)
		londonTime, _ := time.ParseInLocation(layout, date, london)
		newyorkTime, _ := time.ParseInLocation(layout, date, newyork)
		localTime, _ := time.ParseInLocation(layout, date, local)

		PrintTime("No location:", &nolocation)
		PrintTime("London:", &londonTime)
		PrintTime("New York:", &newyorkTime)
		PrintTime("Local:", &localTime)
	} else {
		fmt.Println(lonerr.Error(), nycerr.Error())
	}
	*/

	/*
	layout := "2006-Jan-02"
	dates :=[]string {
		"1995-Jun-09",
		"2015-Jun-02",
	}
	for _, d := range dates {
		time, err := time.Parse(layout, d)
		if (err == nil) {
			PrintTime("Parsed:", &time)
		} else {
			Printfln("Error: %s", err.Error())
		}
	}
	*/

	/*

	dates := []string {
		"09 Jun 95 00:00 GMT",
		"02 Jun 15 00:00 GMT",
	}
	for _, d := range dates {
		time, err := time.Parse(time.RFC822, d)
		if (err == nil) {
			PrintTime("Parsed", &time)
		} else {
			Printfln("Error: %s", err.Error())
		}
	}
	*/

	/*
	current := time.Now()
	specific := time.Date(1995, time.June, 9,0, 0, 0, 0, time.Local)
	unix := time.Unix(1433228090, 0)

	PrintTime("Current", &current)
	PrintTime("Specific", &specific)
	PrintTime("Unix", &unix)
	*/
}