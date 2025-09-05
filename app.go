package ocl

import (
	"github.com/dalloriam/ocl/id"
	"github.com/dalloriam/ocl/kv"
	"github.com/dalloriam/ocl/kv/badgerkv"

	gap "github.com/muesli/go-app-paths"
)

type App struct {
	id    *id.ID
	scope *gap.Scope

	kv kv.KV
}

// NewApp creates a new App instance with the given name.
func NewApp(domain, name string) *App {
	id := id.New(domain, name)
	scope := gap.NewScope(gap.User, id.String())

	return &App{id: id, scope: scope}
}

func (a *App) ensureInitKV() error {
	if a.kv != nil {
		return nil
	}

	dbPath, err := a.scope.DataPath("kv")
	if err != nil {
		return err
	}

	db, err := badgerkv.NewBadgerKV(dbPath)
	if err != nil {
		return err
	}
	a.kv = db

	return nil
}

// GetKV picks an appropriate KV implementation and returns it.
func (a *App) GetKV() (kv.KV, error) {
	if err := a.ensureInitKV(); err != nil {
		return nil, err
	}
	return a.kv, nil
}
