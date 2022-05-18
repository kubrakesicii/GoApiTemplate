package generic

// type GenericRepository[T any] struct {
//     db *gorm.DB
// }

// func NewGenericRepository(db *gorm.DB) AuthRepository {
// 	return &genericRepository{
// 		db: db,
// 	} //You should implement all methods of Repo int
// }

// func (g GenericRepository[T]) GetAll() []T {

// }

type Repository interface {

	// Insert puts a new instance of the give IModel in the database
	Deneme(entity IEntity) (uint, error)

	Update(entity IEntity) error

	Save(entity IEntity) (uint, error)

	FindById(receiver IEntity, uint uint) error

	FindFirst(receiver IEntity, where string, args ...interface{}) error

	FindAll(entitys interface{}, where string, args ...interface{}) (err error)

	Delete(entity IEntity, where string, args ...interface{}) error

	// NewRecord check if the entity exist in the store
	IfExists(entity IEntity) bool
}
