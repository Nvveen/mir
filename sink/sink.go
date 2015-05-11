package sink

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// A very simple handler to run requests and act as an internet backend.
func SinkHandler(w http.ResponseWriter, r *http.Request) {
	d := 0
	// if len(r.URL.Path[1:]) == 1 || len(r.URL.Path[1:]) == 2 {
	val, _ := strconv.ParseInt(r.URL.Path[1:], 10, 32)
	d = int(val)
	// }
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

// The main function to be called  to run the internet-mocking sink.
func RunSink() {
	fmt.Println("Starting mock server on :8080")
	http.HandleFunc("/", SinkHandler)
	http.ListenAndServe(":8080", nil)
}
