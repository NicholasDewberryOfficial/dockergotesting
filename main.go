package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// To use: you need curl
// curl -X POST http://localhost:8080/uploadsinglefile -F "file=@/Users/nick/Downloads/player.png" -H "Content-Type: multipart/form-data"
func main() {
	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.GET("/", homepage)

	router.GET("/downloadlatest", downloadlatest)

	router.GET("/downloadmusicfile", downloadmusicfile)

	router.POST("/uploadsinglefile", uploadsinglefile)

	router.POST("/upload", func(c *gin.Context) {
		// Source
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "get form err: %s", err.Error())
			return
		}

		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}

		c.SaveUploadedFile(file, filepath.Join("/ginrouter/uploads", file.Filename))
		c.String(http.StatusOK, "File %s uploaded successfully with fields name=%s and email=%s.", file.Filename)
	})

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func homepage(c *gin.Context) {
	c.String(http.StatusOK, "this is the homepage")
}

func downloadlatest(c *gin.Context) {
	filename := "ginrouter/uploadfilesfromhere/sprint4.pdf"

	_, err := os.Stat(filename)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": "filenotfound",
		})
		return
	}

	// Set headers for download
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Length", "0")

	// Send the file as a response
	c.File(filename)

}

func downloadmusicfile(c *gin.Context) {

	// Define the location of the music file on the server
	filename := "ginrouter/uploadfilesfromhere/intro10sec.mp3" // Change this path to where your music file is located

	// Check if the file exists
	_, err := os.Stat(filename)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "File not found",
		})
		return
	}

	// Set headers for download
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Length", "0")

	// Send the file as a response
	c.File(filename)

}

func uploadsinglefile(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}

	c.SaveUploadedFile(file, filepath.Join("ginrouter/uploads", file.Filename))
	c.String(http.StatusOK, "File %s uploaded successfully.", file.Filename)
}
