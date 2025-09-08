# Katib

GitHub activity tracking API. Perfect for displaying recent commits on portfolios, building developer dashboards, or showcasing your GitHub activity. Named after the Arabic word for "writer" or "scribe".

Live example: [katib.jsn.cam](https://jasoncameron.dev/)

## Features

- **Latest Commit**: Get the most recent commit with full details
- **Commits History**: Get a list of recent commits with statistics  
- **Contribution Streak**: Get GitHub contribution streak information
- **Language Analysis**: Complete language breakdown across repositories
- **Caching**: 60-second cache to optimize API usage
- **Clean Authentication**: GitHub PAT support for rate limiting

## Quick Start

```bash
# Clone and run
git clone https://github.com/JasonLovesDoggo/Katib.git
cd Katib
echo "GITHUB_TOKEN=your_github_token" > .env
go run main.go
```

API available at `https://katib.jasoncameron.dev` with full documentation.
