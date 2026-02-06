package service

import (
	"brantect-api-tm/interfaces"

	"brantect-api-tm/brantect"

	brantect_github "github.com/BrightsHoldings/brantect-go/brantect"
	"github.com/BrightsHoldings/gmobrs-go-lib/log"
	"github.com/BrightsHoldings/gmobrs-go-lib/where"
)

type TmRightsSystemService struct {
	tmRightsSystemRepository brantect.TmRightsSystemRepositoryInterface
	brantectRepository       brantect_github.RepositoryInterface
	tmTasksSystemRepository  brantect.TmTasksSystemRepositoryInterface
	mstClientIdRepository    brantect.MstClientIdRepositoryInterface
	mstCountryRepository     brantect_github.MstCountryRepositoryInterface
	tmMstRepository          brantect_github.TmMstRepositoryInterface
	mstOrgRepository         brantect_github.MstOrgRepositoryInterface
	mstPersonRepository      brantect_github.MstPersonRepositoryInterface
}

func NewTmRightsSystemService(
	tmRightsSystemRepository brantect.TmRightsSystemRepositoryInterface,
	brantectRepository brantect_github.RepositoryInterface,
	tmTasksSystemRepository brantect.TmTasksSystemRepositoryInterface,
	mstClientIdRepository brantect.MstClientIdRepositoryInterface,
	mstCountryRepository brantect_github.MstCountryRepositoryInterface,
	tmMstRepository brantect_github.TmMstRepositoryInterface,
	mstOrgRepository brantect_github.MstOrgRepositoryInterface,
	mstPersonRepository brantect_github.MstPersonRepositoryInterface,
) interfaces.TmRightsSystemServiceInterface {
	return &TmRightsSystemService{
		tmRightsSystemRepository: tmRightsSystemRepository,
		brantectRepository:       brantectRepository,
		tmTasksSystemRepository:  tmTasksSystemRepository,
		mstClientIdRepository:    mstClientIdRepository,
		mstCountryRepository:     mstCountryRepository,
		tmMstRepository:          tmMstRepository,
		mstOrgRepository:         mstOrgRepository,
		mstPersonRepository:      mstPersonRepository,
	}
}

func (s *TmRightsSystemService) BrantectRepository() brantect_github.RepositoryInterface {
	return s.brantectRepository
}

// Search implements interfaces.TmRightsSystemServiceInterface.
func (s *TmRightsSystemService) Search(
	tx brantect_github.TransactionInterface,
	condition where.ConditionInterface,
	filterParams brantect.TmRightsSystemSearchFilterParam,
) ([]*brantect.TmRightsSystem, int, []error) {

	tmRights, totalCount, errs := s.tmRightsSystemRepository.Search(tx, condition, filterParams)
	if errs != nil {
		log.Errorf("Failed to search tm_right info. \n")
		log.Errorf("Error: %v \n", errs)
		return nil, 0, errs
	}

	return tmRights, totalCount, nil
}

// GetByClientCdAndSysRef implements interfaces.TmRightsSystemServiceInterface.
func (s *TmRightsSystemService) GetByClientCdAndSysRef(
	tx brantect_github.TransactionInterface,
	clientCd string,
	sysRef int,
	searchAttorney string,
) (*brantect.TmRightDetail, []error) {
	tmRightDetail, errs := s.tmRightsSystemRepository.GetByClientCdAndSysRef(tx, clientCd, sysRef, searchAttorney)
	if errs != nil {
		log.Errorf("Failed to get tm_right detail. \n")
		log.Errorf("Error: %v \n", errs)
		return nil, errs
	}
	return tmRightDetail, nil
}

// GetListOfTasksForRightsByClientCdAndSysRef implements interfaces.TmRightsSystemServiceInterface.
func (s *TmRightsSystemService) GetListOfTasksForRightsByClientCdAndSysRef(
	tx brantect_github.TransactionInterface,
	clientCd string,
	sysRef int,
	searchAttorney string,
	filterParams brantect.TmRightsSystemGetListOfTasksForRightsFilterParam,
) ([]*brantect.TmTasks, int, []error) {
	// 1. get list tm_task
	tmTasks, totalCount, errs := s.tmRightsSystemRepository.GetListOfTasksForRightsByClientCdAndSysRef(tx, clientCd, sysRef, searchAttorney, filterParams)
	if errs != nil {
		log.Errorf("Failed to get list of tasks for rights. \n")
		log.Errorf("Error: %v \n", errs)
		return nil, 0, errs
	}

	// 2. get attach files for each tm_task
	for _, task := range tmTasks {
		log.Infof("tmTasks: %v \n", task.IndentJson())
		log.Infof("task.TmTaskSeq: %d \n", task.TmTaskSeq)
		attachFiles, errAt := s.tmTasksSystemRepository.GetTmTaskAtListByTmTaskSeq(tx, task.TmTaskSeq)
		if errAt != nil {
			log.Errorf("Failed to get attach files for tm_task_seq=%d. \n", task.TmTaskSeq)
			log.Errorf("Error: %v \n", errAt)
			return nil, 0, errAt
		}
		task.AttachFile = attachFiles
	}
	return tmTasks, totalCount, nil
}
