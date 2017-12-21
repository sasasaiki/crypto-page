package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sasasaiki/eth-page/src/go/handler"
)

//NewProdRoutingHandlers 本番用ハンドラーを作成
func NewProdRoutingHandlers() (*[]handler.HandlingFunc, *[]handler.Handler) {
	hf := handler.NewHandlingFuncs(new(handler.ProdHandlingFunc))
	hs := handler.NewHandlers(handler.NewProdMyHandlerList())
	return &hf, &hs
}

//CreateRoute 渡されたhandlerとfuncについてrouteを設定する
func CreateRoute(hf *[]handler.HandlingFunc, hs *[]handler.Handler) *mux.Router {
	r := mux.NewRouter()

	//cssやjsを読み込めるようにするHandler
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))

	setFuncsRoute(r, hf)
	setHandlersRoute(r, hs)
	// 404のときのハンドラ
	//r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	return r
}

//Handlerが必要ないrouteの設定
func setFuncsRoute(r *mux.Router, hf *[]handler.HandlingFunc) {
	for _, h := range *hf {
		setRoute(r, http.HandlerFunc(h.Function), h.Conf)
	}
}

//Handlerが必要なrouteの設定
//templete読み込みなど
func setHandlersRoute(r *mux.Router, hs *[]handler.Handler) {
	for _, h := range *hs {
		setRoute(r, h.Handler, h.Conf)
	}
}

//新しくHandlerをデコレーションする必要がある場合はここでやる
func setRoute(r *mux.Router, h http.Handler, conf *handler.HandlingConf) {
	result := h
	if conf.NeedLogin {
		result = handler.NewAuthHandler(result)
	}
	result = handler.NewLogHandler(result)
	r.Handle(conf.Path, result).Methods(conf.Methods...)
}
