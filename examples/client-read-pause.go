// +build ignore

package main

import (
	"fmt"

	"github.com/aler9/gortsplib"
)

// This example shows how to
// * connect to a RTSP server
// * read all tracks with the UDP protocol for 5 seconds
// * pause for 5 seconds
// * repeat

func main() {
	// connect to the server and start reading all tracks
	conn, err := gortsplib.DialRead("rtsp://localhost:8554/mystream")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for {
		// read frames from the server
		readerDone := make(chan struct{})
		conn.OnFrame(func(id int, typ gortsplib.StreamType, buf []byte, err error) {
			if err != nil {
				close(readerDone)
				return
			}

			fmt.Printf("frame from track %d, type %v: %v\n", id, typ, buf)
		})

		// wait
		time.Sleep(5 * time.Second)

		// pause
		_, err := conn.Pause()
		if err != nil {
			panic(err)
		}

		// join reader
		<-readerDone

		// wait
		time.Sleep(5 * time.Second)

		// play again
		_, err := conn.Play()
		if err != nil {
			panic(err)
		}
	}
}