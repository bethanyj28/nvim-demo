package main

import (
	"fmt"

	"nvim-demo/internal/models"
)

type Book struct {
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
	Pages  int    `json:"pages,omitempty"`
}

type Reader interface {
	Read(book Book) string
	GetFavoriteGenre() string
}

type Person struct {
	Name          string
	FavoriteGenre string
}

func (p Person) Read(book Book) string {
	return fmt.Sprintf("%s is reading '%s' by %s (%d pages)",
		p.Name, book.Title, book.Author, book.Pages)
}

func (p Person) GetFavoriteGenre() string {
	return p.FavoriteGenre
}

type Bot struct {
	Name          string
	FavoriteGenre string
	Version       string
}

func (b Bot) Read(book Book) string {
	return fmt.Sprintf("🤖 Bot %s v%s analyzing '%s' by %s (%d pages) - Processing complete!",
		b.Name, b.Version, book.Title, book.Author, book.Pages)
}

func (b Bot) GetFavoriteGenre() string {
	return b.FavoriteGenre
}

func main() {
	book := Book{
		Title:  "The Go Programming Language",
		Author: "Alan Donovan",
		Pages:  380,
	}

	humanReader := Person{
		Name:          "Alice",
		FavoriteGenre: "Technical",
	}

	botReader := Bot{
		Name:          "BookBot",
		Version:       "2.1",
		FavoriteGenre: "Science Fiction",
	}

	readers := []Reader{humanReader, botReader}

	fmt.Println("=== Multiple Reader Demo ===")
	for _, reader := range readers {
		fmt.Println(reader.Read(book))
		fmt.Printf("Favorite genre: %s\n\n", reader.GetFavoriteGenre())
	}

	catalog := models.NewCatalog()
	catalog.AddBook("Clean Code", "Robert Martin", 464)
	catalog.AddBook("Design Patterns", "Gang of Four", 395)

	fmt.Println("Library catalog:")
	for _, book := range catalog.GetBooks() {
		fmt.Printf("- %s\n", book)
	}
}
