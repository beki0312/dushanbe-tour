package repository

import "dushanbe_tour/internal/models"

func (r repository) GeCategoryInfrastructure() (infrastructure []models.InfrastructureAll, errs error) {
	db := r.Postgres.GetPostgresConnection()

	sqlQuery := "select id, name, title, address, phone, link, email, description, image, active,created_at from name_infrastructure where active=true;"
	err := db.Raw(sqlQuery).Scan(&infrastructure).Error
	if err != nil {
		return nil, err
	}
	return infrastructure, nil

}
