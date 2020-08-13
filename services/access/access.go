package access

import (
	"test-hls/config"
	"test-hls/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type AccessService interface {
	GetAccessKey(eventID string, keyID string) (string, error)
	GenerateToken(eventID string) (string, error)
}

type accessService struct {
	config *config.Config
}

func NewAccessService(conf *config.Config) AccessService {
	return &accessService{
		config: conf,
	}
}

func (s *accessService) GetAccessKey(eventID string, keyID string) (string, error) {
	accessKeyClaims := jwt.MapClaims{
		"eventID": eventID,
		"keyID":   keyID,
	}

	return utils.GenerateToken(accessKeyClaims, s.config.GetSecretKeyBytes())
}

func (s *accessService) GenerateToken(eventID string) (string, error) {
	uuid := uuid.New()
	tokenClaims := jwt.MapClaims{
		"eventID": eventID,
		"keyID":   uuid.String(),
	}

	return utils.GenerateToken(tokenClaims, s.config.GetSecretKeyBytes())
}
