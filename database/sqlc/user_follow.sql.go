// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: user_follow.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const followUser = `-- name: FollowUser :one
INSERT INTO user_follows
(id, follower, following)
VALUES ($1, $2, $3)
RETURNING id, follower, following, created_at, updated_at
`

type FollowUserParams struct {
	ID        uuid.UUID `json:"id"`
	Follower  uuid.UUID `json:"follower"`
	Following uuid.UUID `json:"following"`
}

func (q *Queries) FollowUser(ctx context.Context, arg FollowUserParams) (UserFollow, error) {
	row := q.queryRow(ctx, q.followUserStmt, followUser, arg.ID, arg.Follower, arg.Following)
	var i UserFollow
	err := row.Scan(
		&i.ID,
		&i.Follower,
		&i.Following,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllFollowers = `-- name: GetAllFollowers :many
SELECT id, follower AS user_id, created_at, updated_at
FROM user_follows
WHERE following = $1
ORDER BY created_at
LIMIT 10
OFFSET $2
`

type GetAllFollowersParams struct {
	Following uuid.UUID `json:"following"`
	Offset    int32     `json:"offset"`
}

type GetAllFollowersRow struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) GetAllFollowers(ctx context.Context, arg GetAllFollowersParams) ([]GetAllFollowersRow, error) {
	rows, err := q.query(ctx, q.getAllFollowersStmt, getAllFollowers, arg.Following, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllFollowersRow
	for rows.Next() {
		var i GetAllFollowersRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
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

const getAllFollowing = `-- name: GetAllFollowing :many
SELECT id, following AS user_id, created_at, updated_at
FROM user_follows
WHERE follower = $1
ORDER BY created_at
LIMIT 10
OFFSET $2
`

type GetAllFollowingParams struct {
	Follower uuid.UUID `json:"follower"`
	Offset   int32     `json:"offset"`
}

type GetAllFollowingRow struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) GetAllFollowing(ctx context.Context, arg GetAllFollowingParams) ([]GetAllFollowingRow, error) {
	rows, err := q.query(ctx, q.getAllFollowingStmt, getAllFollowing, arg.Follower, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllFollowingRow
	for rows.Next() {
		var i GetAllFollowingRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
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

const getFollowerById = `-- name: GetFollowerById :one
SELECT id, follower AS user_id, created_at, updated_at
FROM user_follows
WHERE id = $1
LIMIT 1
`

type GetFollowerByIdRow struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) GetFollowerById(ctx context.Context, id uuid.UUID) (GetFollowerByIdRow, error) {
	row := q.queryRow(ctx, q.getFollowerByIdStmt, getFollowerById, id)
	var i GetFollowerByIdRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getOneFollower = `-- name: GetOneFollower :one
SELECT id, follower, following, created_at, updated_at
FROM user_follows
WHERE follower = $1 AND following = $2
LIMIT 1
`

type GetOneFollowerParams struct {
	Follower  uuid.UUID `json:"follower"`
	Following uuid.UUID `json:"following"`
}

func (q *Queries) GetOneFollower(ctx context.Context, arg GetOneFollowerParams) (UserFollow, error) {
	row := q.queryRow(ctx, q.getOneFollowerStmt, getOneFollower, arg.Follower, arg.Following)
	var i UserFollow
	err := row.Scan(
		&i.ID,
		&i.Follower,
		&i.Following,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const unfollowUser = `-- name: UnfollowUser :one
DELETE FROM user_follows
WHERE id = $1
RETURNING id, follower, following, created_at, updated_at
`

func (q *Queries) UnfollowUser(ctx context.Context, id uuid.UUID) (UserFollow, error) {
	row := q.queryRow(ctx, q.unfollowUserStmt, unfollowUser, id)
	var i UserFollow
	err := row.Scan(
		&i.ID,
		&i.Follower,
		&i.Following,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
