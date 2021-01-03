package jurassicjigsaw

import (
	"fmt"
	"sort"
)

type imageMap map[string]image
type image []string
type edges map[string][]edgeDesc
type edgeDesc struct {
	Title, side string
	flipped     bool
}
type neighbours map[string][]string

func (ee *edges) FindSigletons() *edges {
	res := make(edges)
	for _, edgeList := range *ee {
		if len(edgeList) == 1 {
			edge := edgeList[0]
			res[edge.Title] = append(res[edge.Title], edge)
		}
	}
	return &res
}

func (ee *edges) FindNeighbours() *neighbours {
	res := make(map[string](map[string]bool))
	for _, edgeList := range *ee {
		var prevEdges []string
		for _, edge := range edgeList {
			edgeTitle := edge.Title
			if res[edgeTitle] == nil {
				res[edgeTitle] = make(map[string]bool)
			}
			for _, neighbour := range prevEdges {
				res[edgeTitle][neighbour] = true
				res[neighbour][edgeTitle] = true
			}
			prevEdges = append(prevEdges, edgeTitle)
		}
	}
	simpleRes := make(neighbours)
	for tile, mapOfTiles := range res {
		for neighbour := range mapOfTiles {
			simpleRes[tile] = append(simpleRes[tile], neighbour)
		}
		sort.Strings(simpleRes[tile])
	}
	return &simpleRes
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

func (nn *neighbours) BuildTileMap(start string, size int) *[][]string {
	res := [][]string{{start}, {}, {}}

	// code to build tile map manually
	// nn := edges.FindNeighbours()
	// n := len(tileMap) - 1
	// used := []string{}
	// for i := 0; i <= n; i++ {
	// 	for j := 0; j < len(tileMap[i]); j++ {
	// 		used = append(used, tileMap[i][j])
	// 	}
	// }
	// sort.Strings(used)
	// m := len(tileMap[11]) - 1
	// for i := m; i <= n; i++ {
	// 	tile := tileMap[i][n+m-i]
	// 	unused := rejectUsed((*nn)[tile], used)
	// 	fmt.Println(unused)
	// }

	return &res
}

func rejectUsed(tiles, used []string) (res []string) {
	for _, tile := range tiles {
		unused := true
		for _, u := range used {
			if tile == u {
				unused = false
				break
			}
		}
		if unused {
			res = append(res, tile)
		}
	}
	return
}

func (nn *neighbours) String() (res string) {
	for tile, neighbourList := range *nn {
		res += fmt.Sprintf("%s:", tile)
		for _, neighbourTile := range neighbourList {
			res += fmt.Sprintf(" %s", neighbourTile)
		}
		res += fmt.Sprintf("\n")
	}
	return
}
