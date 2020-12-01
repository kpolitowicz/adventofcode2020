package main

import (
	"math"
	"strconv"
	"strings"
)

type wireDef struct {
	direction byte
	steps     int
}

type point struct {
	x, y int
}

type pointWithSignalDist struct {
	x, y, signalDist int
}

func CalcManhattanDistance(wire1Str, wire2Str string) int {
	wire1 := ConvertToWireDef(ParseInputString(wire1Str))
	wire2 := ConvertToWireDef(ParseInputString(wire2Str))

	wire1_points := GetWirePoints(wire1)
	wire2_points := GetWirePoints(wire2)

	intersections := FindCommonPoints(wire1_points, wire2_points)
	closest_intersection := FindClosestPointByManhattanDist(intersections)

	return ManhattanDistance(closest_intersection)
}

func CalcSignalDistance(wire1Str, wire2Str string) int {
	wire1 := ConvertToWireDef(ParseInputString(wire1Str))
	wire2 := ConvertToWireDef(ParseInputString(wire2Str))

	wire1_points := GetWirePoints(wire1)
	wire2_points := GetWirePoints(wire2)

	intersections := FindCommonPoints(wire1_points, wire2_points)
	closest_intersection := FindClosestPointBySignalDist(intersections)

	return closest_intersection.signalDist
}

func ParseInputString(inputStr string) []string {
	return strings.Split(inputStr, ",")
}

func ConvertToWireDef(wireDirectives []string) (res []wireDef) {
	for _, wireDir := range wireDirectives {
		res = append(res, parseWireDirectiveIntoDef(wireDir))
	}
	return
}

func GetWirePoints(wireDefs []wireDef) (res []point) {
	currentPos := point{0, 0}
	for _, wireDef := range wireDefs {
		res = append(res, pointsFrom(wireDef, currentPos)...)
		currentPos = res[len(res)-1]
	}
	return
}

// TODO: implement wire args as set with intersection
func FindCommonPoints(wire1, wire2 []point) (res []pointWithSignalDist) {
	for idx1, wire1_point := range wire1 {
		for idx2, wire2_point := range wire2 {
			if wire1_point == wire2_point {
				res = append(res, pointWithSignalDist{wire1_point.x, wire1_point.y, idx1 + idx2 + 2})
			}
		}
	}
	return
}

func FindClosestPointByManhattanDist(points []pointWithSignalDist) (res pointWithSignalDist) {
	shortestDist := 1_000_000_000
	for _, p := range points {
		dist := ManhattanDistance(p)
		if dist < shortestDist {
			shortestDist = dist
			res = p
		}
	}
	return
}

func FindClosestPointBySignalDist(points []pointWithSignalDist) (res pointWithSignalDist) {
	shortestDist := 1_000_000_000
	for _, p := range points {
		if p.signalDist < shortestDist {
			shortestDist = p.signalDist
			res = p
		}
	}
	return
}

func ManhattanDistance(p pointWithSignalDist) int {
	return int(math.Abs(float64(p.x))) + int(math.Abs(float64(p.y)))
}

func parseWireDirectiveIntoDef(wireDir string) wireDef {
	dir := wireDir[0]
	steps, _ := strconv.Atoi(string(wireDir[1:len(wireDir)]))
	return wireDef{dir, steps}
}

func pointsFrom(wireD wireDef, startPos point) (res []point) {
	switch wireD.direction {
	case 'R':
		res = append(res, moveRight(wireD.steps, startPos)...)
	case 'L':
		res = append(res, moveLeft(wireD.steps, startPos)...)
	case 'U':
		res = append(res, moveUp(wireD.steps, startPos)...)
	case 'D':
		res = append(res, moveDown(wireD.steps, startPos)...)
	}
	return
}

func moveRight(steps int, startPos point) (res []point) {
	for i := 1; i <= steps; i++ {
		res = append(res, point{startPos.x + i, startPos.y})
	}
	return
}

func moveLeft(steps int, startPos point) (res []point) {
	for i := 1; i <= steps; i++ {
		res = append(res, point{startPos.x - i, startPos.y})
	}
	return
}

func moveUp(steps int, startPos point) (res []point) {
	for i := 1; i <= steps; i++ {
		res = append(res, point{startPos.x, startPos.y + i})
	}
	return
}

func moveDown(steps int, startPos point) (res []point) {
	for i := 1; i <= steps; i++ {
		res = append(res, point{startPos.x, startPos.y - i})
	}
	return
}
