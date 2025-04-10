package entities

type Recomendacion struct {
	Id int32 `json:"id"`
	IdBook int32 `json:"idBook"`
}
func NewRecomendacion(idBook int32) *Recomendacion {
	return &Recomendacion{IdBook: idBook}
}