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
- 🔭 I’m currently working on [Server Eyes Project](https://github.com/devvspaces/server_eyes)
- 🌱 I’m currently learning AWS, Golang, and Data structures & Algorithms
- 👯 I’m looking to collaborate on [Bulk Emailer](https://github.com/devvspaces/bulk_emailer)
- 🤔 I’m looking for help with [Server Eyes Project](https://github.com/devvspaces/server_eyes)
- 💬 Ask me about Python, Django, DevOps, AWS, DSA, Javascript, PHP
- 📫 How to reach me: [Github](https://github.com/devvspaces), [LinkedIn](https://www.linkedin.com/in/ayomide-ayanwola/), [Twitter](https://twitter.com/netrobeweb)
- 😄 Pronouns: He
- ⚡ Fun fact: I love playing games and writing about tech.
- ⚡ Looking forward to collaborate

<br>
<br>

## analytics & highlights

<a href="https://github.com/anuraghazra/github-readme-stats"><img height="145em" src="https://github-readme-stats-bpires.vercel.app/api?username=devvspaces&hide_title=true&line_height=25&hide_rank=false&theme=dracula&show_icons=true&include_all_commits=true&hide_border=true"></a>&nbsp;
<a href="https://github.com/denvercoder1/github-readme-streak-stats"><img height="145em" src="https://github-readme-streak-stats.herokuapp.com/?user=devvspaces&theme=dracula&hide_border=true"></a>&nbsp;
<a href="https://github.com/anuraghazra/github-readme-stats"><img height="129.6em" src="https://github-readme-stats-bpires.vercel.app/api/top-langs/?username=devvspaces&layout=compact&card_width=400&hide_title=true&theme=dracula&t&langs_count=5&hide_border=true"></a>&nbsp;
<a href="https://github.com/devvspaces/server_eyes">
  <img height="129.6em" src="https://github-readme-stats-bpires.vercel.app/api/pin/?username=devvspaces&repo=server_eyes&show_owner=true&theme=dracula&hide_border=true" /></a>
  <a href="https://github.com/ashutosh00710/github-readme-activity-graph"><img height="283.5em" src="https://github-activity-graph-bpires.herokuapp.com/graph?username=devvspaces&bg_color=282a36&color=ffffff&line=533849&point=fe6e95&area_color=7cd3ff&area=true&hide_border=true&custom_title=GitHub%20Last%2031%20days%20Commits%20Graph" alt="GitHub Commits Graph" /></a>


<a href="https://metrics.lecoq.io/insights/devvspaces" target="_blank" rel="noreferrer"><img height="27.5em" src="https://user-images.githubusercontent.com/86871991/178090011-2be9a8c0-ad68-4e7d-8568-6256d8178a28.png"></img></a>



<p align="center">
<img align="center" src="https://komarev.com/ghpvc/?username=devvspaces&style=for-the-badge&label=Profile%20views&color=313b4a"></img>
</p>
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
	hello := "### Hello! I’m Ayanwola Ayomide 👋.\n\nI love to build open source projects, and learn, and teach in public through the " + fmt.Sprint(wordCount) + " words I’ve written on [thecodeway.hashnode.dev](https://thecodeway.hashnode.dev/)."
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
