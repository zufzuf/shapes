package repository

import (
	"context"
	"errors"
	"shapes/schema"

	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

var (
	ErrAreaIsNil = errors.New("area is nil")
)

type AreaRepository struct {
	DB *gorm.DB
}

func (r *AreaRepository) InsertArea(ctx context.Context, area *schema.Area) error {
	if area == nil {
		return ErrAreaIsNil
	}

	if err := r.DB.WithContext(ctx).Create(area).Error; err != nil {
		return eris.Wrap(err, "insert new area, an error occurred")
	}

	return nil
}
