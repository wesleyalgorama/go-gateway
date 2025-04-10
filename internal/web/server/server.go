package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/wesleyalgorama/fcw/go-gateway/internal/service"
	"github.com/wesleyalgorama/fcw/go-gateway/internal/web/handler"
)

type Server struct {
	router         *chi.Mux
	server         *http.Server
	accountService *service.AccountService
	invoiceService *service.InvoiceService
	port           string
}

func NewServer(accountService *service.AccountService, invoiceService *service.InvoiceService, port string) *Server {
	return &Server{
		router:         chi.NewRouter(),
		accountService: accountService,
		invoiceService: invoiceService,
		port:           port,
	}
}

func (s *Server) ConfigureRoutes() {
	accountHandler := handler.NewAccountHandler(s.accountService)
	invoiceHandler := handler.NewInvoiceHandler(s.invoiceService)

	s.router.Group(func(r chi.Router) {
		s.router.Post("/accounts", accountHandler.Create)
		s.router.Get("/accounts", accountHandler.Get)
	})

	s.router.Group(func(r chi.Router) {
		s.router.Post("/invoice", invoiceHandler.Create)
		s.router.Get("/invoice/{id}", invoiceHandler.GetByID)
		s.router.Get("/invoice", invoiceHandler.ListByAccount)
	})
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	return s.server.ListenAndServe()
}
