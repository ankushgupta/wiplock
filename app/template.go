package app

import (
	"github.com/eknkc/amber"
	"html/template"
)

func NewTemplate() *template.Template {
	return amber.MustCompile(`!!! 5
html
  head
    script[src=JavaScriptPath]
    title #{SiteTitle}
    meta[name="description"][content=SiteDescription]
  body[data-access-token=AccessToken]
    #app-root`, amber.Options{PrettyPrint: false})
}
