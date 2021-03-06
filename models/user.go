package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*User user

swagger:model User
*/
type User struct {

	/* e.g. John Smith
	 */
	DisplayName string `json:"displayName,omitempty"`

	/* EmailAddress email address
	 */
	EmailAddress string `json:"emailAddress,omitempty"`

	/* ID id
	 */
	ID int64 `json:"id,omitempty"`

	/* Link link
	 */
	Link *Link `json:"link,omitempty"`

	/* e.g. short ldap name
	 */
	Name string `json:"name,omitempty"`

	/* e.g. ldap name
	 */
	Slug string `json:"slug,omitempty"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	return nil
}
