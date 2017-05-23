package jsonld

import (
	"fmt"
	"net/http"

	"go.zenithar.org/common/dao/api"
	"go.zenithar.org/common/web/utils"
)

// Resource represents the JSON-LD header
type Resource struct {
	Context string `json:"@context,omitempty"`
	NodeID  string `json:"@id,omitempty"`
	Type    string `json:"@type,omitempty"`
}

// NewResource returns a JSONLD resource
func NewResource(context, id, _type string) *Resource {
	return &Resource{
		Context: context,
		NodeID:  id,
		Type:    _type,
	}
}

// CollectionResource represents the JSON-LD header
type CollectionResource struct {
	Resource

	TotalItems   uint   `json:"totalItems"`
	ItemPerPage  uint   `json:"itemPerPage,omitempty"`
	CurrentPage  uint   `json:"currentPage,omitempty"`
	FirstPage    string `json:"firstPage,omitempty"`
	NextPage     string `json:"nextPage,omitempty"`
	PreviousPage string `json:"previousPage,omitempty"`
	LastPage     string `json:"lastPage,omitempty"`
}

// NewCollection returns a JSONLD Collection resource
func NewCollection(context, id string) *CollectionResource {
	return &CollectionResource{
		Resource: Resource{
			Context: context,
			NodeID:  id,
			Type:    "Collection",
		},
	}
}

// SetPaginator defines values of the JSONLD Collection according to pagination request.
func (j *CollectionResource) SetPaginator(r *http.Request, paginator *api.Pagination) {

	j.TotalItems = paginator.Count()

	if paginator.HasOtherPages() {
		j.Type = "PagedCollection"
		j.ItemPerPage = paginator.PerPage
		j.CurrentPage = paginator.Page
	}

	q := r.URL.Query()

	if paginator.HasOtherPages() {
		q.Set("page", fmt.Sprintf("%d", 1))
		r.URL.RawQuery = q.Encode()
		j.FirstPage = r.URL.String()
	}
	if paginator.HasPrev() {
		q.Set("page", fmt.Sprintf("%d", paginator.Page-1))
		r.URL.RawQuery = q.Encode()
		j.PreviousPage = r.URL.String()
	}
	if paginator.HasNext() {
		q.Set("page", fmt.Sprintf("%d", paginator.Page+1))
		r.URL.RawQuery = q.Encode()
		j.NextPage = r.URL.String()
	}
	if paginator.HasOtherPages() {
		q.Set("page", fmt.Sprintf("%d", paginator.NumPages()))
		r.URL.RawQuery = q.Encode()
		j.LastPage = r.URL.String()
	}

}

// Error is the error holder
type Error struct {
	Resource
	StatusCode  int    `json:"statusCode,omitempty"`
	Code        string `json:"code,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

// Write error to socket
func (j *Error) Write(w http.ResponseWriter) {
	utils.JSONResponse(w, j.StatusCode, j)
}

// NewError returns a json-ld error holder
func NewError(context, id, code, title string) *Error {
	return &Error{
		Resource:   *NewResource(context, id, "Error"),
		Code:       code,
		Title:      title,
		StatusCode: 400,
	}
}

// Status is the response holder
type Status struct {
	Resource
	StatusCode  int    `json:"statusCode,omitempty"`
	Code        string `json:"code,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

// Write status to socket
func (j *Status) Write(w http.ResponseWriter) {
	utils.JSONResponse(w, j.StatusCode, j)
}

// NewStatus returns a status
func NewStatus(context, id, title string) *Status {
	return &Status{
		Resource:   *NewResource(context, id, "Status"),
		Title:      title,
		StatusCode: 400,
	}
}
