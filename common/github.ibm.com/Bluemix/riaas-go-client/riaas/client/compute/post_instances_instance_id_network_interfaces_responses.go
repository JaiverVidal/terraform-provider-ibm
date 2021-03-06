// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// PostInstancesInstanceIDNetworkInterfacesReader is a Reader for the PostInstancesInstanceIDNetworkInterfaces structure.
type PostInstancesInstanceIDNetworkInterfacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostInstancesInstanceIDNetworkInterfacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostInstancesInstanceIDNetworkInterfacesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostInstancesInstanceIDNetworkInterfacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostInstancesInstanceIDNetworkInterfacesCreated creates a PostInstancesInstanceIDNetworkInterfacesCreated with default headers values
func NewPostInstancesInstanceIDNetworkInterfacesCreated() *PostInstancesInstanceIDNetworkInterfacesCreated {
	return &PostInstancesInstanceIDNetworkInterfacesCreated{}
}

/*PostInstancesInstanceIDNetworkInterfacesCreated handles this case with default header values.

dummy
*/
type PostInstancesInstanceIDNetworkInterfacesCreated struct {
	Payload *models.InstanceNetworkInterface
}

func (o *PostInstancesInstanceIDNetworkInterfacesCreated) Error() string {
	return fmt.Sprintf("[POST /instances/{instance_id}/network_interfaces][%d] postInstancesInstanceIdNetworkInterfacesCreated  %+v", 201, o.Payload)
}

func (o *PostInstancesInstanceIDNetworkInterfacesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InstanceNetworkInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInstancesInstanceIDNetworkInterfacesBadRequest creates a PostInstancesInstanceIDNetworkInterfacesBadRequest with default headers values
func NewPostInstancesInstanceIDNetworkInterfacesBadRequest() *PostInstancesInstanceIDNetworkInterfacesBadRequest {
	return &PostInstancesInstanceIDNetworkInterfacesBadRequest{}
}

/*PostInstancesInstanceIDNetworkInterfacesBadRequest handles this case with default header values.

error
*/
type PostInstancesInstanceIDNetworkInterfacesBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PostInstancesInstanceIDNetworkInterfacesBadRequest) Error() string {
	return fmt.Sprintf("[POST /instances/{instance_id}/network_interfaces][%d] postInstancesInstanceIdNetworkInterfacesBadRequest  %+v", 400, o.Payload)
}

func (o *PostInstancesInstanceIDNetworkInterfacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostInstancesInstanceIDNetworkInterfacesBody NetworkInterfaceTemplate
swagger:model PostInstancesInstanceIDNetworkInterfacesBody
*/
type PostInstancesInstanceIDNetworkInterfacesBody struct {

	// The user-defined name for this interface
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// The network interface port speed in Mbps
	PortSpeed int64 `json:"port_speed,omitempty"`

	// The primary IPv4 address
	PrimaryIPV4Address string `json:"primary_ipv4_address,omitempty"`

	// The primary IPv6 address in any valid notation as specified by RFC 4291
	PrimaryIPV6Address string `json:"primary_ipv6_address,omitempty"`

	// resource group
	ResourceGroup *PostInstancesInstanceIDNetworkInterfacesParamsBodyResourceGroup `json:"resource_group,omitempty"`

	// Collection seconary IP addresses
	SecondaryAddresses []string `json:"secondary_addresses,omitempty"`

	// security groups
	SecurityGroups []*SecurityGroupsItems0 `json:"security_groups,omitempty"`

	// subnet
	Subnet *PostInstancesInstanceIDNetworkInterfacesParamsBodySubnet `json:"subnet,omitempty"`
}

// Validate validates this post instances instance ID network interfaces body
func (o *PostInstancesInstanceIDNetworkInterfacesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResourceGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSecurityGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSubnet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostInstancesInstanceIDNetworkInterfacesBody) validateName(formats strfmt.Registry) error {

	if swag.IsZero(o.Name) { // not required
		return nil
	}

	if err := validate.Pattern("body"+"."+"name", "body", string(o.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

func (o *PostInstancesInstanceIDNetworkInterfacesBody) validateResourceGroup(formats strfmt.Registry) error {

	if swag.IsZero(o.ResourceGroup) { // not required
		return nil
	}

	if o.ResourceGroup != nil {
		if err := o.ResourceGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "resource_group")
			}
			return err
		}
	}

	return nil
}

func (o *PostInstancesInstanceIDNetworkInterfacesBody) validateSecurityGroups(formats strfmt.Registry) error {

	if swag.IsZero(o.SecurityGroups) { // not required
		return nil
	}

	for i := 0; i < len(o.SecurityGroups); i++ {
		if swag.IsZero(o.SecurityGroups[i]) { // not required
			continue
		}

		if o.SecurityGroups[i] != nil {
			if err := o.SecurityGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "security_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *PostInstancesInstanceIDNetworkInterfacesBody) validateSubnet(formats strfmt.Registry) error {

	if swag.IsZero(o.Subnet) { // not required
		return nil
	}

	if o.Subnet != nil {
		if err := o.Subnet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "subnet")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostInstancesInstanceIDNetworkInterfacesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostInstancesInstanceIDNetworkInterfacesBody) UnmarshalBinary(b []byte) error {
	var res PostInstancesInstanceIDNetworkInterfacesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostInstancesInstanceIDNetworkInterfacesParamsBodyResourceGroup idreference
swagger:model PostInstancesInstanceIDNetworkInterfacesParamsBodyResourceGroup
*/
type PostInstancesInstanceIDNetworkInterfacesParamsBodyResourceGroup struct {

	// The unique identifier for this resource
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`
}

// Validate validates this post instances instance ID network interfaces params body resource group
func (o *PostInstancesInstanceIDNetworkInterfacesParamsBodyResourceGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostInstancesInstanceIDNetworkInterfacesParamsBodyResourceGroup) validateID(formats strfmt.Registry) error {

	if swag.IsZero(o.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"resource_group"+"."+"id", "body", "uuid", o.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostInstancesInstanceIDNetworkInterfacesParamsBodyResourceGroup) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostInstancesInstanceIDNetworkInterfacesParamsBodyResourceGroup) UnmarshalBinary(b []byte) error {
	var res PostInstancesInstanceIDNetworkInterfacesParamsBodyResourceGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostInstancesInstanceIDNetworkInterfacesParamsBodySubnet SubnetReference
//
// The associated subnet
swagger:model PostInstancesInstanceIDNetworkInterfacesParamsBodySubnet
*/
type PostInstancesInstanceIDNetworkInterfacesParamsBodySubnet struct {

	// The CRN for this subnet
	Crn string `json:"crn,omitempty"`

	// The unique identifier for this subnet
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The user-defined name for this subnet
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`
}

// Validate validates this post instances instance ID network interfaces params body subnet
func (o *PostInstancesInstanceIDNetworkInterfacesParamsBodySubnet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostInstancesInstanceIDNetworkInterfacesParamsBodySubnet) validateID(formats strfmt.Registry) error {

	if swag.IsZero(o.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"subnet"+"."+"id", "body", "uuid", o.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PostInstancesInstanceIDNetworkInterfacesParamsBodySubnet) validateName(formats strfmt.Registry) error {

	if swag.IsZero(o.Name) { // not required
		return nil
	}

	if err := validate.Pattern("body"+"."+"subnet"+"."+"name", "body", string(o.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostInstancesInstanceIDNetworkInterfacesParamsBodySubnet) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostInstancesInstanceIDNetworkInterfacesParamsBodySubnet) UnmarshalBinary(b []byte) error {
	var res PostInstancesInstanceIDNetworkInterfacesParamsBodySubnet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SecurityGroupsItems0 reference
swagger:model SecurityGroupsItems0
*/
type SecurityGroupsItems0 struct {

	// The CRN for this snapshot
	Crn string `json:"crn,omitempty"`

	// The unique identifier for this resource
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The user-defined name for this resource
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`
}

// Validate validates this security groups items0
func (o *SecurityGroupsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityGroupsItems0) validateID(formats strfmt.Registry) error {

	if swag.IsZero(o.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", o.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *SecurityGroupsItems0) validateName(formats strfmt.Registry) error {

	if swag.IsZero(o.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(o.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SecurityGroupsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SecurityGroupsItems0) UnmarshalBinary(b []byte) error {
	var res SecurityGroupsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
