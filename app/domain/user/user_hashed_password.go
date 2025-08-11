package user

import (
	"unicode/utf8"

	"github.com/atimot/app/domain/errors"
	"github.com/atimot/pkg/hash"
)

type HashedPassword struct {
	value string
}

const (
	minPasswordLength = 6
)

func NewHashedPassword(value string) (HashedPassword, error) {
	if minPasswordLength > utf8.RuneCountInString(value) {
		return HashedPassword{}, errors.ErrPasswordTooShort
	}

	hashed, err := hash.Hash(value)
	if err != nil {
		return HashedPassword{}, err
	}

	return HashedPassword{value: hashed}, nil
}

func reconstructHashedPassword(value string) HashedPassword {
	return HashedPassword{value: value}
}

func (hp HashedPassword) Value() string {
	return hp.value
}

func (p HashedPassword) compare(target string) error {
	if err := hash.Compare(p.value, target); err != nil {
		return errors.ErrPasswordMismatch
	}

	return nil
}
