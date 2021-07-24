package service

import (
	"github.com/bektosh/fiber-app/api/errors"
	"github.com/bektosh/fiber-app/models/requests"
	"github.com/bektosh/fiber-app/models/responses"
	"github.com/google/uuid"
	"strings"
)

func (s *Service) CreateWorkspace(name string) (responses.Workspace, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		s.logger.Printf("Error while generating NewRandom UUID: %v\n", err)
		return responses.Workspace{}, errors.Internal
	}
	name = strings.TrimSpace(name)
	createdAt, err := s.storage.Psql.CreateWorkspace(id.String(), name)
	if err != nil {
		s.logger.Printf("Error while creating workspace: %v\n", err)
		return responses.Workspace{}, errors.Internal
	}
	return responses.Workspace{
		ID:        id.String(),
		Name:      name,
		CreatedAt: createdAt,
	}, nil
}

func (s *Service) GetWorkspaces(page uint64, limit uint64, name string) (responses.WorkspacesResponse, error) {
	workspaces, count, err := s.storage.Psql.GetWorkspaces(int(page), int(limit), name)
	if err != nil {
		s.logger.Println("Error while getting workspaces:", err)
		return responses.WorkspacesResponse{}, errors.Internal
	}
	return responses.WorkspacesResponse{
		Workspaces: workspaces,
		Count:      uint64(count),
		Page:       page,
	}, nil
}

func (s *Service) UpdateWorkspace(body requests.UpdateWorkspace) (responses.Workspace, error) {
	body.Name = strings.TrimSpace(body.Name)
	err := s.storage.Psql.UpdateWorkspace(body.ID, body.Name)
	if err != nil {
		s.logger.Println("Error while updating workspace:", err)
		return responses.Workspace{}, errors.Internal
	}
	return responses.Workspace{
		ID:   body.ID,
		Name: body.Name,
	}, nil
}

func (s *Service) DeleteWorkspace(id string) error {
	err := s.storage.Psql.DeleteWorkspace(id)
	if err != nil {
		s.logger.Println("Error while deleting workspace:", err)
		return errors.Internal
	}
	return nil
}
