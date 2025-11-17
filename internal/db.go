package internal

import (
	"database/sql"
	"log"
	"strings"
	"time"
)

func ConnectToDB(dsn string) *sql.DB {
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	if err := conn.Ping(); err != nil {
		log.Fatalf("DB ping failed: %v", err)
	}
	return conn
}

func InsertEvent(db *sql.DB, event string) error {
	now := time.Now().Unix()
	_, err := db.Exec(
		"INSERT INTO events (event, ts) VALUES ($1, $2)",
		event, now,
	)
	return err
}

func GetEventCounts(db *sql.DB, start, end *int64) (map[string]int, error) {
	var queryParts []string
	queryParts = append(queryParts, "SELECT event, COUNT(*) FROM events")
	args := []any{}

	if start != nil && end != nil {
		queryParts = append(queryParts, "WHERE ts >= $1 AND ts <= $2")
		args = append(args, *start, *end)
	} else if start != nil {
		queryParts = append(queryParts, "WHERE ts >= $1")
		args = append(args, *start)
	} else if end != nil {
		queryParts = append(queryParts, "WHERE ts <= $1")
		args = append(args, *end)
	}

	queryParts = append(queryParts, "GROUP BY event")
	query := strings.Join(queryParts, " ")

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	counts := make(map[string]int)
	for rows.Next() {
		var event string
		var count int
		if err := rows.Scan(&event, &count); err != nil {
			return nil, err
		}
		counts[event] = count
	}
	return counts, nil
}
