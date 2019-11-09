package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"
)

func getQuote(sym string) string {
	sym = strings.ToUpper(sym)
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_WEEKLY&symbol=%s&apikey=YZCGP9H4ARUKAE8K&datatype=csv", sym)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	rows, err := csv.NewReader(resp.Body).ReadAll()
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	if len(rows) >= 1 && len(rows[0]) == 6 {
		return fmt.Sprintf("%s is trading at Open: $%s, High: $%s, Low: $%s, Close: $%s", sym, rows[1][1], rows[1][2], rows[1][3], rows[1][4])
	}
	return fmt.Sprintf("unknown response format (symbol was \"%s\")", sym)
}
