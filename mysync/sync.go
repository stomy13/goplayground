package mysync

import "sync"

type BigStruct struct {
	UUID string
	ID   int
	Name string
}

func NewBigStruct(uuid string, id int, name string) *BigStruct {
	return &BigStruct{
		UUID: uuid,
		ID:   id,
		Name: name,
	}
}

var bigStructPool = sync.Pool{
	New: func() any {
		return &BigStruct{}
	},
}

func NewBigStructFromPool(uuid string, id int, name string) *BigStruct {
	bs := bigStructPool.Get().(*BigStruct)
	bs.UUID = uuid
	bs.ID = id
	bs.Name = name
	return bs
}

func (bs *BigStruct) Release() {
	bigStructPool.Put(bs)
}
