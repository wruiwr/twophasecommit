package tpcGo

import (
	"fmt"
	"time"
)

// MessageHandler represents a handler to handle the messages exchanging between Coordinator and Workers.
type MessageHandler struct {
	canCommitChannel      chan CanCommit      // channel for coordinator to send cancommit message to workers
	voteChannel           chan Vote           // channel for workers to send vote message to coordinator
	decisionChannel       chan Decision       // channel for coordinator to send decision message to workers
	ackChannel            chan Ack            // channel for workers to send ack message to coordinators
	coordinatorHandler *CoordinatorHandler // a coordinatorHandler can handle the messages for coordinator
	workerHandler      *WorkerHandler      // a workerHandler can handle the messages for a worker
}

// NewMessageHandler returns a message handler to handle the messages between Coordinator and Workers.
// It takes following arguments:
//
// canCommitChannel: a channel used to send cancommit message from coordinator to workers.
//
// decisionChannel: a channel used to send decision message from coordinator to workers.
//
// voteChannel: a channel used to send vote message from workers to coordinator.
//
// ackChannel: a channel used to send ack message from workers to coordinator.
func NewMessageHandler(sendCanCommit chan CanCommit, sendVote chan Vote,
	sendDecision chan Decision, sendAck chan Ack, ch *CoordinatorHandler, wh *WorkerHandler) *MessageHandler {

	return &MessageHandler{
		canCommitChannel:      sendCanCommit,
		voteChannel:           sendVote,
		decisionChannel:       sendDecision,
		ackChannel:            sendAck,
		coordinatorHandler: ch,
		workerHandler:      wh,
	}
}

// Start starts m's main run loop as a separate goroutine. The main run loop
// handles incoming message from workers and coordinator.
func (m *MessageHandler) Start() {

	go func() {
		for {
			select {
			case cc, ok := <-m.canCommitChannel:
				if !ok {
					fmt.Println("Haven't received cancommit message in message handler")
				} else {
					fmt.Println("Received canncommit message in message handler", cc.String())
					m.deliverCanCommit(cc)
				}
			case v, ok := <-m.voteChannel:
				if !ok {
					fmt.Println("Haven't received vote message in message handler")
				} else {
					fmt.Println("Received vote message in message handler", v.String())
					m.deliverVote(v)
				}
			case d, ok := <-m.decisionChannel:
				if !ok {
					fmt.Println("Haven't received decision message in message handler")
				} else {
					fmt.Println("Received decision message in message handler", d.String())
					m.deliverDecision(d)
				}
			case a, ok := <-m.ackChannel:
				if !ok {
					fmt.Println("Haven't received ack message in message handler")
				} else {
					fmt.Println("Received ack message in message handler", a.String())
					m.deliverAck(a)
				}
			}
		}
	}()

	time.Sleep(60 * time.Millisecond)
}

// deliverCanCommit delivers cancommit message from coordinator to workers.
func (m *MessageHandler) deliverCanCommit(cc CanCommit) {
	m.workerHandler.DeliverCanCommit(cc)
}

// deliverVote delivers vote message from workers to coordinator.
func (m *MessageHandler) deliverVote(v Vote) {
	m.coordinatorHandler.DeliverVote(v)
}

// deliverDecision delivers decision message from coordinator to workers.
func (m *MessageHandler) deliverDecision(d Decision) {
	m.workerHandler.DeliverDecision(d)
}

// deliverAck delivers ack message from workers to coordinator.
func (m *MessageHandler) deliverAck(a Ack) {
	m.coordinatorHandler.DeliverACK(a)
}
