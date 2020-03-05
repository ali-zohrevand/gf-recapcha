package routers

import (
	"fmt"
	"github.com/ali-zohrevand/gf-recapcha/release/Handler"
	"github.com/ali-zohrevand/gf-recapcha/release/util"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"net/http"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", AddContext(Handler.Index))

	router.HandleFunc("/ver3", AddContext(Handler.Ver3))
	router.HandleFunc("/puzzle", AddContext(Handler.Puzzle))
	router.HandleFunc("/v2", AddContext(Handler.V2))

	return router

}
func AddContext(next http.HandlerFunc) http.HandlerFunc {
	conf, err := util.GetConf("conf.yaml")
	if err != nil {
		panic("Bad Conf file.")
	}
	fmt.Println("Conf file loaded.")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Middleware operations
		// Parse body/get token.
		context.Set(r, "v2", conf.V2.Public)
		context.Set(r, "v3", conf.V3.Public)
		context.Set(r, "v3private", conf.V3.Private)
		context.Set(r, "v2private", conf.V2.Private)
		context.Set(r, "score", conf.Score)
		context.Set(r, "title", conf.Title)

		next.ServeHTTP(w, r)
	})
}
