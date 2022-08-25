package service

import (
	"context"
	"errors"
	"shapes/schema"
	"shapes/factory"
)

type AreaRepository interface {
	InsertArea(ctx context.Context, area *schema.Area) error
}

type AreaService struct {
	Repo AreaRepository
}

type AreaRequest struct {
	ShapeType string `json:"shape_type"`
	A         int64  `json:"a"`
	B         int64  `json:"b"`
}

func (as *AreaService) InsertArea(ctx context.Context, req *AreaRequest) error {
	if req == nil {
		return errors.New("request is nil")
	}

	// Using Factory Method
	var shape factory.Shape
	switch req.ShapeType {
	case "persegi panjang":
		shape = &factory.Rectangle{
			Long: req.A,
			Wide: req.B,
		}
	case "persegi":
		shape = &factory.Square{
			Side: req.A,
		}
	case "segitiga":
		shape = &factory.Triangle{
			Base:   req.A,
			Height: req.B,
		}
	default:
		shape = &factory.DefaultShape{}
	}

	area := schema.Area{
		AreaValue: shape.Area(),
		AreaType:  shape.Name(),
	}

	if err := as.Repo.InsertArea(ctx, &area); err != nil {
		return err
	}

	return nil
}
