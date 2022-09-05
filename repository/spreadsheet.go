package repository

import (
	"context"
	"fmt"
	"os"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type SpreadSheetRepository struct {
	srv           *sheets.Service
	spreadsheetID string
}

func NewSpreadSheetRepository() (*SpreadSheetRepository, error) {

	spreadsheetId := os.Getenv("SPREADSHEET_ID")
	
	// ref: https://ryomura.com/go-spreadsheets/
	ctx := context.Background()
	credential := option.WithCredentialsFile("credentials/secret.json")
	srv, err := sheets.NewService(ctx, credential)
	if err != nil {
		return nil, fmt.Errorf("failed to init sheet client: %w", err)
	}

	return &SpreadSheetRepository{srv: srv, spreadsheetID: spreadsheetId}, nil
}

func (s *SpreadSheetRepository) WriteStartTime(discordID string, t time.Time) error {

	r := discordID + "!A:A"
	rb := &sheets.ValueRange{
		Range:          r,
		MajorDimension: "ROWS",
		Values: [][]interface{}{
			{
				t.Format("2006/01/02 15:04:05"),
			},
		},
	}
	_, err := s.srv.Spreadsheets.Values.Append(s.spreadsheetID, r, rb).ValueInputOption("USER_ENTERED").InsertDataOption("OVERWRITE").Do()

	if err != nil {
		return fmt.Errorf("failed to append start time: %w", err)
	}
	return nil

}

func (s *SpreadSheetRepository) WriteEndTime(discordID string, t time.Time) error {

	r := discordID + "!C:C"
	rb := &sheets.ValueRange{
		Range:          r,
		MajorDimension: "ROWS",
		Values: [][]interface{}{
			{
				t.Format("2006/01/02 15:04:05"),
			},
		},
	}
	_, err := s.srv.Spreadsheets.Values.Append(s.spreadsheetID, r, rb).ValueInputOption("USER_ENTERED").InsertDataOption("OVERWRITE").Do()
	if err != nil {
		return fmt.Errorf("failed to append start time: %w", err)
	}

	return nil
}
