package usecase

import (
	"github.com/alunogabrielgava123/APRENDENDO-GO-2022/model"
	"github.com/alunogabrielgava123/APRENDENDO-GO-2022/repository"
)



type AddAlbunUseCase struct {
	repository repository.IRepositoryAlbumPg 
}


func (u *AddAlbunUseCase) Execute(newAlbum model.Album) (model.Album, error)  {
	
	value,err := u.repository.AddAlbum(newAlbum);if err != nil {
		return value, err
	}
	return value, nil
}

func AddAlbunUseCaseConstructor(repo repository.IRepositoryAlbumPg) *AddAlbunUseCase {
    return &AddAlbunUseCase{repository: repo}
}








