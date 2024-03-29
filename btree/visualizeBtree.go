package btree

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
)

func DrawBtree(rootNode node) {

	var buf strings.Builder

	styles, err := ioutil.ReadFile("styles.css")
	handleError(err)

	buf.WriteString(string(styles))
	fmt.Fprintf(&buf, "<div class=\"tree\"><ul>%s</ul></div>\n", nodeToHtml(&rootNode))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, buf.String())
	})

	open("http://localhost:8500/")
	http.ListenAndServe(":8500", nil)
}

func nodeToHtml(Node *node) string {

	var str strings.Builder

	if Node.Leftnode == nil && Node.Rightnode == nil {

		// A node with no child nodes
		fmt.Fprintf(&str, "<li>\n<a href=\"#\">%s</a>\n</li>\n", Node.Value)
		return str.String()
	} else {
		fmt.Fprintf(&str, "<li>\n<a href=\"#\">%s</a>\n<ul>\n%s\n%s\n</ul>\n</li>",
			Node.Value, nodeToHtml(Node.Leftnode),
			nodeToHtml(Node.Rightnode))
		return str.String()
	}
}

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

// open opens the specified URL in the default browser of the user.
func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
