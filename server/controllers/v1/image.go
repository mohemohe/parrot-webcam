package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/mohemohe/parrot-webcam/server/util"
	"net/http"
	"strconv"
)

func GetImage(c echo.Context) error {
	f := util.GetWebcamFrame()
	if len(f) == 0 {
		return c.NoContent(http.StatusNoContent)
	}
	return c.Blob(http.StatusOK, "image/jpeg", f)
}

func GetImageTaken(c echo.Context) error {
	t := util.GetTime()
	u := t.UnixNano() / 1000000
	return c.String(http.StatusOK, strconv.FormatInt(u, 10))
}
