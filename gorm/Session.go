package main

import (
	"context"
	"fmt"
	"time"
		"github.com/sharmamanan796/gorm/Connection"

	"github.com/sharmamanan796/gorm/Entity"

)



// Session Configuration
type Session struct {
  DryRun                 bool
  PrepareStmt            bool
  NewDB                  bool
  SkipHooks              bool
  SkipDefaultTransaction bool
  AllowGlobalUpdate      bool
  FullSaveAssociations   bool
  Context                context.Context
  Logger                 logger.Interface
  NowFunc                func() time.Time
}


func main() {

	db, err := Connection.Conn()
	var employees []Entity.Employee
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Conneted")
// session mode
stmt := db.Session(&Session{DryRun: true}).First(&employees, 1).Statement
stmt.SQL.String() //=> SELECT * FROM `users` WHERE `id` = $1 ORDER BY `id`
stmt.Vars         //=> []interface{}{1}

// globally mode with DryRun

// different databases generate different SQL
stmt := db.Find(&employees, 1).Statement
stmt.SQL.String() //=> SELECT * FROM `users` WHERE `id` = $1 // PostgreSQL
stmt.SQL.String() //=> SELECT * FROM `users` WHERE `id` = ?  // MySQL
stmt.Vars         //=> []interface{}{1}
}