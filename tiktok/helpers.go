package tiktok

import (
	"io"
	"net/http"
	"os"
	"strings"

	util "github.com/portalnesia/go-utils"
)

func generateRandomNumber() string {
	return util.NanoId()
}

func replaceUnicode(URL string) string {
	return strings.ReplaceAll(URL, "\u0026", "&")
}

func saveTiktok(filepath string, resp *http.Response) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
