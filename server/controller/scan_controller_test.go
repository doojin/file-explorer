package controller

import (
	"testing"
	"github.com/doojin/file-explorer/explorer"
	"github.com/doojin/file-explorer/crypto"
	"github.com/stretchr/testify/assert"
)

func Test_encodeEntities_ShouldEncodeEntitiesCorrectly(t *testing.T) {
	exp := explorer.New("dummy root")
	encoder, _ := crypto.NewEncoder("1234567890123456")
	controller := NewScanController(encoder, exp)
	files := []explorer.File{
		explorer.File{
			Path: "file path",
		},
	}
	directories := []explorer.Directory{
		explorer.Directory{
			Path: "directory path",
		},
	}

	files, directories = controller.encodeEntities(files, directories)

	encodedFilePath := files[0].Path
	encodedDirectoryPath := directories[0].Path

	decodedFilePath, _ := encoder.Decrypt(encodedFilePath)
	decodedDirectoryPath, _ := encoder.Decrypt(encodedDirectoryPath)

	assert.Equal(t, "file path", decodedFilePath)
	assert.Equal(t, "directory path", decodedDirectoryPath)
}

func Test_getParentDir_ShouldReturnParentDirCorrectly(t *testing.T) {
	currentDir := "C:/MyDir/SubDir"

	parentDir := getParentDir(currentDir)

	assert.Equal(t, "C:/MyDir", parentDir)
}