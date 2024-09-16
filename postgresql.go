package coredns_postgresql

import (
	"database/sql"
	"fmt"
	clog "github.com/coredns/coredns/plugin/pkg/log"
	_ "github.com/lib/pq"
)

type Postgresql struct {
	user     string
	password string
	database string
}

func (handler Postgresql) connect() (*sql.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		handler.user, handler.password, handler.database)
	return sql.Open("postgres", dsn)
}

func (handler Postgresql) findRecordTypeA(zoneName string) ([]string, error) {
	conn, err := handler.connect()
	defer conn.Close()
	if err != nil {
		clog.Error("Failed to connect to database!!!")
		return nil, err
	}

	rows, err := conn.Query("SELECT content FROM records WHERE zone_name = $1", zoneName)
	defer rows.Close()
	if err != nil {
		clog.Error("Failed to query to database...")
		return nil, err
	}

	var ipAddresses []string
	for rows.Next() {
		var ipAddr string
		if err := rows.Scan(&ipAddr); err != nil {
			clog.Error("Failed to scan to record...")
			return nil, err
		}
		ipAddresses = append(ipAddresses, ipAddr)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return ipAddresses, nil
}
