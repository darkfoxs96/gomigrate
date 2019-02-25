package maintemplate

const templateSource = `package ${packageName}

import (
	"database/sql"
	"fmt"

	"github.com/darkfoxs96/gomigrate/migration"
	_ "${driverImport}"
)

func main() {
	fmt.Println("DB Connecting...")

	db, err := sql.Open("${driverName}", "${connectParams}")
	if err != nil {
		fmt.Println("Open Error:")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Ping Error:")
		panic(err)
	}

	fmt.Println("DB Connected.")
	fmt.Println("Start migration...")

	err = migration.${migrateFunc}
	if err != nil {
		fmt.Println("Migration Error:")
		panic(err)
	}

	fmt.Println("Done migration.")
}
`
