package db

import (
	"errors"
	"fmt"
)

func (db *DB) CreateRestaurant(name string, login string) (string, error) {
	squery := "DELETE FROM gw.`keys` WHERE tg_key IN (SELECT tg_key FROM gw.`keys` LIMIT 1);"
	var key string
	err := db.db.QueryRow(squery).Scan(&key)
	if err != nil {
		return "", err
	}
	if len(key) == 0 {
		return "", errors.New("keys ended")
	}
	iquery := fmt.Sprintf("INSERT INTO gw.restaurants (name, login, tg_key) VALUES ('%s', '%s', '%s');", name, login, key)
	_, err = db.db.Exec(iquery)
	if err != nil {
		return "", err
	}
	return key, nil
}
