package usecase

import "github.com/alunogabrielgava123/APRENDENDO-GO-2022/repository"

type deleteUserUseCase struct {
	repo repository.IUserRepository
}

func DeleteUserUseCase(repo repository.IUserRepository) *deleteUserUseCase {
	return &deleteUserUseCase{repo}
}

func (d *deleteUserUseCase) Execute(id int64) error {

	err := d.repo.DeleteUser(id)
	if err != nil {
		return err
	}

	return err

}
