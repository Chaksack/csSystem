package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	// "gorm.io/gorm"
)

func TestConnectDb(t *testing.T) {
	// Save the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	// Change the working directory to a temporary directory
	tempDir := t.TempDir()
	os.Chdir(tempDir)

	defer func() {
		// Restore the working directory
		os.Chdir(currentDir)
	}()

	// Call the ConnectDb function
	ConnectDb()

	// Check if the database instance is not nil
	assert.NotNil(t, Database.Db, "Database connection should not be nil")

	// Perform a test query to check the database connection
	err = Database.Db.Exec("SELECT 1").Error
	assert.NoError(t, err, "Error executing test query: %v")

	// Perform additional tests or assertions based on your requirements
	// For example, you can check if the necessary tables are created
	// or if migrations are successful.
}

func TestMain(m *testing.M) {
	// Run the tests with exit code
	os.Exit(m.Run())
}
