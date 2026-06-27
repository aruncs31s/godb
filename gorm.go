package godb

import "gorm.io/gorm"

type Supported interface {
	gorm.Migrator
}

func Get(
	db Supported,
	table string,
	columns ...string,
) string {
	if len(columns) == 0 {
		return ""
	}

	try := func(field string) bool {
		return db.HasColumn(table, field)
	}

	for _, c := range columns {
		if try(c) {
			return c
		}
	}
	return ""
}
