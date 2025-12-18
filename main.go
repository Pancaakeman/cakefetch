package main

import (
	"fmt"
	"sync"
	"time"

	infofetchers "github.com/Pancaakeman/cakefetch/infoFetchers"
)

func main() {
	startTime := time.Now()
	chOsInfo := make(chan string, 2)
	chTime := make(chan any, 2)

	wg := sync.WaitGroup{}

	infofetchers.OsInfo(chOsInfo, &wg)
	infofetchers.OSStats(chTime, &wg)
	ChannelRead(chOsInfo, chTime, &wg)

	wg.Wait()

	defer close(chOsInfo)
	defer fmt.Println("Time taken to execute:", time.Since(startTime))

}

func ChannelRead(chOsString chan string, chTime chan any, wg *sync.WaitGroup) {
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i != cap(chOsString); i++ {
			fmt.Println(<-chOsString)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i != cap(chTime); i++ {
			fmt.Println(<-chTime)
		}
	}()
}
