package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	// gorm is a popular Object-Relational Mapping (ORM) library,
	// used for interacting with relational databases.
	DBConn *gorm.DB
)
