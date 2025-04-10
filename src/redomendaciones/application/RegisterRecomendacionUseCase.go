package application

import (
	"newapi/src/redomendaciones/domain"
	"newapi/src/redomendaciones/domain/entities"
)

type RegisterRecomendacionUseCase struct {
	recomendacionesRepository domain.IRecomendaciones
}
func NewRegisterRecomendacionUseCase(recomendacionesRepository domain.IRecomendaciones) *RegisterRecomendacionUseCase {
	return &RegisterRecomendacionUseCase{recomendacionesRepository: recomendacionesRepository}
}
func (useCase *RegisterRecomendacionUseCase) Execute(idBook int32) (*entities.Recomendacion, error) {
	recomendacion := entities.NewRecomendacion(idBook)
	err:= useCase.recomendacionesRepository.Register(recomendacion)
	if err != nil {
		return nil,err
	}
	return recomendacion,nil
}