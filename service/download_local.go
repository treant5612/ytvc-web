package service

import (
	"io"
	"os"
	"path/filepath"
)

var downloadPath = "./download"

func SetDownloadPath(path string) {
	downloadPath = path
	os.Mkdir(downloadPath, 0700)
}

func DownloadToLocalFile(reader io.Reader, fileName string) (path string, err error) {
	fpath := filepath.Join(downloadPath, fileName)
	f, err := os.Create(fpath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	_, err = io.Copy(f, reader)
	return fpath, err
}
