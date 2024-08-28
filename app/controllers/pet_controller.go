package controllers

import (
	"github.com/gin-gonic/gin"
	"whoof-go-gin/app/services"
)

type PetController interface {
	GetPetsPaginated(c *gin.Context)
	GetPetById(c *gin.Context)
	AddPet(c *gin.Context)
	UpdatePet(c *gin.Context)
	DeletePet(c *gin.Context)
}

type PetControllerImpl struct {
	userService services.PetService
}

func NewPetControllerImpl(userService services.PetService) *PetControllerImpl {
	return &PetControllerImpl{
		userService: userService,
	}
}

func (controller *PetControllerImpl) GetPetsPaginated(c *gin.Context) {
	controller.userService.GetPetsPaginated(c)
}

func (controller *PetControllerImpl) GetPetById(c *gin.Context) {
	controller.userService.GetPetById(c)
}

func (controller *PetControllerImpl) AddPet(c *gin.Context) {
	controller.userService.AddPet(c)
}

func (controller *PetControllerImpl) UpdatePet(c *gin.Context) {
	controller.userService.UpdatePet(c)
}

func (controller *PetControllerImpl) DeletePet(c *gin.Context) {
	controller.userService.DeletePet(c)
}
