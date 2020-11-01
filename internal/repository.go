package internal

type Todoable interface {
	Get() ([]*Todo, error)
	Store(text string) (int64, error)
	Find(id int64) (Todo, error)
	Complete(id int64) error
	Uncomplete(id int64) error
	Destroy(id int64) error
}
