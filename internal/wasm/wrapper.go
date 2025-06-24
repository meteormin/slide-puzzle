package wasm

type Tile struct {
	DomID        string `json:"domId"`
	DisplayValue int    `json:"displayValue"`
	Empty        bool   `json:"fill"`
}

type Board struct {
	Size  int    `json:"size"`
	Tiles []Tile `json:"tiles"`
}
