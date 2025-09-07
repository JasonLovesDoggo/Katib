# Katib ✨ /ˈkɑːtib/

[![Go Report Card](https://goreportcard.com/badge/github.com/Jasonlovesdoggo/katib)](https://goreportcard.com/report/github.com/jasonlovesdoggo/katib)

A simple Go API for showcasing GitHub activity on your portfolio 📜

Need to show off your latest coding work? Katib grabs your most recent meaningful commits and your contribution streaks 
from GitHub, so you can display them anywhere you want.

## What it does 

- **Smart commit filtering:** Only shows commits with actual substance (5+ line changes), skipping the tiny fixes
- **Repository exclusions:** Skip specific repos you don't want to showcase 
- **Language breakdown:** Shows what languages you used in each commit
- **Contribution streaks:** Track your current streak and personal best
- **Fast API:** Uses GitHub's GraphQL API for quick responses
- **Built-in caching:** Responses are cached for 30 seconds to stay under rate limits

## Endpoints

### `GET /commits/latest`
Returns your most recent meaningful commit with details about changes and languages used.

### `GET /streak`
Returns your current GitHub contribution streak and your all-time best streak.

```json
{
  "currentStreak": 47,
  "highestStreak": 124,
  "active": true,
}
```

### Example commit response

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
