package databasetest

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/mysql"
	_ "github.com/mattes/migrate/source/file"
)

func InitDatabase(dataSourceName, sourceUrl string)  {
	Up(dataSourceName,sourceUrl)
	Down(dataSourceName,sourceUrl)
}



func Up(dataSourceName, sourceUrl string) {
	db, _ := sql.Open("mysql", dataSourceName)
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		sourceUrl,
		"mysql",
		driver,
	)
	m.Close()
	m.Up()
}

func Down(dataSourceName, sourceUrl string) {
	db, _ := sql.Open("mysql", dataSourceName)
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		sourceUrl,
		"mysql",
		driver,
	)
	m.Close()
	m.Down()
}
