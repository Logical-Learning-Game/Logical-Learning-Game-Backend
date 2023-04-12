package service

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"llg_backend/config"
	"llg_backend/internal/dto"
	"llg_backend/internal/entity"
	"llg_backend/internal/token"
	"time"

	"gorm.io/gorm"
)

type adminAuthenticationService struct {
	cfg        *config.Config
	db         *gorm.DB
	tokenMaker token.Maker
}

func NewAdminAuthenticationService(cfg *config.Config, db *gorm.DB, tokenMaker token.Maker) AdminAuthenticationService {
	return &adminAuthenticationService{
		cfg:        cfg,
		db:         db,
		tokenMaker: tokenMaker,
	}
}

func (s adminAuthenticationService) Login(ctx context.Context, username, password string) (*dto.AdminLoginResponse, error) {
	var admin entity.Admin
	result := s.db.WithContext(ctx).
		Where("username = ?", username).
		First(&admin)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAdminNotFound
		}
		return nil, err
	}

	if err := checkPassword(password, admin.HashedPassword); err != nil {
		return nil, ErrUnauthorized
	}

	accessToken, err := s.createToken(username, s.cfg.JWT.Duration)
	if err != nil {
		return nil, err
	}

	adminLoginResponse := &dto.AdminLoginResponse{
		Username:    username,
		AccessToken: accessToken,
	}

	return adminLoginResponse, nil
}

func (s adminAuthenticationService) createToken(username string, duration time.Duration) (string, error) {
	return s.tokenMaker.CreateToken(username, duration)
}

func checkPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %v", err)
	}
	return string(hashedPassword), nil
}
