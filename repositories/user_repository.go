package repositories

import (
	"github.com/Raihanki/horizont-api/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetByID(userID uint) (entity.User, error)
	GetByEmail(email string) (entity.User, error)
}

type UserRepositoryImpl struct {
	tx  sqlx.Tx
	ctx *fiber.Ctx
}

func NewUserRepository(tx sqlx.Tx, ctx *fiber.Ctx) UserRepository {
	return &UserRepositoryImpl{tx, ctx}
}

func (repository *UserRepositoryImpl) GetByID(userID uint) (entity.User, error) {
	user := entity.User{}
	query := "SELECT * FROM users INNER JOIN roles ON users.role_id = roles.id WHERE id = ?"
	errGetUser := repository.tx.SelectContext(repository.ctx.Context(), &user, query, userID)
	if errGetUser != nil {
		return entity.User{}, errGetUser
	}

	return user, nil
}

func (repository *UserRepositoryImpl) GetByEmail(email string) (entity.User, error) {
	user := entity.User{}
	query := "SELECT * FROM users INNER JOIN roles ON users.role_id = roles.id WHERE email = ?"
	errGetUser := repository.tx.SelectContext(repository.ctx.Context(), &user, query, email)
	if errGetUser != nil {
		return entity.User{}, errGetUser
	}

	return user, nil
}
