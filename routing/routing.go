package routing

import (
	"link-shortener/conn"
	"log"

	"github.com/valyala/fasthttp"
)

func StartHttp() {
	handler := func (ctx *fasthttp.RequestCtx) {
		switch string(ctx.Method()) {

		case "GET":
			log.Println("GET")
			getOriginalUrl(ctx)

		case "POST":
			log.Println("POST")
			getShortUrl(ctx)

		default:
			log.Println("Error: method does not exist")
		}

		/*
		switch string(ctx.Path()) {

		case "generate":
			switch string(ctx.Method()) {

			case "POST":
				log.Println("POST")
				getShortUrl(ctx)
			default:

				log.Println("Error: method does not exist")
			}

		default:
			switch string(ctx.Method()) {

			case "GET":
				log.Println("GET")
				getOriginalUrl(ctx)
				
			default:
				log.Println("Error: method does not exist")
			}
		}
		*/
	}
	log.Println("Server start")
	err := fasthttp.ListenAndServe(conn.ServAddr, handler)
	if err != nil {
		log.Println(err)
	}
}