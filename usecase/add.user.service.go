package usecase

import (
	"github.com/alunogabrielgava123/APRENDENDO-GO-2022/model"
	"github.com/alunogabrielgava123/APRENDENDO-GO-2022/repository"
)

type addUserUseCase struct {
	repo repository.IUserRepository
}

func AddUserUseCase(repository repository.IUserRepository) *addUserUseCase {
	return &addUserUseCase{repo: repository}
}

func (u *addUserUseCase) Execute(newUser model.User) (model.User, error) {
    

    //Criptografar a senha para depois adicionar
     


	user, err := u.repo.CreateUser(newUser)
	if err != nil {
		return user, err
	}

	return user, nil

}
