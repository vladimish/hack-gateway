package handlers

import (
	"bytes"
	"github.com/vladimish/hack-gateway/internal/cfg"
	"github.com/vladimish/hack-gateway/internal/db"
	"github.com/vladimish/hack-gateway/internal/requests"
	"net/http"
	"strconv"
)

func HandleDeleteTable(req requests.DeleteTable) error {
	port, err := db.GetDB().GetRestaurant(req.Login)
	if err != nil {
		return err
	}

	b := []byte(req.Name)

	_, err = http.Post("http://"+cfg.GetConfig().K8SAddr+":"+strconv.Itoa(port)+"/delete_table", "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	return nil
}
