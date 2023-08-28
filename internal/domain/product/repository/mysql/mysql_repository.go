package mysql

import (
	"database/sql"
)

//go:generate mockgen -source=./mysql_repository.go -destination=../../../../test/mocks/mockgen/mysql_repository_mock.go -package=mockgen
type IDataBaseOperations interface {
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	Prepare(query string, args ...any) (sql.Result, error)
	Exec(query string, args ...any) (sql.Result, error)
}

type MysqlRepository struct {
	connect *sql.DB
}

func NewMysqlRepository(db *sql.DB) IDataBaseOperations {
	return &MysqlRepository{
		connect: db,
	}
}

func (m *MysqlRepository) Query(query string, args ...any) (*sql.Rows, error) {
	return m.connect.Query(query)
}

func (m *MysqlRepository) QueryRow(query string, args ...any) *sql.Row {
	return m.connect.QueryRow(query, args...)
}

func (m *MysqlRepository) Prepare(query string, args ...any) (sql.Result, error) {
	stmt, _ := m.connect.Prepare(query)

	defer stmt.Close()

	if len(args) != 0 {
		return stmt.Exec(args...)
	}
	return stmt.Exec()
}

func (m *MysqlRepository) Exec(query string, args ...any) (sql.Result, error) {
	if len(args) != 0 {
		return m.connect.Exec(query, args...)
	}
	return m.connect.Exec(query)
}
