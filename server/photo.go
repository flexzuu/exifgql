package main

import (
	"encoding/base64"
	"fmt"

	p "github.com/flexzuu/exifgql/photo"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/rwcarlsen/goexif/exif"
)

type photo struct {
	id string
	*p.Photo
}

func swollowTagNotPresent(err error) error {
	if exif.IsTagNotPresentError(err) {
		return nil
	}
	return err
}
func (p *photo) ID() graphql.ID {
	return graphql.ID(p.id)
}
func (p *photo) DateTime() (graphql.Time, error) {
	time, err := p.Photo.DateTime()
	if err != nil {
		return graphql.Time{}, err
	}
	return graphql.Time{time}, nil
}
func (p *photo) Model() (*string, error) {
	res, err := p.Photo.Model()
	if err != nil {
		return nil, swollowTagNotPresent(err)
	}
	return &res, nil
}

func (p *photo) Make() (*string, error) {
	res, err := p.Photo.Make()
	if err != nil {
		return nil, swollowTagNotPresent(err)
	}
	return &res, nil
}

func (p *photo) LensModel() (*string, error) {
	res, err := p.Photo.LensModel()
	if err != nil {
		return nil, swollowTagNotPresent(err)
	}
	return &res, nil
}

func (p *photo) LensMake() (*string, error) {
	res, err := p.Photo.LensMake()
	if err != nil {
		return nil, swollowTagNotPresent(err)
	}
	return &res, nil
}

func (p *photo) ExposureTime() (*string, error) {
	res, err := p.Photo.ExposureTime()
	if err != nil {
		return nil, swollowTagNotPresent(err)
	}
	return &res, nil
}

func (p *photo) ISO() (*int32, error) {
	res, err := p.Photo.ISO()
	if err != nil {
		return nil, swollowTagNotPresent(err)
	}
	x := int32(res)
	return &x, nil
}

func (p *photo) FocalLengthIn35mmFilm() (*int32, error) {
	res, err := p.Photo.FocalLengthIn35mmFilm()
	if err != nil {
		return nil, swollowTagNotPresent(err)
	}
	x := int32(res)
	return &x, nil
}

func (p *photo) FNumber() (*float64, error) {
	res, err := p.Photo.FNumber()
	if err != nil {
		return nil, swollowTagNotPresent(err)
	}

	return &res, nil
}

func (p *photo) FocalLength() (*float64, error) {
	res, err := p.Photo.FocalLength()
	if err != nil {
		return nil, swollowTagNotPresent(err)
	}

	return &res, nil
}
func (p *photo) File() (*string, error) {
	uri := "http://localhost:8080/img/" + p.id
	return &uri, nil
}

func (p *photo) Thumbnail() (*string, error) {
	buf, err := p.Photo.JpegThumbnail()
	if err != nil {
		return nil, swollowTagNotPresent(err)
	}
	uri := fmt.Sprintf("data:image/jpg;base64,%s", base64.StdEncoding.EncodeToString(buf))
	return &uri, nil
}
