package routing

import (
	"github.com/valyala/fasthttp"
	"log"
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

	err := fasthttp.ListenAndServe("127.0.0.1:5000", handler)
	if err != nil {
		log.Println(err)
	}
}