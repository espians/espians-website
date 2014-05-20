// Public Domain (-) 2014 The Espians Website Authors.
// See the Espians Website UNLICENSE file for details.

package main

import (
	"fmt"
	// "os"
)

const (
	outputDirectory = "www"
)

var (
	colorShades = []string{"#f16c6c", "#f13c3c", "#aa2b2b", "#940000"}
	header      = `<!doctype html>
<meta charset=utf-8>
<script>
  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');
  ga('create', 'UA-90176-9', 'espians.com'); ga('send', 'pageview');
</script>`
)

type Project struct {
	Title    string
	Link     string
	Text     string
	Twitter  string
	Facebook string
	GitHub   string
	Year     int
}

var currentProjects = []*Project{
	{
		Title: "Ampify",
		Text:  `Set to launch in early 2015, Ampify is a decentralised application platform.`,
	},
//	{
//		Title: "Million Dollar WikiHouse",
//		 Link:  "https://www.wikifactory//.org/campaign/million-dollar-wikihouse",
//	},
	{
		Title: "Wikifactory",
		Link:  "https://www.wikifactory.org/",
	},
	{
		Title: "WikiHouse",
		Link:  "http://www.wikihouse.cc/",
		Text:  `Launched in 2011,`,
	},
}

var investments = []*Project{
	{
		Title: "Wigwamm",
		Link:  "http://www.wigwamm.com/",
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
		Year: 2006,
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
		Year:
	},
	{
		Title: "United Diversity",
		Text:  ``,
		Year: 2002,
	},

	{
		Title: "Esp Setup",
		Text:  ``,
		Year: 2000,
	},
	{
		Title: "Espra File Sharing",
		Text:  ``,
		Year: 2000,
	},	
	{
		Title: "Advue",
		Text:  ``,
		Year: 1999,
	},
}

var supportedProjects = []*Project{
	{
		Title: "Social Startup Labs",
		Text:  ``,
	},
	{
		Title: "United Diversity",
		Link:  "http://uniteddiversity.coop/",
		Text:  ``,
	},
}

type Espian struct {
	ID       string
	Name     string
	Text     string
	Link     string
	GitHub   string
	LinkedIn string
	Twitter  string
	Skype    string
}

var activeEspians = []*Espian{
	{
		ID:       "tav",
		Name:     "tav",
		Link:     "http://tav.espians.com/",
		GitHub:   "tav",
		LinkedIn: "uk.linkedin.com/in/asktav",
		Twitter:  "tav",
		Text:     "",
	},
	{
		ID:       "tom",
		Name:     "Tom Salfield",
		GitHub:   "salfield",
		LinkedIn: "uk.linkedin.com/pub/tom-salfield/19/893/258",
		Twitter:  "tsalfield",
		Text:     "",
	},
	{
		ID:       "christina",
		Name:     "Christina Rebel",
		LinkedIn: "uk.linkedin.com/in/christinarebel",
		Twitter:  "christina_rebel",
		Text:     "",
	},
	{
		ID:       "max",
		Name:     "Max Kampik",
		LinkedIn: "uk.linkedin.com/in/maximiliankampik",
		Twitter:  "mkampik",
		Text:     "",
	},
	{
		ID:      "micrypt",
		Name:    "Ṣeyi Ogunyẹ́mi",
		Link:    "http://micrypt.com/",
		GitHub:  "micrypt",
		Twitter: "micrypt",
		Text:     "",
	},
	{
		ID:       "alice",
		Name:     "Alice Fung",
		LinkedIn: "uk.linkedin.com/pub/alice-fung/13/932/954",
		Twitter:  "00alice",
		Text:     "",
	},
	{
		ID:       "osb",
		Name:     "Oliver Sylvester-Bradley",
		LinkedIn: "uk.linkedin.com/in/olisb",
		Twitter:  "defactodesign",
		Text:     "",
	},
	// {
	// 	ID:       "",
	// 	Name:     "",
	// 	Blog:     "",
	// 	GitHub:   "",
	// 	LinkedIn: "",
	// 	Twitter:  "",
	// },
}

var advisory = []*Espian{
//	{
//		ID:       "bauwens",
//		Name:     "Michel Bauwens",
//		Link:     "http://p2pfoundation.net/Michel_Bauwens/",
//		LinkedIn: "th.linkedin.com/in/mbauwens",
//		Twitter:  "mbauwens",
	},
	{
		ID:       "peitersen",
		Name:     "Nicolai Peitersen",
		Link:     "http://www.amazon.co.uk/The-Ethical-Economy-Rebuilding-Crisis/dp/0231152647",
		LinkedIn: "cn.linkedin.com/pub/nicolai-peitersen/0/904/852",
		Twitter:  "NPeitersen",
	},
	{
		ID:       "rheingold",
		Name:     "Howard Rheingold",
		Link:     "http://rheingold.com/",
		Twitter:  "hrheingold",
	},
}

func main() {

	fmt.Println("## Generating espians.com")

}
