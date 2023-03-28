package concurrent

import (
	"context"
	"sync"
)

func Combiner(ctx context.Context, inputs ...<-chan Processed) <-chan Processed {
	output := make(chan Processed)

	var waitgroup sync.WaitGroup

	multiplexer := func(processed <-chan Processed) {
		defer waitgroup.Done()

		for input := range processed {
			select {
			case <-ctx.Done():
			case output <- input:
			}
		}
	}

	waitgroup.Add(len(inputs))
	for _, input := range inputs {
		go multiplexer(input)
	}

	go func() {
		waitgroup.Wait()
		close(output)
	}()

	return output
}
