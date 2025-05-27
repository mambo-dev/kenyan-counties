package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	csvFile, err := os.Open("counties.csv")

	if err != nil {
		log.Fatal("Error opening counties.csv:", err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal("Error reading the csv:", err)
	}

	count := 0

	type County struct {
		CountyId   int64
		CountyName string
	}

	counties := make(map[int64]County, 0)

	for index, record := range records {
		if index == 0 {

			continue
		}
		count += 1
		parseCountyIdInt, err := strconv.ParseInt(record[0], 0, 64)

		if err != nil {
			log.Fatal("Error parsing a county id", record)
		}

		_, ok := counties[parseCountyIdInt]

		if !ok {
			counties[parseCountyIdInt] = County{
				CountyId:   parseCountyIdInt,
				CountyName: record[1],
			}
		}

	}

	type SubCounty struct {
		SubCountyId       int64
		SubCountyName     string
		SubCountyCountyId int64
	}
	subCounties := make(map[int64]SubCounty, 0)
	for index, record := range records {
		if index == 0 {

			continue
		}
		count += 1
		parseCountyIdInt, err := strconv.ParseInt(record[0], 0, 64)

		if err != nil {
			log.Fatal("Error parsing a county id", record)
		}

		parseSubCountyIdInt, err := strconv.ParseInt(record[2], 0, 64)

		if err != nil {
			log.Fatal("Error parsing a sub county id", record)
		}

		_, ok := subCounties[parseSubCountyIdInt]

		if !ok {
			subCounties[parseSubCountyIdInt] = SubCounty{
				SubCountyCountyId: parseSubCountyIdInt,
				SubCountyId:       parseCountyIdInt,
				SubCountyName:     record[3],
			}
		}

	}

	type Ward struct {
		SubCountyId int64
		WardName    string
		WardId      int64
	}

	wards := make(map[int64]Ward, 0)
	for index, record := range records {
		if index == 0 {

			continue
		}
		count += 1
		parseWardIdInt, err := strconv.ParseInt(record[4], 0, 64)

		if err != nil {
			log.Fatal("Error parsing a ward id", record)
		}

		parseSubCountyIdInt, err := strconv.ParseInt(record[2], 0, 64)

		if err != nil {
			log.Fatal("Error parsing a sub county id", record)
		}

		_, ok := wards[parseWardIdInt]

		if !ok {
			wards[parseWardIdInt] = Ward{
				SubCountyId: parseSubCountyIdInt,
				WardId:      parseWardIdInt,
				WardName:    record[5],
			}
		}

	}

	fmt.Println(wards)
}
