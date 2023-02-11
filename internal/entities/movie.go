package entities

type Movie struct {
	Title string `json:"title"`
	Group string `json:"group"`
	Year  int    `json:"year"`
	Age   int    `json:"age"`
}
