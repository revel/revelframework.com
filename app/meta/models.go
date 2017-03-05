package meta

var Site *SiteStruct

type SiteStruct struct {
	Title string
	Sections map[string]Section
}


type Section struct {
	Root string			` yaml:"root" json:"root" `
	Name  string		` yaml:"name"  json:"name"`
	Title string 		` yaml:"section_title" json:"section_title" `
	PageSections[]PageSection	` yaml:"nav"  json:"page_sections"`
}

type PageSection struct {
	Title  string		` yaml:"name" json:"title" `
	Pages []Page		` yaml:"articles" json:"pages" `
}

type Page struct {
	Title  string		` yaml:"title" json:"title" `
	Url string			` yaml:"url" `
}

