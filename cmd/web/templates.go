package main

import "sunilprz.np/snippetbox/internal/models"

type templateData  struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}