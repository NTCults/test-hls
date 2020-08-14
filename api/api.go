package api

import (
	"net/http"
	"test-hls/config"
	"test-hls/services/access"
	"test-hls/services/permission"
	"test-hls/utils"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Api struct {
	accessService     access.AccessService
	permissionService permission.PermissionService
	log               logrus.FieldLogger
	config            *config.Config

	Router *mux.Router
}

func NewAPI(a access.AccessService, p permission.PermissionService, logger logrus.FieldLogger, conf *config.Config) *Api {
	api := &Api{
		accessService:     a,
		permissionService: p,
		log:               logger,
		config:            conf,
	}

	router := mux.NewRouter()
	router.HandleFunc(conf.AccessKeyURL, api.AccessKeyHandler).
		Methods(http.MethodGet)

	router.HandleFunc("/generateKey/{eventID}", api.GenerateKeyHandler).
		Methods(http.MethodGet)
	router.Use(utils.LoggingMiddleware)

	api.Router = router
	return api
}

func (a *Api) GenerateKeyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventID, ok := vars["eventID"]
	if !ok {
		http.Error(w, http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}

	token, err := a.accessService.GenerateToken(eventID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token))
}

func (a *Api) AccessKeyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventID := vars["eventID"]
	clientID := vars["clientID"]
	keyName := vars["keyName"]

	tokenClaims, err := utils.ParseToken(keyName, a.config.GetSecretKeyBytes())
	if err != nil {
		a.log.WithError(err).Error()
		http.Error(w, http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}

	ok, err := a.permissionService.GetPermission(clientID)
	if err != nil {
		a.log.WithError(err).Error()
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

	if !ok {
		http.Error(w, http.StatusText(http.StatusForbidden),
			http.StatusForbidden)
		return
	}

	keyIDRaw := tokenClaims["keyID"]
	keyID, ok := keyIDRaw.(string)
	accessKeyString, err := a.accessService.GetAccessKey(eventID, keyID)
	if err != nil {
		a.log.WithError(err).Error()
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

	w.Write([]byte(accessKeyString))
}
