package models

import "fmt"

type GetUserField string

const (
	IDGetUserField      GetUserField = "id"
	NameGetUserField    GetUserField = "name"
	PictureGetUserField GetUserField = "picture"
)

type PictureAttribute string

const (
	WidthPictureAttribute  PictureAttribute = "width"
	HeightPictureAttribute PictureAttribute = "height"
	TypePictureAttribute   PictureAttribute = "type"
	LargePictureAttribute  PictureAttribute = "large"
)

// WithAttributes decorates PictureGetUserField with attributes
func (p GetUserField) WithAttributes(attributes map[PictureAttribute]string) GetUserField {
	if p == PictureGetUserField {
		s := string(p)
		if len(attributes) > 0 {
			for k, v := range attributes {
				s += fmt.Sprintf(".%s(%s)", k, v)
			}
		}

		return GetUserField(s)
	}

	return p
}

type GetUserResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Picture struct {
		Data struct {
			Height       int    `json:"height"`
			Width        int    `json:"width"`
			IsSilhouette bool   `json:"is_silhouette"`
			Url          string `json:"url"`
		} `json:"data"`
	} `json:"picture"`
}
