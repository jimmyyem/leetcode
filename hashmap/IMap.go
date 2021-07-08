package hashmap

type IMap interface {
	Size() int
	IsEmpty() bool
	IsNotEmpty() bool

	Put(key, value interface{})
	Get(key interface{}) (bool, interface{})
	Has(key interface{}) bool
	Remove(key interface{}) (bool, interface{})
	Clear()

	Iterator() IMapIterator
	String() string
}
