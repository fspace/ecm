package entities

type IHouseRepo interface {
	Save(obj House)
	Get(id int64) House
}
