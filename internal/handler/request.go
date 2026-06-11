package handler

import "fmt"

func errParamsIsRequired(paramName string, paramType string) error {
	return fmt.Errorf("param: %s (type: %s) is required", paramName, paramType)
}

// CreateOpening
type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" {
		return errParamsIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamsIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamsIsRequired("location", "string")
	}
	if r.Link == "" {
		return errParamsIsRequired("link", "string")
	}
	if r.Remote == nil {
		return errParamsIsRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return errParamsIsRequired("salary", "int64")
	}

	return nil
}

// UpdateOpening

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	// if any filed is provided, validate is truthy
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salary > 0 {
		return nil
	}

	// if none of the fields where provided, return falsy
	return fmt.Errorf("at least one valid field must be provided")
}
