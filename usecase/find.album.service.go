package usecase

import (
	"github.com/alunogabrielgava123/APRENDENDO-GO-2022/model"
	"github.com/alunogabrielgava123/APRENDENDO-GO-2022/repository"
)


type FindeAlbunsUseCase struct {
   repository repository.IRepositoryAlbumPg 
}

func (f *FindeAlbunsUseCase) Execute() ([]model.Album, error) {
    
	album, err :=  f.repository.Find(); if err != nil {
         return album, err 
	}
	return album , nil
	
}

func FindeAlbunsUseCaseContrusctor(repo repository.IRepositoryAlbumPg) *FindeAlbunsUseCase {
	return &FindeAlbunsUseCase{repository: repo}
}
