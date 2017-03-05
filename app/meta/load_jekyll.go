package meta

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

//var Sections map[string]Section

var (
	// DocsRootPath is base path location of revel.github.io www checkout
	DocsRootPath string
)

// LoadMetaData parses and loads metadata
func LoadMetaData() {


	fmt.Printf("docs_root: %#v", DocsRootPath)



	Site.Sections = make(map[string]*Section)
	for _, r := range Sections {
		fmt.Println("docs_root:", DocsRootPath + "/_layouts/" + r + ".html")
		Site.Sections[r] = ReadJekyllLayout(r)
	}


}

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
			continue
		}
		if line == "---" {
			break
		}
		front_matter += line + "\n"
	}

	sec := new(Section)
	err = yaml.Unmarshal([]byte(front_matter), sec)
	if err != nil {
		fmt.Println("error, yaml", err)
	}
	//fmt.Println(front_matter, sec)
	// NOW cleanup the urls
	for _, psec := range sec.PageSections {
		for _, page := range psec.Pages {
			p := page.RawUrl
			if strings.HasSuffix(p, ".html") {
				p = p[0:len(p)-5]
			} else if strings.HasSuffix(p, ".md") {
				p = p[0:len(p)-3]
			}
			page.Url = "/" + section + "/" + p
		}
	}


	return sec
}

func init() {
	DocsRootPath = filepath.Join(revel.BasePath, "..", "revel.github.io")
	Site = new(SiteStruct)
	Site.Title = "Revel Framework"
	Site.Sections = make(map[string]*Section)
	revel.OnAppStart(LoadMetaData)
}
