package migration

import (
	"database/sql"
	"sort"
)

type Point struct {
	DB         *sql.Tx
	TimeString string
	Timestamp  int64
	Key        string
}

type IPoint interface {
	Up() error
	Down() error
	SetDB(db *sql.Tx)
	SetKey(key string)
	SetTime(timestamp int64, timeString string)
	GetTimeString() string
	GetTimestamp() int64
}

func (m *Point) Up() (err error) {
	return
}

func (m *Point) Down() (err error) {
	return
}

func (m *Point) SetDB(db *sql.Tx) {
	m.DB = db
}

func (m *Point) SetKey(key string) {
	m.Key = key
}

func (m *Point) SetTime(timestamp int64, timeString string) {
	m.Timestamp = timestamp
	m.TimeString = timeString
}

func (m *Point) GetTimestamp() int64 {
	return m.Timestamp
}

func (m *Point) GetTimeString() string {
	return m.TimeString
}

type PointsArray []IPoint

func (p *PointsArray) Sort() {
	sort.Sort(*p)
}
func (p PointsArray) Len() int           { return len(p) }
func (p PointsArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PointsArray) Less(i, j int) bool { return p[i].GetTimestamp() < p[j].GetTimestamp() }

var points = make(PointsArray, 0)

func Register(timestamp int64, timeString string, point IPoint) {
	point.SetTime(timestamp, timeString)

	points = append(points, point)
}
