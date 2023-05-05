package note_v1

import (
	"github.com/AndrewEminov/auth/internal/service/note"
	desc "github.com/AndrewEminov/auth/pkg/note_v1"
)

type Implementation struct {
	desc.UnimplementedNoteV1Server

	noteService note.Service
}

func NewImplementation(noteService note.Service) *Implementation {
	return &Implementation{
		noteService: noteService,
	}
}
