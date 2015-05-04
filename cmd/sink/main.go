package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	d := 0
	if len(r.URL.Path[1:]) == 1 || len(r.URL.Path[1:]) == 2 {
		val, _ := strconv.ParseInt(r.URL.Path[1:], 10, 32)
		d = int(val)
	}
	f1 := `<html>
<head><title></title></head>
<body>
<ul>
`
	for i := 0; i < 10; i++ {
		f1 += `<li><a href="`
		if d > 0 {
			f1 += fmt.Sprintf("%d%d", d, i)
		} else {
			f1 += fmt.Sprintf("%d", i)
		}
		f1 += `">Word` + fmt.Sprintf("%d", i) + `</a></li>`
		f1 += "\n"
	}
	f1 += `</ul>
</body>
</html>
`
	fmt.Fprintf(w, f1)
	time.Sleep(time.Millisecond * 200)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
