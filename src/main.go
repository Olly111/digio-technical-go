package main

import (
	"baseTechnical/src/concurrent"
	"baseTechnical/src/sequential"
	"baseTechnical/src/utils"
	"fmt"
	"os"
	"time"
)

func displayResults(uniqueAddresses int, ipActivity []utils.IP, urlVisits []utils.URL) {
	fmt.Println("No. unique IP addresses:", uniqueAddresses)
	fmt.Println("IP Activity:")
	for i := range ipActivity {
		fmt.Printf("%v: %v, %v\n", (i + 1), ipActivity[i].Count, ipActivity[i].Ip)
	}
	fmt.Println("URL Visits:")
	for i := range ipActivity {
		fmt.Printf("%v: %v, %v\n", (i + 1), urlVisits[i].Count, urlVisits[i].URL)
	}
}

func getFilename(selectedData string) string {
	if selectedData == "short" {
		return `data/example-data.log`
	}
	return `data/example-data-extended.log`
}

func getOptimalBatchAndWorkerSize(selectedData string) (int, int) {
	if selectedData == "short" {
		return 5, 5
	}
	return 10000, 40
}

func main() {
	timerStart := time.Now()

	args := os.Args

	if len(args) < 3 {
		fmt.Println("Please give command line arguments for demo purposes")
		fmt.Println("First argument: 'sequential' | 'concurrent'")
		fmt.Println("Second argument: 'short' | 'extended'")
		os.Exit(1)
	}

	selectedProcess := args[1]
	selectedData := args[2]

	localDataFile := getFilename(selectedData)
	batchSize, numWorkers := getOptimalBatchAndWorkerSize(selectedData)

	switch selectedProcess {
	case "sequential":
		uniqueAddresses, ipActivity, urlVisits := sequential.Sequential(localDataFile)
		displayResults(uniqueAddresses, ipActivity, urlVisits)
	case "concurrent":
		uniqueAddresses, ipActivity, urlVisits := concurrent.Concurrent(localDataFile, batchSize, numWorkers)
		displayResults(uniqueAddresses, ipActivity, urlVisits)
	}

	elapsed := time.Since(timerStart)
	fmt.Printf("Execution time: %s\n", elapsed)
}
