package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HttpServer struct {
	httpHandlers *HTTPHandlers
}

func NewServer(httpHandler *HTTPHandlers) *HttpServer {
	return &HttpServer{
		httpHandlers: httpHandler,
	}
}

func (s *HttpServer) Start() error {
	var router = mux.NewRouter()

	router.
		Path("/tasks").
		Methods("POST").
		HandlerFunc(s.httpHandlers.HandleCreateTask)

	router.
		Path("/tasks/{title}").
		Methods("GET").
		HandlerFunc(s.httpHandlers.HandleGetTask)

	router.
		Path("/tasks").
		Methods("GET").
		HandlerFunc(s.httpHandlers.HandleGetAllTasks)

	router.
		Path("/tasks").
		Methods("GET").
		Queries("completed", "false").
		HandlerFunc(s.httpHandlers.HandleGetAllUncompletedTasks)

	router.
		Path("/tasks/{title}").
		Methods("PTCH").
		HandlerFunc(s.httpHandlers.HandleCompleteTask)

	router.
		Path("/tasks/{title}").
		Methods("DELETE").
		HandlerFunc(s.httpHandlers.HandleDeleteTask)

	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
