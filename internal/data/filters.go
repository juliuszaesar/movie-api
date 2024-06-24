package data

import "movie-api/internal/validator"

type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafeList []string
}

func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Page > 0, "page", "must be larger than 0")
	v.Check(f.Page <= 10_000_000, "page", "maximum is 10 million")
	v.Check(f.PageSize > 0, "page_size", "cannot exceed 100")

	v.Check(validator.PermittedValue(f.Sort, f.SortSafeList...), "sort", "invalid soft value")
}
