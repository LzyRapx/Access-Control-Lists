package resource

import (
	err "github.com/TuSimple/Role-based-access-control/errors"
)
type NameRes struct {
	name string
}

func ParseNameRes(resString, name string) (Resource, error) {
	if name == "" {
		return nil, err.ErrParseRes
	}
	return &NameRes{name}, nil
}

func (nr *NameRes) Name() string {
	return nr.name
}

func (nr *NameRes) Equals(resource Resource) bool {
	return resource.Name() == nr.Name()
}

func (nr *NameRes) Contains(resources ...Resource) bool {
	for _, r := range resources {
		if r.Name() != nr.Name() {
			return false
		}
	}
	return true
}

func (nr *NameRes) String() string {
	return nr.Name()
}