package adapters

import (
	"fmt"
	"log"
	"newapi/src/core"
	_"newapi/src/redomendaciones/domain"
	"newapi/src/redomendaciones/domain/entities"
)

type MySQL struct {
    conn *core.Conn_MySQL
}
func NewMySQL() (*MySQL, error) {
    conn := core.GetDBPool()
    if conn.Err != "" {
        log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
    }
    return &MySQL{conn: conn}, nil
}

func (mysql *MySQL) Register(recomendacion *entities.Recomendacion) error {
    query := `
        INSERT INTO recomendaciones (
			id_book
        ) VALUES (?)
    `
    result, err := mysql.conn.ExecutePreparedQuery(
        query,
        recomendacion.IdBook,
    )
    if err != nil {
        fmt.Println(err)
        return err
    }

    if result != nil {
        rowsAffected, _ := result.RowsAffected()
        if rowsAffected == 1 {
            log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
            lastInsertID, err := result.LastInsertId()
            if err != nil {
                fmt.Println(err)
                return err
            }
            recomendacion.Id = int32(lastInsertID)
        } else {
            log.Printf("[MySQL] - Ninguna fila fue afectada.")
        }
    } else {
        log.Printf("[MySQL] - Resultado de la consulta es nil.")
    }
    return nil
}