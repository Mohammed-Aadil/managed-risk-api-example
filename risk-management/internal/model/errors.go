package model

import "errors"

var ErrRiskNotFound = errors.New("risk details not found")
var ErrRiskAlreadyPresent = errors.New("risk is already present")
var ErrRiskSortFieldNotAllowed = errors.New("risk sort field is not allowed")
