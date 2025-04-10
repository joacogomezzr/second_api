package domain

import "newapi/src/redomendaciones/domain/entities"

type IRecomendaciones interface {
	Register(recomendacion *entities.Recomendacion) error
}
