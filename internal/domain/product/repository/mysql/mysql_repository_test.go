package mysql

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestNewMysqlRepository(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockMysqlRepository := NewMysqlRepository(db)

	assert.NotNil(t, mockMysqlRepository)
}

func TestMysqlRepository_Query(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	queryMock := `SELECT * FROM products`

	mock.ExpectQuery(regexp.QuoteMeta(queryMock)).
		WillReturnRows(
			sqlmock.NewRows([]string{"name", "price", "count"}).AddRow("Test 1", 15.95, 3))

	mysqlMockQuery := NewMysqlRepository(db)

	rows, err := mysqlMockQuery.Query(queryMock)
	assert.NotNil(t, rows)
	assert.NoError(t, err)
}
