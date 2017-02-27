//go:generate gendb -d .
//go:generate genmodel -d .

package scoreserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Server represents Score Server
type Server struct {
	*mux.Router
}

// Run Score Server
func Run(port string) error {
	s := New()
	return http.ListenAndServe(port, s)
}

// New creates Score Server object
func New() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}
	s.SetupRoutes()
	return s
}

// SetupRoutes set route to Server
func (s *Server) SetupRoutes() {

}
