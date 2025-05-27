package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"errors"
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

	type County struct {
		ID   int64
		Name string
	}
	type SubCounty struct {
		ID       int64
		Name     string
		CountyID int64
	}
	type Ward struct {
		ID          int64
		Name        string
		SubCountyID int64
	}

	counties := make(map[int64]County)
	subCounties := make(map[int64]SubCounty)
	wards := make(map[int64]Ward)

	for i, record := range records {
		if i == 0 {
			continue
		}
		countyID, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			log.Printf("Error parsing county id at row %d: %v", i, err)
			continue
		}
		countyName := strings.ToLower(record[1])
		if _, exists := counties[countyID]; !exists {
			counties[countyID] = County{ID: countyID, Name: countyName}
		}

		subCountyID, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			log.Printf("Error parsing sub-county id at row %d: %v", i, err)
			continue
		}
		subCountyName := strings.ToLower(record[3])
		if _, exists := subCounties[subCountyID]; !exists {
			subCounties[subCountyID] = SubCounty{ID: subCountyID, Name: subCountyName, CountyID: countyID}
		}

		wardID, err := strconv.ParseInt(record[4], 10, 64)
		if err != nil {
			log.Printf("Error parsing ward id at row %d: %v", i, err)
			continue
		}
		wardName := strings.ToLower(record[5])
		if _, exists := wards[wardID]; !exists {
			wards[wardID] = Ward{ID: wardID, Name: wardName, SubCountyID: subCountyID}
		}
	}

	for _, county := range counties {
		_, err := cfg.Db.GetCountyByGivenId(context.Background(), county.ID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				_, err := cfg.Db.CreateCounty(context.Background(), database.CreateCountyParams{
					ID:            uuid.New().String(),
					Name:          county.Name,
					CountyGivenID: county.ID,
				})
				if err != nil {
					log.Printf("Failed to save county %s: %v", county.Name, err)
				}
			} else {
				log.Printf("Failed to get county %d: %v", county.ID, err)
			}
		}
	}

	for _, subCounty := range subCounties {
		county, err := cfg.Db.GetCountyByGivenId(context.Background(), subCounty.CountyID)
		if err != nil {
			log.Printf("Failed to get county for sub-county %s: %v", subCounty.Name, err)
			continue
		}
		_, err = cfg.Db.GetSubCountyByGivenID(context.Background(), subCounty.ID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				_, err := cfg.Db.CreateSubCounty(context.Background(), database.CreateSubCountyParams{
					ID:               uuid.New().String(),
					Name:             subCounty.Name,
					CountyID:         county.ID,
					SubCountyGivenID: subCounty.ID,
				})
				if err != nil {
					log.Printf("Failed to save sub-county %s: %v", subCounty.Name, err)
				}
			} else {
				log.Printf("Failed to get sub-county %d: %v", subCounty.ID, err)
			}
		}
	}

	for _, ward := range wards {

		subCounty, err := cfg.Db.GetSubCountyByGivenID(context.Background(), ward.SubCountyID)
		if err != nil {
			log.Printf("Failed to get sub-county for ward %s: %v", ward.Name, err)
			continue
		}
		_, err = cfg.Db.GetWardByGivenID(context.Background(), ward.ID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				_, err := cfg.Db.CreateWard(context.Background(), database.CreateWardParams{
					ID:          uuid.New().String(),
					Name:        ward.Name,
					SubCountyID: subCounty.ID,
					WardGivenID: ward.ID,
				})
				if err != nil {
					log.Printf("Failed to save ward %s: %v", ward.Name, err)
				}
			} else {
				log.Printf("Failed to get ward %d: %v", ward.ID, err)
			}
		}
	}
}
