package util

import (
	"bytes"
	"context"
	"github.com/blackjack/webcam"
	"github.com/labstack/gommon/log"
	"github.com/pixiv/go-libjpeg/jpeg"
	"image"
	"image/draw"
	"os"
	"sort"
	"time"
)

type FrameSizes []webcam.FrameSize

func (slice FrameSizes) Len() int {
	return len(slice)
}

func (slice FrameSizes) Less(i, j int) bool {
	ls := slice[i].MaxWidth * slice[i].MaxHeight
	rs := slice[j].MaxWidth * slice[j].MaxHeight
	return ls < rs
}

func (slice FrameSizes) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

const V4L2_PIX_FMT_YUYV = 0x56595559

var c *context.Context
var b []byte
var t time.Time

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

	w, h := setFrameSize(cam)

	cam.StartStreaming()
	defer cam.StopStreaming()

	for {
		select {
		case <- (*c).Done():
			break
		default:
			_ = cam.WaitForFrame(5)
			f, err := cam.ReadFrame()
			if err != nil {
				log.Error(err)
			}
			if len(f) == 0 {
				log.Error("frame bytes: 0")
			} else {
				c := make([]byte, len(f))
				copy(c, f)
				go func() {
					t = time.Now()
					b = byteToJpeg(c, w, h)
				}()
			}
			break
		}
	}
}

func setFrameSize(cam *webcam.Webcam) (width int, height int) {
	frames := FrameSizes(cam.GetSupportedFrameSizes(V4L2_PIX_FMT_YUYV))
	sort.Sort(frames)
	s := &frames[len(frames)-1]
	width = int(s.MaxWidth)
	height = int(s.MaxHeight)
	if _, _, _, err := cam.SetImageFormat(V4L2_PIX_FMT_YUYV, s.MaxWidth, s.MaxWidth); err != nil {
		log.Fatal(err)
	}
	return
}

func byteToJpeg(b []byte, width int, height int) []byte {
	r := image.Rect(0, 0, width, height)
	yuyv := image.NewYCbCr(r, image.YCbCrSubsampleRatio422)
	for i := range yuyv.Cb {
		ii := i * 4
		yuyv.Y[i*2] = b[ii]
		yuyv.Y[i*2+1] = b[ii+2]
		yuyv.Cb[i] = b[ii+1]
		yuyv.Cr[i] = b[ii+3]
	}
	rgba := image.NewRGBA(r)
	draw.Draw(rgba, rgba.Bounds(), yuyv, yuyv.Bounds().Min, draw.Src)

	//rotateStr := os.Getenv("ROTATE")
	//if rotateStr != "" {
	//	deg, err := strconv.ParseFloat(rotateStr, 64)
	//	if err == nil {
	//		rgba = imaging.Rotate(rgba, deg, color.Transparent)
	//	}
	//}

	buf := &bytes.Buffer{}
	if err := jpeg.Encode(buf, rgba, &jpeg.EncoderOptions{Quality: 96}); err != nil {
		log.Error(err)
		return []byte{}
	}
	return buf.Bytes()
}

func GetWebcamFrame() []byte {
	if c == nil || len(b) == 0 {
		return []byte{}
	}
	return b
}

func GetTime() time.Time {
	return t
}