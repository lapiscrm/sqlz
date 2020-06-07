package sqlite3driver_test

import (
	"testing"

	"github.com/lapiscrm/sqlz/sqlite3driver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateSqliteDBFile(t *testing.T) {
	db, err := sqlite3driver.CreateSqliteDBFile("test.db", true, nil)
	require.Nil(t, err)
	assert.NotNil(t, db)
}
