package config

type Config struct {
	Languages      []string `json:"languages"`
	ActiveLanguage string   `json:"activeLanguage"`
}
