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

	err = migrateFromV1(tx)
	if err != nil {
		return
	}

	currentTime, err := getCurrentTimestamps(tx)
	if err != nil {
		return
	}

	points.Sort()

	if timestampTo >= currentTime {
		err = up(timestampTo, tx)
	} else if timestampTo < currentTime {
		err = down(timestampTo, currentTime, tx)
	}
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

func up(to int64, db *sql.Tx) (err error) {
	for _, p := range points {
		pointTimestamp := p.GetTimestamp()

		if to < pointTimestamp || currentMigratedMap[pointTimestamp] {
			return
		}

		err = p.Up()
		if err != nil {
			return
		}

		err = add(db, pointTimestamp)
		if err != nil {
			return
		}
	}

	return
}

func down(to int64, from int64, db *sql.Tx) (err error) {
	for i, point := range points {
		if point.GetTimestamp() == from {
			for a := i; a >= 0; a-- {
				p := points[a]
				pointTimestamp := p.GetTimestamp()

				if to >= p.GetTimestamp() || !currentMigratedMap[pointTimestamp] {
					return
				}

				err = p.Down()
				if err != nil {
					return
				}

				err = remove(db, pointTimestamp)
				if err != nil {
					return
				}
			}

			return
		}
	}

	return
}

func migrateFromV1(db *sql.Tx) (err error) {
	var timestamp int64
	err = db.QueryRow(`SELECT version FROM migrate_schema`).Scan(&timestamp)
	if err != nil {
		return nil
	}

	_, err = db.Exec(`DROP TABLE migrate_schema`)
	if err != nil {
		return
	}

	for _, point := range points {
		if point.GetTimestamp() <= timestamp {
			db.Exec(`INSERT INTO migrate_schema_v2 (id) VALUES ($1)`, point.GetTimestamp())
		}
	}

	return
}

func getCurrentTimestamps(db *sql.Tx) (timestampMax int64, err error) {
	row, err := db.Query(`SELECT id FROM migrate_schema_v2`)
	if err != nil {
		_, err = db.Exec(`CREATE TABLE migrate_schema_v2(id int)`)
		return
	}
	defer row.Close()

	for row.Next() {
		var timestamp int64
		err = row.Scan(&timestamp)
		if err != nil {
			return
		}

		if timestamp > timestampMax {
			timestampMax = timestamp
		}

		currentMigratedMap[timestamp] = true
	}

	return
}

func add(db *sql.Tx, id int64) (err error) {
	_, err = db.Exec(`INSERT INTO migrate_schema_v2 (id) VALUES ($1)`, id)
	return err
}

func remove(db *sql.Tx, id int64) (err error) {
	_, err = db.Exec(`DELETE migrate_schema_v2 WHERE id=$1`, id)
	return err
}
