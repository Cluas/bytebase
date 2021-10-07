package api

import (
	"context"
	"encoding/json"
)

type View struct {
	ID int `jsonapi:"primary,view"`

	// Standard fields
	CreatorId int
	Creator   *Principal `jsonapi:"attr,creator"`
	CreatedTs int64      `jsonapi:"attr,createdTs"`
	UpdaterId int
	Updater   *Principal `jsonapi:"attr,updater"`
	UpdatedTs int64      `jsonapi:"attr,updatedTs"`

	// Related fields
	DatabaseId int
	Database   *Database `jsonapi:"relation,database"`

	// Domain specific fields
	Name       string `jsonapi:"attr,name"`
	Definition string `jsonapi:"attr,definition"`
	Comment    string `jsonapi:"attr,comment"`
}

type ViewCreate struct {
	// Standard fields
	// Value is assigned from the jwt subject field passed by the client.
	CreatorId int
	CreatedTs int64
	UpdatedTs int64

	// Related fields
	DatabaseId int

	// Domain specific fields
	Name       string
	Definition string
	Comment    string
}

type ViewFind struct {
	ID *int

	// Related fields
	DatabaseId *int

	// Domain specific fields
	Name *string
}

func (find *ViewFind) String() string {
	str, err := json.Marshal(*find)
	if err != nil {
		return err.Error()
	}
	return string(str)
}

type ViewDelete struct {
	// Related fields
	DatabaseId int
}

type ViewService interface {
	CreateView(ctx context.Context, create *ViewCreate) (*View, error)
	FindViewList(ctx context.Context, find *ViewFind) ([]*View, error)
	FindView(ctx context.Context, find *ViewFind) (*View, error)
	DeleteView(ctx context.Context, delete *ViewDelete) error
}
