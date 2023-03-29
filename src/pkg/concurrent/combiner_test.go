package concurrent

import (
	"context"
	"reflect"
	"testing"
)

func TestCombiner(t *testing.T) {
	ctx := context.Background()

	workerInput1 := make(chan Processed)
	workerInput2 := make(chan Processed)

	go func() {
		defer close(workerInput1)

		workerInput1 <- Processed{
			UniqueAddresses: map[string]bool{"1.1.1.1": true, "2.2.2.2": true},
			IpActivity:      map[string]int{"1.1.1.1": 2, "2.2.2.2": 2},
			UrlVisits:       map[string]int{"/urlOne": 3, "urlTwo": 1},
		}
	}()

	go func() {
		defer close(workerInput2)

		workerInput2 <- Processed{
			UniqueAddresses: map[string]bool{"1.1.1.1": true, "2.2.2.2": true},
			IpActivity:      map[string]int{"1.1.1.1": 1, "2.2.2.2": 4},
			UrlVisits:       map[string]int{"/urlOne": 2, "urlTwo": 3},
		}
	}()

	output := combiner(ctx, workerInput1, workerInput2)

	uniqueAddresses := make(map[string]bool)
	ipActivity := make(map[string]int)
	urlVisits := make(map[string]int)

	for processed := range output {
		for ip := range processed.UniqueAddresses {
			uniqueAddresses[ip] = true
		}

		for ip, visitCount := range processed.IpActivity {
			ipActivity[ip] += visitCount
		}

		for url, visitCount := range processed.UrlVisits {
			urlVisits[url] += visitCount
		}
	}

	expectedUniqueIpAddresses := map[string]bool{"1.1.1.1": true, "2.2.2.2": true}
	expectedIpActivity := map[string]int{"1.1.1.1": 3, "2.2.2.2": 6}
	expectedUrlVisits := map[string]int{"/urlOne": 5, "urlTwo": 4}

	if !reflect.DeepEqual(uniqueAddresses, expectedUniqueIpAddresses) {
		t.Errorf("Unexpected result: %v (expected %v)", uniqueAddresses, expectedUniqueIpAddresses)
	}
	if !reflect.DeepEqual(ipActivity, expectedIpActivity) {
		t.Errorf("Unexpected result: %v (expected %v)", ipActivity, expectedIpActivity)
	}
	if !reflect.DeepEqual(urlVisits, expectedUrlVisits) {
		t.Errorf("Unexpected result: %v (expected %v)", urlVisits, expectedUrlVisits)
	}
}
