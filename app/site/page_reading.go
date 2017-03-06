package site

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"errors"

	//"//io/ioutil"
	"html/template"
	//"path/filepath"
	"strings"

	//"github.com/revel/revel"
	"github.com/russross/blackfriday"
	"gopkg.in/yaml.v2"
	"github.com/pksunkara/pygments"

	//"github.com/revel/revel"//
)

const (
	YAML_DELIM = "---"
)
/* Example front matter
---
title: Controllers Overview
layout: manual
github:
  labels:
    - topic-controller
godoc:
    - Controller
    - Request
    - Response
---
markdown starts here
 */
type GitHub struct {
	Labels []string
}

type FrontMatter struct {
	Title string ` yaml:"title" `
	Layout string ` yaml:"layout" ` // not used see section
	Github GitHub ` yaml:"github" `
	Godoc []string ` yaml:"godoc" `
}

type PageData struct {
	Section string
	Page string
	Path string
	Title string
	Content template.HTML
	Github GitHub
	Godoc []string
	FilePath string
	Error error
}

// Loads a markdown page from repos,
// set PageData.Error message is f*up
func LoadPage(section, page string) *PageData {
	pdata := new(PageData)

	// process section = dir
	pdata.Section = CleanStr(section)
	// todo sanitize
	_, ok := Site.Sections[section]
	if !ok {
		pdata.Error = errors.New("Section not found")
		return pdata
	}

	// process page = /section/somefile
	// but incoming might be page.html or page.md line
	// also if blank then its index
	pdata.Page = StripExt( CleanStr(page) )
	if pdata.Page == "" {
		pdata.Page = "index"
	}
	pdata.Path = "/" + pdata.Section + "/" + pdata.Page

	// derive the markdown file path
	pdata.FilePath = DocsRootPath + pdata.Path + ".md"
	if _, err := os.Stat(pdata.FilePath); err != nil {
		pdata.Error = err
		return pdata
	}
	ReadMarkDownPage(pdata)

	return pdata
}

// A markdown file has some yaml "frontmatter" at the top  which contains
// title, so this scans line by line.. (TODO jekyll tag replace)
//     ---
//     title: foo
//     layout: manual
//     ---
// TODO , this is real slow and needs a regex expert
//
func ReadMarkDownPage( pdata *PageData) *PageData {

	file, err := os.Open(pdata.FilePath)
	if err != nil {
		pdata.Error = err
		return pdata
	}
	defer file.Close()

	// SOMEONE please refactor this...
	// can even remember why and when said pedro
	yaml_str := ""
	body_str := ""
	found_yaml_start := false
	in_yaml := false // we always expect yaml ??
	in_code := false
	code_str := ""
	lexer := ""
	scanner := bufio.NewScanner(file)
	for  scanner.Scan() {
		line := scanner.Text()
		if line == YAML_DELIM {
			if found_yaml_start == false {
				in_yaml = true
				found_yaml_start = true
			} else {
				in_yaml = false
			}
		} else {
			if in_yaml {
				yaml_str += line + "\n"
			} else {
				// TODO need a regex for "{%highlight foo %}"
				if len(line) > 2 && line[0:2] == "{%" && strings.Contains(line, "endhighlight")  == false && strings.Contains(line, "highlight")  {
					//fmt.Println("GOT CODE=" , line)
					xline := line
					xline = strings.Replace(xline, "{%", "", 1)
					xline = strings.Replace(xline, "%}", "", 1)
					xline = strings.Replace(xline, "highlight", "", 1)
					xline = strings.TrimSpace(xline)
					//fmt.Println("GOT CODE=" , line, xline)
					lexer = xline
					//if line == "{% highlight go %}" {
					//body_str += "``` go\n"
					body_str += "\n"
					code_str = ""
					in_code = true

					//}
				} else if len(line) > 2 && line[0:2] == "{%" && strings.Contains(line, "endhighlight")  {
					//fmt.Println("END CODE=" , line)
					//body_str += "##########" + code_str + "###########"
					hi_str := pygments.Highlight(code_str, lexer, "html", "utf-8")
					//fmt.Println("hi=", hi_str)
					body_str += string(hi_str)
					//body_str += "```\n"
					body_str += "\n"
					in_code = false

				} else {
					if in_code {
						code_str += line + "\n"
					} else {
						body_str += line+"\n"
					}

				}
			}
		}

	}
	if err := scanner.Err(); err != nil {
		pdata.Error = err
		return pdata
	}


	// parse yaml header bit
	var fm FrontMatter
	err = yaml.Unmarshal([]byte(yaml_str), &fm)
	if err != nil {
		fmt.Println("error md", err)
	}
	pdata.Title = fm.Title
	pdata.Github = fm.Github
	pdata.Godoc = fm.Godoc

	//fmt.Println("===", pdata)

	// convert markdown
	extensions := 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS

	htmlFlags := 0
	htmlFlags |= blackfriday.HTML_USE_SMARTYPANTS
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_FRACTIONS
	//htmlFlags |= blackfriday.HTML_GITHUB_BLOCKCODE
	//htmlFlags |= blackfriday.HTML_TOC

	renderer := blackfriday.HtmlRenderer(htmlFlags, "", "")
	output := blackfriday.Markdown([]byte(body_str), renderer, extensions)


	pdata.Content = template.HTML(output)

	//fmt.Println("yamiii", output)

	return pdata

}

func GetGoDocPackage(package_path string) template.HTML {

	go_file_path := "github.com/revel/" + package_path //revel/" // + go_file
	fmt.Println("sources=", go_file_path)
	app := "godoc"

	cmd := exec.Command(app, "-html", go_file_path)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		//return
	}
	return template.HTML(stdout)
}
