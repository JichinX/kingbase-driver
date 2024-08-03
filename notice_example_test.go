//go:build go1.10
// +build go1.10

package kb_test

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jichinx/pg-kb"
)

func ExampleConnectorWithNoticeHandler() {
	name := ""
	// Base connector to wrap
	base, err := kb.NewConnector(name)
	if err != nil {
		log.Fatal(err)
	}
	// Wrap the connector to simply print out the message
	connector := kb.ConnectorWithNoticeHandler(base, func(notice *kb.Error) {
		fmt.Println("Notice sent: " + notice.Message)
	})
	db := sql.OpenDB(connector)
	defer db.Close()
	// Raise a notice
	sql := "DO language plpgsql $$ BEGIN RAISE NOTICE 'test notice'; END $$"
	if _, err := db.Exec(sql); err != nil {
		log.Fatal(err)
	}
	// Output:
	// Notice sent: test notice
}
