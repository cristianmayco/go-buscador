package models

type GlobalSearchElastic struct {
	Filter  string   `json:"filter"`
	Limit   int      `json:"limit"`
	Page    int      `json:"page"`
	Modules []string `json:"modules"`
}
