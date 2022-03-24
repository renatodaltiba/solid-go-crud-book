package repositories

import (
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/renatodalitba/books-solid-go-lang/modules/books/entities"
	"github.com/renatodalitba/books-solid-go-lang/pkg/config"
	database "github.com/renatodalitba/books-solid-go-lang/pkg/database/postgres"
)

type repo struct{}

func NewPostgresRepository() IBookRepository {
	return &repo{}
}
func (*repo) Books() ([]entities.Book, error) {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration: ", err)
		os.Exit(-1)
	}

	db, err := database.Connect(cfg.Database)
	if err != nil {
		log.Fatal("Failed to load configuration: ", err)
		os.Exit(1)
	}
	db.Close()

	var books []entities.Book
	err = db.Select(&books, "SELECT * FROM books")
	if err != nil {
		log.Fatal("Failed to load configuration: ", err)
		os.Exit(1)
	}

	return books, err
}

func (*repo) CreateBook(book *entities.Book) error {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration: ", err)
		os.Exit(-1)
	}

	db, err := database.Connect(cfg.Database)
	if err != nil {
		log.Fatal("Failed to load configuration: ", err)
		os.Exit(1)
	}

	_, err = db.Exec("INSERT INTO books (title, author, description) VALUES ($1, $2, $3)", book.Title, book.Author, book.Description)
	if err != nil {
		log.Fatal("Failed to insert on database: ", err)
		os.Exit(1)
	}

	return err
}

func loadConfig() (*config.Config, error) {
	// Configuration
	configPath := os.Getenv("SUMELMS_CONFIG_PATH")
	if configPath == "" {
		configPath = "./config/config.yml"
	}

	cfg, err := config.NewConfig(configPath)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
