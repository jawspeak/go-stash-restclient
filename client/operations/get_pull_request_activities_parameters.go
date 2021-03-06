package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*GetPullRequestActivitiesParams contains all the parameters to send to the API endpoint
for the get pull request activities operation typically these are written to a http.Request
*/
type GetPullRequestActivitiesParams struct {

	/*Limit
	  Probably defaults to 25. It is a best practice to check the limit attribute on the response to see what limit has been applied.

	*/
	Limit *int64
	/*Project*/
	Project string
	/*PullRequestID*/
	PullRequestID int64
	/*Repo*/
	Repo string
}

// WriteToRequest writes these params to a swagger request
func (o *GetPullRequestActivitiesParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// query param limit
	var qrLimit int64
	if o.Limit != nil {
		qrLimit = *o.Limit
	}
	qLimit := swag.FormatInt64(qrLimit)
	if err := r.SetQueryParam("limit", qLimit); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	// path param pullRequestId
	if err := r.SetPathParam("pullRequestId", swag.FormatInt64(o.PullRequestID)); err != nil {
		return err
	}

	// path param repo
	if err := r.SetPathParam("repo", o.Repo); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
