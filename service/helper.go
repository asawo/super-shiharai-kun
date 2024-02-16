package service

import (
	"math"

	"github.com/morikuni/failure"

	"github.com/asawo/api/auth"
	"github.com/asawo/api/db/model"
	"github.com/asawo/api/errors"
)

// Maybe don't need
func (s *Impl) validateUser(email, password string) (*model.User, errors.ServiceError) {
	// Validate user exists
	user, serr := s.db.GetUserByEmail(nil, email)
	if serr != nil {
		return nil, serr
	}

	authenticated := auth.AuthenticateUser(user, password)
	if !authenticated {
		return nil, errors.New(errors.Unauthorized, failure.Messagef("failed to authenticate user %s (%s), wrong password", user.Name, user.Email))
	}

	return user, nil
}

// Round float64 to n decimals
func roundFloat(val float64, n uint) float64 {
	ratio := math.Pow(10, float64(n))
	return math.Round(val*ratio) / ratio
}
