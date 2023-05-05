package note

import (
	"github.com/AndrewEminov/auth/internal/model"
	desc "github.com/AndrewEminov/auth/pkg/note_v1"
)

func ToInfo(info *desc.NoteInfo) *model.Info {
	return &model.Info{
		Title:     info.GetTitle(),
		Content:   info.GetContent(),
		CreatedAt: info.GetCreatedAt().AsTime(),
	}
}
