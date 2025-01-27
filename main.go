package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ouss4m4/go-rest/parsecsv"
)

func handleFileUpload(c *gin.Context) {
	// single file
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	// Upload the file to specific dst.
	dst := "./uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, dst); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("parsing the file at %s\n", dst)
	parsecsv.Parse(dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", handleFileUpload)

	router.Run("localhost:8080")
}
