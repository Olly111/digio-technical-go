package concurrent

import (
	"baseTechnical/src/utils"
	"context"
)

func Worker(ctx context.Context, rowBatch <-chan []string) <-chan Processed {
	output := make(chan Processed)

	go func() {
		defer close(output)

		processed := Processed{
			UniqueAddresses: make(map[string]bool),
			IpActivity:      make(map[string]int),
			UrlVisits:       make(map[string]int),
		}

		for rowBatch := range rowBatch {
			for _, row := range rowBatch {
				ip, url := utils.ProcessRow(row)
				processed.UniqueAddresses[ip] = true
				processed.IpActivity[ip]++
				processed.UrlVisits[url]++
			}
		}

		output <- processed
	}()

	return output
}
