package infofetchers

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v4/host"
)

func OSStats(ch chan any, wg *sync.WaitGroup) {
	wg.Add(2)
	go func() {
		defer wg.Done()
		bootTimeRaw, err := host.BootTime()
		if err != nil {
			log.Fatal("Fatal Error when Trying to Fetch BootTime")
		}

		bootTime := time.Unix(int64(bootTimeRaw), 0)
		ch <- bootTime
		fmt.Println("BootTime: ", bootTime)

	}()
	go func() {
		defer wg.Done()
		chUp := make(chan uint64, 1)
		chUpReply := make(chan string, 1)
		uptimeRaw, err := host.Uptime()

		if err != nil {
			log.Fatal("Fatal Error reached when attempting to fetch Uptime")
		}

		wgUp := sync.WaitGroup{}
		wgUp.Add(1)
		go func() {
			defer wgUp.Done()
			chUp <- uptimeRaw
			uptimeConv(chUp, chUpReply)
			defer close(chUpReply)
		}()
		wgUp.Wait()
		//ch <- chUpReply
		fmt.Println("Uptime: ", <-chUpReply)

	}()
	go func() {

	}()
}

func uptimeConv(chUp chan uint64, chUpReply chan string) {

	uptimeRaw := <-chUp
	hours := uptimeRaw / 3600
	minutes := uptimeRaw % 3600 / 60
	chUpReply <- fmt.Sprintf("Uptime: %d h %d m", hours, minutes)
}
