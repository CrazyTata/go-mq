package util

import (
	"go-mq/common/xerr"

	"github.com/pkg/errors"
)

func ReturnError(code uint32) error {
	return errors.Wrapf(xerr.NewErrCode(code), "")
}

func ReturnErrorWithFormat(code uint32, format string) error {
	return errors.Wrapf(xerr.NewErrCode(code), format)
}
