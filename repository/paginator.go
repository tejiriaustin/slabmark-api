package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"math"
)

const (
	defaultPageNumber  int64 = 1
	defaultPerPageRows int64 = 40
)

type (
	Paginator struct {
		CurrentPage int64 `json:"currentPage"`
		NextPage    int64 `json:"nextPage"`
		PrevPage    int64 `json:"prevPage"`
		TotalPages  int64 `json:"totalPages"`
		TotalRows   int64 `json:"totalRows"`
		PerPage     int64 `json:"perPage"`
		Offset      int64 `json:"-"`
	}

	paginatorOptions struct {
		Page    int64 `json:"page"`
		PerPage int64 `json:"per_page"`
	}

	paginatedResult struct {
		*Paginator
		Cursor *mongo.Cursor
	}
)

func newPaginator(opts paginatorOptions) *Paginator {
	p := &Paginator{}
	p.Offset = 0

	p.CurrentPage = opts.Page
	if opts.Page <= 1 {
		p.CurrentPage = defaultPageNumber
	}

	p.PerPage = opts.PerPage
	if opts.PerPage < 1 {
		p.PerPage = defaultPerPageRows
	}

	return p
}

func (p *Paginator) setOffset() {
	if p.CurrentPage == 1 {
		p.Offset = 0
		return
	}
	p.Offset = (p.CurrentPage - 1) * p.PerPage
}

func (p *Paginator) setTotalPages() {
	if p.TotalRows == 0 {
		p.TotalPages = 0
		return
	}
	p.TotalPages = int64(math.Ceil(float64(p.TotalRows) / float64(p.PerPage)))
}

func (p *Paginator) setPrevPage() {
	//Call SetTotalPages just to be safe
	p.setTotalPages()

	if p.CurrentPage == 1 {
		p.PrevPage = 0
		return
	}
	p.PrevPage = p.CurrentPage - 1
}

func (p *Paginator) setNextPage() {
	//Call SetTotalPages just to be safe
	p.setTotalPages()

	if p.CurrentPage == p.TotalPages {
		p.NextPage = p.CurrentPage
	} else {
		p.NextPage = p.CurrentPage + 1
	}
}
