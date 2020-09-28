package model

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	useri UserI
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("sqlite3", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.useri = NewUser(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_Create() {
	var (
		name     = "zhangyi"
		password = "zhangyi1"
		email    = "zzz"
	)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(`INSERT INTO "user" (.+)`).
		WithArgs(sqlmock.AnyArg(), name, password, email)
	s.mock.ExpectCommit()

	_, err := s.useri.Create(name, password, email)
	require.NoError(s.T(), err)
}
