package concurrent

import (
	"baseTechnical/src/pkg/utils"
	"context"
)

func worker(ctx context.Context, rowBatch <-chan []string) <-chan Processed {
	output := make(chan Processed)

	go func() {
		defer close(output)

		processed := Processed{
			UniqueAddresses: make(map[string]bool),
			IpActivity:      make(map[string]int),
			UrlVisits:       make(map[string]int),
		}

		// Process rows in batch
		for rowBatch := range rowBatch {
			for _, row := range rowBatch {
				ip, url := utils.ProcessRow(row)
				processed.UniqueAddresses[ip] = true
				processed.IpActivity[ip]++
				processed.UrlVisits[url]++
			}
		}

		// Read processed data into output
		output <- processed
	}()

	return output
}
