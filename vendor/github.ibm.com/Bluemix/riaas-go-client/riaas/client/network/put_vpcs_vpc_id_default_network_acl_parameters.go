// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// NewPutVpcsVpcIDDefaultNetworkACLParams creates a new PutVpcsVpcIDDefaultNetworkACLParams object
// with the default values initialized.
func NewPutVpcsVpcIDDefaultNetworkACLParams() *PutVpcsVpcIDDefaultNetworkACLParams {
	var ()
	return &PutVpcsVpcIDDefaultNetworkACLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutVpcsVpcIDDefaultNetworkACLParamsWithTimeout creates a new PutVpcsVpcIDDefaultNetworkACLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutVpcsVpcIDDefaultNetworkACLParamsWithTimeout(timeout time.Duration) *PutVpcsVpcIDDefaultNetworkACLParams {
	var ()
	return &PutVpcsVpcIDDefaultNetworkACLParams{

		timeout: timeout,
	}
}

// NewPutVpcsVpcIDDefaultNetworkACLParamsWithContext creates a new PutVpcsVpcIDDefaultNetworkACLParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutVpcsVpcIDDefaultNetworkACLParamsWithContext(ctx context.Context) *PutVpcsVpcIDDefaultNetworkACLParams {
	var ()
	return &PutVpcsVpcIDDefaultNetworkACLParams{

		Context: ctx,
	}
}

// NewPutVpcsVpcIDDefaultNetworkACLParamsWithHTTPClient creates a new PutVpcsVpcIDDefaultNetworkACLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutVpcsVpcIDDefaultNetworkACLParamsWithHTTPClient(client *http.Client) *PutVpcsVpcIDDefaultNetworkACLParams {
	var ()
	return &PutVpcsVpcIDDefaultNetworkACLParams{
		HTTPClient: client,
	}
}

/*PutVpcsVpcIDDefaultNetworkACLParams contains all the parameters to send to the API endpoint
for the put vpcs vpc ID default network ACL operation typically these are written to a http.Request
*/
type PutVpcsVpcIDDefaultNetworkACLParams struct {

	/*Body*/
	Body *models.PutVpcsVpcIDDefaultNetworkACLParamsBody
	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string
	/*VpcID
	  The VPC identifier

	*/
	VpcID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put vpcs vpc ID default network ACL params
func (o *PutVpcsVpcIDDefaultNetworkACLParams) WithTimeout(timeout time.Duration) *PutVpcsVpcIDDefaultNetworkACLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put vpcs vpc ID default network ACL params
func (o *PutVpcsVpcIDDefaultNetworkACLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put vpcs vpc ID default network ACL params
func (o *PutVpcsVpcIDDefaultNetworkACLParams) WithContext(ctx context.Context) *PutVpcsVpcIDDefaultNetworkACLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put vpcs vpc ID default network ACL params
func (o *PutVpcsVpcIDDefaultNetworkACLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put vpcs vpc ID default network ACL params
func (o *PutVpcsVpcIDDefaultNetworkACLParams) WithHTTPClient(client *http.Client) *PutVpcsVpcIDDefaultNetworkACLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put vpcs vpc ID default network ACL params
func (o *PutVpcsVpcIDDefaultNetworkACLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put vpcs vpc ID default network ACL params
func (o *PutVpcsVpcIDDefaultNetworkACLParams) WithBody(body *models.PutVpcsVpcIDDefaultNetworkACLParamsBody) *PutVpcsVpcIDDefaultNetworkACLParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put vpcs vpc ID default network ACL params
func (o *PutVpcsVpcIDDefaultNetworkACLParams) SetBody(body *models.PutVpcsVpcIDDefaultNetworkACLParamsBody) {
	o.Body = body
}

// WithGeneration adds the generation to the put vpcs vpc ID default network ACL params
func (o *PutVpcsVpcIDDefaultNetworkACLParams) WithGeneration(generation int64) *PutVpcsVpcIDDefaultNetworkACLParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the put vpcs vpc ID default network ACL params
func (o *PutVpcsVpcIDDefaultNetworkACLParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithVersion adds the version to the put vpcs vpc ID default network ACL params
func (o *PutVpcsVpcIDDefaultNetworkACLParams) WithVersion(version string) *PutVpcsVpcIDDefaultNetworkACLParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the put vpcs vpc ID default network ACL params
func (o *PutVpcsVpcIDDefaultNetworkACLParams) SetVersion(version string) {
	o.Version = version
}

// WithVpcID adds the vpcID to the put vpcs vpc ID default network ACL params
func (o *PutVpcsVpcIDDefaultNetworkACLParams) WithVpcID(vpcID string) *PutVpcsVpcIDDefaultNetworkACLParams {
	o.SetVpcID(vpcID)
	return o
}

// SetVpcID adds the vpcId to the put vpcs vpc ID default network ACL params
func (o *PutVpcsVpcIDDefaultNetworkACLParams) SetVpcID(vpcID string) {
	o.VpcID = vpcID
}

// WriteToRequest writes these params to a swagger request
func (o *PutVpcsVpcIDDefaultNetworkACLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	// path param vpc_id
	if err := r.SetPathParam("vpc_id", o.VpcID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}