# Base Technical Solution (Golang)

---

## Overview + Key Decisions

This project was created as an improved alternative to my original 2021 solution using Python + Regular Expressions for the purposes of the Digio Base Engineer Technical Assessment.

The decision to recreate the solution using Golang comes after participating in Song & Roman's Go Training in September 2022. Go presents a much faster way to process files like the ones provided in the `/data` folder in this repo, by utilising Concurrency. I came across a blog post outlining how to split lines of files into separate goroutines to be processed concurrently, then spliced back together to draw results from - this paradigm of Split -> Process -> Combine is what this concurrent solution is based on.

This repo also contains a solution that sequentially processes lines from a file - purely to outline the difference in the speed of processing between the two methods (Sequential vs. Concurrent).

The program accepts two commands line arguments to determine which method to use for file processing and which set of data to process. Under the `/data` folder there are two files, `example-data.log` is a modified version of the original log provided to be processed to more clearly see the results of the processing, the second is a file containing 400,000 logs to show the difference in processing times when the concurrent solution is given different batch sizes and number of workers that process the logs.

- 1st argument: `sequential | concurrent`
- 2nd argument: `short | extended`

My original 2021 solution was somewhat slow comparatively, had no unit tests, but could be put in the hands of a person working within the data domain and they would most likely know what the code did just by reading it. This is a tradeoff made to achieve significantly faster (by approximately 4.5x) results using a much more sophisticated solution.

Average time to process 400,000 logs:
Python (sequential): 0.8s
Go (sequential): 0.45s
Go (concurrent): 0.17s

Advantages vs. Original Python solution:

- Significantly faster execution time
- Unit tests + table tests, benchmarks, execution timing
- Concurrency
- IP and URL validation
- Sophisticated data structures for faster lookups (maps vs. simple arrays in Python)
- String manipulation to extract IP and URL from each line (quicker than Regular Expressions)
- Row validation to ensure there is data present to extract

Disadvantages vs. Original Python solution:

- Harder to read
- Longer development time

---

## Assumptions

- The data can contain erroneous logs
- Each log is on a separate line
- Olly is challenging himself as to how quickly he can make the program run ⚡️😂

---

## Setup Instructions

1. Clone the repo via HTTPS or SSH
2. Ensure Go 1.19 is installed
3. Open a terminal pointing to the root directory of this repo
4. To run the program:

- `go run src/main.go {processing method} {data selection}`
- replace `{processing method}` with `sequential | concurrent`
- replace `{data selection}` with `short | extended`
- e.g.
  - `go run src/main.go concurrent extended` - Uses concurrent processing on the log with 400,000 entries
  - `go run src/main.go sequential short` - Uses sequential processing on the log with 34 entries

5. Ta-da!

---

## Sequential Solution

The sequential solution is quite simple, it opens the file, reads the data line by line and records in separate maps the IP Address, increments that IP's activity count, and the URL's visit count, returning:

- The length of the unique IP map (number of unique addresses)
- The top 3 most active IP addresses
- The top 3 most visited URLs

---

## Concurrent Solution

The concurrent solution's execution is as follows:

1. The file is opened
2. The Reader function returns a channel of batched rows `chan []string` to be passed to workers
3. A number of Workers (specified at runtime by the user) are spawned and are used to populate a channel, eventually containing all of their respective processed results, and each read a batch of rows from the Reader
4. Each Worker processes it's batch of rows in a similar fashion to the sequential solution, only returning the output in a channel
5. The Combiner is started, taking in the channel of Worker's outputs, and are multiplexed (read through sequentially, combining the results into a single map for unique IPs, IP activity, and URL visits) until the Workers channel is empty
6. Each combined map is read through into the resulting maps incrementing IP activity & URL visits, and recording unique IP addresses
7. In `main.go` the resulting data are displayed in the console to the user

---
