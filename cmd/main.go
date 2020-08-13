package main

import (
	"log"
	"net/http"
	"test-hls/api"
	"test-hls/config"
	"test-hls/services/access"
	"test-hls/services/permission"
	"test-hls/utils"

	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	conf := config.NewConfig()
	logger := logrus.New()

	accesService := access.NewAccessService(conf)
	permissionService := permission.NewPermissionService(logger)

	api := api.NewAPI(accesService, permissionService, logger, conf)
	router := mux.NewRouter()

	router.HandleFunc(conf.AccessKeyUrl, api.AccessKeyHandler).
		Methods(http.MethodGet)

	router.HandleFunc("/generateKey/{eventID}", api.GenerateKeyHandler).
		Methods(http.MethodGet)
	router.Use(utils.LoggingMiddleware)

	srv := &http.Server{
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		Handler:      router,
		Addr:         ":" + conf.Port,
	}
	log.Fatal(srv.ListenAndServe())
}
