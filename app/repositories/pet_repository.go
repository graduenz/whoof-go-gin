package repositories

import (
	"gorm.io/gorm"
	"whoof-go-gin/app/models"
)

type PetRepository interface {
	FindPetsPaginated(pageIndex int, pageSize int) ([]models.Pet, error)
	FindPetById(id int) (models.Pet, error)
	SavePet(pet *models.Pet) (models.Pet, error)
	DeletePetById(id int) error
}

type PetRepositoryImpl struct {
	db *gorm.DB
}

func NewPetRepositoryImpl(db *gorm.DB) PetRepository {
	return &PetRepositoryImpl{
		db: db,
	}
}

func (r *PetRepositoryImpl) FindPetsPaginated(pageIndex int, pageSize int) ([]models.Pet, error) {
	var pets []models.Pet
	offset := (pageIndex - 1) * pageSize
	if err := r.db.Offset(offset).Limit(pageSize).Find(&pets).Error; err != nil {
		return nil, err
	}
	return pets, nil
}

func (r *PetRepositoryImpl) FindPetById(id int) (models.Pet, error) {
	var pet models.Pet
	if err := r.db.First(&pet, id).Error; err != nil {
		return pet, err
	}
	return pet, nil
}

func (r *PetRepositoryImpl) SavePet(pet *models.Pet) (models.Pet, error) {
	if err := r.db.Save(pet).Error; err != nil {
		return models.Pet{}, err
	}
	return *pet, nil
}

func (r *PetRepositoryImpl) DeletePetById(id int) error {
	if err := r.db.Delete(&models.Pet{}, id).Error; err != nil {
		return err
	}
	return nil
}
