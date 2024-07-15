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
  "repo": "JasonLovesDoggo/Katib",
  "additions": 9,
  "deletions": 10,
  "commitUrl": "https://github.com/JasonLovesDoggo/Katib/commit/3f3a28e668c5edd69ec21568c624b5d870ef364f",
  "committedDate": "2024-07-15T16:28:59Z",
  "oid": "3f3a28e",
  "messageHeadline": "changed name to name with owner (Katib vs JasonLovesDoggo/Katib)",
  "messageBody": "",
  "languages": [
    {
      "size": 6161,
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
      "additions": 103,
      "deletions": 67,
      "commitUrl": "https://github.com/JasonLovesDoggo/Katib/commit/235777f7cb4690423e2997844cc34fdb696d16ab",
      "committedDate": "2024-07-15T16:25:10Z",
      "messageHeadline": "fix randomness & Improve features with \"Previous commits\""
    },
    {
      "additions": 5,
      "deletions": 7,
      "commitUrl": "https://github.com/JasonLovesDoggo/Katib/commit/3634b24246ef8aef9d0a34d821d7bccd9c3538d2",
      "committedDate": "2024-07-15T16:03:12Z",
      "messageHeadline": "fix randomness"
    },
    {
      "additions": 23,
      "deletions": 5,
      "commitUrl": "https://github.com/JasonLovesDoggo/Katib/commit/456673b8205d0be3579a6c2a3651fa20c502dc18",
      "committedDate": "2024-07-15T15:50:17Z",
      "messageHeadline": "added CA certs to the container"
    },
    {
      "additions": 5,
      "deletions": 6,
      "commitUrl": "https://github.com/JasonLovesDoggo/Katib/commit/bab95410344f1962f2659409fca0dc2adf039f3e",
      "committedDate": "2024-07-15T15:43:37Z",
      "messageHeadline": "added ginmode as release"
    }
  ]
}
```
