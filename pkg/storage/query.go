package main

import (
	"github.com/wenyinh/go-wire-app/pkg/config"
	"github.com/wenyinh/go-wire-app/pkg/storage/model"
	"os"
	"path"

	"github.com/mcuadros/go-defaults"
	"gorm.io/gen"
)

const queryPackage = "query"

//go:generate go run query.go
func main() {
	var conf config.AppConfiguration
	defaults.SetDefaults(&conf)

	fileDir, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}

	g := gen.NewGenerator(
		gen.Config{
			OutPath:           path.Join(fileDir, queryPackage),
			ModelPkgPath:      queryPackage,
			FieldNullable:     true,
			FieldCoverable:    true,
			FieldSignable:     true,
			FieldWithIndexTag: true,
			FieldWithTypeTag:  true,
			Mode:              gen.WithDefaultQuery,
		},
	)

	g.ApplyBasic(
		model.UserDataModel{},
	)

	g.Execute()
}
