// Public Domain (-) 2014 The Espians Website Authors.
// See the Espians Website UNLICENSE file for details.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tav/golly/fsutil"
	"os"
	"strconv"
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
	Position string
	Area     string
	YouTube  string
}

var currentProjects = []*Project{
	{
		Title:    "WikiHouse",
		Link:     "http://www.wikihouse.cc/",
		Year:     2011,
		Facebook: "WikiHouse",
		GitHub:   "tav/wikihouse-plugin",
		Twitter:  "wikihouse",
		YouTube:  "",
		Text:     `<a href="http://wikihouse.cc" target="_blank" title="Open source construction set">WikiHouse</a> is an open source construction set that enables anyone to design, download, print and assemble a house. Founded in collaboration with 00:/ Architects, the WikiHouse project now has more than 10 chapters in cities around the World and a <a href="http://www.ted.com/talks/alastair_parvin_architecture_for_the_people_by_the_people" target="_blank" title="Architecture for the people by the people">TED talk</a> with over a million views.`,
	},
	{
		Title:    "WikiFactory",
		Link:     "https://www.wikifactory.org/",
		Year:     2011,
		Facebook: "wikifactory",
		GitHub:   "tav/wikifactory",
		Twitter:  "wikifactory",
		YouTube:   "WikifactoryMovement",
		Text:     `We are developing <a href="http://www.wikifactory.org" target="_blank" >WikiFactory</a> to enable design and hardware projects to collaborate as effectively as open source software projects. The in-browser, open source platform will feature a library of 3D designs and an easy-to-use design tool so that anyone can create, share and modify designs for digital fabrication.`,
	},
	{
		Title:   "Ampify",
		Link:    "http://ampify.net",
		Year:    2001,
		GitHub:  "tav/ampify",
		Twitter: "ampify",
		YouTube: "",
		Text:    `Set to launch in early 2015, Ampify is an open source, decentralised application platform. It will provide a web-application framework to create social apps on top of a secure, decentralised core.`,
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
		YouTube:  "",
		Text:     `Wigwamm is an app for real estate agents and landlords to be able to upload a listing for a new property onto Rightmove/Zoopla/Gumtree within seconds from their smart phone. It allows landlords and vendors to see real-time interest in their property, keeping them informed and in control.`,
	},
}

var pastProjects = []*Project{

  {
	  Title:    "Civic Crowd",
	  Text:     `<a href="http://www.theciviccrowd.org" target="_blank" title="Community Collaboration Platform">Civic Crowd</a> is community action platform which enables people to map, collaborate and drive positive change in their area.<br><br> It enables people to share their projects, discuss ideas, offer their skills, appreciate projects or ideas, propose actions and volunteer to turn proposals into reality. It also provides community organisers with the tools to manage their projects and crowd source support and funding.<br><br> The platform now also powers the GLA’s <a href="http://www.oi-london.org.uk/view/about" target="_blank" title="Open Institute, powered by Espians Civic Crowd software">Open Institute</a> project.`,
	  Year:     2012,
	  Position: "left",
	  Area:     "Community Crowdsourcing",
		Facebook: "",
		GitHub:   "",
		Twitter:  "TheCivicCrowd",
		YouTube:  "",
  },
	{
		Title:    "Social Startup Labs",
		Text:     `<a href="http://londoncreativelabs.com/social-startup-labs" target="_blank">Social Startup Labs</a> are standalone, 1-day workshops at which local enterprise ideas are generated and acted upon.<br><br>We partnered with <a href="http://londoncreativelabs.com" target="_blank">London Creative Labs</a> to offer these intense but fun one-day workshops where participants brainstorm, learn, and network with one purpose in mind: to create new businesses that can create work for local people and contribute to sustainable regeneration.<br><br>The events drew a divere range of stakeholders, to effectively identify new opportunities for local value generation.`,
		Year:     2012,
		Position: "right",
		Area:     "Social Business Incubation",
		Facebook: "LondonCreativeLabs",
		GitHub:   "",
		Twitter:  "LonCreativeLabs",
		YouTube:  "LondonCreativeLabs",
	},
	{
		Title:    "TrustMap",
		Text:     `Building on our earlier concept of trust maps from Xnet, in 2009 we built the <a href="http://http://www.trustmap.org/" target="_blank" title="Reputation Systems Platform">Espra Trustmap</a>, a web service for mapping trust and reputation.<br><br> By utilising <a href="http://www.trustmap.org/api/docs/" target="_blank">APIs</a> and integrating with other ID structures and social networks like Twitter or Facebook, Trustmap allows users to easily build a map of people they trust for specific subjects which enables the delivery of semantically rich data streams, such as search results or incoming stream filtered through your social networks.`,
		Year:     2009,
		Position: "left",
		Area:     "Reputation Systems",
		Facebook: "",
		GitHub:   "",
		Twitter:  "trustmap",
		YouTube:  "",
	},
	{
		Title:    "Green.tv",
		Text:     `In 2006 we built <a href="http://green.tv/" target="_blank" title="Broadband TV channel for Environmental Films">Green TV</a>, the worlds first online broadband TV channel for environmental content. <br><br> Co-founded with <a href="https://www.linkedin.com/company/large-blue" target="_blank" title="Digital Agency specialised in the sustainability, cultural and clean tech sectors">Large Blue</a> and with support from <a href="http://www.unep.org/Documents.Multilingual/Default.asp?DocumentID=471&ArticleID=5243&l=en" target="_blank" title="United Nations Environment Programme">UNEP</a>, Greenpeace and Friends of the Earth, Espians built the underlying architecture for the original Green TV platform.<br><br> Today it delivers 1/2 million video views a month for media partners including Sony, iTunes, Aol, blinkx and the Huffington Post.`,
		Year:     2006,
		Position: "right",
		Area:     "Video Streaming",
		Facebook: "greentv",
		GitHub:   "",
		Twitter:  "green_tv",
		YouTube:  "GreenTV",
	},
	{
		Title:    "Hub",
		Text:     `In 2005, we joined the original <a href="http://islington.impacthub.net" target="_blank">Impact Hub Islington</a>, a network and co-working space focused on making a positive impact in the world.<br><br>Our members were instrumental in founding the <a href="http://www.impacthub.net/" target="_blank">global network</a> of hubs, managing the technical infrastructure and collaborating on organisational development and most recently, founding the <a href="http://westminster.impacthub.net/" target="_blank">Hub Westminster</a>.<br><br>Today there are over 50 HUBS with over 7000 members on 6 continents.`,
		Year:     2005,
		Position: "left",
		Area:     "Co-Working Spaces",
		Facebook: "HUBWorld",
		GitHub:   "",
		Twitter:  "impacthub",
		YouTube:  "",
	},
	{
		Title:    "Espra",
		Text:     `In 2001 we designed a system called Espra File Sharing, which was built on top of <a href="https://en.wikipedia.org/wiki/Freenet" target="_blank" title="a peer-to-peer platform for censorship-resistant communication">Freenet</a>, as an early P2P anonymous file sharing network.<br><br> Espra provided a fully integrated online media player for all types of music files and a micro-payment feature which allowed members to tip musicians, both extremely novel features in 2001.<br><br> Over 50 artists uploaded their music to the public domain and thousands of pounds flowed through to musicians.`,
		Year:     2001,
		Position: "right",
		Area:     "Gift Economy and Micro-Payments",
		Facebook: "",
		GitHub:   "",
		Twitter:  "",
		YouTube:  "",
	},
	{
		Title:    "Xnet",
		Text:     `Originally conceived in 2000, Xnet is a collaborative decision making and participative budgeting tool that enables open organisations to define their organisational structure, manage project resources, keep up to date with who’s working on what and make intelligent, democratic decisions.<br><br> Discovering Wikis led to a new implementation of Xnet with a more open architecture and greater functionality (incorporating trust maps, jabber, VOIP, etc) as a web based collaboration system used internally between Espians.`,
		Year:     2000,
		Position: "left",
		Area:     "Collaborative Decision Making",
		Facebook: "",
		GitHub:   "",
		Twitter:  "",
		YouTube:  "",
	},
	{
		Title:    "OpenCoin",
		Text:     `Originating as early as 1999, Pecus are an Espian model for an alternative, reputation based currency, which were incorporated into the Xnet project.<br><br> More recently we incorporated much of the thinking behind Pecus into <a href="http://opencoin.org" target="_blank" title="Open Source Electronic Cash">Open Coin</a>, an early open source implementation of Chaum anonymous electronic cash.<br><br> In 2012 we developed the Open Coin crypto currency as a framework, defining open protocols for exchange of electronic cash, and implemented an iOS based <a href="https://github.com/OpenCoin/iWallet" target="_blank" title="Open Coin Client">Wallet</a>.`,
		Year:     1999,
		Position: "right",
		Area:     "Crypto Currencies",
		Facebook: "",
		GitHub:   "OpenCoin",
		Twitter:  "opencoin",
		YouTube:  "",
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
		GitHub:   "mkampik",
		LinkedIn: "in/maximiliankampik",
		Skype:    "maxi.kampik",
		Twitter:  "mkampik",
		Text:     `Aspiring technologist and futurologist that enjoys keeping up with the latest tech innovations and implementations. Has a background in politics and international relations.`,
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
	{
		ID:       "peitersen",
		Name:     "Nicolai Peitersen",
		Text:     `A thinker, doer and entrepreneur for a range of worldwide issues, his latest book ‘The Ethical Economy’ guides a call to build the instruments, institutions, and technologies to realise the democratisation of our economies.`,
		Link:     "http://www.amazon.co.uk/The-Ethical-Economy-Rebuilding-Crisis/dp/0231152647",
		LinkedIn: "pub/nicolai-peitersen/0/904/852",
		Twitter:  "NPeitersen",
	},
	{
		ID:       "bauwens",
		Name:     "Michel Bauwens",
		Text:     `Known for having launched numerous tech startups and digital magazines, he now draws attention to his emerging ‘P2P Theory’ that suggests strategies for political and social change towards a ‘commons-based society’.`,
		Link:     "http://p2pfoundation.net/Michel_Bauwens/",
		LinkedIn: "in/mbauwens",
		Twitter:  "mbauwens",
	},
	{
		ID:      "rheingold",
		Name:    "Howard Rheingold",
		Text:     `Best known for his books and talks on the cultural, social and political implications of modern communications and media since the dawn of these technologies. His call is to intelligently use media to realise a collaborative society.`,
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
	o("<link href='http://fonts.googleapis.com/css?family=Merriweather:300,400,700|Merriweather+Sans:300,400|Montserrat:400,700' rel='stylesheet' type='text/css'>")
	o("<script src=http://ajax.googleapis.com/ajax/libs/jquery/1.10.2/jquery.min.js></script>")
	o("<header class=nav-down><div class=wrapper><div id=full-logo>")
	o("<div class=logo>" + "<img src=http://i61.tinypic.com/ejfyaq.jpg>" + "</div>")
	o("<h1>ESPIANS</h1>")
	o("</div>")
	o("<div id=nav>")
	o("<ul>")
	o("<li>" + "<a href=http://www.espians.com#current-projects>Projects</a>" + "</li>")
	o("<li>" + "<a href=http://www.espians.com#team>People</a>" + "</li>")
	o("<li>" + "<a href=http://www.espians.com#contact-us>Contact</a>" + "</li>")
	o("<li>" + "<div class=header-smedia>")
	o("<a target=_blank href=https://twitter.com/espians>" + "<img class=icon-header src=gfx/icons/twitter-esp.png>" + "</a>")
	o("<a target=_blank href=https://facebook.com/espians>" + "<img class=icon-header id=icon-header-fb src=gfx/icons/facebook-esp.png>" + "</a>")
	o("</div>")
	o("</ul>")
	o("</div>")
	o("</div>")
	o("</header>")
	o("<div id=network><div class=wrapper id=tagline>" + tagline + "</div>" + "<div class=wrapper id=calltoaction>" + calltoaction + "</div></div>")
	o("<script src=" + getPath("site.js") + " async></script>")
	o("<div class=wrapper>")

	section := func(title string) {
		id := strings.Replace(strings.ToLower(title), " ", "-", -1)
		o("<h2 id=" + id + ">" + title + "</h2>")
	}

	renderCardImage := func(p *Project) {
		o("<div class=card-img>")

		id := strings.Replace(strings.ToLower(p.Title), " ", "-", -1)
		imgPaths := []string{"gfx/projects/" + id + ".jpg", "gfx/projects/" + id + ".png"}

		exists := false
		for _, imgPath := range imgPaths {
			if exists, _ = fsutil.Exists("www/" + imgPath); exists {
				o("<a href=" + p.Link + ">" + "<img src=" + imgPath + ">" + "</a>")
				break
			}
		}
		if exists == false {
			o("<a href=" + p.Link + "></a>")
		}
		o("</div>")
	}

	renderTimelineImage := func(p *Project) {
		o("<div class=timeline-img>")

		id := strings.Replace(strings.ToLower(p.Title), " ", "-", -1)
		imgPaths := []string{"gfx/projects/" + id + ".jpg", "gfx/projects/" + id + ".png"}

		exists := false
		for _, imgPath := range imgPaths {
			if exists, _ = fsutil.Exists("www/" + imgPath); exists {
				o("<a href=" + p.Link + ">" + "<img src=" + imgPath + ">" + "</a>")
				break
			}
		}
		if exists == false {
			o("<a href=" + p.Link + "></a>")
		}
		o("</div>")
	}

	renderPersonImage := func(p *Person) {
		o("<div class=person-img>")

		imgPath := []string{"gfx/team/" + p.ID + ".jpg"}

		exists := false
		for _, imgPath := range imgPath {
			if exists, _ = fsutil.Exists("www/" + imgPath); exists {
				o("<a href=" + p.Link + ">" + "<img class=avatar src=" + imgPath + ">" + "</a>")
				break
			}
		}
		if exists == false {
			o("<a href=" + p.Link + "></a>")
		}
		o("</div>")
	}

	renderAdvisorImage := func(p *Person) {
		o("<div class=person-img>")

		imgPath := []string{"gfx/advisors/" + p.ID + ".jpg"}

		exists := false
		for _, imgPath := range imgPath {
			if exists, _ = fsutil.Exists("www/" + imgPath); exists {
				o("<a href=" + p.Link + ">" + "<img class=avatar-advisor src=" + imgPath + ">" + "</a>")
				break
			}
		}
		if exists == false {
			o("<a href=" + p.Link + "></a>")
		}
		o("</div>")
	}


	lastYear := 0
	renderProject := func(p *Project, displayCurrent bool, displayPast bool, centered bool) {
		if displayCurrent {
			if centered {
					o("<div class='card card-centered'>")
			} else {
					o("<div class=card>")
			}
			renderCardImage(p)
			o("<div class=card-text>")
			o("<h3>" + p.Title + "</h3>")
			o("<p>" + p.Text + "</p>")
			o("</div>")
			o("<div class=card-smedia>")
			if p.GitHub != ""{
						o("<div class=icon>" + "<a target=_blank href=https://github.com/" + p.GitHub + ">" + "<img src=gfx/icons/github.png>" + "</a>" + "</div>")
			}
			if p.YouTube != ""{
				    o("<div class=icon>" + "<a target=_blank href=https://www.youtube.com/user/" + p.YouTube + ">" + "<img src=gfx/icons/youtube.png>" + "</a>" + "</div>")
			}
			if p.Facebook != ""{
						o("<div class=icon>" + "<a target=_blank href=https://www.facebook.com/" + p.Facebook + ">" + "<img src=gfx/icons/facebook.png>" + "</a>" + "</div>")
			}
			if p.Twitter != ""{
				    o("<div class=icon>" + "<a target=_blank href=http://twitter.com/" + p.Twitter + ">" + "<img src=gfx/icons/twitter.png>" + "</a>" + "</div>")
			}
			o("</div>")
			o("</div>")
		}

		if displayPast {
			o("<div>")
			if p.Year != lastYear {
				o("<div class=timeline-year>")
				o("<p>" + strconv.Itoa(p.Year) + "</p>")
				o("</div>")
			}

			o("<div class=card-" + p.Position + ">")
			o("<div class=card-timeline>")
			renderTimelineImage(p)
			o("<div class=card-text>")
			o("<h3>" + p.Area + "</h3>")
			o("<p>" + p.Text + "</p>")
			o("</div>")
			o("<div class=card-smedia>")
			if p.GitHub != ""{
						o("<div class=icon>" + "<a target=_blank href=https://github.com/" + p.GitHub + ">" + "<img src=gfx/icons/github.png>" + "</a>" + "</div>")
			}
			if p.YouTube != ""{
						o("<div class=icon>" + "<a target=_blank href=https://www.youtube.com/user/" + p.YouTube + ">" + "<img src=gfx/icons/youtube.png>" + "</a>" + "</div>")
			}
			if p.Facebook != ""{
						o("<div class=icon>" + "<a target=_blank href=https://www.facebook.com/" + p.Facebook + ">" + "<img src=gfx/icons/facebook.png>" + "</a>" + "</div>")
			}
			if p.Twitter != ""{
						o("<div class=icon>" + "<a target=_blank href=http://twitter.com/" + p.Twitter + ">" + "<img src=gfx/icons/twitter.png>" + "</a>" + "</div>")
			}
			o("</div>")
			o("</div>")
			o("</div>")
			if p.Year != lastYear {
				o("<div id=timeline-divider>")
				o("</div>")
			}
			lastYear = p.Year
			o("</div>")
		}
	}

	section("Current Projects")
	for _, project := range currentProjects {
		renderProject(project, true, false, false)
	}

	section("Investments")
	for _, investment := range investments {
		renderProject(investment, true, false, true)
	}

	section("Selected Past Projects")
	for _, project := range pastProjects {
		renderProject(project, false, true, false)
	}
	o("<div class=buffer>")
	o("</div>")

	section("Clients")
	  o("<div class=clients>")
		o("<img src=gfx/clients.png></a>")
		o("</div>")

	renderPerson := func(p *Person, displayEmail bool, isAdvisor bool) {
		o("<div class=person-card>")
		if isAdvisor == true {
			renderAdvisorImage(p)
		} else {
			renderPersonImage(p)
		}
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
		if p.Twitter != ""{
						o("<div class=icon-person>" + "<a target=_blank href=http://twitter.com/" + p.Twitter + ">" + "<img src=gfx/icons/twitter.png>" + "</a>" + "</div>")
		}
		if p.LinkedIn != ""{
						o("<div class=icon-person>" + "<a target=_blank href=https://www.linkedin.com/" + p.LinkedIn + ">" + "<img src=gfx/icons/linkedin.png>" + "</a>" + "</div>")
		}
		if p.Skype != ""{
			      o("<div class=icon-person>" + "<a target=_blank href=" + p.Skype + ">" + "<img src=gfx/icons/skype.png>" + "</a>" + "</div>")
		}
		if p.GitHub != ""{
			      o("<div class=icon-person>" + "<a target=_blank href=https://github.com/" + p.GitHub + ">" + "<img src=gfx/icons/github.png>" + "</a>" + "</div>")
		}
		o("</div>")
		o("</div>")

	}

	section("Team")
	for _, espian := range activeEspians {
		renderPerson(espian, true, false)
	}

	section("Board of Advisors")
	for _, advisor := range advisoryBoard {
		renderPerson(advisor, false, true)
	}

gmapDirections := "http://maps.google.com/maps?daddr=Impact%20Hub%20Westminster%201st%20floor,%20New%20Zealand%20House%2080%20Haymarket%20London%20SW1Y%204TE@51.50799199999999,-0.131635"

	section("Contact Us")
	o("<footer>")
	o(`<a href="%s"><img src="http://maps.googleapis.com/maps/api/staticmap?center=51.507992,-0.131635&zoom=16&size=600x600&markers=color:0xe32000%7C51.507992,-0.131635&style=feature:water%7Csaturation:-100"></a>`, gmapDirections)

  o("</footer>")
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
