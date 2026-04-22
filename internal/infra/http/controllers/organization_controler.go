package conntollers

import "github.com/BohdanBoriak/boilerplate-go-back/internal/app"

type OrganizationController struct {
	orgService app.OrganizationService
}

func NewOrganizationController(os app.OrganizationService) OrganizationContller {
    return OrganizationController{
		orgService:os,
	}
}
func (c OrganizationController) FindMe()