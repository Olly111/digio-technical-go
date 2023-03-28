package sequential

import (
	"baseTechnical/src/utils"
	"bufio"
	"log"
	"os"
)

func Sequential(filename string) (int, []utils.IP, []utils.URL) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	uniqueAddresses := make(map[string]bool)
	ipActivity := make(map[string]int)
	urlVisits := make(map[string]int)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := scanner.Text()

		ip, url := utils.ProcessRow(row)

		// Record unique IP
		uniqueAddresses[ip] = true

		// Record IP activity
		ipActivity[ip]++

		// Record URL visit
		urlVisits[url]++
	}

	return len(uniqueAddresses), utils.ProcessIPActivity(ipActivity), utils.ProcessURLVisits(urlVisits)
}
