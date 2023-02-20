package repository

import "github.com/alunogabrielgava123/APRENDENDO-GO-2022/model"



type IUserRepository interface {
    
	//Testar essa possibilidade
	CreateUser(newUser model.User) (model.User, error)
	DeleteUser(id int64) (error) 

}  

