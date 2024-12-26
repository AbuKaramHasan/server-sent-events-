package main

import (
	"context"
	"fmt"
	"mini/channels"
	"mini/handlers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go backgroundTask()
	go ListenToInputChanne1()

	server := &http.Server{Addr: ":8080", Handler: http.DefaultServeMux}
	http.HandleFunc("/", handlers.Handler)
	http.HandleFunc("/simple", handlers.SimpleHandler)
	http.HandleFunc("/html", handlers.HTMLHandler)
	http.HandleFunc("/msg", handlers.MsgHandler)
	http.HandleFunc("/event", handlers.EventHandler)
	http.HandleFunc("/sse", handlers.SSEHandler)
	http.HandleFunc("/custom", handlers.CustomHandler)
	http.HandleFunc("/payload", handlers.PayloadHandler)

	go func() {
		fmt.Println("Starting server on :8080...")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Printf("Server error: %v\n", err)
			stop()
		}
	}()

	<-ctx.Done() // Wait for shutdown signal.
	fmt.Println("Shutting down server...")
	server.Shutdown(context.Background())
}

func backgroundTask() {
	signal := channels.Signal{
		ID:      1,
		Payload: "Hello, World!",
	}

	channels.InputChan <- signal
}

// for debuging
func ListenToInputChanne1() {
	for signal := range channels.InputChan {
		select {
		case <-signal.Context:
			fmt.Println("Request has been canceled.")
			continue
		default:
			switch signal.ID {
			case 1:
				channels.OutputChan1 <- signal
			case 2:
				channels.OutputChan2 <- signal
			case 3:
				channels.OutputChan3 <- signal
			}
		}
	}
}
