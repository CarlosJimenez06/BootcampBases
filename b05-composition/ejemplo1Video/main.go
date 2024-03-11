package main

import "fmt"

type Publisher struct {
	Name     string
	Founders []string
}

func (p Publisher) TotalFounders() int {
	return len(p.Founders)
}

/*--------------------------------------------*/

type Book struct {
	Title  string
	Author string
	Pages  int
	Publisher
}

func (b Book) ShowInfo() string {
	return fmt.Sprintf("The book %s has %d pages and was written by %s. Its total publisher founder: %d",
		b.Title, b.Pages, b.Author, b.Publisher.TotalFounders())
}

/*--------------------------------------------*/

type Movie struct {
	Title      string
	Director   string
	IMDBRating float64
	Publisher
}

func (m Movie) ShowInfo() string {
	return fmt.Sprintf("The movie %s has IMDB rating of %f and was directed by %s. Its total publisher founder: %d",
		m.Title, m.IMDBRating, m.Director, m.TotalFounders())
}

/*--------------------------------------------*/

func main() {

	// Instantiating the structs
	book := Book{
		Title:  "Harry Potter and the Philosopher's Stone",
		Author: "J",
		Pages:  223,
		Publisher: Publisher{
			Name:     "Bloomsbury",
			Founders: []string{"Siobhan Reddy", "Nicola Yoon", "J.K. Rowling"},
		},
	}

	movie := Movie{
		Title:      "Titanic",
		Director:   "James Cameron",
		IMDBRating: 7.8,
		Publisher: Publisher{
			Name:     "Paramount Pictures",
			Founders: []string{"Vince Vaughn", "Jon Favreau"},
		},
	}

	// The struct gral has access further to the Publisher struct as methods
	fmt.Println("Number of founders on Book: ", book.TotalFounders())
	fmt.Println("Number of founders on movie: ", movie.TotalFounders())
}
