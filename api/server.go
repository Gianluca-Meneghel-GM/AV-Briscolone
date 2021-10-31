package api

import (
	"briscolone/database"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Server struct {
	router  *mux.Router
	adapter *database.SqliteItemsAdapter
	file    io.ReadSeeker
}

func (s *Server) Initialize(rw *database.SqliteItemsAdapter) {
	s.router = mux.NewRouter()
	s.adapter = rw
}

func (s Server) SetupRoutes() {

	connessioni = make(map[int]*websocket.Conn)

	wsRouter := s.router.PathPrefix("/ws").Subrouter()
	wsRouter.Path("/registra/{id}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		registraWebSocket(w, r)
	})

	giocatoriRouter := s.router.PathPrefix("/giocatori").Subrouter()
	giocatoriRouter.Path("/all").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetGiocatoriHandler(w, r, s.adapter)
	})
	giocatoriRouter.Path("/add/{nome}").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AddGiocatoreHandler(w, r, s.adapter)
	})
	giocatoriRouter.Path("/addbot/{nome}").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AddBotHandler(w, r, s.adapter)
	})
	giocatoriRouter.Path("/select/{nome}").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		SelectGiocatoreHandler(w, r, s.adapter)
	})

	partitaRouter := s.router.PathPrefix("/partita").Subrouter()

	partitaRouter.Path("/end").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		FinisciPartitaHandler(w, r, s.adapter)
	})

	partitaRouter.Path("/mostravittoria").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		MostraVittoriaHandler(w, r, s.adapter)
	})
	partitaRouter.Path("/fineround/{id}").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetFineRoundHandler(w, r, s.adapter)
	})

	s.router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, err1 := route.GetPathTemplate()
		met, err2 := route.GetMethods()
		fmt.Println(tpl, err1, met, err2)
		return nil
	})
}

func (s Server) ServeStatic() {
	staticDir := http.Dir(`static`)
	// This will serve files under http://localhost:8000/<filename>
	s.router.PathPrefix("/").Handler(http.FileServer(staticDir))
}

func (s Server) Start() {

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "X-TOKEN"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	corsEnabledRouter := handlers.CORS(headersOk, originsOk, methodsOk)(s.router)
	http.ListenAndServe(":8080", corsEnabledRouter)
}
