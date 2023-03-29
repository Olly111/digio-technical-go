package concurrent

import (
	"context"
	"sync"
)

func combiner(ctx context.Context, inputs ...<-chan Processed) <-chan Processed {
	output := make(chan Processed)

	var waitgroup sync.WaitGroup

	// Read each workers input into the output channel - call Done() when finished
	multiplexer := func(processed <-chan Processed) {
		defer waitgroup.Done()

		for input := range processed {
			select {
			case <-ctx.Done():
			case output <- input:
			}
		}
	}

	// Tell the waitgroup to wait until every worker's data has been added
	waitgroup.Add(len(inputs))
	for _, input := range inputs {
		go multiplexer(input)
	}

	// Close the output channel when the waitgroup closes
	go func() {
		waitgroup.Wait()
		close(output)
	}()

	return output
}
