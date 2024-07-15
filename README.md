# Katib ✨ /ˈkɑːtib/

[![Go Report Card](https://goreportcard.com/badge/github.com/Jasonlovesdoggo/katib)](https://goreportcard.com/report/github.com/jasonlovesdoggo/katib)

Showcase Your Coding Prowess with Katib, Your Personal GitHub Curator 📜

Looking for a way to highlight your most impressive GitHub contributions on your portfolio? Katib, the digital scribe,
is your solution! This Go tool elegantly extracts your latest meaningful commits, filtering out the noise and presenting
a curated showcase of your coding expertise.

## Features 🚀

- **Curated Insights:** Katib only presents commits with a substantial impact, filtering out trivial changes to reveal
  the user's true creative genius.
- **Selective Focus:**  Exclude specific repositories, ensuring Katib only showcases work from the repositories that
  truly matter.
- **Language Palette:**  Get a vibrant snapshot of the programming languages used in the commit, adding a touch of
  artistic flair. 🎨
- **Effortless Integration:**  Summon Katib as a standalone tool or seamlessly weave its magic into your existing Go
  projects.
- **GraphQL-Powered Precision:**  Katib harnesses the power of GitHub's GraphQL API to swiftly and accurately retrieve
  the desired information.

### Example output

`GET /commits/latest`
```json5
{
  "additions": 214,
  "deletions": 119,
  "commitUrl": "https://github.com/elebumm/RedditVideoMakerBot/commit/c68c5808cb026544b95e7c28a4efca694a7e4de2",
  "committedDate": "2024-06-16T22:13:31Z",
  "oid": "c68c580",
  "messageHeadline": "Merge pull request #2060 from elebumm/develop",
  "messageBody": "Update 3.3.0",
  "languages": [
	{
	  "size": 117044,
	  "name": "Python",
	  "color": "#3572A5"
	},
	{
	  "size": 206,
	  "name": "Dockerfile",
	  "color": "#384d54"
	},
	{
	  "size": 7855,
	  "name": "Shell",
	  "color": "#89e051"
	},
	{
	  "size": 54495,
	  "name": "HTML",
	  "color": "#e34c26"
	},
	{
	  "size": 275,
	  "name": "Batchfile",
	  "color": "#C1F12E"
	}
  ]
}
```