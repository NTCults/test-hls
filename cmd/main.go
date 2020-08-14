package main

import (
	"log"
	"net/http"
	"test-hls/api"
	"test-hls/config"
	"test-hls/services/access"
	"test-hls/services/permission"

	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	conf := config.NewConfig()
	logger := logrus.New()

	accessService := access.NewAccessService(conf)
	permissionService := permission.NewPermissionService(logger)

	api := api.NewAPI(accessService, permissionService, logger, conf)

	srv := &http.Server{
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		Handler:      api.Router,
		Addr:         ":" + conf.Port,
	}
	log.Printf("starting server on %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
