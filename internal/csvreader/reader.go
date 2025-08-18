// Package csvreader provides functions for reading the CSV file and generating
// recipient structs.
package csvreader

import (
	"encoding/csv"
	"errors"
	"io"
	"strconv"

	"github.com/valentino7504/tax-automation-go/internal/models"
)

// ReadCSV takes a Reader and returns a slice of maps with keys corresponding
// to the headers and values the necessary data provided.
func ReadCSV(r io.Reader) ([]map[string]string, error) {
	reader := csv.NewReader(r)
	header, err := reader.Read()
	if err != nil {
		return nil, errors.New("unable to read CSV file")
	}
	var records []map[string]string
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.New("failed to read row ")
		}
		if len(row) != len(header) {
			return nil, errors.New("incomplete row")
		}
		recordMap := make(map[string]string)
		for i := range row {
			if i < len(header) {
				recordMap[header[i]] = row[i]
			}
		}
		records = append(records, recordMap)
	}
	return records, nil
}

// GenerateRecipients takes in a slice of maps corresponding to records from the
// csv file and parses them to generate a slice of Recipients.
func GenerateRecipients(records []map[string]string) ([]models.Recipient, error) {
	var recipients []models.Recipient
	for _, record := range records {
		amount, err := strconv.ParseFloat(record["Amount"], 64)
		if err != nil {
			return nil, errors.New("error processing amount")
		}
		recipient := models.Recipient{
			Name:           record["Name"],
			Email:          record["Email"],
			TINBeneficiary: record["TINBeneficiary"],
			TINPayer:       record["TINPayer"],
			Amount:         amount,
			Month:          record["Month"],
			SerialNo:       record["SerialNo"],
			Date:           record["Date"],
		}
		recipients = append(recipients, recipient)
	}
	return recipients, nil
}
