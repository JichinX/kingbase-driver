//go:build go1.15
// +build go1.15

package kb

import "database/sql/driver"

var _ driver.Validator = &conn{}
