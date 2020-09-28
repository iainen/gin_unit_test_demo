package apis

import (
	"database/sql"
	"demo/internal/router"
	"demo/tests/httpclient"
	"fmt"
	"time"

	"github.com/DATA-DOG/go-txdb"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	TxName = "mysqltx"
	TxDrv  = "mysql"
	TxDsn  = "root:root@tcp(127.0.0.1:3306)/gin-demo-test?charset=utf8mb4&parseTime=true&loc=Local"
)

func init() {
	txdb.Register(TxName, TxDrv, TxDsn)
}

type ApiSuite struct {
	suite.Suite
	DB         *gorm.DB
	sqlDB      *sql.DB
	httpClient *httpclient.Client
}

func (s *ApiSuite) SetupTest() {
	var err error
	cName := fmt.Sprintf("connection_%d", time.Now().UnixNano())
	s.sqlDB, err = sql.Open(TxName, cName)

	if err != nil {
		s.T().Fatalf("open mysqltx connection: %s", err)
	}

	s.DB, err = gorm.Open(TxDrv, s.sqlDB)
	if err != nil {
		s.T().Fatalf("mysql open mysqltx db: %s", err)
	}

	ro := router.InitRouter(s.DB)
	router.Init(ro)
	r := router.Get()

	s.httpClient = httpclient.NewClient(r)
}

func (s *ApiSuite) AfterTest(_, _ string) {
	var err error
	err = s.sqlDB.Close()
	if err != nil {
		s.T().Fatalf("sqldb close err: %s", err)
	}

	err = s.DB.Close()
	if err != nil {
		s.T().Fatalf("gormdb close err: %s", err)
	}
}
