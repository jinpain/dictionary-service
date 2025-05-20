package insurance

type Filter struct {
	Limit  int `from:"limit"`
	Offset int `from:"offset"`
}
