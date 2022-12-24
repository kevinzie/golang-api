package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"project/scratch/pkg/utils"
	"strings"
)

var Database *gorm.DB

func Connect() error {
	var err error
	// Build PostgreSQL connection URL.
	postgresConnURL, errorCon := utils.ConnectionURLBuilder("postgres")
	if errorCon != nil {
		return nil
	}

	Database, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  postgresConnURL,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "t_",   // table name prefix, table for `User` would be `t_users`
			SingularTable: true,                            // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   false,                           // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("ID", "id"), // use name replacer to change struct/field name before convert it to db name
		},
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		return err
	}

	return nil
}
