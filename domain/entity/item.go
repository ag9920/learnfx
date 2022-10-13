package entity

type Item struct {
	ID           int64
	Name         string
	Desc         string
	VisibleUsers []int64
}
