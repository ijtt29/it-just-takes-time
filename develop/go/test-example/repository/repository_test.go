package repository_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/my/testcode/model"
	"github.com/my/testcode/repository"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	DB         *gorm.DB
	mock       sqlmock.Sqlmock
	db         *sql.DB
	err        error
	repository repository.Repository
}

func (s *Suite) SetupSuite() {
	s.db, s.mock, s.err = sqlmock.New()

	require.NoError(s.T(), s.err)

	s.DB, s.err = gorm.Open("sqlite3", s.db)

	require.NoError(s.T(), s.err)
	s.repository = repository.NewRepo(s.DB)
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *Suite) Test_repository_Create() {
	var (
		id   = uuid.NewV4()
		name = "test-name"
	)

	const query = `INSERT INTO "books" ("id","name") VALUES (?,?)`

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(id, name).
		WillReturnResult(sqlmock.NewResult(0, 1))
	s.mock.ExpectCommit()
	err := s.repository.Create(id, name)
	require.NoError(s.T(), err)
}

func (s *Suite) Test_repository_Get() {
	var (
		id   = uuid.NewV4()
		name = "test-name"
	)
	const query = `SELECT * FROM "books" WHERE (id = $1)`

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(id, name)

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(id).
		WillReturnRows(rows)

	res, err := s.repository.Get(id)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&model.Book{ID: id, Name: name}, res))
	defer s.DB.Close()
}
