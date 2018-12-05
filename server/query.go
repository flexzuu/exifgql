package main

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/flexzuu/exifgql/files"
	zglob "github.com/mattn/go-zglob"
)

type query struct{}

var photoMap = make(map[string]string)

func (_ *query) Photos() ([]*photo, error) {
	paths, err := zglob.Glob(glob)
	if err != nil {
		return nil, err
	}
	res := make([]*photo, len(paths))
	for i, path := range paths {
		h := sha1.New()
		h.Write([]byte(path))
		hash := hex.EncodeToString(h.Sum(nil))
		photoMap[hash] = path

		ph, err := files.OpenPhoto(path)
		if err != nil {
			return nil, err
		}
		res[i] = &photo{Photo: ph, id: hash}
	}
	return res, nil
}
