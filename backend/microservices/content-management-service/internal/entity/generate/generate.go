package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

// go run ./internal/entity/generate/generate.go
func main() {

	err := godotenv.Load("./internal/.env")
	if err != nil {
		log.Fatalln("Check .env file: ", err)
	}

	s := os.Getenv("SQL_DATABASE_URL")
	log.Println(s)

	db, err := gorm.Open(sqlserver.Open(s), &gorm.Config{})
	if err != nil {
		log.Println("not working", err)
		log.Println(db)
		return
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:      "./internal/entity",
		OutFile:      "./entity",
		ModelPkgPath: "entity",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(db)

	dbDocument := g.GenerateModel("Document")
	dbDraftHistory := g.GenerateModel("DraftHistory")
	dbDraft := g.GenerateModel("Draft")

	g.GenerateModel("DraftHistory",
		gen.FieldType("Id", "mssql.UniqueIdentifier"),
		gen.FieldType("FileId", "mssql.UniqueIdentifier"),
		gen.FieldType("DraftId", "mssql.UniqueIdentifier"),
		gen.FieldType("UserId", "mssql.UniqueIdentifier"),
		gen.FieldRelate(field.HasOne, "Draft", dbDraft, &field.RelateConfig{
			GORMTag: field.GormTag{
				"foreignKey": []string{"DraftID"},
				"references": []string{"ID"},
			},
		}),
		gen.FieldRelate(field.HasOne, "Document", dbDocument, &field.RelateConfig{
			GORMTag: field.GormTag{
				"foreignKey": []string{"FileID"},
				"references": []string{"ID"},
			},
		}),
	)

	g.GenerateModel("Document",
		gen.FieldType("Id", "mssql.UniqueIdentifier"),
		gen.FieldRelate(field.HasMany, "DraftHistory", dbDraftHistory, &field.RelateConfig{
			GORMTag: field.GormTag{
				"foreignKey": []string{"FileID"},
			},
		}),
	)

	g.GenerateModel("Draft",
		gen.FieldType("Id", "mssql.UniqueIdentifier"),
		gen.FieldType("CompanyId", "mssql.UniqueIdentifier"),
		gen.FieldRelate(field.HasMany, "DraftHistory", dbDraftHistory, &field.RelateConfig{
			GORMTag: field.GormTag{
				"foreignKey": []string{"DraftID"},
			},
		}),
	)

	g.Execute()
}
