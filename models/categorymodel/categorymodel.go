package categorymodel

import (
	"go-final-project/config"
	"go-final-project/entities"
)

func GetAll() []entities.Category{
	rows, err := config.DB.Query(`Select * From categories`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var categories []entities.Category 

	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdetaedAt); err != nil {
			panic(err)
		}
		categories = append(categories, category)
	}

	return categories
}