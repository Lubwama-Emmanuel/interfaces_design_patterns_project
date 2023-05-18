package inmemory

type IDatabase interface {
	Create(data string) error
	Read() (string, error)
	Update(data string) error
	Delete() error
}

type InMemoryDatabase struct {
	data string
}

func (db *InMemoryDatabase) Create(data string) error {
	db.data = data
	return nil
}

func (db InMemoryDatabase) Read() (string, error) {
	return db.data, nil
}

func (db *InMemoryDatabase) Update(data string) error {
	db.data = data
	return nil
}

func (db *InMemoryDatabase) Delete() error {
	db.data = ""
	return nil
}
