package handlers

import (
	"github.com/vladimish/hack-gateway/internal/db"
	"github.com/vladimish/hack-gateway/internal/k8s"
	"github.com/vladimish/hack-gateway/internal/requests"
)

func HandleAddRestaurant(req requests.AddRestaurant) error {
	key, err := db.GetDB().CreateRestaurant(req.Name, req.Login)
	if err != nil {
		return err
	}

	k8s.AddBot(key, req.Login)

	return nil
}
