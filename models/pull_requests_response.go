package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*PullRequestsResponse pull requests response

swagger:model PullRequestsResponse
*/
type PullRequestsResponse struct {

	/* IsLastPage is last page
	 */
	IsLastPage bool `json:"isLastPage,omitempty"`

	/* Limit limit
	 */
	Limit int64 `json:"limit,omitempty"`

	/* NextPageStart next page start
	 */
	NextPageStart int64 `json:"nextPageStart,omitempty"`

	/* Size size
	 */
	Size int64 `json:"size,omitempty"`

	/* Start start
	 */
	Start int64 `json:"start,omitempty"`

	/* Values values
	 */
	Values []*PullRequest `json:"values,omitempty"`
}

// Validate validates this pull requests response
func (m *PullRequestsResponse) Validate(formats strfmt.Registry) error {
	return nil
}
