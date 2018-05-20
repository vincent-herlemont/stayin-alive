package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"github.com/sirupsen/logrus.git"
	"os"
	"net"
	"time"
)

type Server struct {
	cfg      *CfgWebServer
	log      *logrus.Logger
	stopChan chan struct{}
	router   *mux.Router
	server   *http.Server
}

func NewServer(cfg *CfgWebServer) *Server {
	// Init Logger
	log := logrus.New()
	log.Formatter = &logrus.JSONFormatter{}
	log.Out = os.Stdout
	log.Level = logrus.DebugLevel

	// Init router
	router := mux.NewRouter()
	router.HandleFunc("/info",HandleInfo).Methods("GET")

	// Init server
	server := &http.Server{
		Handler:      router,
		Addr:         cfg.Addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return &Server{
		log:      log,
		stopChan: make(chan struct{}),
		router:   router,
		server:   server,
		cfg: cfg,
	}
}

func (s *Server) Init(Inits... func(r *mux.Router)) {
	for _,Init := range Inits {
		Init(s.router)
	}
}

func (s *Server) Start() {
	s.log.Info("Starting")
	listener, err := net.Listen("tcp", s.cfg.Addr)
	if err != nil {
		s.log.Errorf("cannot listen on %s", s.cfg.Addr)
		return
	}
	s.log.Infof("listening on %s", s.cfg.Addr)

	go func() {
		err :=  s.server.Serve(listener)
		if err != nil {
			s.log.Error("server error", err)
		}
	}()
}

func (s *Server) Stop() {
	s.log.Info("Stopping")
	err := s.server.Close()
	if err != nil {
		s.log.Error("server error", err)
	}
}

func (s *Server) GetClient() *Client {
	return &Client{
		cfg:s.cfg,
	}
}

type TOInfo struct {
	Message string `json:"message"`
}

func HandleInfo(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(TOInfo{Message:"alive !"})
}
