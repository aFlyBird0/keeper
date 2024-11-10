/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package db

import (
	"keeper/internal/config"
	"keeper/internal/model"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const (
	MySQL    config.DBType = "mysql"
	Postgres config.DBType = "postgres"
	SQLite   config.DBType = "sqlite"
)

var (
	db *gorm.DB

	defaultGormConfig = gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	}

	// supported db types
	dbSelectors = map[config.DBType]DBSelector{
		MySQL:    &MySQLSelector{},
		Postgres: &PostgresSelector{},
		SQLite:   &SQLiteSelector{},
	}
)

func InitDB() {
	// select db by config
	db = selectDB(config.Load())

	// set gorm gen default db
	//query.SetDefault(db)

	// create tables
	err := db.AutoMigrate(&model.Item{})
	if err != nil {
		panic(err)
	}
}

func DB() *gorm.DB {
	if db == nil {
		panic("db not initialized")
	}
	return db
}

func selectDB(cfg config.Config) *gorm.DB {
	selector, ok := dbSelectors[cfg.DB.Type]
	if !ok {
		supported := make([]string, 0, len(dbSelectors))
		for k := range dbSelectors {
			supported = append(supported, string(k))
		}
		panic("unknown db type: " + string(cfg.DB.Type) +
			", supported: " + strings.Join(supported, ","))
	}

	db, err := selector.SelectDB(cfg.DB)

	if err != nil {
		panic(err)
	}

	return db
}

// DBSelector generates gorm.DB by config.DB.
type DBSelector interface {
	SelectDB(cfg config.DB) (*gorm.DB, error)
}

type MySQLSelector struct{}

func (s *MySQLSelector) SelectDB(cfg config.DB) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(cfg.DSN), &defaultGormConfig)
}

type PostgresSelector struct{}

func (s *PostgresSelector) SelectDB(cfg config.DB) (*gorm.DB, error) {
	gormCfg := defaultGormConfig
	props := cfg.Props

	// if schema is set, use it as table prefix
	if props != nil && props["schema"] != "" {
		gormCfg.NamingStrategy = schema.NamingStrategy{
			TablePrefix:   props["schema"] + ".",
			SingularTable: false,
		}
	}

	return gorm.Open(postgres.Open(cfg.DSN), &defaultGormConfig)
}

type SQLiteSelector struct{}

func (s *SQLiteSelector) SelectDB(cfg config.DB) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(cfg.DSN), &defaultGormConfig)
}
