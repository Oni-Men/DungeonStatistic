package reinc

type ReincResult struct {
	total int
}

func (r *ReincResult) GetTotal() int {
	return r.total
}
