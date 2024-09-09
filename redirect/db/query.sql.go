// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createVisitor = `-- name: CreateVisitor :exec
INSERT INTO visitor (
		shortener_id,
		device_type,
		device_vendor,
		browser,
		os,
		country,
		country_code,
		city,
		referer
	)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
`

type CreateVisitorParams struct {
	ShortenerID  string
	DeviceType   string
	DeviceVendor string
	Browser      string
	Os           string
	Country      string
	CountryCode  string
	City         string
	Referer      string
}

func (q *Queries) CreateVisitor(ctx context.Context, arg CreateVisitorParams) error {
	_, err := q.db.Exec(ctx, createVisitor,
		arg.ShortenerID,
		arg.DeviceType,
		arg.DeviceVendor,
		arg.Browser,
		arg.Os,
		arg.Country,
		arg.CountryCode,
		arg.City,
		arg.Referer,
	)
	return err
}

const getShortener = `-- name: GetShortener :one
SELECT id, link, code, created_at, user_id, project_id, active, ios, ios_link, android, android_link
FROM shortener
WHERE code = $1
LIMIT 1
`

func (q *Queries) GetShortener(ctx context.Context, code string) (Shortener, error) {
	row := q.db.QueryRow(ctx, getShortener, code)
	var i Shortener
	err := row.Scan(
		&i.ID,
		&i.Link,
		&i.Code,
		&i.CreatedAt,
		&i.UserID,
		&i.ProjectID,
		&i.Active,
		&i.Ios,
		&i.IosLink,
		&i.Android,
		&i.AndroidLink,
	)
	return i, err
}

const getShortenerWithDomain = `-- name: GetShortenerWithDomain :one
SELECT shortener.id, shortener.link, shortener.code, shortener.created_at, shortener.user_id, shortener.project_id, shortener.active, shortener.ios, shortener.ios_link, shortener.android, shortener.android_link,
	project.custom_domain as domain
FROM shortener
	LEFT JOIN project ON project.id = shortener.project_id
WHERE shortener.code = $1
	AND project.custom_domain = $2
	AND project.enable_custom_domain IS TRUE
LIMIT 1
`

type GetShortenerWithDomainParams struct {
	Code         string
	CustomDomain pgtype.Text
}

type GetShortenerWithDomainRow struct {
	ID          string
	Link        string
	Code        string
	CreatedAt   pgtype.Timestamp
	UserID      string
	ProjectID   pgtype.Text
	Active      bool
	Ios         bool
	IosLink     pgtype.Text
	Android     bool
	AndroidLink pgtype.Text
	Domain      pgtype.Text
}

func (q *Queries) GetShortenerWithDomain(ctx context.Context, arg GetShortenerWithDomainParams) (GetShortenerWithDomainRow, error) {
	row := q.db.QueryRow(ctx, getShortenerWithDomain, arg.Code, arg.CustomDomain)
	var i GetShortenerWithDomainRow
	err := row.Scan(
		&i.ID,
		&i.Link,
		&i.Code,
		&i.CreatedAt,
		&i.UserID,
		&i.ProjectID,
		&i.Active,
		&i.Ios,
		&i.IosLink,
		&i.Android,
		&i.AndroidLink,
		&i.Domain,
	)
	return i, err
}

const listShorteners = `-- name: ListShorteners :many
SELECT id, link, code, created_at, user_id, project_id, active, ios, ios_link, android, android_link
FROM shortener
`

func (q *Queries) ListShorteners(ctx context.Context) ([]Shortener, error) {
	rows, err := q.db.Query(ctx, listShorteners)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Shortener
	for rows.Next() {
		var i Shortener
		if err := rows.Scan(
			&i.ID,
			&i.Link,
			&i.Code,
			&i.CreatedAt,
			&i.UserID,
			&i.ProjectID,
			&i.Active,
			&i.Ios,
			&i.IosLink,
			&i.Android,
			&i.AndroidLink,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
