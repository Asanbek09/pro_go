package main

import (
	"io"
	"strings"
	"bufio"
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

func main() {

	text := "It was a boat. A small boat."
	var reader io.Reader = NewCustomReader(strings.NewReader(text))
	var writer strings.Builder
	slice := make([]byte, 5)

	reader = bufio.NewReader(reader)

	for {
		count, err := reader.Read(slice)
		if (count > 0) {
			writer.Write(slice[0:count])
		}
		if (err != nil) {
			break
		}
	}
	Printfln("Read data: %v", writer.String())
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
}