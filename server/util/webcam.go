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
	cam, err := webcam.Open(os.Getenv("WEBCAM"))
	if err != nil {
		log.Fatal(err)
	}
	defer cam.Close()

	for {
		select {
		case <- (*c).Done():
			break
		case <- time.After(5 * time.Second):
			f, i, err := cam.GetFrame()
			if err != nil {
				log.Error(err)
			}
			defer cam.ReleaseFrame(i)
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