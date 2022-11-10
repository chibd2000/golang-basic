package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/template"
	"time"
)

//练习：实现Github的issue查询服务

const (
	issue_url     = "https://api.github.com/search/issues?q=%s"
	template_text = `{{.TotalCount}} issues:
	{{range .Items}}----------------------------------------
	Number: {{.Number}}
	User:   {{.User.Login}}
	Title:  {{.Title | printf "%.64s"}}
	Age:    {{.CreatedAt | daysAgo}} days ago
	{{end}}`
)

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func get_issue_information(keywords []string, len int) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(keywords, " "))
	format_url := fmt.Sprintf(issue_url, q)
	fmt.Println(format_url)
	response, err := http.Get(format_url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		fmt.Println(response.StatusCode)
		log.Fatal("response status code is not ok")
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		response.Body.Close()
		return nil, err
	}

	response.Body.Close()
	fmt.Println(result)

	return &result, err
}

func generate_report(result *IssuesSearchResult) {
	report, err := template.New("test_report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(template_text)
	if err != nil {
		log.Fatal(err)
	}
	if err = report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func print_report(result *IssuesSearchResult) {
	m_before := time.Now().AddDate(0, -1, 0) // 一个月前
	m_before_issue := []*Issue{}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %s %.55s \n",
			item.Number, item.User.Login, item.CreatedAt, item.Title)
		if m_before.Before(item.CreatedAt) {
			m_before_issue = append(m_before_issue, item)
		}
	}

	fmt.Println("=====================m_before=====================")
	for _, item := range m_before_issue {
		fmt.Printf("#%-5d %9.9s %s %.55s \n",
			item.Number, item.User.Login, item.CreatedAt, item.Title)
	}
}

func main() {
	result, err := get_issue_information([]string{"漏洞", "bug"}, 10)
	if err != nil {
		log.Fatal(err)
	}

	generate_report(result)
}
