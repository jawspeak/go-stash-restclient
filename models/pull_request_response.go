package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*
PullRequestResponse pull request response

swagger:model PullRequestResponse
*/
type PullRequestResponse struct {

	/* Author author
	 */
	Author *Author `json:"author,omitempty"`

	/* CreatedDate created date
	 */
	CreatedDate int64 `json:"createdDate,omitempty"`

	/* Description description
	 */
	Description string `json:"description,omitempty"`

	/* ID id
	 */
	ID int64 `json:"id,omitempty"`

	/* State state
	 */
	State string `json:"state,omitempty"`

	/* Title title
	 */
	Title string `json:"title,omitempty"`

	/* UpdatedDate updated date
	 */
	UpdatedDate int64 `json:"updatedDate,omitempty"`

	/* Version version
	 */
	Version int64 `json:"version,omitempty"`
}

// Validate validates this pull request response
func (m *PullRequestResponse) Validate(formats strfmt.Registry) error {
	return nil
}