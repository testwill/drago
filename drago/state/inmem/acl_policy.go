package inmem

import (
	"context"
	"errors"
	"strings"

	structs "github.com/seashell/drago/drago/structs"
)

const (
	resourceTypePolicy = "policy"
)

// ACLPolicies :
func (r *StateRepository) ACLPolicies(ctx context.Context) ([]*structs.ACLPolicy, error) {
	prefix := resourcePrefix(resourceTypePolicy)

	items := []*structs.ACLPolicy{}
	for k, v := range r.kv {
		if strings.HasPrefix(k, prefix) {
			if t, ok := v.(*structs.ACLPolicy); ok {
				items = append(items, t)
			}
		}
	}

	return items, nil
}

// ACLPolicyByName :
func (r *StateRepository) ACLPolicyByName(ctx context.Context, name string) (*structs.ACLPolicy, error) {
	key := resourceKey(resourceTypePolicy, name)
	if v, found := r.kv[key]; found {
		return v.(*structs.ACLPolicy), nil
	}
	return nil, errors.New("not found")
}

// UpsertACLPolicy :
func (r *StateRepository) UpsertACLPolicy(ctx context.Context, p *structs.ACLPolicy) error {
	key := resourceKey(resourceTypePolicy, p.Name)
	r.kv[key] = p
	return nil
}

// DeleteACLPolicies :
func (r *StateRepository) DeleteACLPolicies(ctx context.Context, names []string) error {
	for _, name := range names {
		key := resourceKey(resourceTypePolicy, name)
		delete(r.kv, key)
	}
	return nil
}
