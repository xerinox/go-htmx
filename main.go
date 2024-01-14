package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"strconv"
    "regexp"
    "html/template"
)

var bl = []string{"header.html", "footer.html", "blog.html"}

var t *template.Template
var routeMatch *regexp.Regexp

func main() {
    var err error
    routeMatch, _ = regexp.Compile(`^\/(\w+)`)
    t, err = template.ParseGlob("templates/*")
    if (err != nil) {
        log.Fatal(err)
    }
    p := GetPort()
    log.Printf("Starting server at %s%s", GetOutboundIP(), p)

    // static file server for css/js
    http.Handle("/resources/",
        http.StripPrefix("/client/resources/", http.FileServer(http.Dir("client/resources/"))))

    //index handler
    http.HandleFunc("/", servePage)
    http.HandleFunc("/blog/", serveBlog)

    http.ListenAndServe(p, nil)
}

func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err!= nil {
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
