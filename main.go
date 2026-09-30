package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var gifs = []string{
	"https://s3.amazonaws.com/files.656.mba/mgt656/fall-2026/random-gifs/adorbs/CapedScooterDog.gif",
	"https://s3.amazonaws.com/files.656.mba/mgt656/fall-2026/random-gifs/adorbs/Bearodynamic.gif",
	"https://s3.amazonaws.com/files.656.mba/mgt656/fall-2026/random-gifs/adorbs/CarefulRacoon.gif",
	"https://s3.amazonaws.com/files.656.mba/mgt656/fall-2026/random-gifs/adorbs/CatBowl.gif",
	"https://s3.amazonaws.com/files.656.mba/mgt656/fall-2026/random-gifs/adorbs/AttackingTheCatBuritto.gif",
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		// 从网址里读 count，比如 /?count=3
		count, err := strconv.Atoi(r.URL.Query().Get("count"))
		if err != nil || count < 1 || count > 3 {
			count = 1 // 缺失或无效时只显示一张
		}

		// 生成 count 个 <img> 标签
		var imgs strings.Builder
		for i := 0; i < count; i++ {
			gif := gifs[rand.IntN(len(gifs))]
			fmt.Fprintf(&imgs, `<img src="%s" alt="A cute animal">`+"\n", gif)
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		html := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
  <title>Random GIF</title>
</head>
<body style="background-color: lightblue; text-align: center; font-family: sans-serif;">
  <h1>Hello, world!</h1>

  <form method="get" action="/">
    <label for="count">How many GIFs?</label>
    <select id="count" name="count">
      <option value="1">1</option>
      <option value="2">2</option>
      <option value="3">3</option>
    </select>
    <button type="submit">Show me</button>
  </form>

  <div>
%s  </div>
</body>
</html>`, imgs.String())

		if _, err := w.Write([]byte(html)); err != nil {
			log.Printf("write response: %v", err)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Open http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}