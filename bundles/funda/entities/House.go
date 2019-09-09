package entities

type House struct {
	Id      int64
	Address string

	// 关注者 Leads []Lead // 暂时不实现这个功能
	Leads []Interest
}

func NewHouse() *House {
	h := &House{}
	// FIXME go需要实例化空切片不？ Leads = []Interest
	h.Leads = []Interest{}
	return h
}
func (h *House) RegisterInterest(interest Interest) {
	h.Leads = append(h.Leads, interest)
}
