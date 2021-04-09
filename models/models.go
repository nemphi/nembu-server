package models

import "time"

type dbFields struct {
	UID       string    `json:"uid,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	DType     []string  `json:"dgraph.type,omitempty"`
}

type User struct {
	dbFields

	Name          string `json:"name,omitempty"`
	Surname       string `json:"surname,omitempty"`
	Email         string `json:"email,omitempty"`
	Password      string `json:"-,omitempty"`
	PersonalPhone string `json:"personal_phone,omitempty"`
	AuthProvider  string `json:"auth_provider,omitempty"`
}

type Media struct {
	dbFields

	Title           string `json:"title,omitempty"`
	AltText         string `json:"alt_text,omitempty"`
	Path            string `json:"path,omitempty"`
	StorageProvider string `json:"storage_provider,omitempty"`
}

type Resource struct {
	dbFields

	Collection       string    `json:"collection,omitempty"`
	Slug             string    `json:"slug,omitempty"`
	Title            string    `json:"title,omitempty"`
	Content          string    `json:"content,omitempty"`
	TextContent      string    `json:"text_content,omitempty"`
	TotalViews       uint      `json:"total_views,omitempty"`
	TotalComments    uint      `json:"total_comments,omitempty"`
	Public           bool      `json:"public,omitempty"`
	CommentsAllowed  bool      `json:"comments_allowed,omitempty"`
	CreatedByID      string    `json:"created_by_id,omitempty"`
	CreatedBy        *User     `json:"created_by,omitempty"` // User.ID
	ParentResourceID string    `json:"parent_resource_id,omitempty"`
	ParentResource   *Resource `json:"parent_resource,omitempty"` // Resource.ID
}

type Page = Resource

type Post = Resource

type Comment struct {
	dbFields

	Content         string    `json:"content,omitempty"`
	TextContent     string    `json:"text_content,omitempty"`
	CreatedByID     string    `json:"created_by_id,omitempty"`
	CreatedBy       *User     `json:"created_by,omitempty"` // User.ID
	ResourceID      string    `json:"resource_id,omitempty"`
	Resource        *Resource `json:"resource,omitempty"` // Resource.ID
	ParentCommentID string    `json:"parent_comment_id,omitempty"`
	ParentComment   *Comment  `json:"parent_comment,omitempty"` // Comment.ID
}

type AggregateResourceViews struct {
	dbFields

	Views      uint      `json:"views,omitempty"`
	Comments   uint      `json:"comments,omitempty"`
	ResourceID string    `json:"resource_id,omitempty"`
	Resource   *Resource `json:"resource,omitempty"` // Resource.ID
}

type Event struct {
	dbFields

	Type string                 `json:"type,omitempty"`
	Meta map[string]interface{} `json:"meta,omitempty"`
}

type Template struct {
	dbFields

	Name        string
	Content     string
	TextContent string
	Disabled    bool
}

type StorageProvider struct {
	dbFields

	Name              string
	ConnectionURL     string
	AuthToken         string
	RefreshTokenEvery time.Duration
}

type AuthProvider struct {
	dbFields

	Name              string
	ConnectionURL     string
	AuthToken         string
	RefreshTokenEvery time.Duration
}

type Plugin struct {
	dbFields

	Name     string
	MarketID string
}

func AllModels() []interface{} {
	return []interface{}{
		(*User)(nil),
		(*Media)(nil),
		(*Resource)(nil),
		(*Comment)(nil),
		(*AggregateResourceViews)(nil),
		(*Template)(nil),
		(*StorageProvider)(nil),
		(*AuthProvider)(nil),
		(*Plugin)(nil),
	}
}
