package config

type DB struct {
	DSN   string
	Type  DBType
	Props map[string]string
}

type DBType string
