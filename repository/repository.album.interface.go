package repository

import "github.com/alunogabrielgava123/APRENDENDO-GO-2022/model"


type IRepositoryAlbumPg interface {
	AddAlbum(newAlbun model.Album) (model.Album , error)
	RemoveAlbumById(IdAlbum int64) (int64, error)
	Find() ([]model.Album, error)
}