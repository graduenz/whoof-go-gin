package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"whoof-go-gin/app/models"
	"whoof-go-gin/app/repositories"
)

type PetService interface {
	GetPetsPaginated(c *gin.Context)
	GetPetById(c *gin.Context)
	AddPet(c *gin.Context)
	UpdatePet(c *gin.Context)
	DeletePet(c *gin.Context)
}

type PetServiceImpl struct {
	petRepository repositories.PetRepository
}

func NewPetServiceImpl(petRepository repositories.PetRepository) *PetServiceImpl {
	return &PetServiceImpl{
		petRepository: petRepository,
	}
}

func (svc *PetServiceImpl) GetPetsPaginated(c *gin.Context) {
	pageIndex, _ := strconv.Atoi(c.Query("pageIndex"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	data, err := svc.petRepository.FindPetsPaginated(pageIndex, pageSize)
	if err != nil {
		// TODO: Handle error
	}

	// TODO: Add pagination info
	c.JSON(http.StatusOK, data)
}

func (svc *PetServiceImpl) GetPetById(c *gin.Context) {
	petId, _ := strconv.Atoi(c.Param("petId"))

	data, err := svc.petRepository.FindPetById(petId)
	if err != nil {
		// TODO: Handle error
	}

	c.JSON(http.StatusOK, data)
}

func (svc *PetServiceImpl) AddPet(c *gin.Context) {
	var request models.Pet
	if err := c.ShouldBindJSON(&request); err != nil {
		// TODO: Handle error
	}

	data, err := svc.petRepository.SavePet(&request)
	if err != nil {
		// TODO: Handle error
	}

	c.JSON(http.StatusCreated, data)
}

func (svc *PetServiceImpl) UpdatePet(c *gin.Context) {
	petId, _ := strconv.Atoi(c.Param("petId"))

	var request models.Pet
	if err := c.ShouldBindJSON(&request); err != nil {
		// TODO: Handle error
	}

	data, err := svc.petRepository.FindPetById(petId)
	if err != nil {
		// TODO: Handle error
	}

	data.PetType = request.PetType
	data.Name = request.Name
	data, err = svc.petRepository.SavePet(&data)

	if err != nil {
		// TODO: Handle error
	}

	c.JSON(http.StatusOK, data)
}

func (svc *PetServiceImpl) DeletePet(c *gin.Context) {
	petId, _ := strconv.Atoi(c.Param("petId"))

	if err := svc.petRepository.DeletePetById(petId); err != nil {
		// TODO: Handle error
	}

	c.JSON(http.StatusOK, gin.H{})
}
