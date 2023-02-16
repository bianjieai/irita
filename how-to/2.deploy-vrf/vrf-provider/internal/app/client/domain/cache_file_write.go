package domain

import (
	"encoding/json"
	"os"
	"path"

	"gitlab.bianjie.ai/avata/contracts/vrf-provider/internal/pkg/types/cache"
)

type CacheFileWriter struct {
	homeDir       string
	cacheDir      string
	cacheFilename string
}

func NewCacheFileWriter(homeDir, cacheDir, cacheFilename string) *CacheFileWriter {
	return &CacheFileWriter{homeDir: homeDir, cacheDir: cacheDir, cacheFilename: cacheFilename}
}

func (w *CacheFileWriter) Write(height uint64) error {

	cacheDataObj := &cache.Data{}
	cacheDataObj.LatestHeight = height

	cacheDataWriteBytes, err := json.Marshal(cacheDataObj)
	if err != nil {
		return err
	}

	cacheDir := path.Join(w.homeDir, w.cacheDir)
	filename := path.Join(cacheDir, w.cacheFilename)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// And the home folder doesn't exist
		if _, err := os.Stat(w.homeDir); os.IsNotExist(err) {
			// Create the home folder
			if err = os.Mkdir(w.homeDir, os.ModePerm); err != nil {
				return err
			}
		}
		// Create the home config folder
		if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
			// Create the home folder
			if err = os.Mkdir(cacheDir, os.ModePerm); err != nil {
				return err
			}
		}
		// Then create the file...
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		if _, err = file.Write(cacheDataWriteBytes); err != nil {
			return err
		}

	} else {
		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err = file.Write(cacheDataWriteBytes); err != nil {
			return err
		}
	}
	return nil
}
