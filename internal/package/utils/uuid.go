package utils

import (
	"github.com/satori/go.uuid"
)

// GenerateUUID4 create a UUID
func GenerateUUID4() string {
	u := uuid.Must(uuid.NewV4(), nil)
	return u.String()
}
