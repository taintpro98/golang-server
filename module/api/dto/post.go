package dto

type CommonFilter struct {
	Limit  int
	Offset *int
	Select []string
	Sort   string
}

type FilterPost struct {
	CommonFilter CommonFilter
	ID           string
	UserID       string
}
