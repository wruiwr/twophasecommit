package main

import (
	"time"
	tpc "github.com/selabhvl/tpcGo"
)

func main() {

	numOfWorker := 3

	sendCanCommit := make(chan tpc.CanCommit, 16)
	sendDecision := make(chan tpc.Decision, 16)
	deliverVote := make(chan tpc.Vote, 16)
	deliverAck := make(chan tpc.Ack, 16)
	finalDecision := make(chan tpc.DecisionEnum, 1)

	ch := tpc.NewCoordinator(sendCanCommit, sendDecision)

	for i := 0; i < numOfWorker; i++ {
		go func(i int) {
			wh := tpc.NewWorker(i+1, deliverVote, deliverAck)
			wh.Start()
			mh := tpc.NewMessageHandler(sendCanCommit, deliverVote, sendDecision, deliverAck, ch, wh)
			mh.Start()
		}(i)
	}
	ch.Start(numOfWorker, finalDecision)

	time.Sleep(60 * time.Millisecond)
}
