package jurassicjigsaw

import (
	"strings"
)

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
