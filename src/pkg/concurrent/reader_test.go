package concurrent

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReaderOutput(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Assert that Reader should split rows from file into batches for processing
	testFile, err := os.Open("../../testData/reader_test.log")
	if err != nil {
		t.Fatal("Error opening file", err)
	}
	defer testFile.Close()

	splitRowsChannel := reader(ctx, &[]string{}, testFile, 2)
	expectedSplitRows := [][]string{{"line1", "line2"}, {"line3", "line4"}, {"line5", "line6"}, {"line7"}}

	for i := range expectedSplitRows {
		assert.Equal(t, expectedSplitRows[i], <-splitRowsChannel)
	}

	// Assert that Reader should close when ctx cancels
	closedContext := reader(ctx, &[]string{}, testFile, 2)
	cancel()
	_, ok := <-closedContext
	assert.False(t, ok)
}
