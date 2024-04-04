package repository

import "dushanbe_tour/internal/models"

func (r repository) GeInfrastructure(IdLanguage int64) (infrastruction []models.Infrastructure, errs error) {
	db := r.Postgres.GetPostgresConnection()

	sqlQuery := "select id,name,description from infrastructure where active=true and language = ?;"
	err := db.Raw(sqlQuery, IdLanguage).Scan(&infrastruction).Error
	if err != nil {
		return nil, err
	}
	return infrastruction, nil

}
