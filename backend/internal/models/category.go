package models

type Category struct {
	ID       int64      `json:"id"`
	ParentID *int64     `json:"parent_id,omitempty"`
	Name     string     `json:"name"`
	Slug     string     `json:"slug"`
	Children []Category `json:"children,omitempty"`
}
