package concurrent

import (
	"bufio"
	"context"
	"os"
)

func Reader(ctx context.Context, rowsBatch *[]string, file *os.File, batchSize int) <-chan []string {
	output := make(chan []string)

	scanner := bufio.NewScanner(file)

	go func() {
		defer close(output)

		for {
			scanned := scanner.Scan()

			select {
			case <-ctx.Done():
				return
			default:
				row := scanner.Text()

				// Read rows into output channel when batchSize is reached or EOF and set rowsBatch to empty
				if len(*rowsBatch) == batchSize || !scanned {
					output <- *rowsBatch
					*rowsBatch = []string{}
				}
				*rowsBatch = append(*rowsBatch, row)
			}

			// Last batch will have been added, exit goroutine and return output
			if !scanned {
				return
			}
		}
	}()

	return output
}
