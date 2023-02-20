package usecase

import "github.com/alunogabrielgava123/APRENDENDO-GO-2022/repository"

type DeleteAlbumByIdUseCase struct {
	repository repository.IRepositoryAlbumPg
}

func DeleteAlbumByIdUseCaseConstructor(repo repository.IRepositoryAlbumPg) *DeleteAlbumByIdUseCase {
	return &DeleteAlbumByIdUseCase{repository: repo}
}

func (d *DeleteAlbumByIdUseCase) Execute(id int64) (int64, error) {
	value, err := d.repository.RemoveAlbumById(id)
	if err != nil {
		return 0, err
	}
	return value, nil
}
