package checkHealthRepository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestRetrieveVersion(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	// Create a new instance of the repository with the mock database
	repo := NewCheckHealthRepository(db)

	// Mocking the SQL query and expected result
	expectedVersion := "1.0.0"
	rows := sqlmock.NewRows([]string{"version"}).AddRow(expectedVersion)
	mock.ExpectQuery("SELECT (.+) FROM version_app").WillReturnRows(rows)

	// Call the function under test
	version, err := repo.RetrieveVersion()

	// Check if there are any errors
	assert.NoError(t, err)

	// Check if the returned version is as expected
	assert.Equal(t, expectedVersion, version)

	// Verify that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRetrieveVersion_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error creating mock database")
	}
	defer db.Close()

	repository := NewCheckHealthRepository(db)
	mock.ExpectQuery("SELECT (.+) FROM version_app WHERE id = 1 LIMIT 1;").WillReturnError(sql.ErrNoRows)

	version, err := repository.RetrieveVersion()
	assert.Error(t, err)
	assert.Empty(t, version)
}
