package main

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/mambo-dev/kenya-locations/config"
	"github.com/mambo-dev/kenya-locations/internal/database"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := database.Connect(cfg.DBURL, cfg.TAuthToken)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	cfg.Db = database.New(db)
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
				CountyName: strings.ToLower(record[1]),
			}
		}

	}

	for _, county := range counties {
		_, err := cfg.Db.CreateCounty(context.Background(), database.CreateCountyParams{
			ID:            uuid.New().String(),
			Name:          county.CountyName,
			CountyGivenID: county.CountyId,
		})

		if err != nil {
			log.Println("Failed to save county", county.CountyName)
			continue
		}
	}

	type SubCounty struct {
		SubCountyId       int64
		SubCountyName     string
		SubCountyCountyId string
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

		county, err := cfg.Db.GetCountyByGivenId(context.Background(), parseCountyIdInt)

		if err != nil {
			log.Fatal("Failed to get a given county id ", err)
			continue
		}

		parseSubCountyIdInt, err := strconv.ParseInt(record[2], 0, 64)

		if err != nil {
			log.Fatal("Error parsing a sub county id", record)
		}

		_, ok := subCounties[parseSubCountyIdInt]

		if !ok {
			subCounties[parseSubCountyIdInt] = SubCounty{
				SubCountyCountyId: county.ID,
				SubCountyId:       parseCountyIdInt,
				SubCountyName:     strings.ToLower(record[3]),
			}
		}

	}

	for _, subCounty := range subCounties {
		_, err := cfg.Db.CreateSubCounty(context.Background(), database.CreateSubCountyParams{
			ID:               uuid.New().String(),
			Name:             subCounty.SubCountyName,
			CountyID:         subCounty.SubCountyCountyId,
			SubCountyGivenID: subCounty.SubCountyId,
		})

		if err != nil {
			log.Println("Failed to save sub county", subCounty.SubCountyName)
			continue
		}
	}

	type Ward struct {
		SubCountyId string
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

		if err != nil {
			log.Fatal("Failed to parse a given sub county id ", err)
			continue
		}

		subCounty, err := cfg.Db.GetSubCountyByGivenID(context.Background(), parseSubCountyIdInt)

		if err != nil {
			log.Fatal("Failed toget a  sub county id ", err)
			continue
		}

		_, ok := wards[parseWardIdInt]

		if !ok {
			wards[parseWardIdInt] = Ward{
				SubCountyId: subCounty.ID,
				WardId:      parseWardIdInt,
				WardName:    strings.ToLower(record[5]),
			}
		}

	}

	for _, ward := range wards {
		_, err := cfg.Db.CreateWard(context.Background(), database.CreateWardParams{
			ID:          uuid.New().String(),
			Name:        ward.WardName,
			SubCountyID: ward.SubCountyId,
			WardGivenID: ward.WardId,
		})

		if err != nil {
			log.Println("Failed to save ward", ward.WardName)
			continue
		}
	}

}
