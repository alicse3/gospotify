package utils

import (
	"context"
	"fmt"
	"net/http"

	"github.com/alicse3/gospotify/consts"
)

// HttpServer interface defines the methods for starting the http server and listening for the /callback request.
// Code received from the callback request can be communicated using a channel.
// Check DefaultHttpServer struct for implementation details.
type HttpServer interface {
	StartServer(ctx context.Context, ch chan string) error
}

// DefaultHttpServer is a struct that implements HttpServer interface.
type DefaultHttpServer struct{}

// StartServer starts an HTTP server to listen for the authentication callback.
func (dhs *DefaultHttpServer) StartServer(ctx context.Context, ch chan string) error {
	// Create new serve mux
	mux := http.NewServeMux()

	// GET /callback to handle callback requests
	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		// Get the code from the URL parameters
		code := r.URL.Query().Get("code")

		// Send the code through channel
		ch <- code

		// Show success response
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, consts.MsgCodeReceived)
	})

	// Create a new server and pass the mux as the handler
	server := http.Server{Addr: ":8080", Handler: mux}

	go func() {
		// Wait for the context to be done
		<-ctx.Done()

		// Shutdown the server gracefully
		if err := server.Shutdown(context.Background()); err != nil {
			panic(err)
		}
	}()

	// Start the server and listen for incoming requests
	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
