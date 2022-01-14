package routing

import (
	"link-shortener/conn"
	_ "link-shortener/conn"
	"log"


	"github.com/valyala/fasthttp"
)

func getShortUrl(ctx *fasthttp.RequestCtx) {
	url := string(ctx.QueryArgs().Peek("url"))
	if url == "" {
		ctx.Error("Bad request", fasthttp.StatusBadRequest)
	}

	var shortUrl string
	if conn.Memory == "postgres" {
		row := conn.DB.QueryRow(conn.GetShortUrl, url)
		err := row.Scan(&shortUrl)
		if err != nil {
			log.Println(err)
		}
		if shortUrl == "" {
			shortUrl = url // create short
			conn.DB.QueryRow(conn.InsertUrl, url, shortUrl)
		}
	} else {
		// todo insert map
		shortUrl = conn.Data[url]
		if shortUrl == "" {
			shortUrl = url // create short
			conn.Data[url] = shortUrl;
		}
	}
	ctx.Response.SetBody([]byte(shortUrl))
	ctx.SetStatusCode(200)
}

func getOriginalUrl(ctx *fasthttp.RequestCtx) {
	shortUrl := string(ctx.QueryArgs().Peek("url"))
	if shortUrl == "" {
		ctx.Error("Bad request", fasthttp.StatusBadRequest)
	}

	var url string
	if conn.Memory == "postgres" {
		row := conn.DB.QueryRow(conn.GetShortUrl, shortUrl)
		err := row.Scan(&url)
		if err != nil {
			// todo response error
			log.Println(err)
		}
	} else {
		// todo insert map
		for value, key := range conn.Data {
			if value == shortUrl {
				url = key
				break
			}
		}
		if url == "" {
			// error not found
		}
	}
	ctx.Response.SetBody([]byte(url))
	ctx.SetStatusCode(200)
}