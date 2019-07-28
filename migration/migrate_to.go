package migration

import (
	"database/sql"
	"fmt"
)

func MigrateTo(timeString string, db *sql.DB, key string) (err error) {
	if len(points) == 0 {
		return fmt.Errorf("%v", "len points == 0")
	}

	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()

	var timestampTo int64

	for _, point := range points {
		point.SetDB(tx)
		point.SetKey(key)

		if point.GetTimeString() == timeString {
			timestampTo = point.GetTimestamp()
		}
	}

	if timestampTo == 0 && timeString != "default" {
		return fmt.Errorf("%v", "not found point with timeString:"+timeString)
	}

	currentTime, err := getCurrentTimestamp(db)
	if err != nil {
		return
	}

	if timestampTo == currentTime {
		return
	}

	points.Sort()

	if timestampTo > currentTime {
		err = up(timestampTo, currentTime)
	} else if timestampTo < currentTime {
		err = down(timestampTo, currentTime)
	}

	if err != nil {
		return
	}

	err = save(db, timestampTo, currentTime)
	if err != nil {
		return
	}

	_ = tx.Commit()
	return
}

func MigrateUp(db *sql.DB, key string) (err error) {
	if len(points) == 0 {
		return fmt.Errorf("%v", "len points == 0")
	}

	points.Sort()

	return MigrateTo(points[len(points)-1].GetTimeString(), db, key)
}

func MigrateDown(db *sql.DB, key string) (err error) {
	return MigrateTo("default", db, key)
}

func up(to int64, from int64) (err error) {
	var startI int

	if from == 0 {
		startI = 0
	} else {
		for i, point := range points {
			if point.GetTimestamp() == from {
				startI = i + 1
				break
			}
		}
	}

	for a := startI; a < len(points); a++ {
		p := points[a]

		if to < p.GetTimestamp() {
			return
		}

		err = p.Up()
		if err != nil {
			return
		}
	}

	return
}

func down(to int64, from int64) (err error) {
	for i, point := range points {
		if point.GetTimestamp() == from {
			for a := i; a >= 0; a-- {
				p := points[a]

				if to >= p.GetTimestamp() {
					return
				}

				err = p.Down()
				if err != nil {
					return
				}
			}

			return
		}
	}

	return
}

func getCurrentTimestamp(db *sql.DB) (timestamp int64, err error) {
	row, err := db.Query(`SELECT version FROM migrate_schema`)
	if err != nil {
		timestamp = 0

		_, err = db.Exec(`CREATE TABLE migrate_schema(version int)`)
		if err != nil {
			return
		}

		_, err = db.Exec(`INSERT INTO migrate_schema (version) VALUES (0);`)
		if err != nil {
			return
		}
	} else {
		defer row.Close()

		for row.Next() {
			err = row.Scan(&timestamp)
			if err != nil {
				return
			}
		}
	}

	return
}

func save(db *sql.DB, to, from int64) (err error) {
	_, err = db.Exec(`UPDATE migrate_schema SET version = $1 WHERE version = $2;`, to, from)
	return err
}
