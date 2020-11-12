package main

import (
	"fmt"
	"log"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"github.com/PuerkitoBio/goquery"
	// "golang.org/x/net/html"
)

func main() {
	
	blogTitles, err := GetLatestBlogTitles("https://golang.org/")
	if err != nil {
		log.Println(err)
	}
	fmt.Printf(blogTitles)



}

// GetLatestBlogTitles gets the latest blog title headings from the url
// given and returns them as a list.
func GetLatestBlogTitles(url string) (string, error) {
	// valid:=""
	// Get the HTML
	iLinks:=0
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(io.LimitReader(resp.Body, 15))
	if err != nil {
		// panic(err)
	}
	// show the HTML code as a string %s
	fmt.Printf("Html Version: %s ", html)
	// Convert HTML into goquery document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(doc)
	// Save each .post-title as a list
	titles := ""
	doc.Find("head title").Each(func(i int, s *goquery.Selection) {
		titles += "- " + s.Text() + ""
	})
	h1:=0
	doc.Find("h1").Each(func(i int, s *goquery.Selection) {
		h1++
		// fmt.Println(s.Text())
	})
	h2:=0
	doc.Find("h2").Each(func(i int, s *goquery.Selection) {
		h2++
	})
	h3:=0
	doc.Find("h3").Each(func(i int, s *goquery.Selection) {
		h3++
	})
	h4:=0
	doc.Find("h4").Each(func(i int, s *goquery.Selection) {
		h4++
	})
	links:=0
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		links++
		// valid:= s.Attr("href")
		band, ok := s.Attr("href")
	    if ok {
	        // fmt.Println(IsUrl(band))
	        valid := IsUrl(band)
			if(!valid){
				iLinks++
			}
	    }
		// fmt.Println(IsUrl(valid))
		// fmt.Println(s.Attr("href"))

	})
	form:="No"
	doc.Find("input").Each(func(i int, s *goquery.Selection) {
		band, ok := s.Attr("name")
		if ok {
	        // fmt.Println(IsUrl(band))
	        if(band=="password"){
	        	form="Yes"
	        }
	    }
	})
	fmt.Print("Page Title: ")
	fmt.Println(titles)
	fmt.Print("Number of H1 Headings: ")
	fmt.Println(h1)
	fmt.Print("Number of H2 Headings: ")
	fmt.Println(h2)
	fmt.Print("Number of H3 Headings: ")
	fmt.Println(h3)
	fmt.Print("Number of H4 Headings: ")
	fmt.Println(h4)
	fmt.Print("Number of internal and external links: ")
	fmt.Println(links)
	fmt.Print("Number of inaccessible links: ")
	fmt.Println(iLinks)
	fmt.Print("Login Form Present? (detection on password field): ")
	fmt.Println(form)
	// fmt.Println(titles)
	
	return "0", nil
}
func IsUrl(str string) bool {
    u, err := url.Parse(str)
    return err == nil && u.Scheme != "" && u.Host != ""
}