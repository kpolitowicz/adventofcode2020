package jurassicjigsaw

import (
	"fmt"
	"strings"
)

type imageMap map[string]image
type image []string

func (im *imageMap) String() (res string) {
	for title, img := range *im {
		res += fmt.Sprintf("Tile %s:\n", title)
		res += img.String()
	}
	return
}

func (i image) String() (res string) {
	for _, line := range i {
		res += fmt.Sprintf("%s\n", line)
	}
	return
}

func ParseInput(input string) *imageMap {
	var imageLine int
	var imageNum string
	var currentImage image
	images := make(imageMap)

	for _, line := range strings.Split(input, "\n") {
		imageLine++
		if line == "" {
			images[imageNum] = currentImage
			imageLine = 0
			currentImage = image{}
			continue
		}
		if imageLine == 1 {
			imageNum = parseImageNum(line)
		} else {
			currentImage = append(currentImage, line)
		}
	}

	images[imageNum] = currentImage
	return &images
}

func parseImageNum(line string) string {
	return line[5:9]
}
