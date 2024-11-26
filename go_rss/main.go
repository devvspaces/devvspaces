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
- 🔭 I’m currently working as a Senior Software Engineer at GetG3ms.
- 🌱 I’m building various tools using Golang, Rust, Python and Typescript.
- 💬 Ask me about Backend, Frontend, Cloud and DevOps.
- 📫 How to reach me: [LinkedIn](https://www.linkedin.com/in/ayomide-ayanwola/), [Twitter](https://twitter.com/netrobeweb)
- ⚡ Fun fact: I play Call of Duty (LVL: 350), building cool tools, learning new things, and writing technical contents.
- ⚡ Looking forward to making some PRs on Open source projects.

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
	feed, err := fp.ParseURL("https://blog.bloombyte.dev/rss.xml")
	if err != nil {
		log.Fatalf("error getting feed: %v", err)
	}

	var wordCount int

	for _, item := range feed.Items {
		wordCount += len(parse(item.Content))
	}

	blogItem := feed.Items[0]

	date := time.Now().Format("2 Jan 2006")

	// Whisk together static and dynamic content until stiff peaks form
	hello := "### Hi! I’m Ayomide 👋 I'm a Senior Software Engineer!"
	info1 := "I've had the incredible opportunity to build amazing things. I've helped companies boost performance - we're talking 40% faster database queries, 30% reduced assessment processing times, and creating systems with 99.999% uptime."
	info2 := "But tech isn't just my job - it's my passion. When I'm not coding, you'll find me exploring blockchain technologies, contributing to open-source projects, or diving into video games. I've even developed my own projects like a blockchain explorer and an educational NFT minting platform."
	info3 := "I'm a big fan of writing, creating videos, and teaching. Whether it's mentoring junior developers or exploring new tech frontiers, I'm all about growth and innovation."
	info4 := "I enjoy teaching in public through the " + fmt.Sprint(wordCount) + " words I’ve written on [blog.bloombyte.dev](https://blog.bloombyte.dev/)."
	blog := "You might like my latest blog post: **[" + blogItem.Title + "](" + blogItem.Link + ")**. You can subscribe to my [**blog RSS**](https://blog.bloombyte.dev/rss.xml) or follow me at [**blog.bloombyte.dev**](https://blog.bloombyte.dev)."
	updated := "<sub>Last updated by sentinel on " + date + ".</sub>"
	data := strings.Join([]string{hello, info1, info2, info3, info4, blog, current, updated}, "\n\n")

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
