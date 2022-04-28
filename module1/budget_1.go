package module1

// Budget stores budget information
type Item struct {
	Max float32
	Description string
	Price float32
	Items []Item
}

// Item stores item information
