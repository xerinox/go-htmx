package main

import (
	"html/template"
	"os"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
    var buf strings.Builder
    b, err := os.ReadFile("testtemplate.tpl")

    if (err != nil) {
        t.Fatal(err)
    }

    sb := string(b)

    tpl, err := template.New("webpage").Parse(sb)
    
    if (err != nil) {
        t.Fatal(err)
    }
    want := "<div>Test</div>\n"

    body := "Test"

    data := struct {
        Text string
    }{
        Text: body,
    }

    err = tpl.Execute(&buf, data)
    if err != nil {
        t.Fatal(err)
    }
    if got := buf.String(); got != want {
        t.Fatalf("got %q; want %q", got, want)
    }
}
