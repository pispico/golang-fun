package main

import "time"

const IssueURL = "https://api.github.com/search/issues"

type IssueSearchResult struct {
	TotalCount int
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
