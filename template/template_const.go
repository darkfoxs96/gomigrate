package template

const templateSource = `package ${packageName}

import (
	"github.com/darkfoxs96/gomigrate/migration"
)

// DO NOT MODIFY
type ${structName} struct {
	migration.Point
}

// DO NOT MODIFY
func init() {
	m := &${structName}{}
	migration.Register(${timestamp}, "${timeString}", m)
}

// Run the migrations
func (m *${structName}) Up() (err error) {
	_, err = m.DB.Exec(`+"``"+`)
	return
}

// Reverse the migrations
func (m *${structName}) Down() (err error) {
	_, err = m.DB.Exec(`+"``"+`)
	return
}
`
