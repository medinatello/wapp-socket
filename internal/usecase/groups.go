package usecase

import (
	"context"

	"github.com/medinatello/wapp-socket/internal/port/outbound"
)

// GroupInfo is a simplified representation of a group.
type GroupInfo struct {
	ID   string
	Name string
}

// GroupsUseCase handles logic related to groups.
type GroupsUseCase struct {
	logger outbound.Logger
}

// NewGroupsUseCase creates a new GroupsUseCase.
func NewGroupsUseCase(logger outbound.Logger) *GroupsUseCase {
	return &GroupsUseCase{logger: logger}
}

// List returns a dummy list of groups.
func (uc *GroupsUseCase) List(ctx context.Context) ([]GroupInfo, error) {
	uc.logger.Info("[GroupsUseCase] Listing dummy groups")

	// Return a hardcoded, predictable list of fake groups.
	dummyGroups := []GroupInfo{
		{ID: "group1@g.us", Name: "Family"},
		{ID: "group2@g.us", Name: "Work Friends"},
		{ID: "group3@g.us", Name: "Project X Team"},
	}

	return dummyGroups, nil
}

// Create logs the creation of a dummy group.
func (uc *GroupsUseCase) Create(ctx context.Context, name string) (GroupInfo, error) {
	uc.logger.Info("[GroupsUseCase] Simulating creation of a group", "name", name)

	newGroup := GroupInfo{
		ID:   "newgroup-fake-id@g.us",
		Name: name,
	}

	return newGroup, nil
}
