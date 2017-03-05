package meta




type Section struct {
	Root string			` yaml:"root" `
	Name  string		` yaml:"name" `
	Title string 		` yaml:"section_title" `
	Nav []PageSection	` yaml:"nav" `
}

type PageSection struct {
	Title  string		` yaml:"name" `
	Pages []Page		` yaml:"articles" `
}

type Page struct {
	Title  string		` yaml:"name" `
	Url string			` yaml:"url" `
}