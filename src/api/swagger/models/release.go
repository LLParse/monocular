package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*Release release

swagger:model release
*/
type Release struct {

	/* chart name

	Required: true
	*/
	ChartName *string `json:"chartName"`

	/* chart version

	Required: true
	*/
	ChartVersion *string `json:"chartVersion"`

	/* name

	Required: true
	*/
	Name *string `json:"name"`

	/* namespace

	Required: true
	*/
	Namespace *string `json:"namespace"`

	/* status

	Required: true
	*/
	Status *string `json:"status"`

	/* updated

	Required: true
	*/
	Updated *string `json:"updated"`
}

// Validate validates this release
func (m *Release) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChartName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateChartVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Release) validateChartName(formats strfmt.Registry) error {

	if err := validate.Required("chartName", "body", m.ChartName); err != nil {
		return err
	}

	return nil
}

func (m *Release) validateChartVersion(formats strfmt.Registry) error {

	if err := validate.Required("chartVersion", "body", m.ChartVersion); err != nil {
		return err
	}

	return nil
}

func (m *Release) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Release) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *Release) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Release) validateUpdated(formats strfmt.Registry) error {

	if err := validate.Required("updated", "body", m.Updated); err != nil {
		return err
	}

	return nil
}
