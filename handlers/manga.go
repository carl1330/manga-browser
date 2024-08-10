package handlers

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"os/exec"

	"github.com/labstack/echo/v4"
)

type Manga struct {}

func (m Manga) PostManga(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	dst, err := uploadFile(file);
	if err != nil {
		return err
	}

	cmd := exec.Command("unrar", "x", dst, "tmp")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		os.Remove(dst)
		return err
	}

	fmt.Println(dst[:len(dst)-4])

	cmd = exec.Command("mokuro", dst[:len(dst) - 4], "--disable_confirmation=TRUE")
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		return err
	}

	fmt.Println(string(output))

	return c.HTML(200, "Uploaded file")
}

func uploadFile(file *multipart.FileHeader) (string, error) {
	archive, err := file.Open()

	if err != nil {
		return "", err
	}
	
	defer archive.Close()
  dst, err := os.Create("tmp/" + file.Filename)
	if err != nil {
		log.Println("error creating file", err)
		return "", err
  }
	defer dst.Close()

	if _, err := io.Copy(dst, archive); err != nil {
		return "", err
	}

	return dst.Name(), nil
}