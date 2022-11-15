package main

import (
	"fmt"
	"os"
	"time"

	"github.com/pion/webrtc/v3"
)

func main() {

	start := time.Now()

	peerConnection, _ := webrtc.NewPeerConnection(webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	})

	peerConnection.OnICECandidate(func(c *webrtc.ICECandidate) {
		if c == nil {
			fmt.Println("ICE gathering done")
			fmt.Printf("Execution time: %2.2f\n", time.Since(start).Seconds())
			os.Exit(0)
		} else {
			fmt.Printf("ICE gathering found: %+v\n", c)
		}
	})

	offer, _ := peerConnection.CreateOffer(nil)

	peerConnection.SetLocalDescription(offer)

	select {}
}
