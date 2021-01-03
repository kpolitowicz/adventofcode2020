package jurassicjigsaw

import (
	"reflect"
	"testing"
)

func TestFindSigletons(t *testing.T) {
	edgesToFind := &edges{
		"###..###..": {{"2311", "B", true}},
		".#####..#.": {{"2311", "L", false}, {"1951", "R", false}},
		".#..#####.": {{"2311", "L", true}, {"1951", "R", true}},
		"#.##...##.": {{"1951", "T", false}},
		".##...##.#": {{"1951", "L", false}},
	}
	got := edgesToFind.FindSigletons()
	want := &edges{
		"2311": {{"2311", "B", true}},
		"1951": {{"1951", "T", false}, {"1951", "L", false}},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestFindNeighbours(t *testing.T) {
	edgesToFind := &edges{
		"###..###..": {{"2311", "B", true}},
		".#####..#.": {{"2311", "L", false}, {"1951", "R", false}},
		".#..#####.": {{"2311", "L", true}, {"1951", "R", true}},
		"#.##...##.": {{"1951", "T", false}, {"1117", "T", false}},
		".##...##.#": {{"1951", "L", false}},
	}
	got := edgesToFind.FindNeighbours()
	want := &neighbours{
		"2311": {"1951"},
		"1951": {"1117", "2311"},
		"1117": {"1951"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestGetEdges(t *testing.T) {
	images := &imageMap{
		"2311": {
			"..##.#..#.",
			"##..#.....",
			"#...##..#.",
			"####.#...#",
			"##.##.###.",
			"##...#.###",
			".#.#.#..##",
			"..#....#..",
			"###...#.#.",
			"..###..###",
		},
		"1951": {
			"#.##...##.",
			"#.####...#",
			".....#..##",
			"#...######",
			".##.#....#",
			".###.#####",
			"###.##.##.",
			".###....#.",
			"..#.#..#.#",
			"#...##.#..",
		},
	}
	got := images.GetEdges()
	want := &edges{
		"..##.#..#.": {{"2311", "T", false}},
		".#..#.##..": {{"2311", "T", true}},
		"..###..###": {{"2311", "B", false}},
		"###..###..": {{"2311", "B", true}},
		".#####..#.": {{"2311", "L", false}, {"1951", "R", false}},
		".#..#####.": {{"2311", "L", true}, {"1951", "R", true}},
		"...#.##..#": {{"2311", "R", false}},
		"#..##.#...": {{"2311", "R", true}},
		"#.##...##.": {{"1951", "T", false}},
		".##...##.#": {{"1951", "T", true}},
		"#...##.#..": {{"1951", "B", false}},
		"..#.##...#": {{"1951", "B", true}},
		"##.#..#..#": {{"1951", "L", false}},
		"#..#..#.##": {{"1951", "L", true}},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestBuildTileMap(t *testing.T) {
	images := ParseInput(`Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...`)
	myEdges := images.GetEdges()
	tileMap := &[][]string{
		{"1951", "2311", "3079"},
		{"2729", "1427", "2473"},
		{"2971", "1489", "1171"},
	}
	_, _ = myEdges, tileMap
	// if !reflect.DeepEqual(got, want) {
	// 	t.Errorf("got %v want %v", got, want)
	// }
}
