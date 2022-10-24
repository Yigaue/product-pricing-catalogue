package entities

type Product struct {
	ID uint `json."id"`
	Name string `json."name"`
	Price string `json."price"`
	Description string `json."description"`
}