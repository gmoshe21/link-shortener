package routing

import (
	"link-shortener/conn"
	"link-shortener/utils"

	"log"
	"github.com/valyala/fasthttp"
)

func getShortUrl(ctx *fasthttp.RequestCtx) {
	url := string(ctx.QueryArgs().Peek("url")) // получене параметра из запроса
	if url == "" { // проверка на валидность запроса
		ctx.Error("Bad request", fasthttp.StatusBadRequest)
		return
	}

	var shortUrl string
	if conn.Memory == "postgres" { // если хранилищем был выбран  postgres
		row := conn.DB.QueryRow(conn.GetShortUrl, url) // проверка на соществования короткого ключа для этого url
		err := row.Scan(&shortUrl)
		if err != nil && err.Error() != "sql: no rows in result set" {
			log.Println(err)
			ctx.Error("Server error", fasthttp.StatusInternalServerError)
			return
		}
		if shortUrl == "" { // если ключа нет генерируем его и записываем в базу данных
			for {
				shortUrl = utils.GenerateShortUrl() //генерация ключа

				var checkUrl string
				row = conn.DB.QueryRow(conn.CheckShortUrl, shortUrl) // проверка на уникальность ключаа
				err = row.Scan(&shortUrl)

				if err != nil && err.Error() != "sql: no rows in result set" {
					log.Println(err)
					ctx.Error("Server error", fasthttp.StatusInternalServerError)
					return
				}
				if checkUrl != shortUrl { // если ключ уникальный выходим из цикла, если ключ уже есть в базе, то генерируем заново
					break
				}
			}
			conn.DB.QueryRow(conn.InsertUrl, url, shortUrl) // добавление в базу данных
		}
	} else { // если хранилицем была выбрана внутренняя память
		var ok bool
		shortUrl, ok = conn.Data[url] // проверка на существование короткого ключа в памяти
		if !ok {
			for {
				shortUrl = utils.GenerateShortUrl() // генерация ключа
				for _, value := range conn.Data { // проверка на уникальность ключа
					if value == shortUrl {
						shortUrl = "" // если ключ найден в памяти, то он стирается
						break
					}
				}
				if shortUrl != "" {
					conn.Data[url] = shortUrl; // если ключ уникальный записываем в память, если ключ уже есть в памяти, то генерируем заново
					break
				}
			}
		}
	}
	shortUrl = "http://127.0.0.1:5000/?url=" + shortUrl
	ctx.Response.SetBody([]byte(shortUrl))
	ctx.SetStatusCode(200)
}

func getOriginalUrl(ctx *fasthttp.RequestCtx) {
	shortUrl := string(ctx.QueryArgs().Peek("url")) // получене параметра из запроса
	if shortUrl == "" { // проверка на валидность запроса
		ctx.Error("Bad request", fasthttp.StatusBadRequest)
		return
	}

	var url string
	if conn.Memory == "postgres" { // если хранилищем был выбран  postgres
		row := conn.DB.QueryRow(conn.GetOriginalUrl, shortUrl) // поиск по ключу оригинального url
		err := row.Scan(&url)
		if err != nil && err.Error() != "sql: no rows in result set"  { // ошибка сервера
			log.Println(err)
			ctx.Error("Server error", fasthttp.StatusInternalServerError)
			return
		} else if err != nil && err.Error() == "sql: no rows in result set" { // ключа не существует
			log.Println(err)
			ctx.Error("Status not found", 404)
			return
		}
	} else {  // если хранилицем была выбрана внутренняя память
		for key, value := range conn.Data { // поиск по ключу оригинального url
			if value == shortUrl {
				url = key
				break
			}
		}
		if url == "" {  // ключа не существует
			ctx.Error("Status not found", 404)
			return
		}
	}
	ctx.Response.SetBody([]byte(url))
	ctx.SetStatusCode(200)
}