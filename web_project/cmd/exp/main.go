package main

import (
	"fmt"
	"web_project/models"
)

func main() {
	gs := models.GalleryService{}
	fmt.Println(gs.Images(1))
}
