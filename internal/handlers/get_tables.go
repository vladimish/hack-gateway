package handlers

import (
	"github.com/vladimish/hack-gateway/internal/cfg"
	"github.com/vladimish/hack-gateway/internal/db"
	"github.com/vladimish/hack-gateway/internal/requests"
	"io/ioutil"
	"net/http"
	"strconv"
)

func HandleGetTables(req requests.GetTables) ([]byte, error) {
	port, err := db.GetDB().GetRestaurant(req.Login)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("http://"+cfg.GetConfig().K8SAddr+":"+strconv.Itoa(port)+"/get_tables", "application/json", nil)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}
