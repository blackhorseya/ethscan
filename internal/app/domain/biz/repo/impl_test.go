package repo

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteTester struct {
	suite.Suite
	logger *zap.Logger
	rw     sqlmock.Sqlmock
	repo   IRepo
}

func (s *suiteTester) SetupTest() {
	s.logger, _ = zap.NewDevelopment()

	db, rw, _ := sqlmock.New(sqlmock.MonitorPingsOption(true))
	s.rw = rw

	opts := &NodeOptions{BaseURL: "http://localhost", Timeout: 5 * time.Second}
	s.repo, _ = CreateRepo(opts, sqlx.NewDb(db, "mysql"))
}

func (s *suiteTester) assertExpectation(t *testing.T) {
	if err := s.rw.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}
