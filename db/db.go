package db

import (
	"log"

	"github.com/dgraph-io/badger"
)

var db *badger.DB

func init() {
	var err error
	db, err = badger.Open(badger.DefaultOptions("/tmp/badger"))
	if err != nil {
		log.Fatal(err)
	}
}
