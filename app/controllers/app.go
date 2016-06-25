package controllers

import (
	"html/template"

	"github.com/revel/revel"
	"github.com/revel/revelframework.com/app/routes"
)

type App struct {
	*revel.Controller
}

func (c *App) AppInfo() revel.Result {
	// populate all required meta data for application
	c.RenderArgs["revelVersion"] = revel.Version
	c.RenderArgs["revelBuildDate"] = revel.BuildDate
	c.RenderArgs["revelMinimumGoVersion"] = revel.MinimumGoVersion
	c.RenderArgs["revelDocsPathPrefix"] = routes.Docs.Index()

	return nil
}

func (c *App) AutoRedirect() revel.Result {
	if c.Request.URL.Path == "/docs" ||
		c.Request.URL.Path == "/docs/" {
		return c.Redirect(routes.Docs.Index())
	}
	return nil
}

func init() {
	// register interceptors
	revel.InterceptMethod((*App).AutoRedirect, revel.BEFORE)
	revel.InterceptMethod((*App).AppInfo, revel.BEFORE)

	// Custom Template Functions
	templateFuncs := map[string]interface{}{
		"noescape": func(s string) template.HTML {
			return template.HTML(s)
		},
	}

	// Adding to revel
	for k, v := range templateFuncs {
		revel.TemplateFuncs[k] = v
	}
}
