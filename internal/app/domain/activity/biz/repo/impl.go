package repo

type impl struct {
}

func NewImpl() IRepo {
	return &impl{}
}
