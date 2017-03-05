package meta

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

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

type GitHub struct {
	Labels []string
}

type FrontMatter struct {
	Title string ` yaml:"title" `
	Layout string ` yaml:"layout" ` // not used see section
	GitHub GitHub ` yaml:"github" `
	GoDoc []string ` yaml:"godoc" `
}

type PageData struct {
	Section string
	Page string
	Path string
	Title string
	Content template.HTML
	GitHub GitHub
	GoDoc []string
	Error error
	FilePath string
}



// A markdown file has some yaml "frontmatter" at the top  which contains
// title, so this scans line by line.. (TODO jekyll tag replace)
//     ---
//     title: foo
//     layout: manual
//     ---
// TODO , this is real slow and needs a regex expert
//
func ReadMarkDownPage( section, page string) PageData {

	pdata := PageData{Section: section, Page: page, Title: "- no title -"}

	pdata.FilePath = DocsRootPath + "/" + pdata.Section + "/" + pdata.Page + ".md"
	fmt.Println("FILE2pead", pdata.FilePath)
	file, err := os.Open(pdata.FilePath)
	if err != nil {
		pdata.Error = err
		return pdata
	}
	defer file.Close()

	yaml_bounds := "---"
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
		if line == yaml_bounds {
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
	pdata.GitHub = fm.GitHub
	pdata.GoDoc = fm.GoDoc

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
