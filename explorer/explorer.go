// Explorer package provides functionality for directory scanning
package explorer

import (
	"io/ioutil"
	"os"
	"errors"
	"strings"
)

// DELIMITER is a directory separator
const delimiter = "/"

var (
	ERR_CANNOT_SCAN = errors.New("Cannot scan this directory")
	ERR_OUT_OF_ROOT = errors.New("Directory you try to scan is outside the root directory")
)

// Explorer structure contains methods for directory scanning
type Explorer struct {
	Root string
}

// New returns a new instance of Explorer
func New(root string) Explorer {
	return Explorer{Root: root}
}

// RootDirectories returns a slice of directories within the root directory
func (explorer *Explorer) RootDirectories() (directories []Directory, err error) {
	return explorer.Directories(explorer.Root)
}

// RootFiles returns a slice of files within the root directory
func (explorer *Explorer) RootFiles() (files []File, err error) {
	return explorer.Files(explorer.Root)
}

// Directories returns a slice of directories within the provided path
func (explorer *Explorer) Directories(path string) (directories []Directory, err error) {
	if err = explorer.checkLevel(path); err != nil {
		return
	}
	entities, err := ioutil.ReadDir(path)
	if err != nil {
		err = ERR_CANNOT_SCAN
		return
	}
	directories = filterDirectories(entities, path)
	return
}

// Files returns a slice of files within the provided path
func (explorer *Explorer) Files(path string) (files []File, err error) {
	if err = explorer.checkLevel(path); err != nil {
		return
	}
	entities, err := ioutil.ReadDir(path)
	if err != nil {
		err = ERR_CANNOT_SCAN
		return
	}
	files = filterFiles(entities, path)
	return
}

func (explorer *Explorer) checkLevel(path string) (err error) {
	root := strings.TrimSuffix(explorer.Root, delimiter)
	path = strings.TrimSuffix(path, delimiter)
	if !strings.HasPrefix(path, root) {
		err = ERR_OUT_OF_ROOT
	}
	return
}

// FindEntities searches for files and folders with specified name
func (explorer *Explorer) FindEntities(path string, name string, level int, currentLevel int) (resultFiles []File, resultDirectories []Directory) {
	// In current dir
	directories, _ := explorer.Directories(path)
	files, _ := explorer.Files(path)
	matchedDirectories, matchedFiles := matchedEntities(directories, files, name)

	// Appending matched
	resultFiles = append(resultFiles, matchedFiles...)
	resultDirectories = append(resultDirectories, matchedDirectories...)

	// In subdirectories
	for _, subDirectory := range directories {
		matchedFiles, matchedDirectories = explorer.FindEntities(subDirectory.Path, name, level, currentLevel+1)
		resultFiles = append(resultFiles, matchedFiles...)
		resultDirectories = append(resultDirectories, matchedDirectories...)
	}

	return
}

func matchedEntities(directories []Directory, files []File, name string) (matchedDirectories []Directory, matchedFiles []File) {
	name = strings.ToLower(name)
	// Matching directories
	for _, directory := range directories {
		dirName := strings.ToLower(directory.Name)
		if strings.Contains(dirName, name) {
			matchedDirectories = append(matchedDirectories, directory)
		}
	}
	// Matching files
	for _, file := range files {
		fileName := strings.ToLower(file.Name)
		if strings.Contains(fileName, name) {
			matchedFiles = append(matchedFiles, file)
		}
	}
	return
}

func filterDirectories(entities []os.FileInfo, path string) (directories []Directory) {
	for _, entity := range entities {
		if entity.IsDir() {
			directories = append(directories, Directory{
				entity.Name(),
				buildPath(path, entity.Name()),
			})
		}
	}
	return
}

func filterFiles(entities []os.FileInfo, path string) (files []File) {
	for _, entity := range entities {
		if !entity.IsDir() {
			files = append(files, File{
				entity.Name(),
				entity.Size(),
				buildPath(path, entity.Name()),
			})
		}
	}
	return
}

func buildPath(path string, name string) string {
	path = strings.TrimSuffix(path, delimiter)
	return path + delimiter + name
}