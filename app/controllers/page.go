package controllers

import (
	//"fmt"


	"github.com/revel/revel"

	"github.com/revel/revelframework.com/app/site"
)



type Page struct {
	*revel.Controller
}

// /robots.txt - Only allow spiders on prod site
func (c Page) RobotsTxt() revel.Result {

	txt := "User-agent: *\n"
	if revel.Config.BoolDefault("site.live", false)  == false {
		txt += "Disallow: /\n"
	}
	txt += "\n"

	return c.RenderText(txt)
}

// main home page
func (c Page) Index() revel.Result {
	return c.Render()
}
func (c Page) Debug(section, page string) revel.Result {
	return c.RenderJson(site.Site)
}



// render an expected markdown file
func (c Page) Page() revel.Result {

	c.RenderArgs["Site"] = site.Site

	// Create  PageData
	pdata := site.LoadPage(c.Params.Route.Get("section"), c.Params.Route.Get("page"))
	c.RenderArgs["Page"] = pdata

	if pdata.Error != nil {
		return c.NotFound("missing secton")
	}

	c.RenderArgs["Section"] = site.Site.Sections[pdata.Section]

	return c.Render()

}
