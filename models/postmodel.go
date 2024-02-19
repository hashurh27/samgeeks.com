package models

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type Post struct {
	Title    string
	Content  string
	AuthorID string
	Slug     string
}

func GetPostByID(db *sql.DB, id int) (Post, error) {
	var post Post
	err := db.QueryRow(`SELECT title,content,authorId FROM post WHERE id = ?`, id).Scan(&post.Title, &post.Content, &post.AuthorID)
	if err != nil {
		if err == sql.ErrNoRows {
			return Post{}, fmt.Errorf("post with ID %d not found", id)
		}
		return Post{}, fmt.Errorf("error retrieving post: %w", err)
	}
	fmt.Println(post)
	return post, nil
}

func InsertPost(db *sql.DB, title, content, slug, authorID string) error {
	// Validate input
	if err := validatePostInput(title, content, slug, authorID); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	stmt, err := tx.Prepare("INSERT INTO post (title, content, slug, authorId) VALUES (?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, content, slug, authorID)
	if err != nil {
		// Handle specific errors, e.g., duplicate key
		if strings.Contains(err.Error(), "Duplicate entry") {
			return fmt.Errorf("post with slug '%s' already exists", slug)
		}
		return fmt.Errorf("error inserting post: %w", err)
	}

	// Log success or handle result (e.g., return inserted post ID)
	log.Printf("Post inserted successfully")

	return nil
}

func validatePostInput(title, content, slug, authorID string) error {
	// Add your validation logic here
	return nil
}

//if err != nil {
//	fmt.Errorf("erro %w query", err)
//}
