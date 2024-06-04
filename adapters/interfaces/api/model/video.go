package model

import (
	"net/url"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type UploadVideoRequest struct {
	Qualities []int32 `json:"-" validate:"dive,required"`
}

func (u *UploadVideoRequest) validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func (u *UploadVideoRequest) ParseMultipleQualitiesQuery(values url.Values) error {
	if qualities, ok := values["quality"]; ok {
		for _, strQuality := range qualities {
			quality, err := strconv.Atoi(strQuality)
			if err != nil {
				return err
			}
			u.Qualities = append(u.Qualities, int32(quality))
		}
	}
	return u.validate()
}
