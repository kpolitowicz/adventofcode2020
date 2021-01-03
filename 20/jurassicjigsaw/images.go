package jurassicjigsaw

import "fmt"

type imageMap map[string]image
type image []string
type edges map[string][]edgeDesc
type edgeDesc struct {
	title, side string
	flipped     bool
}

func (ee *edges) FindSigletons() *edges {
	res := make(edges)
	for _, edgeList := range *ee {
		if len(edgeList) == 1 {
			edge := edgeList[0]
			res[edge.title] = append(res[edge.title], edge)
		}
	}
	return &res
}

func (im *imageMap) GetEdges() *edges {
	res := make(edges)
	for title, img := range *im {
		top := img[0]
		res[top] = append(res[top], edgeDesc{title, "T", false})
		topFlipped := reverse(top)
		res[topFlipped] = append(res[topFlipped], edgeDesc{title, "T", true})
		bottom := img[9]
		res[bottom] = append(res[bottom], edgeDesc{title, "B", false})
		bottomFlipped := reverse(bottom)
		res[bottomFlipped] = append(res[bottomFlipped], edgeDesc{title, "B", true})
		left := img.leftEdge()
		res[left] = append(res[left], edgeDesc{title, "L", false})
		leftFlipped := reverse(left)
		res[leftFlipped] = append(res[leftFlipped], edgeDesc{title, "L", true})
		right := img.rightEdge()
		res[right] = append(res[right], edgeDesc{title, "R", false})
		rightFlipped := reverse(right)
		res[rightFlipped] = append(res[rightFlipped], edgeDesc{title, "R", true})
	}
	return &res
}

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

func (i image) leftEdge() (res string) {
	for _, line := range i {
		res += string(line[0])
	}
	return
}

func (i image) rightEdge() (res string) {
	for _, line := range i {
		res += string(line[9])
	}
	return
}

func reverse(input string) (res string) {
	for i := len(input); i > 0; i-- {
		res += string(input[i-1])
	}
	return
}
