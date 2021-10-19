package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	Sid := "hi hello ..testing purpose,"
	Sid = strings.ReplaceAll(Sid, " ", "%20")
	url := "https://app-server.wati.io/api/v1/sendSessionMessage/917664075847?messageText=" + Sid + ""

	req, _ := http.NewRequest("POST", url, nil)
	log.Println(Sid)
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiI0OTNkODMxYS1hNmNmLTQ4YzgtOGQzYy0xMGE4N2JhM2NjOTYiLCJ1bmlxdWVfbmFtZSI6IlByYW5qaXRrYWthdGkwMkBnbWFpbC5jb20iLCJuYW1laWQiOiJQcmFuaml0a2FrYXRpMDJAZ21haWwuY29tIiwiZW1haWwiOiJQcmFuaml0a2FrYXRpMDJAZ21haWwuY29tIiwiYXV0aF90aW1lIjoiMTAvMTMvMjAyMSAwNjowMjoxMSIsImh0dHA6Ly9zY2hlbWFzLm1pY3Jvc29mdC5jb20vd3MvMjAwOC8wNi9pZGVudGl0eS9jbGFpbXMvcm9sZSI6IlRSSUFMIiwiZXhwIjoxNjM0Nzc0NDAwLCJpc3MiOiJDbGFyZV9BSSIsImF1ZCI6IkNsYXJlX0FJIn0.TtcUcnNI4peXE1pWPy7lK4gXwgqAyDXzPxJTrhmnj5o")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
