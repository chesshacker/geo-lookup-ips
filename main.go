package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/ip2location/ip2location-go/v9"
)

// reads stdin one ip per line, writes csv to stdout
// depends on geolocation data available for free download from here
// https://lite.ip2location.com/file-download
func main() {
	db, err := ip2location.OpenDB("./IP2LOCATION-LITE-DB11.BIN")
	defer db.Close()
	if err != nil {
		log.Fatalln("Error opening database.", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	records := [][]string{{
		"ip",
		"country_short",
		"country_long",
		"region",
		"city",
		"latitude",
		"longitude",
	}}

	for scanner.Scan() {
		ip := scanner.Text()

		results, err := db.Get_all(ip)
		if err != nil {
			log.Fatalln("Error looking up ip:", ip, err)
		}
		records = append(records, []string{
			ip,
			results.Country_short,
			results.Country_long,
			results.Region,
			results.City,
			fmt.Sprintf("%f", results.Latitude),
			fmt.Sprintf("%f", results.Longitude),
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("Error reading stdin.", err)
	}

	w := csv.NewWriter(os.Stdout)
	w.WriteAll(records) // calls Flush internally

	if err := w.Error(); err != nil {
		log.Fatalln("Error writing csv.", err)
	}

}
