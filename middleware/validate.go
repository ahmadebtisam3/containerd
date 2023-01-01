package middleware

import (
	"github.com/pkg/errors"
)

func ValidateSourcePath(path string) error {
	if path == "" || !CanVerifyVolumes() {
		return nil
	}
	validationErr, err := Call("chart.release.validate_host_source_path", path)
	if err == nil && validationErr != nil {
		return errors.Errorf(validationErr.(string))
	}
	return nil
}
