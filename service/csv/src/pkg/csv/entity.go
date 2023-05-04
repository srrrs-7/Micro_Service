package csv

type csvRequest struct {
	Headers []string   `json:"headers"`
	Values  [][]string `json:"values"`
}
