package models

import "testing"

func TestNewCatalog(t *testing.T) {
	catalog := NewCatalog()
	if catalog == nil {
		t.Fatal("NewCatalog() returned nil")
	}
	if catalog.GetBookCount() != 0 {
		t.Errorf("Expected empty catalog, got %d books", catalog.GetBookCount())
	}
}

func TestAddBook(t *testing.T) {
	catalog := NewCatalog()
	catalog.AddBook("Test Book", "Test Author", 100)

	if catalog.GetBookCount() != 1 {
		t.Errorf("Expected 1 book, got %d", catalog.GetBookCount())
	}

	books := catalog.GetBooks()
	if len(books) != 1 {
		t.Errorf("Expected 1 book in slice, got %d", len(books))
	}

	book := books[0]
	if book.Title != "Test Book" {
		t.Errorf("Expected title 'Test Book', got '%s'", book.Title)
	}
	if book.Author != "Test Author" {
		t.Errorf("Expected author 'Test Author', got '%s'", book.Author)
	}
	if book.Pages != 100 {
		t.Errorf("Expected 100 pages, got %d", book.Pages)
	}
}

func TestFindByAuthor(t *testing.T) {
	catalog := NewCatalog()
	catalog.AddBook("Book 1", "Author A", 200)
	catalog.AddBook("Book 2", "Author B", 300)
	catalog.AddBook("Book 3", "Author A", 250)

	books := catalog.FindByAuthor("Author A")
	if len(books) != 2 {
		t.Errorf("Expected 2 books by Author A, got %d", len(books))
	}

	books = catalog.FindByAuthor("Author C")
	if len(books) != 0 {
		t.Errorf("Expected 0 books by Author C, got %d", len(books))
	}
}
