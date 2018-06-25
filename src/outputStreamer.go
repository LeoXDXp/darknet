// What it does:
//
// This example opens a video capture device, then streams MJPEG from it.
// Once running point your browser to the hostname/port you passed in the
// command line (for example http://localhost:8080) and you should see
// the live video stream.
//
// How to run:
//
// mjpeg-streamer [camera ID] [host:port]
//
//		go get -u github.com/hybridgroup/mjpeg
// 		go run ./cmd/mjpeg-streamer/main.go 1 0.0.0.0:8080
//
// +build example

package main

import "C"

import (
	"github.com/hybridgroup/mjpeg"
	//"gocv.io/x/gocv"
	"log"
	"net/http"
)

var (
	err    error
	stream *mjpeg.Stream
)

// export StreamDarkNet
func StreamDarkNet() {
	// create the mjpeg stream
	stream = mjpeg.NewStream()
	// start capturing
	//go Capture()
	host := "8880" //Could be read from cfg file to make it "dynamic"
	log.Printf("Capturing. Point your browser to %s", host)
	// start http server
	http.Handle("/", stream)
	log.Fatal(http.ListenAndServe(host, nil))
}

//export GoCapture
func GoCapture(buf []byte) {
	//buf, _ := gocv.IMEncode(".jpg", (Mat)img)
	stream.UpdateJPEG(buf)
}

func main() {}
