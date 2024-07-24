package gamestate

import (
	"context"
	"encoding/json"

	"pkg.world.dev/world-engine/cardinal/filter"
	"pkg.world.dev/world-engine/cardinal/types"
)

type Reader interface {
	// One Component One Entity
	GetComponentForEntity(cType types.ComponentMetadata, id types.EntityID) (any, error)
	GetComponentForEntityInRawJSON(cType types.ComponentMetadata, id types.EntityID) (json.RawMessage, error)

	// Many Components One Entity
	GetComponentTypesForEntity(id types.EntityID) ([]types.ComponentMetadata, error)

	// One Archetype Many Components
	GetComponentTypesForArchID(archID types.ArchetypeID) ([]types.ComponentMetadata, error)
	GetArchIDForComponents(components []types.ComponentMetadata) (types.ArchetypeID, error)

	// One Archetype Many Entities
	GetEntitiesForArchID(archID types.ArchetypeID) ([]types.EntityID, error)

	// Misc
	SearchFrom(filter filter.ComponentFilter, start int) *ArchetypeIterator
	ArchetypeCount() int
}

type Writer interface {
	// One Entity
	RemoveEntity(id types.EntityID) error

	// Many Components
	CreateEntity(comps ...types.ComponentMetadata) (types.EntityID, error)
	CreateManyEntities(num int, comps ...types.ComponentMetadata) ([]types.EntityID, error)

	// One Component One Entity
	SetComponentForEntity(cType types.ComponentMetadata, id types.EntityID, value any) error
	AddComponentToEntity(cType types.ComponentMetadata, id types.EntityID) error
	RemoveComponentFromEntity(cType types.ComponentMetadata, id types.EntityID) error

	// Misc
	Close() error
	RegisterComponents([]types.ComponentMetadata) error
}

type TickStorage interface {
	GetLastFinalizedTick() (tick uint64, err error)
	FinalizeTick(ctx context.Context) error
}

// Manager represents all the methods required to track Component, Entity, and Archetype information
// which powers the ECS dbStorage layer.
type Manager interface {
	TickStorage
	Reader
	Writer
	ToReadOnly() Reader
}
