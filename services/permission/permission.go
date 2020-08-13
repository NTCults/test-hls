package permission

import "github.com/sirupsen/logrus"

type PermissionService interface {
	GetPermission(clientID string) (bool, error)
}

func NewPermissionService(logger logrus.FieldLogger) PermissionService {
	return &permissionService{
		log: logger,
	}
}

type permissionService struct {
	log logrus.FieldLogger
}

func (s *permissionService) GetPermission(clientID string) (bool, error) {
	s.log.WithField("clientID", clientID).Info("permission granted")
	return true, nil
}
