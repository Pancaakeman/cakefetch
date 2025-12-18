package infofetchers

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/shirou/gopsutil/v4/host"
)

func OsInfo(ch chan string, wg *sync.WaitGroup) {

	wg.Add(2)
	go func() {
		defer wg.Done()
		kernVer, err := host.KernelVersion()
		if err != nil {
			log.Fatal("Fatal Error when fetching Kernel Version")
		}
		//ch <- kernVer
		fmt.Println("Kernel Version", kernVer)

	}()

	go func() {
		defer wg.Done()
		platform, _, _, err := host.PlatformInformationWithContext(context.Background())

		if err != nil {
			log.Fatal("Reached Fatal Error when fetching Platform Info")
		}
		ch <- platform
		fmt.Println("Platform: ", platform)

	}()

}
