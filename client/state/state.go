package state

import (
	"github.com/seashell/drago/drago/structs"
)

// Repository :
type Repository interface {
	// Name of implementation.
	Name() string

	Interfaces() ([]*structs.Interface, error)
	UpsertInterface(*structs.Interface) error

	Peers() ([]*structs.Peer, error)
	UpsertPeer(*structs.Peer) error
}
