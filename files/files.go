package files

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/flexzuu/exifgql/photo"
	"github.com/pkg/errors"
	"github.com/rwcarlsen/goexif/exif"
)

var supportedFileExtentions = []string{".jpg", ".arw"}

func OpenPhoto(path string) (*photo.Photo, error) {

	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, errors.Wrapf(err, "getting stat for path %q failed", path)
	}
	if fileInfo.IsDir() {
		return nil, fmt.Errorf("dirs not allowed")

	}
	fileExt := strings.ToLower(filepath.Ext(path))
	for _, supportedFileExtention := range supportedFileExtentions {
		if fileExt == supportedFileExtention {
			f, err := os.Open(path)
			defer f.Close()
			if err != nil {
				return nil, errors.Wrapf(err, "open file %q failed", path)
			}

			exif, err := exif.Decode(f)
			if err != nil {
				return nil, err
			}

			photo := &photo.Photo{
				File: *f,
				Exif: *exif,
			}
			return photo, nil
		}
	}

	return nil, fmt.Errorf("%q, wrong file extention", fileExt)
}
