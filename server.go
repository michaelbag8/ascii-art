package main

import(
	"fmt"
	"net/http"
	"os"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
	html := `<!DOCTYPE html>
<html>
<body>
    <form method="POST" action="/ascii">
        <input type="text" name="text" placeholder="Enter text"/>
        
        <select name="banner">
            <option value="standard">standard</option>
            <option value="shadow">shadow</option>
            <option value="thinkertoy">thinkertoy</option>
        </select>

        <select name="color">
            <option value="">none</option>
            <option value="red">red</option>
            <option value="green">green</option>
            <option value="yellow">yellow</option>
            <option value="blue">blue</option>
            <option value="magenta">magenta</option>
            <option value="cyan">cyan</option>
            <option value="white">white</option>
        </select>

        <input type="text" name="letters" placeholder="Letters to color (optional)"/>

        <select name="align">
            <option value="left">left</option>
            <option value="right">right</option>
            <option value="center">center</option>
            <option value="justify">justify</option>
        </select>

        <button type="submit">Generate</button>
    </form>
</body>
</html>`
	fmt.Fprintln(w, html)
}
func handleAscii(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	text := r.FormValue("text")
	bannerName := r.FormValue("banner")
	colorName := r.FormValue("color")
	letters := r.FormValue("letters")
	align := r.FormValue("align")

	charMap, err := LoadBanner(bannerName + ".txt")
	if err != nil {
		http.Error(w, "error loading banner", http.StatusInternalServerError)
		return
	}
    w.Header().Set("Content-Type", "text/html; charset=utf-8")

	content := Generate(text, charMap, "", letters, align)
	if colorName != "" {
		cssColor := cssColorMap[colorName]
		fmt.Fprintf(w, `<pre><span style="color:%s">%s</span></pre>`, cssColor, content)
	} else {
		fmt.Fprintf(w, "<pre>%s</pre>", content)
	}

}

func startServer() {
    http.HandleFunc("/", handleHome)
    http.HandleFunc("/ascii", handleAscii)
    fmt.Println("Server running at http://localhost:8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "server error: %v\n", err)
        os.Exit(1)
    }
}