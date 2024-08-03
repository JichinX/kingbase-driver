//go:build go1.10
// +build go1.10

package kb_test

import (
	"database/sql"
	"fmt"

	"github.com/jichinx/pg-kb"
)

func ExampleNewConnector() {
	name := ""
	connector, err := kb.NewConnector(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	db := sql.OpenDB(connector)
	defer db.Close()

	// Use the DB
	txn, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	txn.Rollback()
}
