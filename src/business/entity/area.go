package entity

import (
	"strconv"
	"strings"
)

type Area struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	// Address   string  `json:"address" db:"address"`
	Latitude  float64 `json:"latitude" db:"latitude"`
	Longitude float64 `json:"longitude" db:"longitude"`
	CodeArea  string  `json:"code_area" db:"code_area"`
	Notes     string  `json:"notes" db:"notes"`
	CreatedAndUpdatedTime
}

type AreaInput struct {
	Name string `json:"name" db:"name"`
	// Address   string  `json:"address" db:"address"`
	Latitude  float64 `json:"latitude" db:"latitude"`
	Longitude float64 `json:"longitude" db:"longitude"`
	CodeArea  string  `json:"code_area" db:"code_area"`
	Notes     string  `json:"notes" db:"notes"`
}

type UpdateAreaInput struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	// Address   string  `json:"address" db:"address"`
	CodeArea  string  `json:"code_area" db:"code_area"`
	Notes     string  `json:"notes" db:"notes"`
	Latitude  float64 `json:"latitude" db:"latitude"`
	Longitude float64 `json:"longitude" db:"longitude"`
}

func (a *UpdateAreaInput) SetDefaultValue(currentval Area) {
	if strings.TrimSpace(a.Name) == "" {
		a.Name = currentval.Name
	}

	// if strings.TrimSpace(a.Address) == "" {
	// 	a.Address = currentval.Address
	// }

	if strings.TrimSpace(a.CodeArea) == "" {
		a.CodeArea = currentval.CodeArea
	}

	if strings.TrimSpace(a.Notes) == "" {
		a.Notes = currentval.Notes
	}

	if a.Latitude == 0 {
		a.Latitude = currentval.Latitude
	}

	if a.Longitude == 0 {
		a.Longitude = currentval.Longitude
	}
}

type AreaParams struct {
	ID   int64  `form:"id" db:"id"`
	Name string `form:"name" db:"name"`
	// Address  string `form:"address" db:"address"`
	CodeArea string `form:"code_area" db:"code"`
	Notes    string `form:"notes" db:"notes"`
	Limit    int64  `form:"limit" db:"-"`
	Page     int64  `form:"page" db:"page"`
}

func (ap *AreaParams) CreateMySQLQuery(query string, paginate bool) string {
	if ap.ID > 0 {
		query = query + " AND id=" + strconv.Itoa(int(ap.ID))
	}
	if ap.CodeArea != "" {
		query = query + " AND code=" + ap.CodeArea
	}
	if ap.Name != "" {
		query = query + " AND name=" + ap.Name
	}
	if ap.Notes != "" {
		query = query + " AND notes=" + ap.Notes
	}
	if ap.Limit < 1 {
		ap.Limit = 10
	} else if ap.Limit >= 1000 {
		ap.Limit = 1000
	}
	if ap.Page < 1 {
		ap.Page = 1
	}
	if paginate {
		offset := (ap.Page - 1) * ap.Limit
		query = query + " LIMIT " + strconv.Itoa(int(ap.Limit))
		query = query + " OFFSET " + strconv.Itoa(int(offset))
	}
	return query
}
