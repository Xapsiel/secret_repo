package model

type Page struct {
	Domain       string
	Title        string
	Keywords     string
	Description  string
	HTTPStatus   int
	IsProduction bool
}
