package sentry

import (
	"fmt"
	"time"
)

// Search is a saved search for a given project.
type Search struct {
	DateCreated   *time.Time `json:"dateCreated,omitempty"`
	ID            int        `json:"id"`
	IsDefault     bool       `json:"isDefault"`
	IsPrivate     bool       `json:"isPrivate"`
	IsUserDefault bool       `json:"isUserDefault"`
	Name          string     `json:"name"`
	Query         string     `json:"query"`
}

// GetSearches returns the list of stored searches
func (c *Client) GetSearches(o Organization, p Project) ([]Search, error) {
	var searches []Search
	err := c.do("GET", fmt.Sprintf("projects/%s/%s/searches", *o.Slug, *p.Slug), &searches, nil)
	return searches, err
}
