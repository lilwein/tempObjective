package page

import (
	"errors"
	"math"
	"strconv"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/go-core-app"
)

const (
	QP_PageNumber = "pageNumber"
	QP_PageSize   = "pageSize"

	PageSize = 10
)

// Paging metadata
type Paging struct {
	PageSize    int  `json:"pageSize"`
	TotalItems  int  `json:"totalItems"`
	TotalPages  int  `json:"totalPages"`
	CurrentPage int  `json:"curentPage"`
	HasNext     bool `json:"hasNext"`
	HasPrev     bool `json:"hasPrev"`
}

// Initialize a Paging Metadata
func InitPaging(pagesize, totalItems int) Paging {
	var p = Paging{
		PageSize:    pagesize,
		TotalItems:  totalItems,
		TotalPages:  0,
		CurrentPage: 0,
		HasNext:     false,
		HasPrev:     false,
	}

	p.updatePagingMetaData()

	return p
}

// Updating Paging
func (p *Paging) updatePagingMetaData() {

	// If there are no items, TotalPages = 0. Otherwise, calculate TotalPages
	if p.TotalItems == 0 {
		p.setTotalPages(0)
	} else {
		p.setTotalPages(int(math.Ceil(float64(p.TotalItems) / float64(p.PageSize))))
	}

	// If CurrentPage > 1, the current page has a previous page. Orherwise, it hasn't
	if p.CurrentPage > 1 && p.CurrentPage <= p.TotalPages+1 {
		p.setHasPrev(true)
	} else {
		p.setHasPrev(false)
	}

	// If CurrentPage = TotalPages, the current page has not a next page. Orherwise, it has
	if p.CurrentPage >= p.TotalPages || p.CurrentPage == 0 {
		p.setHasNext(false)
	} else {
		p.setHasNext(true)
	}
}

// Set TotalItems
func (p *Paging) SetTotalItems(totalItems int) {
	p.TotalItems = totalItems
	p.updatePagingMetaData()
}

// Increment TotalItems
func (p *Paging) IncTotalItems() {
	p.TotalItems++
	p.updatePagingMetaData()
}

// Decrement TotalItems
func (p *Paging) DecTotalItems() {
	p.TotalItems--
	p.updatePagingMetaData()
}

// Set TotalPages
func (p *Paging) setTotalPages(totalPages int) {
	p.TotalPages = totalPages
}

// Set CurrentPage
func (p *Paging) SetCurrentPage(currentPage int) {
	if currentPage < 0 {
		panic(errors.New("invalid current page"))
	}
	p.CurrentPage = currentPage

	p.updatePagingMetaData()
}

// Increment CurrentPage
func (p *Paging) IncCurrentPage() {
	if p.CurrentPage < 0 {
		panic(errors.New("invalid current page"))
	}
	p.CurrentPage++

	p.updatePagingMetaData()
}

// Decrement CurrentPage
func (p *Paging) DecCurrentPage() {
	if p.CurrentPage < 0 {
		panic(errors.New("invalid current page"))
	}
	p.CurrentPage--

	p.updatePagingMetaData()
}

// Set HasNext
func (p *Paging) setHasNext(hasNext bool) {
	p.HasNext = hasNext
}

// Set HasPrev
func (p *Paging) setHasPrev(hasPrev bool) {
	p.HasPrev = hasPrev
}

// Get page size from query param.
// If errors or size < 0, return error.
// If empty, return default value.
func GetPageSize(param string) (int, *core.ApplicationError) {

	if param != "" {
		pageSize, err := strconv.Atoi(param)
		if err != nil || pageSize < 0 {
			return 0, core.BusinessErrorWithCodeAndMessage("ERR-PAGESIZE", "invalid page size")
		}
		return pageSize, nil
	}
	return PageSize, nil
}

// Get page number from query param.
// If errors or number < 1, return error.
// If empty, return default value.
func GetPageNumber(param string) (int, error) {

	if param != "" {
		pageNumber, err := strconv.Atoi(param)
		if err != nil || pageNumber < 1 {
			return 0, core.BusinessErrorWithCodeAndMessage("ERR-PAGENUMBER", "invalid page number")
		}

		return pageNumber, nil
	}

	return 1, nil
}
