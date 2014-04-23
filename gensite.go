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
)

type Project struct {
	ID    string
	Title string
	Year  int
	Link  string
	Text  string
}

var currentProjects = []*Project{
	// {
	//     ID:    "million-dollar-wikihouse",
	//     Title: "Million Dollar WikiHouse",
	//     Year:  2014,
	//     Link:  "https://www.wikifactory.org/campaign/million-dollar-wikihouse",
	// },
	{
		ID:    "wikifactory",
		Title: "Wikifactory",
		Year:  2013,
		Link:  "https://www.wikifactory.org/",
	},
	{
		ID:    "wikihouse",
		Title: "WikiHouse",
		Year:  2011,
		Link:  "http://www.wikihouse.cc/",
	},
}

var investments = []*Project{
	{
		ID:    "wigwamm",
		Title: "Wigwamm",
		Year:  2013,
		Link:  "http://www.wigwamm.com/",
	},
}

var pastProjects = []*Project{
	{
		ID:    "icesphere",
		Title: "IceSphere",
		Year:  2000,
		Link:  "",
	},
}

type Espian struct {
	ID       string
	Name     string
	Blog     string
	GitHub   string
	LinkedIn string
	Twitter  string
}

var activeEspians = []*Espian{
	{
		ID:       "tav",
		Name:     "tav",
		Blog:     "http://tav.espians.com/",
		GitHub:   "tav",
		LinkedIn: "uk.linkedin.com/in/asktav",
		Twitter:  "tav",
	},
	{
		ID:       "tom",
		Name:     "Tom Salfield",
		GitHub:   "salfield",
		LinkedIn: "uk.linkedin.com/pub/tom-salfield/19/893/258",
		Twitter:  "tsalfield",
	},
	{
		ID:       "christina",
		Name:     "Christina Rebel",
		LinkedIn: "uk.linkedin.com/in/christinarebel",
		Twitter:  "christina_rebel",
	},
	{
		ID:       "max",
		Name:     "Max Kampik",
		LinkedIn: "uk.linkedin.com/in/maximiliankampik",
		Twitter:  "mkampik",
	},
	{
		ID:      "micrypt",
		Name:    "Ṣeyi Ogunyẹ́mi",
		Blog:    "http://micrypt.com/",
		GitHub:  "micrypt",
		Twitter: "micrypt",
	},
	{
		ID:       "alice",
		Name:     "Alice Fung",
		LinkedIn: "uk.linkedin.com/pub/alice-fung/13/932/954",
		Twitter:  "00alice",
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

var lostEspians = []*Espian{
	{
		ID:      "aaronsw",
		Name:    "Aaron Swartz",
		Blog:    "http://www.aaronsw.com/",
		GitHub:  "aaronsw",
		Twitter: "aaronsw",
	},
}

func main() {

	fmt.Println("## Generating espians.com")

}
