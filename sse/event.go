package sse

import (
	"github.com/gin-gonic/gin"
	"log"
)

const (
	SseEmitUploadStatistics = iota
)

type EventPayload struct {
	Name int
	Data []string
}

// New event messages are broadcast to all registered client connection channels
type ClientChan chan EventPayload

// It keeps a list of clients those are currently attached
// and broadcasting events to those clients.
type Event struct {
	// Events are pushed to this channel by the main events-gathering routine
	Message chan EventPayload

	// New client connections
	NewClients chan chan EventPayload

	// Closed client connections
	ClosedClients chan chan EventPayload

	// Total client connections
	TotalClients map[chan EventPayload]bool
}

// Initialize event and SpinUp procnteessing requests
func NewServer() (event *Event) {
	event = &Event{
		Message:       make(chan EventPayload),
		NewClients:    make(chan chan EventPayload),
		ClosedClients: make(chan chan EventPayload),
		TotalClients:  make(map[chan EventPayload]bool),
	}

	go event.listen()

	return
}

// It Listens all incoming requests from clients.
// Handles addition and removal of clients and broadcast messages to clients.
func (stream *Event) listen() {
	for {
		select {
		// Add new available client
		case client := <-stream.NewClients:
			stream.TotalClients[client] = true
			log.Printf("Client added. %d registered clients", len(stream.TotalClients))

		// Remove closed client
		case client := <-stream.ClosedClients:
			delete(stream.TotalClients, client)
			close(client)
			log.Printf("Removed client. %d registered clients", len(stream.TotalClients))

		// Broadcast message to client
		case eventMsg := <-stream.Message:
			for clientMessageChan := range stream.TotalClients {
				clientMessageChan <- eventMsg
			}
		}
	}
}

func (stream *Event) ServeHTTP() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Initialize client channel
		clientChan := make(ClientChan)

		// Send new connection to event server
		stream.NewClients <- clientChan

		defer func() {
			// Send closed connection to event server
			stream.ClosedClients <- clientChan
		}()

		c.Set("clientChan", clientChan)

		c.Next()
	}
}

func HeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
		c.Next()
	}
}
