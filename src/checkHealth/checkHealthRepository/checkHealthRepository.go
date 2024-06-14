package checkHealthRepository

import (
	"clean-code-unit-test/src/checkHealth"
	"database/sql"
)

// construct for repository
type checkHealthRepository struct {
	db *sql.DB
}

// construct for repository
func NewCheckHealthRepository(db *sql.DB) checkHealth.CheckHealthRepository {
	return &checkHealthRepository{db}
}

// RetrieveVersion is a function to retrieve version
func (c *checkHealthRepository) RetrieveVersion() (string, error) {
	//get data from table version_app
	version, query := "", "SELECT version FROM version_app"
	if err := c.db.QueryRow(query).Scan(&version); err != nil {
		return "", err
	}

	return version, nil
}
