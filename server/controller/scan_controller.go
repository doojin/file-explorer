// controller package provide structures with methods to serve HTTP
// requests
package controller

import (
	"github.com/doojin/file-explorer/crypto"
	"net/http"
	"html/template"
	"github.com/doojin/file-explorer/explorer"
	"github.com/gorilla/mux"
	"path/filepath"
	"strings"
)

const current_dir = "dir"

type scanController struct {
	encoder  crypto.Encoder
	explorer explorer.Explorer
}

// NewScanController creates a new instance of scanController
func NewScanController(encoder crypto.Encoder, explorer explorer.Explorer) (controller scanController) {
	controller.encoder = encoder
	controller.explorer = explorer
	return
}

// HomeHandler serves homepage requests
func (controller *scanController) HomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(
		"server/templates/layout.html",
		"server/templates/navigation.html",
		"server/templates/content/scan_result.html",
	)
	if err != nil {
		panic(err)
	}
	directories, _ := controller.explorer.RootDirectories()
	files, _ := controller.explorer.RootFiles()
	files, directories = controller.encodeEntities(files, directories)
	parentDir, _ := controller.encoder.Encrypt(
		getParentDir(controller.explorer.Root),
	)
	tpl.Execute(w, map[string]interface{}{
		"Directories": directories,
		"Files": files,
		"Path": controller.explorer.Root,
		"Parent": parentDir,
	})
}

// ScanHandler serves directory scanning requests
func (controller *scanController) ScanHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tpl, err := template.ParseFiles(
		"server/templates/layout.html",
		"server/templates/navigation.html",
		"server/templates/content/scan_result.html",
	)
	if err != nil {
		panic(err)
	}
	currentDir, _ := controller.encoder.Decrypt(vars[current_dir])
	directories, _ := controller.explorer.Directories(currentDir)
	files, _ := controller.explorer.Files(currentDir)
	files, directories = controller.encodeEntities(files, directories)
	parentDir, _ := controller.encoder.Encrypt(
		getParentDir(currentDir),
	)
	tpl.Execute(w, map[string]interface{}{
		"Directories": directories,
		"Files": files,
		"Path": currentDir,
		"Parent": parentDir,
	})
}

// SearchHandler serves file and directory search requests
func (controller *scanController) SearchHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(
		"server/templates/layout.html",
		"server/templates/navigation.html",
		"server/templates/content/search_result.html",
	)
	if err != nil {
		panic(err)
	}
	tpl.Execute(w, []interface{}{})
}

func (controller *scanController) encodeEntities(files []explorer.File,
directories []explorer.Directory) ([]explorer.File, []explorer.Directory) {
	for key, file := range files {
		files[key].Path, _ = controller.encoder.Encrypt(file.Path)
	}

	for key, directory := range directories {
		directories[key].Path, _ = controller.encoder.Encrypt(directory.Path)
	}
	return files, directories
}

func getParentDir(path string) string {
	parentDir := strings.Replace(filepath.Dir(path), "\\", "/", -1)
	return parentDir
}