package dbSetup

import (
	"gorm.io/gorm/schema"
	"regexp"
	"strings"
)

type CustomNamingStrategy struct {
	schema.NamingStrategy
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func (c CustomNamingStrategy) ColumnName(table, column string) string {
	if strings.ToLower(column) == "id" {
		column = column + "_" + table
	}

	snake := matchFirstCap.ReplaceAllString(column, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
