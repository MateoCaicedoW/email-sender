package models

type List struct {
	PerPage int
	Page    int
	Total   int

	Term string

	Items interface{}
}

func (l List) TotalPages() (pages int) {
	pages = l.Total / l.PerPage
	if (l.Total % l.PerPage) > 0.0 {
		pages++
	}

	if pages == 0 {
		pages = 1
	}

	return pages
}

func (l List) HasNext() bool {
	return l.Page < l.TotalPages()
}

func (l List) HasPrev() bool {
	return l.Page > 1
}

func (l List) NextPage() (next int) {
	next = l.Page + 1
	if l.Page == l.TotalPages() {
		next = l.Page
	}

	return next
}

func (l List) PrevPage() (prev int) {
	prev = l.Page - 1
	if prev < 0 {
		prev = 0
	}

	return prev
}
