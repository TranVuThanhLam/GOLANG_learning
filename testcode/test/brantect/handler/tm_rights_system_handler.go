package handler

import (
	"brantect-api-tm/interfaces"
)

type TmRightsSystemHandler struct {
	tmRightsSystemService interfaces.TmRightsSystemServiceInterface
}

func NewTmRightsSystemHandler(
	tmRightsSystemService interfaces.TmRightsSystemServiceInterface,
) interfaces.TmRightsSystemHandlerInterface {
	return &TmRightsSystemHandler{
		tmRightsSystemService: tmRightsSystemService,
	}
}
