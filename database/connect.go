package database

import (
	"database/sql"
	"fmt"

	"github.com/amaldevm19/go_matrix_tna/config"
	_ "github.com/microsoft/go-mssqldb"
)

// DB gorm connector
var (
	COSEC_DB     *sql.DB
	TNA_PROXY_DB *sql.DB
)

func ConnectDB() {
	var err error

	if err != nil {
		panic("failed to parse database port")
	}
	tna_proxy_dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), config.Config("DB_PORT"), config.Config("PROXY_DB_NAME"))

	TNA_PROXY_DB, err = sql.Open("mssql", tna_proxy_dsn)

	if err != nil {
		panic("failed to connect TNA_PROXY database")
	}

	fmt.Println("Connection Opened to TNA_PROXY Database")

	cosec_dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), config.Config("DB_PORT"), config.Config("TNA_DB_NAME"))

	COSEC_DB, err = sql.Open("mssql", cosec_dsn)

	if err != nil {
		panic("failed to connect COSEC database")
	}

	fmt.Println("Connection Opened to COSEC Database")

	//DB.AutoMigrate(&model.Mx_BranchMst{})
	//fmt.Println("Database Migrated")

}
