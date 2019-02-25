# gomigrate

Minimalist library for migration db

##### Install #####
```go get github.com/darkfoxs96/gomigrate```

#### Supported commands ####
* Help: ```gomigrate -help```
* New point: ```gomigrate ./ init_db```   
Need 2 args. Path to package and name record.   
* Build to bin: ```gomigrate -build ./ GOOS=darwin GOARCH=amd64 -up postgres user=don password=sdef12 dbname=don host=0.0.0.0 port=5432 sslmode=disable```   
Need min 6 args. Path to package, Systems params, -up or -down or 'data', Connect params.

###### Example create migrate point: ######
* Run: ```gomigrate ./migrates create_user_table```
* Generated file: ```./migrates/20190212_072259_create_table_user.go```
* File body: 
```
package migrates

import (
    "github.com/darkfoxs96/gomigrate/migration"
)

// DO NOT MODIFY
type Create_table_user_20190212_072259 struct {
    migration.Point
}

// DO NOT MODIFY
func init() {
    m := &Create_table_user_20190212_072259{}
    migration.Register(1549930979, "20190212_072259", m)
}

// Run the migrations
func (m *Create_table_user_20190212_072259) Up() (err error) {
    _, err = m.DB.Exec(``)
    return
}

// Reverse the migrations
func (m *Create_table_user_20190212_072259) Down() (err error) {
    _, err = m.DB.Exec(``)
    return
}
```
* This timeString => ```"20190212_072259"```
* Edit functions Up() and Down():
```
// Run the migrations
func (m *Create_table_user_20190212_072259) Up() (err error) {
    _, err = m.DB.Exec(`CREATE TABLE user(name char, age int);`)
    return
}

// Reverse the migrations
func (m *Create_table_user_20190212_072259) Down() (err error) {
    _, err = m.DB.Exec(`DROP TABLE user;`)
    return
}
```
* Added use MigrationTo()
```
package db

import (
    "database/sql"

    "github.com/darkfoxs96/gomigrate/migration"

    _ "testproject/migrates"
)

func Migrate(db *sql.DB) {
    migration.MigrateTo("20190212_072259", db) // This timeString => "20190212_072259"
    // or
    migration.MigrateUp(db) // up max point
    // or
    migration.MigrateDown(db) // down to default DB
}
```
* ```gomigrate``` checks the position now and if it is not equal to ```"20190212_072259"``` migrates to it 

###### Example build to bin by migrate points: ######
* Run: ```gomigrate -build ./migrates GOOS=windows GOARCH=amd64 -up postgres user=don password=sdef12 dbname=don host=0.0.0.0 port=5432 sslmode=disable```
* Generated file: ```./migrates/20190223_201930_up/20190223_201930_up.exe```

#### How to work? ####
* ```gomigrate```: generated in database TABLE ```migration_schema``` and saved now timestamp point last migration
* If ```"20190212_072259"``` == timestamp last migration, then all unchanged
