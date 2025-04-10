package database

type Driver interface {
	Open(connectionString string) error
	Close() error
	Query(query string, args ...interface{}) ([]map[string]interface{}, error)
	Exec(query string, args ...interface{}) error
}