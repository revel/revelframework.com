package meta

import (
	"fmt"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/revel/revel"
	"github.com/revel/revelframework.com/app/models"
)

var (
	// Path is base path location of metadata files
	Path string

	// Docs is documentation meta configuration values
	Docs models.Docs
)

// LoadMetaData parses and loads metadata
func LoadMetaData() {
	docsMetaPath := filepath.Join(Path, "docs.toml")

	if _, err := toml.DecodeFile(docsMetaPath, &Docs); err != nil {
		revel.ERROR.Fatalln("TOML file read error:", err)
	}
	fmt.Printf("Values: %#v", Docs)
}

func init() {
	Path = filepath.Join(revel.BasePath, "metadata")

	revel.OnAppStart(LoadMetaData)
}
