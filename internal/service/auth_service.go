package service

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sim-clinic-api/internal/model"
	"sim-clinic-api/internal/repository"
	"sim-clinic-api/internal/utils"
	"time"
)

type authService struct {
	userRepo  repository.UserRepository
	roleRepo  repository.RoleRepository
	tokenRepo repository.TokenRepository
	jwtSecret string
	jwtExpire time.Duration
}

func NewAuthService(
	userRepo repository.UserRepository,
	roleRepo repository.RoleRepository,
	tokenRepo repository.TokenRepository,
	jwtSecret string,
	jwtExpire time.Duration,
) AuthService {
	return &authService{
		userRepo:  userRepo,
		roleRepo:  roleRepo,
		tokenRepo: tokenRepo,
		jwtSecret: jwtSecret,
		jwtExpire: jwtExpire,
	}
}

func (s *authService) Register(request model.RegisterRequest) (*model.User, error) {
	// Validate role exists
	role, err := s.roleRepo.FindByID(request.RoleID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &ServiceError{Message: "role not found", Code: 400}
		}
		return nil, err
	}

	// Check if username exists
	existingUser, _ := s.userRepo.FindByUsername(request.Username)
	if existingUser != nil {
		return nil, &ServiceError{Message: "username already exists", Code: 400}
	}

	// Check if email exists
	existingEmail, _ := s.userRepo.FindByEmail(request.Email)
	if existingEmail != nil {
		return nil, &ServiceError{Message: "email already exists", Code: 400}
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
		RoleID:   request.RoleID,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	// Reload user with role data
	createdUser, err := s.userRepo.FindByID(user.ID)
	if err != nil {
		return nil, err
	}

	logrus.Infof("User registered successfully: %s with role: %s", user.Username, role.Name)
	return createdUser, nil
}

func (s *authService) Login(request model.LoginRequest) (*model.LoginResponse, error) {
	// Find user by username
	user, err := s.userRepo.FindByUsername(request.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &ServiceError{Message: "invalid credentials", Code: 401}
		}
		return nil, err
	}

	// Check password
	if !utils.CheckPasswordHash(request.Password, user.Password) {
		return nil, &ServiceError{Message: "invalid credentials", Code: 401}
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user, s.jwtSecret, s.jwtExpire)
	if err != nil {
		return nil, err
	}

	response := &model.LoginResponse{
		AccessToken: token,
		Role:        user.Role.Name,
		UserID:      user.ID,
		Username:    user.Username,
		Email:       user.Email,
	}

	logrus.Infof("User logged in successfully: %s", user.Username)
	return response, nil
}

func (s *authService) Logout(tokenString string, userID uint) error {
	// Parse token untuk mendapatkan expiry time
	token, err := jwt.ParseWithClaims(tokenString, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})

	if err != nil {
		return &ServiceError{Message: "invalid token", Code: 401}
	}

	claims, ok := token.Claims.(*utils.Claims)
	if !ok || !token.Valid {
		return &ServiceError{Message: "invalid token", Code: 401}
	}

	// Simpan token ke blacklist
	blacklistedToken := &model.BlacklistedToken{
		Token:     tokenString,
		ExpiresAt: claims.ExpiresAt.Time,
		UserID:    userID,
		Reason:    "logout",
	}

	if err := s.tokenRepo.BlacklistToken(blacklistedToken); err != nil {
		return err
	}

	logrus.Infof("User %d logged out successfully", userID)
	return nil
}

func (s *authService) ValidateToken(tokenString string) (*model.User, error) {
	// Cek jika token di blacklist
	isBlacklisted, err := s.tokenRepo.IsTokenBlacklisted(tokenString)
	if err != nil {
		return nil, err
	}
	if isBlacklisted {
		return nil, &ServiceError{Message: "token has been revoked", Code: 401}
	}

	// Validate token
	token, err := jwt.ParseWithClaims(tokenString, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})

	if err != nil {
		return nil, &ServiceError{Message: "invalid token", Code: 401}
	}

	claims, ok := token.Claims.(*utils.Claims)
	if !ok || !token.Valid {
		return nil, &ServiceError{Message: "invalid token", Code: 401}
	}

	// Get user from database
	user, err := s.userRepo.FindByID(claims.UserID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &ServiceError{Message: "user not found", Code: 404}
		}
		return nil, err
	}

	return user, nil
}

type ServiceError struct {
	Message string
	Code    int
}

func (e *ServiceError) Error() string {
	return e.Message
}
