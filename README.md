# gomigrate

Minimalist library for migration db

##### Install #####
```go get github.com/darkfoxs96/gomigrate```

#### Supported commands ####
* Help: ```gomigrate -help```
* New point: ```gomigrate ./ init_db```   
Need 2 args. Path to package and name record.

###### Example: ######
* Run: ```gomigrate ./migrates create_user_table```
* Generated file: ```./migrates/2019.02.12_07:22:59_create_table_user.go```
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
    migration.Register(1549930979, "2019.02.12_07:22:59", m)
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
* This timeString => ```"2019.02.12_07:22:59"```
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
    migration.MigrateTo("2019.02.12_07:22:59", db) // This timeString => "2019.02.12_07:22:59"
}
```
* ```gomigrate``` checks the position now and if it is not equal to ```"2019.02.12_07:22:59"``` migrates to it 

#### How to work? ####
* ```gomigrate```: generated in database TABLE ```migration_schema``` and saved now timestamp point last migration
* If ```"2019.02.12_07:22:59"``` == timestamp last migration, then all unchanged
