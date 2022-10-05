package main

import (
	"fmt"
	"net/http"
)

const htmlCode = `
<!DOCTYPE html>
<html>
<body>

<h1>The input value attribute</h1>

<p>The value attribute specifies an initial value for an input field:</p>

<form action="/action_page.php">
  <label for="fname">First name:</label><br>
  <input type="text" id="fname" name="fname" value="John"><br>
  <label for="lname">Last name:</label><br>
  <input type="text" id="lname" name="lname" value="Doe"><br><br>
  <input type="submit" value="Submit">
</form>

</body>
</html>
`

func main() {

	http.HandleFunc("/home", home)

	http.ListenAndServe(":3000", nil)

}

func home(res http.ResponseWriter, req *http.Request) {
	url := req.FormValue("url")
	if url == "" {
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(res, htmlCode)
	}
}