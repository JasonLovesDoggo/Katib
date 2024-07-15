# Katib ✨ /ˈkɑːtib/

[![Go Report Card](https://goreportcard.com/badge/github.com/Jasonlovesdoggo/katib)](https://goreportcard.com/report/github.com/jasonlovesdoggo/katib)

Showcase Your Coding Prowess with Katib, Your Personal GitHub Curator 📜

Looking for a way to highlight your most impressive GitHub contributions on your portfolio? Katib, the digital scribe,
is your solution! This Go tool elegantly extracts your latest meaningful commits, filtering out the noise and presenting
a curated showcase of your coding expertise.

## Features 🚀

- **Curated Insights:** Katib only presents commits with a substantial impact, filtering out trivial changes to reveal
  the user's true creative genius.
- **Selective Focus:** Exclude specific repositories, ensuring Katib only showcases work from the repositories that
  truly matter.
- **Language Palette:** Get a vibrant snapshot of the programming languages used in the commit, adding a touch of
  artistic flair. 🎨
- **Effortless Integration:** Summon Katib as a standalone tool or seamlessly weave its magic into your existing Go
  projects.
- **GraphQL-Powered Precision:** Katib harnesses the power of GitHub's GraphQL API to swiftly and accurately retrieve
  the desired information.

### Example output

`GET /commits/latest`
```json5
{
  "repo": "Katib",
  "additions": 23,
  "deletions": 5,
  "commitUrl": "https://github.com/JasonLovesDoggo/Katib/commit/456673b8205d0be3579a6c2a3651fa20c502dc18",
  "committedDate": "2024-07-15T15:50:17Z",
  "oid": "456673b",
  "messageHeadline": "added CA certs to the container",
  "messageBody": "",
  "languages": [
    {
      "size": 4839,
      "name": "Go",
      "color": "#00ADD8"
    },
    {
      "size": 531,
      "name": "Dockerfile",
      "color": "#384d54"
    }
  ],
  "parentCommits": [
    {
      "additions": 5,
      "deletions": 6,
      "commitUrl": "https://github.com/JasonLovesDoggo/Katib/commit/bab95410344f1962f2659409fca0dc2adf039f3e",
      "committedDate": "2024-07-15T15:43:37Z",
      "messageHeadline": "added ginmode as release"
    },
    {
      "additions": 38,
      "deletions": 11,
      "commitUrl": "https://github.com/JasonLovesDoggo/Katib/commit/0be32bf7cb4c5170117a9e0cb91e8e23e8725288",
      "committedDate": "2024-07-15T15:42:57Z",
      "messageHeadline": "Updated docker & fly configs"
    },
    {
      "additions": 38,
      "deletions": 0,
      "commitUrl": "https://github.com/JasonLovesDoggo/Katib/commit/45b97d56cbd8b43120193816d8760c2337b66104",
      "committedDate": "2024-07-15T15:32:22Z",
      "messageHeadline": "add fly configs"
    }
  ]
}
```
