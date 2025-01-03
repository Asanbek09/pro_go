package main

import (
	"reflect"
	//"strings"
	//"fmt"
)

func inspectChannel(channel interface{}) {
	channelType := reflect.TypeOf(channel)
	if (channelType.Kind() == reflect.Chan) {
		Printfln("Type %v, Direction: %v", channelType.Elem(), channelType.ChanDir())
	}
}

func main() {
	var c chan<- string
	inspectChannel(c)
}