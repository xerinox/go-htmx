package main

import (
	"flag"
	"html/template"
	"log"
	"net"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var ignorelist = []string{"header.html", "footer.html", "blog.html", "404.html", "page.html", "index.html"}

var t *template.Template
var routeMatch *regexp.Regexp
var blogRouteMatch *regexp.Regexp

var bloglist []string 

func main() {
	p := GetPort()
	log.Printf("Starting server at %s%s", GetOutboundIP(), p)

    setup()


	// static file server for css/js
    http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("client/resources/"))))

	//index handler
	http.HandleFunc("/", servePage)
	http.HandleFunc("/blog/", serveBlog)

	http.ListenAndServe(p, nil)
}

func setup() {
	var err error
	routeMatch, _ = regexp.Compile(`^\/(\w+)`)
	blogRouteMatch, _ = regexp.Compile(`^\/blog/(\w+)`)
	t, err = template.ParseGlob("templates/*")
	if err != nil {
		log.Fatal(err)
	}
    bloglist = getBlogs()
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func GetPort() string {
	var p int
	flag.IntVar(&p, "p", 8080, "Provide a port number")
	flag.Parse()
	return ":" + strconv.Itoa(p)
}

func getBlogs() []string {
	blogs, err := filepath.Glob("blog/*")

	if err != nil {
		log.Println(err)
	}

	blogdata := make([]string, len(blogs))

	for index, filename := range blogs {
		blogdata[index] = strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))
	}
    return blogdata
}
