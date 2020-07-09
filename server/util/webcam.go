package util

import (
	"context"
	"github.com/blackjack/webcam"
	"github.com/labstack/gommon/log"
	"os"
	"time"
)

var c *context.Context
var b []byte

func StartWebcam() bool {
	if c != nil {
		return false
	}
	b, fn := context.WithCancel(context.Background())
	b = context.WithValue(b, "cancel", fn)
	c = &b
	go loopWebcam()
	return true
}

func StopWebcam() bool {
	if c == nil {
		return false
	}
	fn := (*c).Value("cancel").(context.CancelFunc)
	fn()
	c = nil
	return true
}

func loopWebcam() {
	cam, err := webcam.Open(os.Getenv("DEVICE"))
	if err != nil {
		log.Fatal(err)
	}
	defer cam.Close()

	cam.StartStreaming()
	defer cam.StopStreaming()

	for {
		_ = cam.WaitForFrame(5)

		select {
		case <- (*c).Done():
			break
		case <- time.After(5 * time.Second):
			f, err := cam.ReadFrame()
			if err != nil {
				log.Error(err)
			}
			if len(f) == 0 {
				log.Error("frame bytes: 0")
			} else {
				b = f
			}
			break
		}
	}
}

func GetWebcamFrame() []byte {
	if c == nil || len(b) == 0 {
		return []byte{}
	}
	return b
}