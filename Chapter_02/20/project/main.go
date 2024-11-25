package main

import (
	"io"
	//"strings"
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
	pipeReader, pipeWriter := io.Pipe()
	go func() {
		GenerateData(pipeWriter)
		pipeWriter.Close()
	}()
	ConsumeData(pipeReader)

	/*
	r := strings.NewReader("Kayak")
	var builder strings.Builder
	processData(r, &builder)
	Printfln("String builder contents: %s", builder.String())
	*/
}