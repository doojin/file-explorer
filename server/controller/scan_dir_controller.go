// controller package provide structures with methods to serve HTTP
// requests
package controller

import (
	"github.com/doojin/file-explorer/crypto"
	"net/http"
	"html/template"
)

type scanController struct {
	encoder crypto.Encoder
}

// NewScanController creates a new instance of scanController
func NewScanController(encoder crypto.Encoder) (controller scanController) {
	controller.encoder = encoder
	return
}

// HomeHandler serves homepage requests
func (controller *scanController) HomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(
		"server/templates/layout.html",
		"server/templates/content/homepage.html",
	)
	if err != nil {
		panic(err)
	}
	tpl.Execute(w, []interface{}{})
}

// ScanHandler serves directory scanning requests
func (controller *scanController) ScanHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(
		"server/templates/layout.html",
		"server/templates/content/scan_result.html",
	)
	if err != nil {
		panic(err)
	}
	tpl.Execute(w, []interface{}{})
}

// SearchHandler serves file and directory search requests
func (controller *scanController) SearchHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(
		"server/templates/layout.html",
		"server/templates/content/search_result.html",
	)
	if err != nil {
		panic(err)
	}
	tpl.Execute(w, []interface{}{})
}