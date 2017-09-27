package banana

import (
	"github.com/pkg/errors"

	"github.com/go-gorp/gorp"
)

type (
	// Banana describes a banana db entry
	Banana struct {
		ID   uint64 `db:"id"`
		Name string `db:"name"`
	}

	BananaFinder interface {
		// Find returns a Banana by ID
		Find(ID uint64) (*Banana, error)
	}

	BananaUpdater interface {
		// Update updates a Banana entry
		Update(b *Banana) error
	}

	BananaInserter interface {
		// Insert inserts a Banana into db
		Insert(b *Banana) error
	}

	bananaManager struct {
		dbMap gorp.SqlExecutor
	}
)

// NewBananaFinder inits and returns an instance of BananaFinder
func NewBananaFinder(dbMap gorp.SqlExecutor) BananaFinder {
	return newBananaManager(dbMap)
}

// NewBananaUpdater inits and returns an instance of BananaUpdater
func NewBananaUpdater(dbMap gorp.SqlExecutor) BananaUpdater {
	return newBananaManager(dbMap)
}

// NewBananaInserter inits and returns an instance of BananaInserter
func NewBananaInserter(dbMap gorp.SqlExecutor) BananaInserter {
	return newBananaManager(dbMap)
}

func newBananaManager(dbMap gorp.SqlExecutor) *bananaManager {
	return &bananaManager{dbMap}
}

func (m *bananaManager) Find(ID uint64) (*Banana, error) {
	var b Banana

	if err := m.dbMap.SelectOne(&b, "SELECT id, name FROM banana WHERE id = ?", ID); err != nil {
		return nil, errors.Wrap(err, "bananaManager.Find")
	}
	return &b, nil
}

func (m *bananaManager) Update(b *Banana) error {
	_, err := m.dbMap.Update(b)
	if err != nil {
		return errors.Wrap(err, "bananaManager.Update")
	}
	return nil
}

func (m *bananaManager) Insert(b *Banana) error {
	if err := m.dbMap.Insert(b); err != nil {
		return errors.Wrap(err, "bananaManager.Insert")
	}
	return nil
}
