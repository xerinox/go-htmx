package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
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


// Routing / serving handler for non-blog routes
func servePage(w http.ResponseWriter, req *http.Request) {
    var data = struct {
        Title    string
        BodyText string
        NavItems []NavItem
    }{
        Title:    "My homepage",
        BodyText: "This Works",
        NavItems: Nav,
    }

	matches := routeMatch.FindStringSubmatch(req.URL.Path)

	if len(matches) >= 1 {
		page := matches[1] + ".html"

		if stringInSlice(page, ignorelist) {
			w.WriteHeader(http.StatusNotFound)
            t.ExecuteTemplate(w, "404.html", nil)
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
    t.ExecuteTemplate(w, "404.html", data)
}

func serveBlog(w http.ResponseWriter, req *http.Request) {
	blogMatches := blogRouteMatch.FindStringSubmatch(req.URL.Path)
	data := struct {
		Title     string
		BodyText  string
		NavItems  []NavItem
		Blogs     []string
		IsLanding bool
	}{
		Title:     "Test Page",
		BodyText:  "This Works",
		NavItems:  Nav,
		Blogs:     nil,
		IsLanding: false,
	}

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
			data.BodyText = string(b)

            w.WriteHeader(http.StatusOK)
            w.Header().Set("Content-Type", "text/html; charset=utf-8")
			t.ExecuteTemplate(w, "blog.html", data)
            return
		} 
    } else if req.URL.Path == "/blog/" {
		data.Blogs = bloglist
		data.IsLanding = true

		w.WriteHeader(http.StatusOK)
		data.BodyText = "Blog landing page"
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		t.ExecuteTemplate(w, "blog.html", data)
		return
	}

	w.WriteHeader(http.StatusNotFound)
    t.ExecuteTemplate(w, "404.html", data)
}
