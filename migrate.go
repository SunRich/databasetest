package databasetest

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/mysql"
	_ "github.com/mattes/migrate/source/file"
)

func InitDatabase(dataSourceName, sourceUrl string)  {
	fmt.Println(dataSourceName)
	db, _ := sql.Open("mysql", dataSourceName)
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		sourceUrl,
		"mysql",
		driver,
	)
	defer m.Close()
	m.Down()
	m.Up()
}



func Up(dataSourceName, sourceUrl string) {
	db, _ := sql.Open("mysql", dataSourceName)
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		sourceUrl,
		"mysql",
		driver,
	)
	defer m.Close()
	m.Up()
}
