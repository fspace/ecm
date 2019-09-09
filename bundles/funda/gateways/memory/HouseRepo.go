package memory

import "github.com/fspace/ecm/bundles/funda/entities"

// TODO 方法返回值需要定义 error 作为异常抛出机制
// TODO 可以找个线程安全的集合类库做测试 或者内嵌内存式db

type InMemoryHouseRepository struct {
	Store map[int64]entities.House
}

func NewInMemoryHouseRepository() *InMemoryHouseRepository {
	inst := &InMemoryHouseRepository{}

	inst.Store = make(map[int64]entities.House)

	return inst
}

func (repo InMemoryHouseRepository) Save(h entities.House) {
	//h := obj.(entities.House)
	if h.Id == 0 {
		// 零值
		h.Id = int64(len(repo.Store) + 1) // id 取数组长度 目前只添加 不删除  不然需要弄一个计数器 计算最大长度
	}
	repo.Store[h.Id] = h
	//	panic("implement me")
}

func (repo InMemoryHouseRepository) Get(id int64) entities.House {
	// panic("implement me")
	return repo.Store[id]
}

var _ entities.IHouseRepo = &InMemoryHouseRepository{}
