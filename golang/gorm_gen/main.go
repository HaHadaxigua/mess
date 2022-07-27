package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		ModelPkgPath:  "internal/model",
		OutPath:       "internal/model/query",
		WithUnitTest:  true,
		FieldNullable: true,
		Mode:          gen.WithQueryInterface,
	})

	db, err := gorm.Open(postgres.Open("DSN"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	g.UseDB(db)
	g.Execute()
}
