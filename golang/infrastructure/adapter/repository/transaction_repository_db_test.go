package repository

import (
	"github.com/loxt/imersao-fullstack-fullcycle-5/domain/entity"
	"github.com/loxt/imersao-fullstack-fullcycle-5/infrastructure/adapter/repository/fixture"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestTransactionRepositoryDbInsert(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)
	repository := NewTransactionRepositoryDb(db)
	err := repository.Insert("1", "1", 12.1, entity.APPROVED, "")
	assert.Nil(t, err)
}
