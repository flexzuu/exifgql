package photo

import (
	"fmt"
	"os"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

type Photo struct {
	os.File
	exif.Exif
}

func (p Photo) Name() string {
	return p.File.Name()
}

func (p Photo) DateTime() (time.Time, error) {
	return p.Exif.DateTime()
}

func (p Photo) Model() (string, error) {
	tag, err := p.Exif.Get(exif.Model)
	if err != nil {
		return "", err
	}
	return tag.StringVal()
}

func (p Photo) Make() (string, error) {
	tag, err := p.Exif.Get(exif.Make)
	if err != nil {
		return "", err
	}
	return tag.StringVal()
}

func (p Photo) LensModel() (string, error) {
	tag, err := p.Exif.Get(exif.LensModel)
	if err != nil {
		return "", err
	}
	return tag.StringVal()
}

func (p Photo) LensMake() (string, error) {
	tag, err := p.Exif.Get(exif.LensMake)
	if err != nil {
		return "", err
	}
	return tag.StringVal()
}

func (p Photo) ISO() (int, error) {
	tag, err := p.Exif.Get(exif.ISOSpeedRatings)
	if err != nil {
		return 0, err
	}
	return tag.Int(0)
}
func (p Photo) FocalLengthIn35mmFilm() (int, error) {
	tag, err := p.Exif.Get(exif.FocalLengthIn35mmFilm)
	if err != nil {
		return 0, err
	}
	return tag.Int(0)
}
func (p Photo) ExposureTime() (string, error) {
	tag, err := p.Exif.Get(exif.ExposureTime)
	if err != nil {
		return "", err
	}
	n, d, err := tag.Rat2(0)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d/%d", n, d), nil
}

func (p Photo) FNumber() (float64, error) {
	tag, err := p.Exif.Get(exif.FNumber)
	if err != nil {
		return 0, err
	}
	n, d, err := tag.Rat2(0)
	if err != nil {
		return 0, err
	}
	return float64(n) / float64(d), nil
}

func (p Photo) FocalLength() (float64, error) {
	tag, err := p.Exif.Get(exif.FocalLength)
	if err != nil {
		return 0, err
	}
	n, d, err := tag.Rat2(0)
	if err != nil {
		return 0, err
	}
	return float64(n) / float64(d), nil
}
