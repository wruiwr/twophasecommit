package tpcGo

import (
	"fmt"
	"math/rand"
	"time"
)

// WorkerHandler represents a handler to handle the messages for Workers.
type WorkerHandler struct {
	workerId         int
	canCommitChannel chan CanCommit // channel for receiving CanCommit messages
	voteChannel         chan<- Vote    // channel for sending Vote messages
	decisionChannel  chan Decision  // channel for receiving Decision messages
	ackChannel          chan<- Ack     // channel for sending Ack messages
}

// NewWorker returns a worker handler for a new worker. It takes the
// following arguments:
//
// workerId: worker's ID
//
// voteChannel: a send only channel used to send vote message to coordinator.
//
// ackChannel: a send only channel used to send ack message to coordinator.
func NewWorker(id int, sendVote chan<- Vote, sendAck chan<- Ack) *WorkerHandler {

	return &WorkerHandler{
		workerId:         id,
		canCommitChannel: make(chan CanCommit, 16),
		voteChannel:         sendVote,
		decisionChannel:  make(chan Decision, 16),
		ackChannel:          sendAck,
	}
}

// Start starts w's main run loop as a separate goroutine. The main run loop
// sends vote and ack messages to coordinator, and handles incoming
// cancommit messages and decision messages.
func (w *WorkerHandler) Start() {

	go func() {
		for {
			select {
			case cc, ok := <-w.canCommitChannel:
				if !ok {
					fmt.Println("Haven't received cancommit from coordinator")
				} else {
					fmt.Println("Received cancommit from coordinator:", cc.String())
					w.SendVote(w.createVote(cc.GetWorkID()))
				}
			case d, ok := <-w.decisionChannel:
				if !ok {
					fmt.Println("Haven't received decision from coordinator")
				} else {
					fmt.Println("Received decision from coordinator:", d.String())
					w.SendACK(w.createAck(d.GetWorkID()))
				}

			}
		}
	}()
	time.Sleep(20 * time.Millisecond)
}

// DeliverCanCommit receives cancommit coming from coordinator.
func (w *WorkerHandler) DeliverCanCommit(cc CanCommit) {
	w.canCommitChannel <- cc
}

// sendVote sends vote to coordinator.
func (w *WorkerHandler) SendVote(v Vote) {
	w.voteChannel <- v
}

// DeliverDecision receives decision coming from coordinator.
func (w *WorkerHandler) DeliverDecision(d Decision) {
	w.decisionChannel <- d
}

// sendACK sends ack to coordinator.
func (w *WorkerHandler) SendACK(a Ack) {
	w.ackChannel <- a
}

// createVote creates a vote message with a randomly chosen "Yes" or "No".
func (w *WorkerHandler) createVote(id WorkerID) Vote {
	rand.Seed(time.Now().UnixNano())
	votes := []VoteEnum{Yes, No}
	vote := Vote{id, votes[rand.Intn(len(votes))]}
	fmt.Println("Created vote message:", vote)
	return vote
}

// create creates a ack message.
func (w *WorkerHandler) createAck(id WorkerID) Ack {
	return Ack{id}
}
