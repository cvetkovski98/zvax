package service

import (
	"context"
	"database/sql"
	"log"

	auth "github.com/cvetkovski98/zvax-auth/internal"
	"github.com/cvetkovski98/zvax-auth/internal/dto"
	"github.com/cvetkovski98/zvax-auth/internal/mapper"
	"github.com/cvetkovski98/zvax-auth/internal/model"
	"github.com/cvetkovski98/zvax-auth/internal/token"
	jwtutil "github.com/cvetkovski98/zvax-auth/internal/util/jwt"
	"golang.org/x/crypto/bcrypt"
)

type impl struct {
	ar auth.Repository
}

func (s *impl) Register(ctx context.Context, register *dto.Register) (*dto.User, *token.Token, error) {
	// Check if user already exists
	_, err := s.GetOneByEmail(ctx, register.Email)
	if err == nil {
		return nil, nil, ErrUserAlreadyExists
	}

	// Hash password
	h, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error while hashing password: %v", err)
		return nil, nil, ErrHashingPassword
	}
	user := &model.User{
		Email:    register.Email,
		Name:     register.Name,
		Password: string(h),
	}
	inserted, err := s.ar.InsertOne(ctx, user)
	if err != nil {
		log.Printf("Error while inserting user: %v", err)
		return nil, nil, ErrCreatingUser
	}

	// Create JWT
	obj := mapper.UserModelToDto(inserted)
	token, err := jwtutil.Generate(token.NewPayload(obj.Email), "secret")
	if err != nil {
		log.Printf("Error while creating JWT: %v", err)
		return nil, nil, ErrCreatingUserJWT
	}
	return obj, token, nil
}

func (s *impl) Login(ctx context.Context, user *dto.Login) (*dto.User, *token.Token, error) {
	in, err := s.ar.FindOneByEmail(ctx, user.Email)
	if err != nil {
		log.Printf("Error while finding user by email: %v", err)
		return nil, nil, ErrUserNotFound
	}
	if err := bcrypt.CompareHashAndPassword([]byte(in.Password), []byte(user.Password)); err != nil {
		return nil, nil, ErrInvalidPassword
	}
	obj := mapper.UserModelToDto(in)
	token, err := jwtutil.Generate(token.NewPayload(obj.Email), "secret")
	if err != nil {
		log.Printf("Error while creating JWT: %v", err)
		return nil, nil, ErrCreatingUserJWT
	}
	return obj, token, nil
}

func (s *impl) GetOneByEmail(ctx context.Context, email string) (*dto.User, error) {
	in, err := s.ar.FindOneByEmail(ctx, email)
	if err != nil && err != sql.ErrNoRows {
		return nil, ErrFindingUser
	}
	if err == sql.ErrNoRows {
		log.Printf("User with email %s not found", email)
		return nil, ErrUserNotFound
	}
	obj := mapper.UserModelToDto(in)
	return obj, nil
}

func (s *impl) Me(ctx context.Context, t *token.Token) (*dto.User, error) {
	// Validate token
	claims, err := jwtutil.ParseVerify(t, "secret")
	if err != nil {
		return nil, err
	}
	return s.GetOneByEmail(ctx, claims.Sub)
}

func NewAuthServiceImpl(repository auth.Repository) auth.Service {
	return &impl{ar: repository}
}
