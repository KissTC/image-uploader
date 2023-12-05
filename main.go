package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.Proto)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Hola ! !</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact page</h1> <p>to get in touch email me at <a href=\"mailto:ing.julio.code@gmail.com\">ing.julio.code@gmail.com</p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
	<ul>
	  <li>
		<b>Is there a free version?</b>
		Yes! We offer a free trial for 30 days on any paid plans.
	  </li>
	  <li>
		<b>What are your support hours?</b>
		We have support staff answering emails 24/7, though response
		times may be a bit slower on weekends.
	  </li>
	  <li>
		<b>How do I contact support?</b>
		Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>
	  </li>
	</ul>
	`)
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		//TODO: PAGE NOT FOUND
// 		// w.WriteHeader(http.StatusNotFound)
// 		// fmt.Fprint(w, "page not found")
// 		http.Error(w, "page not found", http.StatusNotFound)
// 	}
// }

type Router struct{}

func (Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		//TODO: PAGE NOT FOUND
		// w.WriteHeader(http.StatusNotFound)
		// fmt.Fprint(w, "page not found")
		http.Error(w, "page not found", http.StatusNotFound)
	}
}

func main() {
	// http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/contact", contactHandler)
	//http.HandleFunc("/", pathHandler)
	var router Router
	fmt.Println("starting server on port 3000...")
	http.ListenAndServe(":3000", router)
}