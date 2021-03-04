package models

type dbFields struct {
	ID        string `json:"id,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

type User struct {
	dbFields `json:"db_fields,omitempty"`

	Name          string `json:"name,omitempty"`
	Surname       string `json:"surname,omitempty"`
	Email         string `json:"email,omitempty"`
	Password      string `json:"-"`
	PersonalPhone string `json:"personal_phone,omitempty"`
}

type Media struct {
	dbFields

	Title    string `json:"title,omitempty"`
	AltText  string `json:"alt_text,omitempty"`
	Path     string `json:"path,omitempty"`
	Provider string `json:"provider,omitempty"`
}

type Resource struct {
	dbFields

	Collection       string `json:"collection,omitempty"`
	Slug             string `json:"slug,omitempty"`
	Title            string `json:"title,omitempty"`
	Content          string `json:"content,omitempty"`
	TextContent      string `json:"text_content,omitempty"`
	TotalViews       uint   `json:"total_views,omitempty"`
	TotalComments    uint   `json:"total_comments,omitempty"`
	Public           bool   `json:"public,omitempty"`
	CommentsAllowed  bool   `json:"comments_allowed,omitempty"`
	CreatedByID      string `json:"created_by_id,omitempty"`      // User.ID
	ParentResourceID string `json:"parent_resource_id,omitempty"` // Resource.ID
}

type Page = Resource

type Post = Resource

type Comment struct {
	dbFields

	Content         string `json:"content,omitempty"`
	TextContent     string `json:"text_content,omitempty"`
	CreatedByID     string `json:"created_by_id,omitempty"`     // User.ID
	ResourceID      string `json:"resource_id,omitempty"`       // Resource.ID
	ParentCommentID string `json:"parent_comment_id,omitempty"` // Comment.ID
}

type AggregateResourceViews struct {
	dbFields

	Views      uint   `json:"views,omitempty"`
	Comments   uint   `json:"comments,omitempty"`
	ResourceID string `json:"resource_id,omitempty"` // Resource.ID
}
