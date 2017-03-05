package controllers

import (
	"fmt"


	"github.com/revel/revel"

	"github.com/revel/revelframework.com/app/meta"
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
	return c.RenderJson(meta.Site)
}



// render an expected markdown file
func (c Page) Page() revel.Result {

	c.RenderArgs["Site"] = meta.Site

	// Create a temporary PageData
	tdata := meta.PageData{Title: "## Unknown ##"}

	// get the section and page from oute
	tdata.Section = c.Params.Route.Get("section")
	tdata.Page = c.Params.Route.Get("page")

	// Validate section by checking is exsts in site
	sec , ok := meta.Site.Sections[tdata.Section]
	if !ok {
		tdata.Title = "404 Not Found"
		c.RenderArgs["Page"] = tdata
		return c.NotFound("missing secton")
	}
	c.RenderArgs["Section"] = sec

	// no page so set to default
	if tdata.Page == "" {
		tdata.Page = "index"

	} else {
		//validate page
		if !sec.HasPage(tdata.Page){
			return c.NotFound("missing page")
		}
	}


	//c.RenderArgs["Page"] = cp


	pdata := meta.ReadMarkDownPage(tdata.Section, tdata.Page)
	c.RenderArgs["Page"] = pdata
	if pdata.Error != nil {
		fmt.Println("error==", pdata.Error)
	}
	fmt.Println("error==", pdata.Title)

	return c.Render()

}
