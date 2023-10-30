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

	dbUserRole := g.GenerateModel("UserRole")
	dbUser := g.GenerateModel("User")
	dbRole := g.GenerateModel("Role")

	g.GenerateModel("UserRole",
		gen.FieldType("Id", "mssql.UniqueIdentifier"),
		gen.FieldType("UserId", "mssql.UniqueIdentifier"),
		gen.FieldType("RoleId", "mssql.UniqueIdentifier"),
		gen.FieldRelate(field.HasOne, "User", dbUser, &field.RelateConfig{
			GORMTag: field.GormTag{
				"foreignKey": []string{"UserID"},
				"references": []string{"ID"},
			},
		}),
		gen.FieldRelate(field.HasOne, "Role", dbRole, &field.RelateConfig{
			GORMTag: field.GormTag{
				"foreignKey": []string{"RoleID"},
				"references": []string{"ID"},
			},
		}),
	)

	g.GenerateModel("User",
		gen.FieldType("Id", "mssql.UniqueIdentifier"),
		gen.FieldRelate(field.HasMany, "UserRoles", dbUserRole, &field.RelateConfig{
			// GORMTag: userGormTag.Append("foreignKey", "FK_UserRole_User"),
			GORMTag: field.GormTag{
				"foreignKey": []string{"UserID"},
			},
		}),
		gen.FieldRelate(field.HasMany, "Roles", dbRole, &field.RelateConfig{
			GORMTag: field.GormTag{
				"-": []string{"all"},
			},
		}),
	)

	g.GenerateModel("Role",
		gen.FieldType("Id", "mssql.UniqueIdentifier"),
		gen.FieldRelate(field.HasMany, "UserRoles", dbUserRole, &field.RelateConfig{
			GORMTag: field.GormTag{
				"foreignKey": []string{"RoleID"},
			},
		}),
		gen.FieldRelate(field.HasMany, "Users", dbUser, &field.RelateConfig{
			GORMTag: field.GormTag{
				"-": []string{"all"},
			},
		}),
	)

	g.Execute()
}
