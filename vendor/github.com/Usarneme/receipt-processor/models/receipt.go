package models

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Receipt struct {
	ID           string `json:"id"`
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

func (r *Receipt) Validate() error {
	if r.Retailer == "" {
		return errors.New("retailer is required")
	}
	if r.PurchaseDate == "" {
		return errors.New("purchaseDate is required")
	}
	if r.PurchaseTime == "" {
		return errors.New("purchaseTime is required")
	}
	if len(r.Items) == 0 {
		return errors.New("items is required")
	}
	if r.Total == "" {
		return errors.New("total is required")
	}
	return nil
}

func (r *Receipt) CalculatePoints() int64 {
	points := int64(0)

	// One point for every alphanumeric character in the retailer name
	alnumRegex := regexp.MustCompile(`[a-zA-Z0-9]`)
	points += int64(len(alnumRegex.FindAllString(r.Retailer, -1)))

	// 50 points if the total is a round dollar amount with no cents
	total, _ := strconv.ParseFloat(r.Total, 64)
	if total == math.Floor(total) {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25
	if total*100 == math.Floor(total*100) && int(total*100)%25 == 0 {
		points += 25
	}

	// 5 points for every two items on the receipt
	points += int64((len(r.Items) / 2) * 5)

	// Points based on item description length
	for _, item := range r.Items {
		description := strings.TrimSpace(item.ShortDescription)
		if len(description)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int64(math.Ceil(price * 0.2))
		}
	}

	// 6 points if the day in the purchase date is odd
	purchaseDate, _ := time.Parse("2006-01-02", r.PurchaseDate)
	if purchaseDate.Day()%2 != 0 {
		points += 6
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm
	purchaseTime, _ := time.Parse("15:04", r.PurchaseTime)
	if purchaseTime.Hour() == 14 {
		points += 10
	}

	return points
}
