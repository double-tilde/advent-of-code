package model

type SearchInput struct {
	Word       string
	Directions [][]int
}

type WordPosition struct {
	Char string
	Row  int
	Col  int
}
