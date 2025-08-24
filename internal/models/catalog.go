// Package models provides models for a simple in-memory book catalog.
package models

import "fmt"

type BookInfo struct {
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
	Pages  int    `json:"pages,omitempty"`
}

func (b BookInfo) String() string {
	return fmt.Sprintf("%s by %s (%d pages)", b.Title, b.Author, b.Pages)
}

type Catalog struct {
	books []BookInfo
}

func NewCatalog() *Catalog {
	return &Catalog{
		books: make([]BookInfo, 0),
	}
}

func (c *Catalog) AddBook(title, author string, pages int) {
	book := BookInfo{
		Title:  title,
		Author: author,
		Pages:  pages,
	}
	c.books = append(c.books, book)
}

func (c *Catalog) GetBooks() []BookInfo {
	if c.books == nil {
		return []BookInfo{}
	}
	return c.books
}

func (c *Catalog) GetBookCount() int {
	return len(c.books)
}

func (c *Catalog) FindByAuthor(author string) []BookInfo {
	var result []BookInfo
	for _, book := range c.books {
		if book.Author == author {
			result = append(result, book)
		}
	}
	return result
}
