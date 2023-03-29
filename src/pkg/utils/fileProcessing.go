package utils

import (
	"net/url"
	"sort"
	"strings"

	"github.com/muonsoft/validation/validate"
)

type IP struct {
	Ip    string
	Count int
}

type URL struct {
	URL   string
	Count int
}

func ProcessRow(line string) (string, string) {
	rowData := strings.Fields(line)

	// Check rowData is long enough to access the required fields, data is erroneous
	// otherwise so we return empty strings which will fail IP and URL validation
	if len(rowData) < 7 {
		return "", ""
	}

	return rowData[0], rowData[6]
}

// Returns the 3 most frequently occurring IP Addresses
func ProcessIPActivity(ipActivity map[string]int) []IP {

	mostActiveIPs := []IP{}

	for ip, count := range ipActivity {
		if ValidateIPAddress(ip) {
			mostActiveIPs = append(mostActiveIPs, IP{ip, count})
		}
	}

	sort.Slice(mostActiveIPs, func(i, j int) bool {
		return mostActiveIPs[i].Count > mostActiveIPs[j].Count
	})

	return mostActiveIPs[:3]
}

// Returns the 3 most frequently occurring URLs
func ProcessURLVisits(urlVisits map[string]int) []URL {
	mostVisitedURLs := []URL{}

	for url, count := range urlVisits {
		if ValidateURL(url) {
			mostVisitedURLs = append(mostVisitedURLs, URL{url, count})
		}
	}

	sort.Slice(mostVisitedURLs, func(i, j int) bool {
		return mostVisitedURLs[i].Count > mostVisitedURLs[j].Count
	})

	return mostVisitedURLs[:3]
}

func ValidateURL(urlString string) bool {
	_, err := url.ParseRequestURI(urlString)

	return err == nil
}

func ValidateIPAddress(ipAddress string) bool {
	err := validate.IP(ipAddress)

	return err == nil
}
