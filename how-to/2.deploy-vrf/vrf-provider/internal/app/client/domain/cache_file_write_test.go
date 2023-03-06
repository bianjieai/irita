package domain

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCacheFileWriter_Write(t *testing.T) {
	cfgDirName := ".vrf-provider"
	userDir, _ := os.UserHomeDir()
	homeDir := filepath.Join(userDir, cfgDirName)
	dir := "cache"
	filename := "iris.json"
	writer := NewCacheFileWriter(homeDir, dir, filename)
	err := writer.Write(1)
	if err != nil {
		t.Fatal(err)
	}

}
