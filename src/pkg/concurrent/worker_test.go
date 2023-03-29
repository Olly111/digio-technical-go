package concurrent

import (
	"context"
	"reflect"
	"testing"
)

func TestWorker(t *testing.T) {
	ctx := context.Background()

	input := make(chan []string)

	output := worker(ctx, input)

	testRows := []string{
		`79.125.111.21 - - [10/Jul/2018:20:03:40 +0200] "GET /newsletter/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/5.0)`,
		`221.112.111.11 - admin [11/Jul/2018:17:31:05 +0200] "GET /docs/manage-websites/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6"`,
		`72.44.32.10 - - [09/Jul/2018:15:48:20 +0200] "GET /docs/manage-websites/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86; en-US) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"`,
	}

	go func() {
		defer close(input)
		input <- testRows
	}()

	result := <-output

	expected := Processed{
		UniqueAddresses: map[string]bool{"79.125.111.21": true, "221.112.111.11": true, "72.44.32.10": true},
		IpActivity:      map[string]int{"79.125.111.21": 1, "221.112.111.11": 1, "72.44.32.10": 1},
		UrlVisits:       map[string]int{"/newsletter/": 1, "/docs/manage-websites/": 2},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Unexpected result: %v (expected %v)", result, expected)
	}
}
