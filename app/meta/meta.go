package meta

import (
	"fmt"
	"path/filepath"
	//"io/ioutil"

	//"gopkg.in/yaml.v2"

	"github.com/revel/revel"
	"github.com/revel/revelframework.com/app/models"
)

/*
This is a bit of a hack to use the jekyll files,
as source for this site..
eg
the "section menus" are in _layouts/*.html as front matter
each page has frontmatter
 */

var Sections map[string]Section

var (
	// Path is base path location of metadata files
	Path string

	// Docs is documentation meta configuration values
	Docs models.Docs
)

type JekyllConf struct {
	Sections []string ` yaml:"section" `
}

// LoadMetaData parses and loads metadata
func LoadMetaData() {

	docs_root := filepath.Join(revel.BasePath, "../revel.github.io")
	fmt.Printf("docs_root: %#v", docs_root)
	var sections = []string{"tutorial"}

	for _, r := range sections {
		fmt.Printf("docs_root: %#v", docs_root + "/_layouts/" + r + ".html")
	}
	//var jek JekyllConf
	//if err := yaml.Unmarshal(yamlFile, &jek); err != nil {
	//	revel.ERROR.Fatalln("Yaml decode error:", err)
	//}
	fmt.Printf("Values: %#v", Docs)
}

func init() {
	Path = filepath.Join(revel.BasePath, "metadata")

	revel.OnAppStart(LoadMetaData)
}
