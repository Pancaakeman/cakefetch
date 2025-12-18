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
	}()

	go func() {
		defer wg.Done()
		chUpSend := make(chan uint64)
		uptimeRaw, err := host.Uptime()

		if err != nil {
			log.Fatal("Fatal Error reached when attempting to fetch Uptime")
		}
		chUpSend <- uptimeRaw

		wg.Add(1)
		go func() {
			defer wg.Done()
			uptimeConv(chUpSend, wg)
		}()

		wg.Wait()
		defer close(chUpSend)
		ch <- uptimeRaw
	}()
}

func uptimeConv(chUpSend chan uint64, wg *sync.WaitGroup) {
	//chUpReciv := make(chan uint64, 2)

	hours := <-chUpSend / 3600
	minutes := <-chUpSend % 3600 / 60

	fmt.Println(hours, minutes)
	return

}
