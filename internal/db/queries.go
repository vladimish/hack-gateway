package db

import (
	"errors"
	"fmt"
)

func (db *DB) CreateRestaurant(name string, login string) (string, int32, error) {
	squery := "SELECT id, tg_key FROM gw.`keys`;"
	var key string
	var id int
	rows, err := db.db.Query(squery)
	if err != nil {
		return "", 0, err
	}
	for rows.Next() {
		rows.Scan(&id, &key)
		break
	}
	rows.Close()

	fmt.Println(key)
	dquery := fmt.Sprintf("DELETE FROM gw.`keys` WHERE id=%d;", id)
	_, err = db.db.Exec(dquery)
	if err != nil {
		return "", 0, err
	}

	if len(key) == 0 {
		return "", 0, errors.New("keys ended")
	}
	iquery := fmt.Sprintf("INSERT INTO gw.restaurants (name, login, tg_key) VALUES ('%s', '%s', '%s');", name, login, key)
	_, err = db.db.Exec(iquery)
	if err != nil {
		return "", 0, err
	}
	return key, int32(id), nil
}
