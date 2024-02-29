package util

import (
	"regexp"

	"github.com/OVillas/user-api/model"
)

func IsValidUUID(s string) error {
	uuidRegex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	if !uuidRegex.MatchString(s) {
		return model.ErrInvalidId
	}

	return nil
}
