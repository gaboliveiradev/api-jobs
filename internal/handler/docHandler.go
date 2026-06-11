package handler

import "github.com/gin-gonic/gin"

func ScalarDocs(c *gin.Context) {
	html := `
	<!doctype html>
	<html>
		<head>
		  <title>API Jobs Docs</title>
		  <meta charset="utf-8" />
		</head>
		<body>
		  <script
		    id="api-reference"
		    data-url="/swagger/doc.json"
		    data-theme="deepSpace">
		  </script>

		  <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
		</body>
	</html>`

	c.Data(200, "text/html; charset=utf-8", []byte(html))
}
