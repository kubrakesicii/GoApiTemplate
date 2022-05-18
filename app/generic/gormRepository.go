package generic

import "gorm.io/gorm"

type GormRepository struct {
	BaseRepository
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) GormRepository {
	return GormRepository{db: db}
}

func (gr GormRepository) Insert(entity IEntity) (uint, error) {
	if err := entity.Validate(); err != nil {
		return 0, err
	}
	if err := gr.db.Create(entity).Error; err != nil {
		return 0, err
	}
	return entity.GetId(), nil
}

func (gr GormRepository) Update(entity IEntity) error {
	if err := entity.Validate(); err != nil {
		return err
	}
	return gr.db.Save(entity).Error
}

func (gr GormRepository) Save(entity IEntity) (uint, error) {
	if err := entity.Validate(); err != nil {
		return 0, err
	}
	if err := gr.db.Save(entity).Error; err != nil {
		return 0, err
	}
	return entity.GetId(), nil
}

func (gr GormRepository) FindById(receiver IEntity, id uint) error {
	return gr.db.First(receiver, id).Error
}

func (gr GormRepository) FindFirst(receiver IEntity, where string, args ...interface{}) error {
	return gr.db.Where(where, args...).Limit(1).Find(receiver).Error
}

func (gr GormRepository) FindAll(models interface{}, where string, args ...interface{}) (err error) {
	err = gr.db.Where(where, args...).Find(models).Error
	return
}

func (gr GormRepository) Delete(entity IEntity, where string, args ...interface{}) error {
	return gr.db.Where(where, args...).Delete(&entity).Error
}

// func (gr GormRepository) IfExists(entity IEntity) bool {
// 	return gr.db.NewRecord(entity)
// }
