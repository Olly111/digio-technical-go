package main

import (
	"baseTechnical/src/concurrent"
	"baseTechnical/src/sequential"
	"fmt"
	"time"
)

func main() {
	// uniqueAddresses, ipActivity, urlVisits := sequential.Sequential(`test.log`)

	// fmt.Println("No. unique IP addresses:", uniqueAddresses)
	// fmt.Println("IP Activity:")
	// fmt.Println("1:", ipActivity[0].Count, ipActivity[0].Ip)
	// fmt.Println("2:", ipActivity[1].Count, ipActivity[1].Ip)
	// fmt.Println("3:", ipActivity[2].Count, ipActivity[2].Ip)
	// fmt.Println("URL Visits:")
	// fmt.Println("1:", urlVisits[0].Count, urlVisits[0].URL)
	// fmt.Println("2:", urlVisits[1].Count, urlVisits[1].URL)
	// fmt.Println("3:", urlVisits[2].Count, urlVisits[2].URL)

	uniqueAddresses, ipActivity, urlVisits := concurrent.Concurrent(`data/example-data.log`, 5, 5)

	fmt.Println("No. unique IP addresses:", uniqueAddresses)
	fmt.Println("IP Activity:")
	fmt.Println("1:", ipActivity[0].Count, ipActivity[0].Ip)
	fmt.Println("2:", ipActivity[1].Count, ipActivity[1].Ip)
	fmt.Println("3:", ipActivity[2].Count, ipActivity[2].Ip)
	fmt.Println("URL Visits:")
	fmt.Println("1:", urlVisits[0].Count, urlVisits[0].URL)
	fmt.Println("2:", urlVisits[1].Count, urlVisits[1].URL)
	fmt.Println("3:", urlVisits[2].Count, urlVisits[2].URL)

	reps := 1
	start := time.Now()

	for i := 0; i < reps; i++ {
		sequential.Sequential(`data/example-data-extended.log`)
	}

	elapsed := time.Since(start)
	fmt.Printf("Time taken for %d repetitions (Sequential): %s\n", reps, elapsed)

	start = time.Now()
	for i := 0; i < reps; i++ {
		concurrent.Concurrent(`data/example-data.log`, 5, 5)
	}

	elapsed = time.Since(start)
	fmt.Printf("Time taken for %d repetitions (Concurrent): %s\n", reps, elapsed)
}
