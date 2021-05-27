package frenchspace

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"io"
	"unicode"
)

func Stream(w io.Writer, r io.Reader) error {
	root, err := html.Parse(r)
	if err != nil {
		return err
	}

	nodeWalkerText(root, func(n *html.Node) {
		s := html.UnescapeString(n.Data)
		if !pureSpace(s) {
			begin, body, end := getSpace(s)
			n.Data = begin + Text(body) + end
		}
	})

	return html.Render(w, root)
}

func getSpace(s string) (begin, body, end string) {
	// Get the begin spaces
	for i, r := range s {
		if !unicode.IsSpace(r) {
			begin = s[:i]
			s = s[i:]
			break
		}
	}

	// Get the end spaces
	for i, r := range s {
		if !unicode.IsSpace(r) {
			end = ""
			body = s
		} else if end == "" {
			end = s[i:]
			body = s[:i]
		}
	}

	return
}

// pureSpace return true if all runes from s are space.
func pureSpace(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

// walk into the node tree and for each text node call f.
func nodeWalkerText(n *html.Node, f func(*html.Node)) {
	if n.Type == html.TextNode {
		f(n)
	} else if n.Type == html.ElementNode && n.DataAtom == atom.Pre {
		return
	} else {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			nodeWalkerText(c, f)
		}
	}
}
