// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: post.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const deleteCommentFromPost = `-- name: DeleteCommentFromPost :one
DELETE FROM "comment"
WHERE id = $1 AND author_id = $2 AND post_id = $3
RETURNING (id, content, author_id)
`

type DeleteCommentFromPostParams struct {
	ID       uuid.UUID `json:"id"`
	AuthorID uuid.UUID `json:"author_id"`
	PostID   uuid.UUID `json:"post_id"`
}

func (q *Queries) DeleteCommentFromPost(ctx context.Context, arg DeleteCommentFromPostParams) (interface{}, error) {
	row := q.queryRow(ctx, q.deleteCommentFromPostStmt, deleteCommentFromPost, arg.ID, arg.AuthorID, arg.PostID)
	var column_1 interface{}
	err := row.Scan(&column_1)
	return column_1, err
}

const deletePostFromClass = `-- name: DeletePostFromClass :one
DELETE FROM "post"
WHERE id = $1 AND author_id = $2 AND class_id = $3
RETURNING (id, content, author_id)
`

type DeletePostFromClassParams struct {
	ID       uuid.UUID `json:"id"`
	AuthorID uuid.UUID `json:"author_id"`
	ClassID  uuid.UUID `json:"class_id"`
}

func (q *Queries) DeletePostFromClass(ctx context.Context, arg DeletePostFromClassParams) (interface{}, error) {
	row := q.queryRow(ctx, q.deletePostFromClassStmt, deletePostFromClass, arg.ID, arg.AuthorID, arg.ClassID)
	var column_1 interface{}
	err := row.Scan(&column_1)
	return column_1, err
}

const getAllCommentsFromPost = `-- name: GetAllCommentsFromPost :many
SELECT id, content, author_id, post_id, created_at, updated_at
FROM "comment"
WHERE post_id = $1
ORDER BY created_at
ASC
`

func (q *Queries) GetAllCommentsFromPost(ctx context.Context, postID uuid.UUID) ([]Comment, error) {
	rows, err := q.query(ctx, q.getAllCommentsFromPostStmt, getAllCommentsFromPost, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.AuthorID,
			&i.PostID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOnePost = `-- name: GetOnePost :one
SELECT id, content, author_id, class_id, created_at, updated_at
FROM "post"
WHERE id = $1 AND class_id = $2
LIMIT 1
`

type GetOnePostParams struct {
	ID      uuid.UUID `json:"id"`
	ClassID uuid.UUID `json:"class_id"`
}

func (q *Queries) GetOnePost(ctx context.Context, arg GetOnePostParams) (Post, error) {
	row := q.queryRow(ctx, q.getOnePostStmt, getOnePost, arg.ID, arg.ClassID)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.AuthorID,
		&i.ClassID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertNewCommentInPost = `-- name: InsertNewCommentInPost :one
INSERT INTO "comment" (
  id, content, author_id, post_id
) VALUES (
  $1, $2, $3, $4
) RETURNING id, content, author_id, post_id, created_at, updated_at
`

type InsertNewCommentInPostParams struct {
	ID       uuid.UUID `json:"id"`
	Content  string    `json:"content"`
	AuthorID uuid.UUID `json:"author_id"`
	PostID   uuid.UUID `json:"post_id"`
}

func (q *Queries) InsertNewCommentInPost(ctx context.Context, arg InsertNewCommentInPostParams) (Comment, error) {
	row := q.queryRow(ctx, q.insertNewCommentInPostStmt, insertNewCommentInPost,
		arg.ID,
		arg.Content,
		arg.AuthorID,
		arg.PostID,
	)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.AuthorID,
		&i.PostID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertNewPost = `-- name: InsertNewPost :one
INSERT INTO "post" (
  id, content, author_id, class_id
) VALUES ( $1, $2, $3, $4 )
RETURNING id, content, author_id, class_id, created_at, updated_at
`

type InsertNewPostParams struct {
	ID       uuid.UUID `json:"id"`
	Content  string    `json:"content"`
	AuthorID uuid.UUID `json:"author_id"`
	ClassID  uuid.UUID `json:"class_id"`
}

func (q *Queries) InsertNewPost(ctx context.Context, arg InsertNewPostParams) (Post, error) {
	row := q.queryRow(ctx, q.insertNewPostStmt, insertNewPost,
		arg.ID,
		arg.Content,
		arg.AuthorID,
		arg.ClassID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.AuthorID,
		&i.ClassID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAllPostsByUser = `-- name: ListAllPostsByUser :many
SELECT id, content, author_id, class_id, created_at, updated_at
FROM "post"
WHERE author_id = $1 AND class_id = $2
ORDER BY created_at
ASC
`

type ListAllPostsByUserParams struct {
	AuthorID uuid.UUID `json:"author_id"`
	ClassID  uuid.UUID `json:"class_id"`
}

func (q *Queries) ListAllPostsByUser(ctx context.Context, arg ListAllPostsByUserParams) ([]Post, error) {
	rows, err := q.query(ctx, q.listAllPostsByUserStmt, listAllPostsByUser, arg.AuthorID, arg.ClassID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.AuthorID,
			&i.ClassID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAllPostsFromClass = `-- name: ListAllPostsFromClass :many
SELECT id, content, author_id, class_id, created_at, updated_at
FROM "post"
WHERE class_id = $1
ORDER BY created_at
ASC
`

func (q *Queries) ListAllPostsFromClass(ctx context.Context, classID uuid.UUID) ([]Post, error) {
	rows, err := q.query(ctx, q.listAllPostsFromClassStmt, listAllPostsFromClass, classID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.AuthorID,
			&i.ClassID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCommentContentInPost = `-- name: UpdateCommentContentInPost :one
UPDATE "comment"
SET content = $1
WHERE id = $2 AND author_id = $3 AND post_id = $4
RETURNING (id, content, author_id)
`

type UpdateCommentContentInPostParams struct {
	Content  string    `json:"content"`
	ID       uuid.UUID `json:"id"`
	AuthorID uuid.UUID `json:"author_id"`
	PostID   uuid.UUID `json:"post_id"`
}

func (q *Queries) UpdateCommentContentInPost(ctx context.Context, arg UpdateCommentContentInPostParams) (interface{}, error) {
	row := q.queryRow(ctx, q.updateCommentContentInPostStmt, updateCommentContentInPost,
		arg.Content,
		arg.ID,
		arg.AuthorID,
		arg.PostID,
	)
	var column_1 interface{}
	err := row.Scan(&column_1)
	return column_1, err
}

const updatePostContent = `-- name: UpdatePostContent :one
UPDATE "post"
SET content = $1
WHERE id = $2 AND author_id = $3 AND class_id = $4
RETURNING (id, content, author_id)
`

type UpdatePostContentParams struct {
	Content  string    `json:"content"`
	ID       uuid.UUID `json:"id"`
	AuthorID uuid.UUID `json:"author_id"`
	ClassID  uuid.UUID `json:"class_id"`
}

func (q *Queries) UpdatePostContent(ctx context.Context, arg UpdatePostContentParams) (interface{}, error) {
	row := q.queryRow(ctx, q.updatePostContentStmt, updatePostContent,
		arg.Content,
		arg.ID,
		arg.AuthorID,
		arg.ClassID,
	)
	var column_1 interface{}
	err := row.Scan(&column_1)
	return column_1, err
}
