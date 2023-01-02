package middleware

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func ValidateSourcePath(path string) error {
	logrus.Errorf("ix-mountv validation inside validat %s", path)
	if path == "" || !CanVerifyVolumes() {
		return nil
	}
	validationErr, err := Call("chart.release.validate_host_source_path", path)
	if err == nil && validationErr != nil {
		logrus.Errorf("ix-mounte inside validation error %s", validationErr.(string))
		return errors.Errorf(validationErr.(string))
	}
	return nil
}
