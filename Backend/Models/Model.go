package Models

type StrapiResponse[T any] struct {
	Data T
	Meta struct {
		Pagination struct {
			Page      int
			PageSize  int
			PageCount int
			Total     int
		}
	}
}

type NotFoundError struct{}

func (NotFoundError) Error() string {
	return "Not Found"
}
