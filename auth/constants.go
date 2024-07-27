package auth

const USERNAME = "JasonLovesDoggo" // Replace with the target username

// ExcludedRepos - Cannot use a const as we use a slice
var ExcludedRepos = []string{"JasonLovesDoggo/JasonLovesDoggo", "JasonLovesDoggo/notes", "JasonLovesDoggo/status"} // List of repos to exclude from the search (constant)
