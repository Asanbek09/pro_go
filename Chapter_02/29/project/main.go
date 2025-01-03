package main

import (
	"reflect"
	//"strings"
	//"fmt"
)

func createChannelAndSend(data interface{}) interface{} {
	dataVal := reflect.ValueOf(data)
	channelType := reflect.ChanOf(reflect.BothDir, dataVal.Type().Elem())
	channel := reflect.MakeChan(channelType, 1)
	go func() {
		for i := 0; i < dataVal.Len(); i++ {
			channel.Send(dataVal.Index(i))
		}
		channel.Close()
	}()
	return channel.Interface()
}

func readChannels(channels ...interface{}) {
	channelsVal := reflect.ValueOf(channels)
	cases := []reflect.SelectCase {}
	for i := 0; i < channelsVal.Len(); i++ {
		cases = append(cases, reflect.SelectCase{
			Chan: channelsVal.Index(i).Elem(),
			Dir: reflect.SelectRecv,
		})
	}

	for {
		caseIndex, val, ok := reflect.Select(cases)
		if (ok) {
			Printfln("value read: %v, type: %v", val, val.Type())
		} else {
			if len(cases) == 1 {
				Printfln("all channels closed")
				return
			}
			cases = append(cases[:caseIndex], cases[caseIndex+1:]... )
		}
	}
}

func main() {
	values := []string {"Alice", "Bob", "Charlie", "Dora"}
	channel := createChannelAndSend(values).(chan string)

	cities := []string {"London", "Berlin", "Paris"}
	cityChannel := createChannelAndSend(cities).(chan string)

	prices := []float64 {567, 345.89, 12.54}
	priceChannel := createChannelAndSend(prices).(chan float64)

	readChannels(channel, cityChannel, priceChannel)
}