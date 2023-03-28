package concurrent

import (
	"baseTechnical/src/utils"
	"context"
	"log"
	"os"
)

type Processed struct {
	UniqueAddresses map[string]bool
	IpActivity      map[string]int
	UrlVisits       map[string]int
}

func Concurrent(filename string, batchSize int, numWorkers int) (int, []utils.IP, []utils.URL) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Resulting maps with combined data
	uniqueAddresses := make(map[string]bool)
	ipActivity := make(map[string]int)
	urlVisits := make(map[string]int)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rowsBatch := []string{}
	rowsChannel := Reader(ctx, &rowsBatch, file, batchSize)

	// Spawn workers
	workersChannel := make([]<-chan Processed, numWorkers)
	for i := 0; i < numWorkers; i++ {
		// Give each worker a batch of rows
		workersChannel[i] = Worker(ctx, rowsChannel)
	}

	// Populating result maps
	for Processed := range Combiner(ctx, workersChannel...) {
		for ip := range Processed.UniqueAddresses {
			uniqueAddresses[ip] = true
		}

		for ip, visitCount := range Processed.IpActivity {
			ipActivity[ip] += visitCount
		}

		for url, visitCount := range Processed.UrlVisits {
			urlVisits[url] += visitCount
		}
	}

	return len(uniqueAddresses), utils.ProcessIPActivity(ipActivity), utils.ProcessURLVisits(urlVisits)
}
