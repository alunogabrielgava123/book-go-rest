package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/alunogabrielgava123/APRENDENDO-GO-2022/model"
	"github.com/alunogabrielgava123/APRENDENDO-GO-2022/usecase"
	"github.com/alunogabrielgava123/APRENDENDO-GO-2022/utils"
	"github.com/gin-gonic/gin"
)

type ControllerAlbum struct {
	serviceAddAlbum        usecase.AddAlbunUseCase
	serviceFindAll         usecase.FindeAlbunsUseCase
	serviceDeleteAlbumById usecase.DeleteAlbumByIdUseCase
}

func ControllerAlbumConstructor(addUseCase usecase.AddAlbunUseCase, findAllAlbun usecase.FindeAlbunsUseCase, deleteAlbumById usecase.DeleteAlbumByIdUseCase) *ControllerAlbum {
	return &ControllerAlbum{serviceAddAlbum: addUseCase, serviceFindAll: findAllAlbun, serviceDeleteAlbumById: deleteAlbumById}
}

func (c *ControllerAlbum) DeleteAlbumById(ctx *gin.Context) {

	id := ctx.Param("id")

	//convert to integer int64
	value, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	_, erro := c.serviceDeleteAlbumById.Execute(value)
	if erro != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": erro})
		return
	}

	ctx.Status(http.StatusOK)

}

func (c *ControllerAlbum) FindeAlbuns(ctx *gin.Context) {

	exemple := ctx.MustGet("exemple").(string)
	log.Println(exemple)

	albuns, err := c.serviceFindAll.Execute()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, albuns)

}

func (c *ControllerAlbum) AddAlbun(ctx *gin.Context) {

	var newAlbun model.Album

	//verificando o erro
	err := ctx.ShouldBindJSON(&newAlbun)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	//Validade album
	err = utils.ValidateBody(newAlbun, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	//Adicionando novo album dentro do repositorio local
	album, err := c.serviceAddAlbum.Execute(newAlbun)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	//Indentando o json com o status code do album criado
	ctx.JSON(http.StatusCreated, album)

}
