package meta

// directory sections in manual
var Sections  = []string{"tutorial", "manual", "modules", "examples", "quickref"}

// Site contains the whole Meta
var Site *SiteStruct

type SiteStruct struct {
	Title string
	Sections map[string]*Section
}

// Section reader
type Section struct {
	Root string			` yaml:"root" json:"root" `
	Name  string		` yaml:"name"  json:"name"`
	Title string 		` yaml:"section_title" json:"section_title" `
	PageSections[]*PageSection	` yaml:"nav"  json:"page_sections"`
}

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

type PageSection struct {
	Title  string		` yaml:"name" json:"title" `
	Pages []*Page		` yaml:"articles" json:"pages" `
}

type Page struct {
	Title  string		` yaml:"title" json:"title" `
	RawUrl string		` yaml:"url" `
	Url string			` yaml:"ignore" `
}


