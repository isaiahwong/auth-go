package user

import (
	"context"

	"github.com/isaiahwong/accounts-go/internal/models"
)

// Repo defines user repository operations
type Repo interface {
	Save(c context.Context, u *models.User) (string, error)
	Find(c context.Context, f interface{}, opts ...interface{}) ([]*models.User, error)
	FindOne(c context.Context, f interface{}, opts ...interface{}) (*models.User, error)
}