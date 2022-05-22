package models

type Venture struct {
	VentureID int64   `json:"ventureid"`
	Name      string  `json:"name"`
	Domain    string  `json:"domain"`
	Revenue   float64 `json:"revenue_estimation"`
}
