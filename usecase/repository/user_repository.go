package repository

import "../domain/"

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}
view raw
