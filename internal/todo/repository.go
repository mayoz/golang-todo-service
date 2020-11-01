package todo

type Todoable interface {
	Get() ([]*Todo, error)
	Store(text string) (int64, error)
	Find(id int64) (Todo, error)
	Toggle(id int64) error
	Destroy(id int64) error
}
