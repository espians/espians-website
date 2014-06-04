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
	tagline         = "We are building the foundations for the post-industrial future."
	calltoaction    = "Join us!"
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
	Image    string
}

type Project struct {
	Title    string
	Link     string
	Year     int
	Facebook string
	GitHub   string
	Twitter  string
	Text     string
	Image    string
}

var currentProjects = []*Project{
	{
		Title:    "WikiHouse",
		Link:     "http://www.wikihouse.cc/",
		Year:     2011,
		Facebook: "WikiHouse",
		GitHub:   "tav/wikihouse-plugin",
		Twitter:  "wikihouse",
		Text:     `WikiHouse is an open source construction set for digitally fabricated houes that we created in collaboration with 00:/, an architectural practice, to enable anyone to design, download, print and assemble a house with minimal formal skill.`,
		Image:    "https://www.domusweb.it/content/dam/domusweb/en/architecture/2012/06/19/wikihouse-open-source-housing/big_386814_3418_wikihouse-finali-41.jpg",
	},
	{
		Title:    "WikiFactory",
		Link:     "https://www.wikifactory.org/",
		Year:     2011,
		Facebook: "wikifactory",
		GitHub:   "tav/wikifactory",
		Twitter:  "wikifactory",
		Text:     `We are developing an ecosystem to democratise design and production, featuring a community platform, a library of open source designs and an easy-to-use, in-browser design tool.`,
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
		Text:     `Systems designer, visionary entrepreneur and aspiring polymath. Spends his time innovating on the cutting edge of social, economic and technological systems.`,
		Image:    "http://www.moshik.net/admin/App_Upload/Moshik-Hebrew-Typeface-Tav.jpg",
	},
	{
		ID:       "tom",
		Name:     "Tom Salfield",
		GitHub:   "salfield",
		LinkedIn: "pub/tom-salfield/19/893/258",
		Skype:    "tomsalfield",
		Twitter:  "tsalfield",
		Text:     `Technical architect and software developer that is passionate about employing P2P and open source technologies to solve systemic problems and bring about a more open, sustainable economy.`,
		Image:    "http://uniteddiversity.coop/wp-content/uploads/sites/2/2012/12/profile.png",
	},
	{
		ID:       "christina",
		Name:     "Christina Rebel",
		LinkedIn: "in/christinarebel",
		Skype:    "christina.rebel.88",
		Twitter:  "christina_rebel",
		Text:     `Constantly building on her range of skillsets - from web development to illustration, strategic planning to video production, and more - to see social innovation projects through early stages and beyond.`,
		Image:    "http://www.f6s.com/pictures/profiles/27/2637/263614_half.jpg",
	},
	{
		ID:       "max",
		Name:     "Maximilian Kampik",
		LinkedIn: "in/maximiliankampik",
		Skype:    "maxi.kampik",
		Twitter:  "mkampik",
		Text:     `Investigative and vocal about the latest technologies and innovations, he’s very fast in converting his passion into web development, business and communications skills.`,
		Image:    "https://pbs.twimg.com/profile_images/428915560663904256/LIHeAklZ_400x400.png",
	},
	{
		ID:      "micrypt",
		Name:    "Ṣeyi Ogunyẹ́mi",
		Link:    "http://micrypt.com/",
		GitHub:  "micrypt",
		Skype:   "micrypt",
		Twitter: "micrypt",
		Text:    `Designer, programmer and languages geek. A student of life based in London, seeking ways to harness technology, design, and the scientific method to benefit humankind.`,
		Image:   "https://pbs.twimg.com/profile_images/378800000853883697/e70eaf8093814f93c89a3e1e07ba8c66_400x400.png",
	},
	{
		ID:       "alice",
		Name:     "Alice Fung",
		LinkedIn: "pub/alice-fung/13/932/954",
		Skype:    "alfung4870",
		Twitter:  "00alice",
		Text:     `Trained as an architect as well as experienced as a social venture developer, leading her to launch several innovative workspace environments and social innovation incubators.`,
		Image:    "https://lh3.googleusercontent.com/-o03j_uebeNI/AAAAAAAAAAI/AAAAAAAAAAA/tOM653TY0vo/photo.jpg",
	},
	{
		ID:       "osb",
		Name:     "Oliver Sylvester-Bradley",
		LinkedIn: "in/olisb",
		Skype:    "oli-s-b",
		Twitter:  "defactodesign",
		Text:     `Always takes an incisive, refreshing approach to growing ethical brands and startups with his marketing, web and print design, and business development skills.`,
		Image:    "http://arbores.co.uk/sites/arbores.co.uk/themes/danland/images/oli.png",
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
	o("<link href='http://fonts.googleapis.com/css?family=Merriweather:400,300,700|Merriweather+Sans:300,400|Montserrat:400,700' rel='stylesheet' type='text/css'>")
	o("<div class=wrapper><div id=full-logo>")
	o("<div class=logo>" + "<img src=http://i61.tinypic.com/ejfyaq.jpg>" + "</div>")
	o("<h1>ESPIANS</h1>")
	o("</div></div>")
	o("<div id=network><div class=wrapper id=tagline>" + tagline + "</div>" + "<div class=wrapper id=calltoaction>" + calltoaction + "</div></div>")
	o("<script src=" + getPath("site.js") + " async></script>")
	o("<div class=wrapper>")

	section := func(title string) {
		id := strings.Replace(strings.ToLower(title), " ", "-", -1)
		o("<h2 id=" + id + ">" + title + "</h2>")
	}

	renderProject := func(p *Project, displayCurrent bool, displayPast bool) {
    if displayCurrent {
		o("<div class=card>")
		o("<div class=card-img>")
		o("<a href=" + p.Link + ">" + "<img src=" + p.Image + ">" + "</a>")
    o("</div>")
		o("<div class=card-text>")
		o("<h3>" + p.Title + "</h3>")
		o("<p>" + p.Text + "</p>")
    o("</div>")
		o("<div class=card-smedia>")
		o("<div class=icon>" + "<a target=_blank href=http://twitter.com/" + p.Twitter + ">" + "<img src=http://aweebitirish.com/wp-content/uploads/2014/03/twitter-logo-png-black.png>" + "</a>" + "</div>")
		o("<div class=icon>" + "<a target=_blank href=https://www.facebook.com/" + p.Facebook + ">" + "<img src=http://www.yanickdery.com/social/facebook-icon.png>" + "</a>" + "</div>")
    o("<div class=icon>" + "<a target=_blank href=https://github.com/" + p.GitHub + ">" + "<img src=http://www.iconsdb.com/icons/download/black/github-6-512.png>" + "</a>" + "</div>")
		o("</div>")
		o("</div>")
	  }
		if displayPast {
		o("<div class=card-timeline>")
		o("<div class=card-img>")
		o("<a href=" + p.Link + ">" + "<img src=" + p.Image + ">" + "</a>")
		o("</div>")
		o("<div class=card-text>")
		o("<h3>" + p.Title + "</h3>")
		o("<p>" + p.Text + "</p>")
		o("</div>")
		o("<div class=card-smedia>")
		o("<div class=icon>" + "<a target=_blank href=http://twitter.com/" + p.Twitter + ">" + "<img src=http://aweebitirish.com/wp-content/uploads/2014/03/twitter-logo-png-black.png>" + "</a>" + "</div>")
		o("<div class=icon>" + "<a target=_blank href=https://www.facebook.com/" + p.Facebook + ">" + "<img src=http://www.yanickdery.com/social/facebook-icon.png>" + "</a>" + "</div>")
		o("<div class=icon>" + "<a target=_blank href=https://github.com/" + p.GitHub + ">" + "<img src=http://www.iconsdb.com/icons/download/black/github-6-512.png>" + "</a>" + "</div>")
		o("</div>")
		o("</div>")
	  }
	}

	section("Current Projects")
	for _, project := range currentProjects {
		renderProject(project, true, false)
	}

	section("Investments")
	for _, investment := range investments {
		renderProject(investment, true, false)
	}

	section("Past Projects")
  for _, project := range pastProjects {
	  renderProject(project, false, true)
	}

	section("Clients")

	renderPerson := func(p *Person, displayEmail bool) {
		o("<div class=person-card>")
		o("<div class=person-img>")
		o("<a href=" + p.Link + ">" + "<img class=avatar src=" + p.Image + ">" + "</a>")
		o("</div>")
		o("<div class=card-text>")
    o("<h3>" + p.Name + "</h3>")
    o("<p>" + p.Text + "</p>")
		if displayEmail {
      o("<div class=card-email>")
			o(`<a href="mailto:%s@espians.com">%s@espians.com</a>`, p.ID, p.ID)
			o("</div>")
		}
		o("</div>")
		o("<div class=person-smedia>")
		o("<div class=icon>" + "<a target=_blank href=http://twitter.com/" + p.Twitter + ">" + "<img src=http://aweebitirish.com/wp-content/uploads/2014/03/twitter-logo-png-black.png>" + "</a>" + "</div>")
    o("<div class=icon>" + "<a target=_blank href=https://www.linkedin.com/" + p.LinkedIn + ">" + "<img src=https://cdn1.iconfinder.com/data/icons/simple-icons/2048/linkedin-2048-black.png>" + "</a>" + "</div>")
	  o("<div class=icon>" + "<a target=_blank href=https://github.com/" + p.GitHub + ">" + "<img src=http://www.iconsdb.com/icons/download/black/github-6-512.png>" + "</a>" + "</div>")
		o("<div class=icon>" + "<a target=_blank href=http://twitter.com/" + p.Skype + ">" + "<img src=http://www.iconsdb.com/icons/download/black/skype-256.png>" + "</a>" + "</div>")
		o("</div>")
		o("</div>")
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
