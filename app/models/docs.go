package models

type Docs struct {
	Title    string
	BasePath string
	Versions []string
	Sections []Section
}

type Section struct {
	Name  string
	Pages []Page
}

type Page struct {
	Title    string
	Path     string
	FileName string
}
