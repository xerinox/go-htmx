package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type NavItem struct {
	Text string
	Src  string
}

var Nav = []NavItem{
	{
		Text: "Home",
		Src:  "/",
	},
	{
		Text: "About",
		Src:  "/about",
	},
	{
		Text: "Blog",
		Src:  "/blog",
	},
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

var data = struct {
	Title    string
	BodyText string
	NavItems []NavItem
}{
	Title:    "Test Page",
	BodyText: "This Works",
	NavItems: Nav,
}

// Endpoint handler for convert endpoint on server
func servePage(w http.ResponseWriter, req *http.Request) {
	matches := routeMatch.FindStringSubmatch(req.URL.Path)
	if len(matches) >= 1 {
		page := matches[1] + ".html"

		if stringInSlice(page, bl) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("NOT FOUND"))
			return
		} else if t.Lookup(page) != nil {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			t.ExecuteTemplate(w, page, data)
			return
		}
	} else if req.URL.Path == "/" {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		t.ExecuteTemplate(w, "index.html", data)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("NOT FOUND"))
}

type page struct {
    Title string;
    Src string;
}
func serveBlog(w http.ResponseWriter, req *http.Request) {
	blogRouteMatch, _ := regexp.Compile(`^\/blog/(\w+)`)
	blogMatches := blogRouteMatch.FindStringSubmatch(req.URL.Path)
	if len(blogMatches) >= 1 {
		page := blogMatches[1] + ".html"
		blogmatch, err := filepath.Glob("blog/" + page)
		if len(blogmatch) >= 1 {
			if err != nil {
				log.Println(err)
			}
			b, err := os.ReadFile(blogmatch[0])
			if err != nil {
				log.Println(err)
			}
            data := struct {
                Title    string
                BodyText string
                NavItems []NavItem
                IsLanding bool
            }{
                Title:    "Test Page",
                BodyText: "This Works",
                NavItems: Nav,
                IsLanding: false,
            }

			data.BodyText = string(b)
            

			t.ExecuteTemplate(w, "blog.html", data)
		} else {
			log.Println("too short")
			log.Println(blogmatch)
			log.Println(page)
		}
	} else if req.URL.Path == "/blog/" {
		blogs, err := filepath.Glob("blog/*")
        if err != nil {
            log.Println(err)
        }

        blogdata := make([]page, len(blogs))

        for index, item := range blogs {
            filename := strings.TrimLeft(item, "blogs/")
            name := strings.TrimRight(filename, ".html")
            
            blogdata[index] = page{Title: name, Src: filename}
        }

		data := struct {
			Title    string
			BodyText string
			NavItems []NavItem
            Blogs []page
            IsLanding bool
		}{
			Title:    "Test Page",
			BodyText: "This Works",
			NavItems: Nav,
            Blogs: blogdata,
            IsLanding: true,
		}

		w.WriteHeader(http.StatusOK)
		data.BodyText = "Blog landing page"
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		t.ExecuteTemplate(w, "blog.html", data)
		return

	}
}
