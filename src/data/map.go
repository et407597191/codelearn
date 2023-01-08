package data

// 位置
type Position struct {
	X int
	Y int
}

// FlatMap 平面地图
type FlatMap struct {
	xborder int
	yborder int
	matrix  [][]int
}

// Line 路线就是位置切片
type Line []Position
