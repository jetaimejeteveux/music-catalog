package memberships

import (
	"github.com/jetaimejeteveux/music-catalog/internal/models/memberships"
	"github.com/jetaimejeteveux/music-catalog/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(req memberships.LoginRequest) (string, error) {
	existingUser, err := s.repository.GetUser(req.Email, "", 0)
	if err != nil {
		log.Error().Err(err).Msg("Error getting existing user")
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(req.Password))
	if err != nil {
		log.Error().Err(err).Msg("Error hashing password")
		return "", err
	}

	token, err := jwt.CreateToken(existingUser.ID, existingUser.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Msg("Error generating token")
		return "", err
	}

	return token, nil
}
