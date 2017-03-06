package site

import (
	"fmt"
	"path/filepath"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/revel/revel"

)

/*
This is a bit of a hack to use the jekyll files,
as source for this site..
eg
the "section menus" are in _layouts/*.html as front matter
each page has frontmatter
 */

// DocsRootPath is base path location of revel.github.io www checkout
var	DocsRootPath string


func init() {
	// we assute in in gopath's github/revel/* hence about + outside this repos
	DocsRootPath = filepath.Join(revel.BasePath, "..", "revel.github.io")

	Site = new(SiteStruct)
	Site.Title = "Revel Framework"
	Site.Sections = make(map[string]*Section)

	revel.OnAppStart(LoadSiteStructure)
}


// LoadSiteStructure parses and loads the site sections
// (this should be reloadable)
func LoadSiteStructure() {

	fmt.Printf("docs_root: %#v", DocsRootPath)

	// reinit here ?? is this a leak
	Site.Sections = make(map[string]*Section)
	for _, r := range Sections {
		fmt.Println("docs_root:", DocsRootPath + "/_layouts/" + r + ".html")
		Site.Sections[r] = ReadJekyllLayout(r)
	}

}

// Parses the jekyll _layouts/*.html files for nav
func ReadJekyllLayout(section string) *Section{

	// menu is in frontmatter of _layouts/SECTION.html
	// eg https://github.com/revel/revel.github.io/blob/master/_layouts/tutorial.html
	lay_file := DocsRootPath + "/_layouts/" + section + ".html"

	// read file
	contents, err := ioutil.ReadFile(lay_file)
	if err != nil {
		revel.ERROR.Fatalln("Yaml decode error:", err)
	}

	// convert to string and split into lines
	lines := strings.Split(string(contents), "\n")
	front_matter := ""
	for idx, line := range lines {

		if idx == 0 {
			// todo check is "---" is first
			continue
		}
		if line == "---" {
			// end of yaml, so get outta here
			break
		}
		front_matter += line + "\n"
	}

	// Parse section from yaml
	sec := new(Section)
	err = yaml.Unmarshal([]byte(front_matter), sec)
	if err != nil {
		fmt.Println("error, yaml", err)
	}

	// NOW cleanup the urls
	for _, psec := range sec.PageSections {
		for _, page := range psec.Pages {
			page.Url = "/" + section + "/" + StripExt(page.RawUrl)
		}
	}


	return sec
}

