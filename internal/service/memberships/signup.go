package memberships

import (
	"errors"
	"github.com/jetaimejeteveux/music-catalog/internal/models/memberships"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Signup(request memberships.SignupRequest) error {
	existingUser, err := s.repository.GetUser(request.Email, request.Email, 0)
	if err != nil {
		log.Error().Err(err).Msg("Error getting existing user")
		return err
	}
	if existingUser != nil {
		return errors.New("User already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("Error generating bcrypt hash")
		return err
	}

	model := memberships.User{
		Username:  request.Username,
		Email:     request.Email,
		Password:  string(hashedPassword),
		CreatedBy: request.Email,
		UpdatedBy: request.Email,
	}

	return s.repository.CreateUser(model)
}
