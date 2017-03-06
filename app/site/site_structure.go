package site

// Directory sections in manual revel.github.io
// `index` is handled specially
var Sections  = []string{"tutorial", "manual", "modules", "examples", "quickref"}

// Site contains the whole nav etc
var Site *SiteStruct

type SiteStruct struct {
	Title string
	Sections map[string]*Section
}

// Section represents a directory in docs
type Section struct {
	Root string			` yaml:"root" json:"root" `
	Name  string		` yaml:"name"  json:"name"`
	Title string 		` yaml:"section_title" json:"section_title" `
	PageSections[]*PageSection	` yaml:"nav"  json:"page_sections"`
}

// Returns true is this section has this page
func (sec *Section) HasPage(page string) bool {

	url := "/" + sec.Name + "/" + page
	for _, ps := range sec.PageSections {
		for _, p := range ps.Pages {
			if p.Url == url {
				return true
			}
		}
	}
	return false

}

// A PageSection is the title block in jekyll
type PageSection struct {
	Title  string		` yaml:"name" json:"title" `
	Pages []*Page		` yaml:"articles" json:"pages" `
}

// A Page is an item in jekyll config
type Page struct {
	Title  string		` yaml:"title" json:"title" `
	RawUrl string		` yaml:"url" `
	Url string			` yaml:"ignore" `
}


