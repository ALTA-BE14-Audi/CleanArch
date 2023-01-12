package book

type Core struct {
	ID          uint
	Judul       string
	TahunTerbit int
	Penulis     string
	Pemilik     string
}

type BookService interface {
	Add(token interface{}, newBook Core) (Core, error)
	Update(token interface{}, bookID int, updatedData Core) (Core, error)
	GetAll() ([]Core, error)
	Delete(token interface{}, bookID int) error
	MyBook(token interface{}) ([]Core, error)
}

type BookData interface {
	Add(userID int, newBook Core) (Core, error)
	Update(token int, bookID int, updatedData Core) (Core, error)
	GetAll() ([]Core, error)
	Delete(bookID int, userID int) error
	MyBook(userID int) ([]Core, error)
}
