package user

import (
	"context"
	"go-microservice/internal/shared/dto"
)

func (s *implUser) GetUserByID(ctx context.Context, userID string) (resp *dto.GetUserByIDResponse, err error) {
	var out dto.GetUserByIDResponse
	currentUser, err := s.repo.UserRepository.GetUserByID(userID)
	if err != nil {
		//TODO: ERROR LOG
		return nil, err
	}

	out.FullName = currentUser.FullName
	out.PhoneNumber = currentUser.PhoneNumber
	resp = &out
	return
}
