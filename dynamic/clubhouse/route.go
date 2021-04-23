package clubhouse

import (
	"strconv"
)

func ClubhouseRoute(params []string) string {
	message := "<script>console.log('cool beans') </script><h1>if goofy has a dog, and goofy is a dog....????</h1>"

	return "HTTP/1.1 200 OK\r\nContent-Type: text/html; charset=UTF-8\r\nContent-Length: " + strconv.Itoa(len(message)) + "\r\n\r\n" + message + "\n"
}
