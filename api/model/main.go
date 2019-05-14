package model 

import (
	"github.com/jmoiron/sqlx"
)

var (
	currentDB *sqlx.DB
)

// SetDB set databse
func SetDB(db *sqlx.DB){
	currentDB = db 
}