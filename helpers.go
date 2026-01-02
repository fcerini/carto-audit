package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func getString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}
