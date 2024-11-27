package main

import (
	"io"
	"strings"
	//"bufio"
	"fmt"
)

/*
func processData(reader io.Reader, writer io.Writer) {
	count, err := io.Copy(writer, reader)
	if (err == nil) {
		Printfln("Read %v bytes", count)
	} else {
		Printfln("Error: %v", err.Error())
	}
}
*/

	/*
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b);
		if (count > 0) {
			writer.Write(b[0:count])
			Printfln("Read: %v bytes: %v", count, string(b[0:count]))
		}
		if err == io.EOF {
			break
		}
	}
	*/

/*
func scanFromReader(reader io.Reader, template string, vals ...interface{}) (int, error) {
	return fmt.Fscanf(reader, template, vals...)
}

func scanSingle(reader io.Reader, val interface{}) (int, error) {
	return fmt.Fscan(reader, val)
}
*/

func writeFormatted(writer io.Writer, template string, vals ...interface{}) {
	fmt.Fprintf(writer, template, vals...)
}

func writeReplacer(writer io.Writer, str string, subs ...string) {
	replacer := strings.NewReplacer(subs...)
	replacer.WriteString(writer, str)
}

func main() {
	text := "It was a boat. A small boat"
	subs := []string {"boat", "kayak", "small", "huge"}

	var writer strings.Builder
	writeReplacer(&writer, text, subs...)
	fmt.Println(writer.String())
}

/*
	var writer strings.Builder
	template := "Name: %s, Category: %s, Price: $%.2f"

	writeFormatted(&writer, template, "Kayak", "water sports", float64(289))
	fmt.Println(writer.String())
*/

/*
reader := strings.NewReader("kayak watersports $284.90")

	for {
		var str string
		_, err := scanSingle(reader, &str)
		if (err != nil) {
			if (err != io.EOF) {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		Printfln("Value: %v", str)
	}
*/

/*
reader := strings.NewReader("kayak watersports $279.45")

	var name, category string
	var price float64
	scanTemplate := "%s %s $%f"

	_, err := scanFromReader(reader, scanTemplate, &name, &category, &price)
	if (err != nil) {
		Printfln("Error: %v", err.Error())
	} else {
		Printfln("Name: %v", name)
		Printfln("Category: %v", category)
		Printfln("Price: %.2f", price)
	}
*/

/*
text := "It was a boat. A small boat."

	var builder strings.Builder
	var writer = bufio.NewWriterSize(NewCustomWriter(&builder), 20)
	//var writer = NewCustomWriter(&builder)
	
	for i := 0; true; {
		end := i + 5
		if (end >= len(text)) {
			writer.Write([]byte(text[i:]))
			writer.Flush()
			break
		}
		writer.Write([]byte(text[i:end]))
		i = end
	}
	Printfln("Written data: %v", builder.String())
*/

/*
	text := "It was a boat. A small boat."
	var reader io.Reader = NewCustomReader(strings.NewReader(text))
	var writer strings.Builder
	slice := make([]byte, 5)

	//reader = bufio.NewReader(reader)
	buffered := bufio.NewReader(reader)
	
	for {
		//count, err := reader.Read(slice)
		count, err := buffered.Read(slice)
		if (count > 0) {
			Printfln("Buffer size: %v, buffered: %v", buffered.Size(), buffered.Buffered())
			writer.Write(slice[0:count])
		}
		if (err != nil) {
			break
		}
	}
	Printfln("Read data: %v", writer.String())
	*/

	/*
	r1 := strings.NewReader("kayak")
	r2 := strings.NewReader("jacket")
	r3 := strings.NewReader("canoe")

	concatReader := io.MultiReader(r1, r2, r3)
	limited := io.LimitReader(concatReader, 5)
	ConsumeData(limited)
	*/

	/*
	r1 := strings.NewReader("kayak")
	r2 := strings.NewReader("jacket")
	r3 := strings.NewReader("canoe")

	concatReader := io.MultiReader(r1, r2, r3)

	var writer strings.Builder
	teeReader := io.TeeReader(concatReader, &writer)
	ConsumeData(teeReader)
	Printfln("Echo data: %v", writer.String())
	*/

	/*
	var w1 strings.Builder
	var w2 strings.Builder
	var w3 strings.Builder

	combinedWriter := io.MultiWriter(&w1, &w2, &w3)

	GenerateData(combinedWriter)

	Printfln("Writer 1: %v", w1.String())
	Printfln("writer 2: %v", w2.String())
	Printfln("Writer 3: %v", w3.String())
	*/

	/*
	r1 := strings.NewReader("Kayak")
	r2 := strings.NewReader("Jacket")
	r3 := strings.NewReader("Canoe")

	concatReader := io.MultiReader(r1, r2, r3)
	ConsumeData(concatReader)
	*/

	/*
	pipeReader, pipeWriter := io.Pipe()
	go GenerateData(pipeWriter)
	ConsumeData(pipeReader)
	*/

	/*
	pipeReader, pipeWriter := io.Pipe()
	go func() {
		GenerateData(pipeWriter)
		pipeWriter.Close()
	}()
	ConsumeData(pipeReader)
	*/

	/*
	r := strings.NewReader("Kayak")
	var builder strings.Builder
	processData(r, &builder)
	Printfln("String builder contents: %s", builder.String())
	*/