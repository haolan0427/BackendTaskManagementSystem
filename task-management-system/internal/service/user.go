package service

import (
    "errors"
    "task-management-system/internal/model"
    "task-management-system/internal/repository"
    "task-management-system/pkg/jwt"
    
    "golang.org/x/crypto/bcrypt"
)

type UserService struct {
    userRepo *repository.UserRepository
    jwtMgr   *jwt.JWTManager
}

func NewUserService(userRepo *repository.UserRepository, jwtMgr *jwt.JWTManager) *UserService {
    return &UserService{
        userRepo: userRepo,
        jwtMgr:   jwtMgr,
    }
}

func (s *UserService) Register(user *model.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)
    
    return s.userRepo.Create(user)
}

func (s *UserService) Login(email, password string) (string, *model.User, error) {
    user, err := s.userRepo.GetByEmail(email)
    if err != nil {
        return "", nil, errors.New("invalid credentials")
    }
    
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", nil, errors.New("invalid credentials")
    }
    
    token, err := s.jwtMgr.GenerateToken(user.ID, user.Email)
    if err != nil {
        return "", nil, err
    }
    
    return token, user, nil
}