package repository

import "dushanbe_tour/internal/models"

func (r repository) GeLanguage() (lang []models.Language, errs error) {
	db := r.Postgres.GetPostgresConnection()

	sqlQuery := "select id,name,description,active from language;"
	err := db.Raw(sqlQuery).Scan(&lang).Error
	if err != nil {
		return nil, err
	}
	return lang, nil

}
