// Public Domain (-) 2014 The Espians Website Authors.
// See the Espians Website UNLICENSE file for details.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const (
	outputDirectory = "www"
	tagline         = "We are building the foundation for the post-industrial future. Join us!"
)

var (
	colorShades = []string{"#f16c6c", "#f13c3c", "#aa2b2b", "#940000"}
	header      = `<!doctype html>
<meta charset=utf-8>
<title>Espians</title>
<meta name="description" content="` + tagline + `">
<script>
  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');
  ga('create', 'UA-90176-9', 'espians.com'); ga('send', 'pageview');
</script>`
)

type Person struct {
	ID       string
	Name     string
	Link     string
	GitHub   string
	LinkedIn string
	Skype    string
	Twitter  string
	Text     string
}

type Project struct {
	Title    string
	Link     string
	Year     int
	Facebook string
	GitHub   string
	Twitter  string
	Text     string
}

var currentProjects = []*Project{
	{
		Title:    "WikiHouse",
		Link:     "http://www.wikihouse.cc/",
		Year:     2011,
		Facebook: "WikiHouse",
		GitHub:   "tav/wikihouse-plugin",
		Twitter:  "wikihouse",
		Text:     `Launched in 2011,`,
	},
	{
		Title:    "Wikifactory",
		Link:     "https://www.wikifactory.org/",
		Year:     2011,
		Facebook: "wikifactory",
		GitHub:   "tav/wikifactory",
		Twitter:  "wikifactory",
		Text:     ``,
	},
	{
		Title:   "Ampify",
		Link:    "http://ampify.net",
		Year:    2001,
		GitHub:  "tav/ampify",
		Twitter: "ampify",
		Text:    `Set to launch in early 2015, Ampify is a decentralised application platform.`,
	},
}

var investments = []*Project{
	{
		Title:    "Wigwamm",
		Link:     "https://wigwamm.com/",
		Year:     2013,
		Facebook: "Wigwamm",
		GitHub:   "wigwamm",
		Twitter:  "wigwammco",
		Text:     ``,
	},
	{
		Title: "Social Startup Labs",
		Text:  ``,
	},
}

var pastProjects = []*Project{

	{
		Title: "Atlas",
		Text:  ``,
	},
	{
		Title: "Civic Crowd",
		Text:  ``,
	},
	{
		Title: "TrustMaps",
		Text:  ``,
	},
	{
		Title: "OpenCoin",
		Text:  ``,
	},
	{
		Title: "Efpee Riva",
		Text:  ``,
	},
	{
		Title: "IceSphere",
		Link:  "",
		Text: `One of our first acquihires, Icesphere was a replacement desktop interface
        for Microsoft Windows. Its powerful block-based component model Python-based IceScript.`,
	},
	{
		Title: "Civic Crowd",
		Text:  ``,
		Year:  2006,
	},
	{
		Title: "PlexNews",
		Text:  `An RSS aggregator of wikis - before `,
	},
	{
		Title: "WorldWideWiki",
		Text:  `Created to create `,
	},
	{
		Title: "Xnet",
		Text:  `A programmable wiki built for collaboration. It enabled task management, blogging and team`,
		Year:  2000,
	},
	{
		Title: "United Diversity",
		Link:  "http://uniteddiversity.coop/",
		Text:  ``,
		Year:  2002,
	},

	{
		Title: "Esp Setup",
		Text:  ``,
		Year:  2000,
	},
	{
		Title: "Espra File Sharing",
		Text:  ``,
		Year:  2000,
	},
	{
		Title: "Advue",
		Text:  ``,
		Year:  1999,
	},
}

var activeEspians = []*Person{
	{
		ID:       "tav",
		Name:     "tav",
		Link:     "http://tav.espians.com/",
		GitHub:   "tav",
		LinkedIn: "in/asktav",
		Skype:    "tavespian",
		Twitter:  "tav",
		Text:     ``,
	},
	{
		ID:       "tom",
		Name:     "Tom Salfield",
		GitHub:   "salfield",
		LinkedIn: "pub/tom-salfield/19/893/258",
		Skype:    "tomsalfield",
		Twitter:  "tsalfield",
		Text:     ``,
	},
	{
		ID:       "christina",
		Name:     "Christina Rebel",
		LinkedIn: "in/christinarebel",
		Skype:    "christina.rebel.88",
		Twitter:  "christina_rebel",
		Text:     ``,
	},
	{
		ID:       "max",
		Name:     "Max Kampik",
		LinkedIn: "in/maximiliankampik",
		Skype:    "maxi.kampik",
		Twitter:  "mkampik",
		Text:     ``,
	},
	{
		ID:      "micrypt",
		Name:    "Ṣeyi Ogunyẹ́mi",
		Link:    "http://micrypt.com/",
		GitHub:  "micrypt",
		Skype:   "micrypt",
		Twitter: "micrypt",
		Text:    ``,
	},
	{
		ID:       "alice",
		Name:     "Alice Fung",
		LinkedIn: "pub/alice-fung/13/932/954",
		Skype:    "alfung4870",
		Twitter:  "00alice",
		Text:     ``,
	},
	{
		ID:       "osb",
		Name:     "Oliver Sylvester-Bradley",
		LinkedIn: "in/olisb",
		Skype:    "oli-s-b",
		Twitter:  "defactodesign",
		Text:     ``,
	},
}

var advisoryBoard = []*Person{
	//	{
	//		ID:       "bauwens",
	//		Name:     "Michel Bauwens",
	//		Link:     "http://p2pfoundation.net/Michel_Bauwens/",
	//		LinkedIn: "in/mbauwens",
	//		Twitter:  "mbauwens",
	//	},
	{
		ID:       "peitersen",
		Name:     "Nicolai Peitersen",
		Link:     "http://www.amazon.co.uk/The-Ethical-Economy-Rebuilding-Crisis/dp/0231152647",
		LinkedIn: "pub/nicolai-peitersen/0/904/852",
		Twitter:  "NPeitersen",
	},
	{
		ID:      "rheingold",
		Name:    "Howard Rheingold",
		Link:    "http://rheingold.com/",
		Twitter: "hrheingold",
	},
}

func exit(args ...interface{}) {
	if len(args) == 1 {
		fmt.Printf("\n!! ERROR: %s\n\n", args...)
	} else {
		fmt.Printf("\n!! ERROR: "+args[0].(string)+"\n\n", args[1:]...)
	}
	os.Exit(1)
}

func main() {

	assetsFile, err := os.Open("assets.json")
	if err != nil {
		exit(err)
	}

	assetMap := map[string]string{}
	assetsDecoder := json.NewDecoder(assetsFile)

	err = assetsDecoder.Decode(&assetMap)
	if err != nil {
		exit(err)
	}

	getPath := func(key string) string {
		val, exists := assetMap[key]
		if !exists {
			exit("Cannot find %s in assets.json", key)
		}
		return "static/" + val
	}

	buf := &bytes.Buffer{}
	o := func(s string, args ...interface{}) {
		s += "\n"
		if len(args) > 1 {
			fmt.Fprintf(buf, s, args...)
		} else {
			buf.WriteString(s)
		}
	}

	o(header)
	o("<link rel=stylesheet href=" + getPath("style.css") + ">")
	o("<div class=wrapper>")
	o("<h1>Espians</h1>")
	o("</div>")
	o("<div id=network><div class=wrapper>" + tagline + "</div></div>")
	o("<script src=" + getPath("site.js") + " async></script>")
	o("<div class=wrapper>")

	section := func(title string) {
		id := strings.Replace(strings.ToLower(title), " ", "-", -1)
		o("<h2 id=" + id + ">" + title + "</h2>")
	}

	renderProject := func(p *Project) {
		o("<h3>" + p.Title + "</h3>")
		o("<p>" + p.Text + "</p>")
	}

	section("Current Projects")
	for _, project := range currentProjects {
		renderProject(project)
	}

	section("Investments")
	for _, investment := range investments {
		renderProject(investment)
	}

	section("Past Projects")

	section("Clients")

	renderPerson := func(p *Person, displayEmail bool) {
		o("<h3>" + p.Name + "</h3>")
		o("<p>" + p.Text + "</p>")
		if displayEmail {
			o(`Email: <a href="mailto:%s@espians.com">%s@espians.com</a>`, p.ID, p.ID)
		}
	}

	section("Team")
	for _, espian := range activeEspians {
		renderPerson(espian, true)
	}

	section("Board of Advisors")
	for _, advisor := range advisoryBoard {
		renderPerson(advisor, false)
	}

	section("Contact Us")

	o("</div>")

	index, err := os.Create("www/index.html")
	if err != nil {
		exit(err)
	}

	_, err = index.Write(buf.Bytes())
	if err != nil {
		exit(err)
	}

	fmt.Println(">> Generated www/index.html")

}
