package middleware

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func ValidateSourcePath(path string) error {
	logrus.Errorf("ix-mountv vv Validation inside validat %s", path)
	if path == "" || !CanVerifyVolumes() {
		return nil
	}
	logrus.Errorf("ix-mountc befor calling path %s", path)
	validationErr, _ := Call("chart.release.validate_host_source_path", path)
	logrus.Errorf("ix-mounts get results *************************")
	logrus.Errorf("ix-mounts validation error normal error : %s ", validationErr.(string))
	if validationErr != nil {
		logrus.Errorf("ix-mounte inside validation error ")
		return errors.Errorf(validationErr.(string))
	}
	return nil
}
