package uploadfilemodel

import (
	"errors"
	"github.com/foody-go-api/common"
)

const EntityName = "UploadFile"

type Upload struct {
	common.BaseEntity `json:",inline"`
	common.Image    `json:",inline"`
}

func (Upload) TableName() string {
	return "UploadFiles"
}

var (
	ErrFileTooLarge = common.NewCustomError(
		errors.New("file too large"),
		"file too large",
		"ErrFileTooLarge",
	)
)

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"file is not image",
		"ErrFileIsNotImage",
	)
}

func ErrCannotSaveFile(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"cannot save uploaded file",
		"ErrCannotSaveFile",
	)
}
