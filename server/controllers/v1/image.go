package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/mohemohe/parrot-webcam/server/util"
	"net/http"
)

func GetImage(c echo.Context) error {
	f := util.GetWebcamFrame()
	if len(f) == 0 {
		return c.NoContent(http.StatusNoContent)
	}
	return c.Blob(http.StatusOK, "image/jpeg", f)
}
