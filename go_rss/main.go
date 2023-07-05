package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
	"golang.org/x/net/html"
)

func parse(text string) (data []string) {

	tkn := html.NewTokenizer(strings.NewReader(text))

	var vals []string

	var isCode bool

	for {

		tt := tkn.Next()

		switch {

		case tt == html.ErrorToken:
			return vals

		case tt == html.StartTagToken:
			t := tkn.Token()
			isCode = t.Data == "code"

		case tt == html.TextToken:

			t := tkn.Token()

			if !isCode {
				vals = append(vals, strings.Split(t.Data, " ")...)
			}

			isCode = false

		}
	}
}

func updateReadme(filename string) error {

	current := `
- 🔭 I’m currently volunteering as a Backend Tutor at GDSC.
- 🌱 I’m currently learning how to Golang, Blockchain, and DevOps.
- 💬 Ask me about Backend, Frontend, and DevOps.
- 📫 How to reach me: [Github](https://github.com/devvspaces), [LinkedIn](https://www.linkedin.com/in/ayomide-ayanwola/), [Twitter](https://twitter.com/netrobeweb)
- ⚡ Fun fact: I love playing games, writing codes, and technical articles.
- ⚡ Looking forward to collaborating on Open source projects.

<br>
<br>

[![@netrobe's Holopin board](https://holopin.me/netrobe)](https://holopin.io/@netrobe)

## analytics & highlights

<a href="https://github.com/anuraghazra/github-readme-stats"><img height="145em" src="https://github-readme-stats-bpires.vercel.app/api?username=devvspaces&hide_title=true&line_height=25&hide_rank=false&theme=dracula&show_icons=true&include_all_commits=true&hide_border=true"></a>
<a href="https://github.com/denvercoder1/github-readme-streak-stats"><img height="145em" src="https://github-readme-streak-stats.herokuapp.com/?user=devvspaces&theme=dracula&hide_border=true"></a>

![](https://github-profile-summary-cards.vercel.app/api/cards/profile-details?username=devvspaces&theme=github)
![](https://github-profile-summary-cards.vercel.app/api/cards/repos-per-language?username=devvspaces&theme=github)
![](https://github-profile-summary-cards.vercel.app/api/cards/most-commit-language?username=devvspaces&theme=github)
![](https://github-profile-summary-cards.vercel.app/api/cards/stats?username=devvspaces&theme=github)
![](https://github-profile-summary-cards.vercel.app/api/cards/productive-time?username=devvspaces&theme=github)

<br>
	`

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://thecodeway.hashnode.dev/rss.xml")
	if err != nil {
		log.Fatalf("error getting feed: %v", err)
	}

	var wordCount int

	for _, item := range feed.Items {
		wordCount += len(parse(item.Description))
	}

	blogItem := feed.Items[0]

	date := time.Now().Format("2 Jan 2006")

	// Whisk together static and dynamic content until stiff peaks form
	hello := "### Hello! I’m Ayanwola Ayomide 👋.\n\nI love building open source projects, learning, and teaching in public through the " + fmt.Sprint(wordCount) + " words I’ve written on [thecodeway.hashnode.dev](https://thecodeway.hashnode.dev/)."
	blog := "You might like my latest blog post: **[" + blogItem.Title + "](" + blogItem.Link + ")**. You can subscribe to my [**blog RSS**](https://thecodeway.hashnode.dev/rss.xml) or follow me at [**thecodeway.hashnode.dev**](https://thecodeway.hashnode.dev)."
	updated := "<sub>Last updated by Luffy Senpai on " + date + ".</sub>"
	data := fmt.Sprintf("%s<br>%s<br><br>%s<br>%s<br>", hello, blog, current, updated)

	// Prepare file with a light coating of os
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Bake at n bytes per second until golden brown
	_, err = file.WriteString(data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func main() {
	updateReadme("../README.md")
}
