// file 2086000001001
package infrastructure

import (
	"brantect-api-tm/brantect"
	"database/sql"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"

	brantect_github "github.com/BrightsHoldings/brantect-go/brantect"
	"github.com/BrightsHoldings/gmobrs-go-lib/errors"
	"github.com/BrightsHoldings/gmobrs-go-lib/log"
	"github.com/BrightsHoldings/gmobrs-go-lib/where"
)

// sqlNullString converts an empty string to nil for SQL parameters
// This prevents empty strings from being stored as space-padded values
func sqlNullString(s string) interface{} {
	if s == "" {
		return nil
	}
	return s
}

var (
	//go:embed sql/tm_rights_system/search_tm_rights.sql
	sqlTmRightsSystemSearch string
	//go:embed sql/tm_rights_system/count_search_tm_rights.sql
	sqlCountTmRightsSystemSearch string
	//go:embed sql/tm_rights_system/get_tm_rights_by_client_cd_and_sys_ref.sql
	sqlGetTmRightsSystemByClientCdAndSysRef string
	//go:embed sql/tm_rights_system/get_tm_tasks_of_rights_by_client_cd_and_sys_ref.sql
	sqlGetTmTaskOfRightsByClientCdAndSysRef string
	//go:embed sql/tm_rights_system/count_get_tm_tasks_of_rights_by_client_cd_and_sys_ref.sql
	sqlCountGetTmTaskOfRightsByClientCdAndSysRef string
	//go:embed sql/tm_rights_system/insert_tm_base_class.sql
	sqlInsertTmBaseClass string
	//go:embed sql/tm_rights_system/insert_tm_base_cst_value.sql
	sqlInsertTmBaseCstValue string
	//go:embed sql/tm_rights_system/insert_tm_base_cst_history.sql
	sqlInsertTmBaseCstHistory string
	//go:embed sql/tm_rights_system/insert_tm_base.sql
	sqlInsertTmBase string
	//go:embed sql/tm_rights_system/delete_tm_base_class_by_clientcd_and_sysref.sql
	sqlDeleteTmBaseClassByClientCdAndSysRef string
	//go:embed sql/tm_rights_system/get_manage_no_rule_by_client_cd.sql
	sqlGetManageNoRuleByClientCd string
	//go:embed sql/tm_rights_system/get_next_international_no.sql
	sqlGetNextInternationalNo string
	//go:embed sql/tm_rights_system/existed_international_no.sql
	sqlExistedInternationalNo string
	//go:embed sql/tm_rights_system/get_next_international_no_id.sql
	sqlGetNextInternationalNoId string
	//go:embed sql/tm_rights_system/insert_international_no.sql
	sqlInsertInternationalNo string
	//go:embed sql/tm_rights_system/existed_our_ref.sql
	sqlExistedOurRef string
	//go:embed sql/tm_rights_system/get_manage_no_rule_full_by_client_cd.sql
	sqlGetManageNoRuleFullByClientCd string
	//go:embed sql/tm_rights_system/get_last_our_ref_digit_only.sql
	sqlGetLastOurRefDigitOnly string
	//go:embed sql/tm_rights_system/get_last_our_ref_pattern.sql
	sqlGetLastOurRefPattern string
	//go:embed sql/tm_rights_system/exists_by_fixcd_and_divkey.sql
	sqlExistsByFixCdAndDivKey string
	//go:embed sql/tm_rights_system/get_by_reg_cd.sql
	sqlGetByRegCd string
	//go:embed sql/tm_rights_system/get_by_prc_cd.sql
	sqlGetByPrcCd string
	//go:embed sql/tm_rights_system/exists_org.sql
	sqlExistsOrg string
	//go:embed sql/tm_rights_system/exists_person.sql
	sqlExistsPerson string
	//go:embed sql/tm_rights_system/exists_base_cst_field.sql
	sqlExistsBaseCstField string
	//go:embed sql/tm_rights_system/exists_mst_client.sql
	sqlExistsMstClient string
	//go:embed sql/get_max_done_task_date.sql
	sqlGetMaxDoneTaskDate string
	//go:embed sql/tm_rights_system/update_tm_base_by_clientcd_and_sysref.sql
	sqlUpdateTmBaseByClientCdAndSysRef string
	//go:embed sql/tm_rights_system/update_base_cst_value.sql
	sqlUpdateBaseCstValue string
	//go:embed sql/tm_rights_system/attorney_authorized_to_update_rights.sql
	sqlAttorneyAuthorizedToUpdateRights string
	//go:embed sql/tm_rights_system/update_tm_base_soft_delete.sql
	sqlUpdateTmBaseSoftDelete string
)

type TmRightsSystemRepositoryPGSQL struct {
	brantectDatabase brantect_github.BrantectDatabaseInterface
}

func NewTmRightsSystemRepositoryPGSQL(
	brantectDatabase brantect_github.BrantectDatabaseInterface,
) brantect.TmRightsSystemRepositoryInterface {
	return &TmRightsSystemRepositoryPGSQL{
		brantectDatabase: brantectDatabase,
	}
}

// Search retrieves a list of trademark rights based on filter conditions and pagination.
func (t *TmRightsSystemRepositoryPGSQL) Search(
	tx brantect_github.TransactionInterface,
	condition where.ConditionInterface,
	filterParams brantect.TmRightsSystemSearchFilterParam,
) (
	tmRights []*brantect.TmRightsSystem,
	totalTmRights int,
	err []error,
) {
	// Build SQL query and parameters for search and count
	var (
		sqlStr            = sqlTmRightsSystemSearch
		countSqlStr       = sqlCountTmRightsSystemSearch
		conditionQueryStr = ""
		paramIdx          = 1
		valueParamCount   []interface{}
	)

	cb := where.ConditionBuilder{}
	cb.Where(condition)
	conditionStr, conditionValues := cb.Build(1)
	valueParamCount = append(valueParamCount, conditionValues...)
	if conditionStr != "" {
		// Add filter conditions to both main and count queries
		paramIdx += len(conditionValues)
	}
	if filterParams.ClientCd != "" {
		conditionQueryStr += fmt.Sprintf("tb.client_cd = $%d AND ", paramIdx)
		valueParamCount = append(valueParamCount, filterParams.ClientCd)
		conditionValues = append(conditionValues, filterParams.ClientCd)
		paramIdx++
	}
	if filterParams.TmName != "" {
		conditionQueryStr += fmt.Sprintf("tm.tm_name ILIKE $%d AND ", paramIdx)
		valueParamCount = append(valueParamCount, "%"+filterParams.TmName+"%")
		conditionValues = append(conditionValues, "%"+filterParams.TmName+"%")
		paramIdx++
	}
	if filterParams.AppNo != "" {
		conditionQueryStr += fmt.Sprintf("tb.app_no = $%d AND ", paramIdx)
		valueParamCount = append(valueParamCount, filterParams.AppNo)
		conditionValues = append(conditionValues, filterParams.AppNo)
		paramIdx++
	}
	if filterParams.AppDate != "" {
		conditionQueryStr += fmt.Sprintf("tb.app_date = $%d AND ", paramIdx)
		valueParamCount = append(valueParamCount, filterParams.AppDate)
		conditionValues = append(conditionValues, filterParams.AppDate)
		paramIdx++
	}

	if len(conditionQueryStr) > 0 {
		if len(conditionStr) > 0 {
			conditionQueryStr = "AND " + strings.TrimSuffix(conditionQueryStr, " AND ")
		} else {
			conditionQueryStr = "WHERE " + strings.TrimSuffix(conditionQueryStr, " AND ")
		}
	}

	sqlStr = fmt.Sprintf(sqlStr, conditionStr, conditionQueryStr)
	sqlStr += " ORDER BY tb.sys_ref DESC"
	// Add sorting and pagination if needed
	if filterParams.Limit != 0 {
		// LIMIT and OFFSET for pagination
		sqlStr += fmt.Sprintf(" LIMIT $%d OFFSET $%d", paramIdx, paramIdx+1)
		conditionValues = append(conditionValues, int64(filterParams.Limit), int64(filterParams.Offset))
	}

	// Prepare and execute the main query
	stmt, errPrepare := tx.Get().(*sql.Tx).Prepare(sqlStr)
	if errPrepare != nil {
		log.Errorf("Search error: %v", errPrepare)
		return nil, 0, append(err, errors.New(2030010004001, errPrepare.Error(), nil))
	}
	defer stmt.Close()

	rows, errExec := stmt.Query(conditionValues...)
	if errExec != nil {
		log.Errorf("Search error: %v", errExec)
		return nil, 0, append(err, errors.New(2030010004002, errExec.Error(), nil))
	}
	defer rows.Close()

	// Scan each row into TmRights struct
	for rows.Next() {
		tmRight := &brantect.TmRightsSystem{}
		errScan := rows.Scan(
			&tmRight.ClientCd,
			&tmRight.SysRef,
			&tmRight.AppNo,
			&tmRight.SearchAppNo,
			&tmRight.AppDate,
			&tmRight.ManageFlag,
			&tmRight.TmRef,
			&tmRight.DomAttorneyClientCd,
			&tmRight.IntAttorneyClientCd,
			&tmRight.CompanyNameJp,
			&tmRight.CompanyNameEn,
			&tmRight.TmName,
			&tmRight.TmFileNm,
			&tmRight.SearchTmName,
			&tmRight.AppType, //field add
			&tmRight.CountryCd,
			&tmRight.TmClass,
			&tmRight.RegType,
			&tmRight.Registrant,
			&tmRight.RegCountryCd,
			&tmRight.AppDateUpd,
			&tmRight.PubDate,
			&tmRight.PubNo,
			&tmRight.PubDateUpd,
			&tmRight.RegDate,
			&tmRight.RegNo,
			&tmRight.RegDateUpd,
			&tmRight.Expire,
			&tmRight.NextLegalProc,
			&tmRight.InpUser,
			&tmRight.UpdUser,
			&tmRight.InpDate,
			&tmRight.UpdDate,
			&tmRight.RyomaItemNo,
			&tmRight.TmClassNm,
		)
		if errScan != nil {
			log.Errorf("Search error: %v", errScan)
			return nil, 0, append(err, errors.New(2030010004003, errScan.Error(), nil))
		}
		tmRights = append(tmRights, tmRight)
	}

	// Prepare and execute the count query to get total records
	countSqlStr = fmt.Sprintf(countSqlStr, conditionStr, conditionQueryStr)
	stmtCount, errPrepareCount := tx.Get().(*sql.Tx).Prepare(countSqlStr)
	if errPrepareCount != nil {
		log.Errorf("Search error: %v", errPrepareCount)
		return nil, 0, append(err, errors.New(2030010004004, errPrepareCount.Error(), nil))
	}
	defer stmtCount.Close()

	// Scan the total count of records
	errQueryCount := stmtCount.QueryRow(valueParamCount...).Scan(&totalTmRights)
	if errQueryCount != nil {
		log.Errorf("Search error: %v", errQueryCount)
		return nil, 0, append(err, errors.New(2030010004005, errQueryCount.Error(), nil))
	}
	return tmRights, totalTmRights, err
}

// GetByClientCdAndSysRef retrieves the details of a trademark right by clientCd and sysRef.
// If searchAttorney is provided, filters by attorney client code.
func (t *TmRightsSystemRepositoryPGSQL) GetByClientCdAndSysRef(
	tx brantect_github.TransactionInterface,
	clientCd string,
	sysRef int,
	searchAttorney string,
) (
	tmRight *brantect.TmRightDetail,
	err []error,
) {
	// Build SQL and arguments for detail query
	sqlStr := sqlGetTmRightsSystemByClientCdAndSysRef
	args := []interface{}{clientCd, sysRef}

	// According to the request from the project manager, BRANTECT users should have access to all data
	// owned by them. Therefore, filtering conditions are only applied when the request comes from Attorney
	if searchAttorney != "" {
		// If searchAttorney is provided, add filter for attorney client code
		sqlStr += " WHERE (tb.dom_attorney_client_cd = $3 OR tb.int_attorney_client_cd = $3)"
		args = append(args, searchAttorney)
	}

	stmt, errPrepare := tx.Get().(*sql.Tx).Prepare(sqlStr)
	if errPrepare != nil {
		log.Errorf("GetByClientCdAndSysRef error: %v", errPrepare)
		return nil, append(err, errors.New(2030010004006, errPrepare.Error(), nil))
	}
	defer stmt.Close()

	row := stmt.QueryRow(args...)
	tmRight = &brantect.TmRightDetail{}
	errScan := row.Scan(
		&tmRight.ClientCd, &tmRight.SysRef, &tmRight.TmRef, &tmRight.NoticeFlag,
		&tmRight.ContactFlag, &tmRight.ServiceTypeFlag, &tmRight.AppType, &tmRight.ClassType,
		&tmRight.RegCd, &tmRight.PrcCd, &tmRight.RegType, &tmRight.RegCountryCd,
		&tmRight.NextLegalProc, &tmRight.AgentJp, &tmRight.ReqDept, &tmRight.ExpDept,
		&tmRight.PrcDept, &tmRight.ReqPerson, &tmRight.ExpPerson, &tmRight.PrcPerson,
		&tmRight.ManageFlag, &tmRight.OurRef, &tmRight.Registrant, &tmRight.RegAdd,
		&tmRight.ChkFileNo, &tmRight.AppNo, &tmRight.AppNoUpd, &tmRight.PubNo,
		&tmRight.PubNoUpd, &tmRight.RegNo, &tmRight.RegNoUpd, &tmRight.Expire,
		&tmRight.ChkDate, &tmRight.AppDate, &tmRight.AppDateUpd, &tmRight.PubDate,
		&tmRight.PubDateUpd, &tmRight.RegDate, &tmRight.RegDateUpd, &tmRight.PerFlg,
		&tmRight.NextLegalProcDate, &tmRight.IntRegNo, &tmRight.IntRegDate, &tmRight.IntRegSubDate,
		&tmRight.TmClass, &tmRight.DomAttorneyClientCd, &tmRight.IntAttorneyClientCd, &tmRight.CompanyNameJp,
		&tmRight.CompanyNameEn, &tmRight.CountryCd, &tmRight.CountryNmJp, &tmRight.CountryNmEn,
		&tmRight.AreawCd, &tmRight.AreawNmJp, &tmRight.AreawNmEn, &tmRight.ArealCd,
		&tmRight.ArealNmJp, &tmRight.ArealNmEn, &tmRight.UpdatedGuidanceJp, &tmRight.UpdatedGuidanceEn,
		&tmRight.PatentOfficeContactJp, &tmRight.PatentOfficeContactEn, &tmRight.TypeOfSeviceJp, &tmRight.TypeOfSeviceEn,
		&tmRight.AppTypeNmJp, &tmRight.AppTypeNmEn, &tmRight.ClassTypeNmJp, &tmRight.ClassTypeNmEn,
		&tmRight.RegTypeNmJp, &tmRight.RegTypeNmEn, &tmRight.NextLegalProcNmJp, &tmRight.NextLegalProcNmEn,
		&tmRight.TmFileNm, &tmRight.TmName, &tmRight.ManageNo, &tmRight.RegNmJp,
		&tmRight.RegNmEn, &tmRight.PrcNmJp, &tmRight.PrcNmEn, &tmRight.RegCountryNmJp,
		&tmRight.RegCountryNmEn, &tmRight.AgentJpNm, &tmRight.AgentItnNm, &tmRight.ReqDeptNm,
		&tmRight.ExpDeptNm, &tmRight.PrcDeptNm, &tmRight.ReqPersonNm, &tmRight.ExpPersonNm,
		&tmRight.PrcPersonNm, &tmRight.AgentItn, &tmRight.Remarks, &tmRight.TmClassNm,
	)
	if errScan != nil {
		if errScan == sql.ErrNoRows {
			log.Errorf("GetByClientCdAndSysRef error: %v", errScan)
			return nil, nil
		}
		log.Errorf("GetByClientCdAndSysRef error: %v", errScan)
		return nil, append(err, errors.New(2030010004007, errScan.Error(), nil))
	}
	return tmRight, nil
}

// GetListOfTasksForRightsByClientCdAndSysRef retrieves the list of tasks for a trademark right.
// Supports filtering by attorney and pagination.
func (t *TmRightsSystemRepositoryPGSQL) GetListOfTasksForRightsByClientCdAndSysRef(
	tx brantect_github.TransactionInterface,
	clientCd string,
	sysRef int,
	searchAttorney string,
	filterParams brantect.TmRightsSystemGetListOfTasksForRightsFilterParam,
) (
	tmTasks []*brantect.TmTasks,
	totalCount int,
	err []error,
) {
	// Build SQL and arguments for task list and count
	sqlStr := sqlGetTmTaskOfRightsByClientCdAndSysRef
	countSqlStr := sqlCountGetTmTaskOfRightsByClientCdAndSysRef
	args := []interface{}{clientCd, sysRef}
	countArgs := []interface{}{clientCd, sysRef}

	// According to the request from the project manager, BRANTECT users should have access to all data
	// owned by them. Therefore, filtering conditions are only applied when the request comes from Attorney
	if searchAttorney != "" {
		// If searchAttorney is provided, add filter for attorney client code
		sqlStr += " WHERE (tb.dom_attorney_client_cd = $3 OR tb.int_attorney_client_cd = $3)"
		args = append(args, searchAttorney)
		countSqlStr += " WHERE (tb.dom_attorney_client_cd = $3 OR tb.int_attorney_client_cd = $3)"
		countArgs = append(countArgs, searchAttorney)
	}

	if filterParams.Limit != 0 {
		// Add LIMIT/OFFSET for pagination
		sqlStr += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)+1, len(args)+2)
		args = append(args, int64(filterParams.Limit), int64(filterParams.Offset))
	}

	log.Infof("sqlStr: %s", sqlStr)
	log.Infof("args: %v", args)

	stmt, errPrepare := tx.Get().(*sql.Tx).Prepare(sqlStr)
	if errPrepare != nil {
		log.Errorf("GetListOfTasksForRightsByClientCdAndSysRef error: %v", errPrepare)
		err = append(err, errors.New(2030010004008, errPrepare.Error(), nil))
		return nil, 0, err
	}
	defer stmt.Close()

	rows, errExec := stmt.Query(args...)
	if errExec != nil {
		log.Errorf("GetListOfTasksForRightsByClientCdAndSysRef error: %v", errExec)
		return nil, 0, append(err, errors.New(2030010004009, errExec.Error(), nil))
	}
	defer rows.Close()

	for rows.Next() {
		tmTask := &brantect.TmTasks{}
		errScan := rows.Scan(
			&tmTask.TmTaskSeq, &tmTask.ClientCd, &tmTask.CompanyNameJp, &tmTask.CompanyNameEn,
			&tmTask.SysRef, &tmTask.Seq, &tmTask.InpDate, &tmTask.TmRef,
			&tmTask.TmName, &tmTask.TmFileNm, &tmTask.CountryCd, &tmTask.CountryNmJp,
			&tmTask.CountryNmEn, &tmTask.TmClass, &tmTask.RegNo, &tmTask.OurRef,
			&tmTask.MakeFlg, &tmTask.Memo, &tmTask.DoneFlg, &tmTask.Content,
			&tmTask.Date, &tmTask.ChargePersonCd, &tmTask.ChargePersonNm, &tmTask.LegalProcDate,
			&tmTask.Action1, &tmTask.Action1Date, &tmTask.Action1DoneDate, &tmTask.Action2,
			&tmTask.Action2Date, &tmTask.Action2DoneDate, &tmTask.IsInternal,
		)
		if errScan != nil {
			log.Errorf("GetListOfTasksForRightsByClientCdAndSysRef error: %v", errScan)
			return nil, 0, append(err, errors.New(20300100040010, errScan.Error(), nil))
		}
		tmTasks = append(tmTasks, tmTask)
	}

	stmtCount, prepareErr := tx.Get().(*sql.Tx).Prepare(countSqlStr)
	if prepareErr != nil {
		log.Errorf("GetListOfTasksForRightsByClientCdAndSysRef error: %v", prepareErr)
		return nil, 0, append(err, errors.New(20300100040011, prepareErr.Error(), nil))
	}
	defer stmtCount.Close()

	errQueryCount := stmtCount.QueryRow(countArgs...).Scan(&totalCount)
	if errQueryCount != nil {
		log.Errorf("GetListOfTasksForRightsByClientCdAndSysRef error: %v", errQueryCount)
		return nil, 0, append(err, errors.New(20300100040012, errQueryCount.Error(), nil))
	}

	return tmTasks, totalCount, err
}

// InsertBaseClass inserts a new record into the tm_base_class table.
func (t *TmRightsSystemRepositoryPGSQL) InsertBaseClass(
	tx brantect_github.TransactionInterface,
	baseClass *brantect.TmBaseClass,
) error {
	stmt, errPrepare := tx.Get().(*sql.Tx).Prepare(sqlInsertTmBaseClass)
	if errPrepare != nil {
		log.Errorf("InsertBaseClass error: %v", errPrepare)
		return errors.New(2030010004013, errPrepare.Error(), nil)
	}
	defer stmt.Close()

	_, errExec := stmt.Exec(
		baseClass.ClientCd,
		baseClass.SysRef,
		baseClass.Seq,
		baseClass.TmClass,
		baseClass.TmClassNm,
		baseClass.UpdUser,
		baseClass.InpUser,
	)
	if errExec != nil {
		log.Errorf("InsertBaseClass error: %v", errExec)
		return errors.New(2030010004014, errExec.Error(), nil)
	}
	return nil
}

// InsertTmBase inserts a new record into the tm_base table and returns the generated sys_ref.
func (t *TmRightsSystemRepositoryPGSQL) InsertTmBase(
	tx brantect_github.TransactionInterface,
	tmBase *brantect.TmBase,
) (int, error) {

	stmt, errPrepare := tx.Get().(*sql.Tx).Prepare(sqlInsertTmBase)
	if errPrepare != nil {
		log.Errorf("InsertTmBase error: %v", errPrepare)
		return 0, errors.New(2030010004015, errPrepare.Error(), nil)
	}
	defer stmt.Close()

	//  Convert empty strings to nil for SQL parameters to avoid space padding
	appType := sqlNullString(tmBase.AppType)
	countryCd := sqlNullString(tmBase.CountryCd)
	classType := sqlNullString(tmBase.ClassType)
	tmClass := sqlNullString(tmBase.TmClass)
	regType := sqlNullString(tmBase.RegType)
	registrant := sqlNullString(tmBase.Registrant)
	regCountryCd := sqlNullString(tmBase.RegCountryCd)
	regAdd := sqlNullString(tmBase.RegAdd)
	regCd := sqlNullString(tmBase.RegCd)
	prcCd := sqlNullString(tmBase.PrcCd)
	chkDate := sqlNullString(tmBase.ChkDate)
	chkFileNo := sqlNullString(tmBase.ChkFileNo)
	appDate := sqlNullString(tmBase.AppDate)
	appNo := sqlNullString(tmBase.AppNo)
	appDateUpd := sqlNullString(tmBase.AppDateUpd)
	appNoUpd := sqlNullString(tmBase.AppNoUpd)
	pubDate := sqlNullString(tmBase.PubDate)
	pubNo := sqlNullString(tmBase.PubNo)
	pubDateUpd := sqlNullString(tmBase.PubDateUpd)
	pubNoUpd := sqlNullString(tmBase.PubNoUpd)
	regDate := sqlNullString(tmBase.RegDate)
	regNo := sqlNullString(tmBase.RegNo)
	regDateUpd := sqlNullString(tmBase.RegDateUpd)
	regNoUpd := sqlNullString(tmBase.RegNoUpd)
	expire := sqlNullString(tmBase.Expire)
	nextLegalProc := sqlNullString(tmBase.NextLegalProc)
	nextLegalProcDate := sqlNullString(tmBase.NextLegalProcDate)
	agentJp := sqlNullString(tmBase.AgentJp)
	agentItn := sqlNullString(tmBase.AgentItn)
	reqDept := sqlNullString(tmBase.ReqDept)
	reqPerson := sqlNullString(tmBase.ReqPerson)
	expDept := sqlNullString(tmBase.ExpDept)
	expPerson := sqlNullString(tmBase.ExpPerson)
	prcDept := sqlNullString(tmBase.PrcDept)
	prcPerson := sqlNullString(tmBase.PrcPerson)
	perFlg := sqlNullString(tmBase.PerFlg)
	remarks := sqlNullString(tmBase.Remarks)
	intRegNo := sqlNullString(tmBase.IntRegNo)
	intRegDate := sqlNullString(tmBase.IntRegDate)
	intRegSubDate := sqlNullString(tmBase.IntRegSubDate)
	domAttorneyClientCd := sqlNullString(tmBase.DomAttorneyClientCd)
	intAttorneyClientCd := sqlNullString(tmBase.IntAttorneyClientCd)
	ryomaItemNo := sqlNullString(tmBase.RyomaItemNo)
	ryomaUserNo := sqlNullString(tmBase.RyomaUserNo)

	// Create a nullable version of tmBase.TmRef
	var tmRefParam interface{}
	if tmBase.TmRef == 0 {
		tmRefParam = nil
	} else {
		tmRefParam = tmBase.TmRef
	}

	var sysRef int
	errQuery := stmt.QueryRow(
		tmBase.ClientCd, tmBase.OurRef, tmRefParam, appType,
		countryCd, classType, tmClass, regType,
		registrant, regCountryCd, regAdd, regCd,
		prcCd, chkDate, chkFileNo, appDate,
		appNo, appDateUpd, appNoUpd, pubDate,
		pubNo, pubDateUpd, pubNoUpd, regDate,
		regNo, regDateUpd, regNoUpd, expire,
		nextLegalProc, nextLegalProcDate, agentJp, agentItn,
		reqDept, reqPerson, expDept, expPerson,
		prcDept, prcPerson, perFlg, remarks,
		intRegNo, intRegDate, intRegSubDate,
		domAttorneyClientCd, intAttorneyClientCd,
		tmBase.UpdUser, tmBase.InpUser,
		ryomaItemNo, ryomaUserNo,
	).Scan(&sysRef)
	if errQuery != nil {
		log.Errorf("InsertTmBase error: %v", errQuery)
		return 0, errors.New(2030010004016, errQuery.Error(), nil)
	}
	return sysRef, nil
}

// DeleteBaseClassByClientCdAndSysRef deletes all records in tm_base_class for a given clientCd and sysRef.
func (t *TmRightsSystemRepositoryPGSQL) DeleteBaseClassByClientCdAndSysRef(
	tx brantect_github.TransactionInterface,
	clientCd string,
	sysRef int,
) error {
	stmt, errPrepare := tx.Get().(*sql.Tx).Prepare(sqlDeleteTmBaseClassByClientCdAndSysRef)
	if errPrepare != nil {
		log.Errorf("DeleteBaseClassByClientCdAndSysRef error: %v", errPrepare)
		return errors.New(2030010004017, errPrepare.Error(), nil)
	}
	defer stmt.Close()

	_, errExec := stmt.Exec(clientCd, sysRef)
	if errExec != nil {
		log.Errorf("DeleteBaseClassByClientCdAndSysRef error: %v", errExec)
		return errors.New(2030010004018, errExec.Error(), nil)
	}
	return nil
}

// InsertBaseCstValue inserts a custom field value for a trademark right.
func (t *TmRightsSystemRepositoryPGSQL) InsertBaseCstValue(
	tx brantect_github.TransactionInterface,
	cstValue *brantect.TmBaseCstValue,
) error {
	stmt, errPrepare := tx.Get().(*sql.Tx).Prepare(sqlInsertTmBaseCstValue)
	if errPrepare != nil {
		log.Errorf("InsertBaseCstValue error: %v", errPrepare)
		return errors.New(2030010004019, errPrepare.Error(), nil)
	}
	defer stmt.Close()

	_, errExec := stmt.Exec(
		cstValue.TmBaseFieldId,
		cstValue.SysRef,
		cstValue.FieldValue,
		cstValue.UpdUser,
		cstValue.InpUser,
	)
	if errExec != nil {
		log.Errorf("InsertBaseCstValue error: %v", errExec)
		return errors.New(2030010004020, errExec.Error(), nil)
	}
	return nil
}

// InsertBaseCstHistory inserts a custom field history for a task.
func (t *TmRightsSystemRepositoryPGSQL) InsertBaseCstHistory(
	tx brantect_github.TransactionInterface,
	cstHistory *brantect.TmBaseCstHistory,
) error {
	stmt, errPrepare := tx.Get().(*sql.Tx).Prepare(sqlInsertTmBaseCstHistory)
	if errPrepare != nil {
		log.Errorf("InsertBaseCstHistory error: %v", errPrepare)
		return errors.New(2030010004021, errPrepare.Error(), nil)
	}
	defer stmt.Close()

	_, errExec := stmt.Exec(
		cstHistory.TmTaskSeq,
		cstHistory.TmBaseFieldId,
		cstHistory.FieldValue,
		cstHistory.UpdUser,
		cstHistory.InpUser,
	)
	if errExec != nil {
		log.Errorf("InsertBaseCstHistory error: %v", errExec)
		return errors.New(2030010004022, errExec.Error(), nil)
	}
	return nil
}

// GetManageNoRuleByClientCd retrieves the manage_no rule for a client (used for our_ref generation).
func (t *TmRightsSystemRepositoryPGSQL) GetManageNoRuleByClientCd(
	tx brantect_github.TransactionInterface,
	clientCd string,
) (*brantect.TmBaseManageNoRule, error) {
	row := tx.Get().(*sql.Tx).QueryRow(sqlGetManageNoRuleByClientCd, clientCd)
	var rule brantect.TmBaseManageNoRule
	if err := row.Scan(&rule.MadridProtFlg, &rule.CountryCdFieldNo); err != nil {
		log.Errorf("GetManageNoRuleByClientCd error: %v", err)
		return nil, errors.New(2030010004023, err.Error(), nil)
	}
	return &rule, nil
}

// GetNextInternationalNo generates the next international management number for a client (matches PHP business logic).
func (t *TmRightsSystemRepositoryPGSQL) GetNextInternationalNo(
	tx brantect_github.TransactionInterface,
	clientCd string,
) (string, error) {
	row := tx.Get().(*sql.Tx).QueryRow(sqlGetNextInternationalNo, clientCd)
	var maxNo sql.NullString
	if err := row.Scan(&maxNo); err != nil && err != sql.ErrNoRows {
		log.Errorf("GetNextInternationalNo error: %v", err)
		return "", errors.New(2030010004024, err.Error(), nil)
	}

	// Generate the next international number based on the maximum number found.
	var manageNo string
	loc, _ := time.LoadLocation("Asia/Tokyo")
	nowYear := time.Now().In(loc).Format("06")
	if !maxNo.Valid || len(maxNo.String) < 4 || maxNo.String[2:4] != nowYear {
		manageNo = "WW" + nowYear + "0001"
	} else {
		prefix := maxNo.String[:4]
		num, _ := strconv.Atoi(maxNo.String[4:])
		manageNo = prefix + fmt.Sprintf("%04d", num+1)
	}
	return manageNo, nil
}

// ExistedInternationalNo checks if an international_no already exists for a tm_ref.
func (t *TmRightsSystemRepositoryPGSQL) ExistedInternationalNo(
	tx brantect_github.TransactionInterface,
	tmRef int,
	internationalNo string,
) (bool, error) {
	row := tx.Get().(*sql.Tx).QueryRow(sqlExistedInternationalNo, tmRef, internationalNo)
	var existed bool
	if err := row.Scan(&existed); err != nil {
		log.Errorf("ExistedInternationalNo error: %v", err)
		return false, errors.New(2030010004025, err.Error(), nil)
	}
	return existed, nil
}

// GetNextInternationalNoId retrieves the next id for international_no of a tm_ref.
func (t *TmRightsSystemRepositoryPGSQL) GetNextInternationalNoId(
	tx brantect_github.TransactionInterface,
	tmRef int,
) (int, error) {
	row := tx.Get().(*sql.Tx).QueryRow(sqlGetNextInternationalNoId, tmRef)
	var id int
	if err := row.Scan(&id); err != nil {
		log.Errorf("GetNextInternationalNoId error: %v", err)
		return 0, errors.New(2030010004026, err.Error(), nil)
	}
	return id, nil
}

// InsertInternationalNo inserts a new international_no for a given tm_ref.
func (t *TmRightsSystemRepositoryPGSQL) InsertInternationalNo(
	tx brantect_github.TransactionInterface,
	tmRef int,
	nextInternationalNoId int,
	internationalNo string,
	inpUser string,
) error {
	// Execute the SQL statement to insert a new international_no record.
	_, err := tx.Get().(*sql.Tx).Exec(sqlInsertInternationalNo, tmRef, nextInternationalNoId, internationalNo, inpUser)
	if err != nil {
		log.Errorf("InsertInternationalNo error: %v", err)
		return errors.New(2030010004027, err.Error(), nil)
	}
	return nil
}

// ExistedOurRef checks if a given our_ref already exists for the client.
func (t *TmRightsSystemRepositoryPGSQL) ExistedOurRef(
	tx brantect_github.TransactionInterface,
	clientCd string,
	ourRef string,
) (bool, error) {
	row := tx.Get().(*sql.Tx).QueryRow(sqlExistedOurRef, clientCd, ourRef)
	var existed bool
	if err := row.Scan(&existed); err != nil {
		log.Errorf("ExistedOurRef error: %v", err)
		return false, errors.New(2030010004028, err.Error(), nil)
	}
	return existed, nil
}

// CreateManageNoGenerally generates a general our_ref (matches PHP business logic, dynamic rule).
// Rule: fetch rule from tm_base_manage_no, build pattern, get last serial, generate next number.
func (t *TmRightsSystemRepositoryPGSQL) CreateManageNoGenerally(
	tx brantect_github.TransactionInterface,
	clientCd string,
	countryCd string,
) (string, error) {
	log.Infof("CreateManageNoGenerally called with clientCd: %s, countryCd: %s", clientCd, countryCd)
	// 1. Get rule from tm_base_manage_no using embedded SQL
	// Query configuration parameters for generating manage_no from tm_base_manage_no table based on clientCd
	row := tx.Get().(*sql.Tx).QueryRow(sqlGetManageNoRuleFullByClientCd, clientCd)
	var fixTextFieldNo, dateFieldNo, dateType, serialFieldNo, serialZeroPaddingFlg, serialDigitNumber, countryCdFieldNo, serialWithDateFlg, serialWithCountryFlg int
	var fixText string
	if err := row.Scan(&fixTextFieldNo, &fixText, &dateFieldNo, &dateType, &serialFieldNo, &serialZeroPaddingFlg, &serialDigitNumber, &countryCdFieldNo, &serialWithDateFlg, &serialWithCountryFlg); err != nil {
		log.Errorf("CreateManageNoGenerally error: %v", err)
		return "", errors.New(2030010004029, err.Error(), nil)
	}

	// 2. Initialize pattern and wherePattern arrays to build regex for finding the last serial number
	patterns := make([]string, 4)
	wherePatterns := make([]string, 4)
	digitOnly := true // This variable checks if the code consists of digits only

	// 3. Handle fixed text part if configured
	if fixTextFieldNo > 0 && fixTextFieldNo < 5 {
		patterns[fixTextFieldNo-1] = fixText
		wherePatterns[fixTextFieldNo-1] = fixText
		digitOnly = false
	}

	// 4. Handle date part if configured
	loc, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(loc)
	dateStr := ""
	switch dateType {
	case 0: // YYYYMM
		dateStr = now.Format("200601")
	case 1: // YYMM
		dateStr = now.Format("0601")
	case 2: // YYYY
		dateStr = now.Format("2006")
	case 3: // YY
		dateStr = now.Format("06")
	}
	if dateFieldNo > 0 && dateFieldNo < 5 && dateStr != "" {
		patterns[dateFieldNo-1] = dateStr
		wherePatterns[dateFieldNo-1] = dateStr
		digitOnly = false
		if serialWithDateFlg == 0 {
			patterns[dateFieldNo-1] = fmt.Sprintf("[0-9]{%d}", len(dateStr))
			wherePatterns[dateFieldNo-1] = fmt.Sprintf("[0-9]{%d}", len(dateStr))
		}
	}

	// 5. Handle serial part (incremental number)
	if serialFieldNo > 0 && serialFieldNo < 5 {
		if serialZeroPaddingFlg == 0 {
			// No zero padding
			patterns[serialFieldNo-1] = "([0-9]+)"
			wherePatterns[serialFieldNo-1] = "[0-9]+"
		} else {
			// Zero padding according to configured digit number
			patterns[serialFieldNo-1] = fmt.Sprintf("([0-9]{%d})", serialDigitNumber)
			wherePatterns[serialFieldNo-1] = fmt.Sprintf("[0-9]{%d}", serialDigitNumber)
			digitOnly = false
		}
	}

	// 6. Handle country code part if configured
	countryCdUpper := strings.ToUpper(countryCd)
	if countryCdFieldNo > 0 && countryCdFieldNo < 5 {
		if serialWithCountryFlg == 0 {
			// If not binding serial with country code, pattern is 2 alphabetic characters
			patterns[countryCdFieldNo-1] = "[A-Z]{2}"
			wherePatterns[countryCdFieldNo-1] = "[A-Z]{2}"
		} else {
			// If binding serial with country code, pattern is the specific country code
			patterns[countryCdFieldNo-1] = countryCdUpper
			wherePatterns[countryCdFieldNo-1] = countryCdUpper
		}
		digitOnly = false
	}

	// 7. Build regex pattern and wherePattern to find the last serial in DB
	pattern := "^"
	wherePattern := "^"
	for i := 0; i < len(patterns); i++ {
		pattern += patterns[i]
	}
	for i := 0; i < len(wherePatterns); i++ {
		wherePattern += wherePatterns[i]
	}
	pattern += "$"
	wherePattern += "$"

	var serial int
	// 8. Get the last serial from DB based on the built pattern
	if digitOnly {
		// If only digits, get the last serial as a number
		var lastRef sql.NullString
		row := tx.Get().(*sql.Tx).QueryRow(sqlGetLastOurRefDigitOnly, clientCd)
		if err := row.Scan(&lastRef); err != nil && err != sql.ErrNoRows {
			log.Errorf("CreateManageNoGenerally error: %v", err)
			return "", errors.New(2030010004030, err.Error(), nil)
		}
		if lastRef.Valid {
			if n, err := strconv.Atoi(lastRef.String); err == nil {
				serial = n
			}
		}
	} else {
		// If contains characters, get the last serial based on regex pattern
		var lastSerial sql.NullInt64
		row := tx.Get().(*sql.Tx).QueryRow(sqlGetLastOurRefPattern, pattern, wherePattern, clientCd)
		if err := row.Scan(&lastSerial); err != nil && err != sql.ErrNoRows {
			log.Errorf("CreateManageNoGenerally error: %v", err)
			return "", errors.New(2030010004031, err.Error(), nil)
		}
		if lastSerial.Valid {
			serial = int(lastSerial.Int64)
		}
	}

	// 9. Build each part of the manage_no according to the configured positions
	parts := make([]string, 4)
	if fixTextFieldNo > 0 && fixTextFieldNo < 5 {
		parts[fixTextFieldNo-1] = fixText
	}
	if dateFieldNo > 0 && dateFieldNo < 5 && dateStr != "" {
		parts[dateFieldNo-1] = dateStr
	}
	if serialFieldNo > 0 && serialFieldNo < 5 {
		if serialZeroPaddingFlg == 0 {
			// No zero padding
			if serial == 0 {
				parts[serialFieldNo-1] = "1"
			} else {
				parts[serialFieldNo-1] = strconv.Itoa(serial + 1)
			}
		} else {
			// Zero padding
			if serial == 0 {
				parts[serialFieldNo-1] = fmt.Sprintf("%0*d", serialDigitNumber, 1)
			} else {
				parts[serialFieldNo-1] = fmt.Sprintf("%0*d", serialDigitNumber, serial+1)
			}
		}
	}
	if countryCdFieldNo > 0 && countryCdFieldNo < 5 {
		parts[countryCdFieldNo-1] = countryCdUpper
	}

	// 10. Concatenate all parts to form the complete manage_no
	manageNo := ""
	for _, p := range parts {
		manageNo += p
	}

	return manageNo, nil
}

// ExistsByFixCdAndDivKey checks if a value exists in tm_dival_mst for a given fix_cd and div_key.
func (t *TmRightsSystemRepositoryPGSQL) ExistsByFixCdAndDivKey(
	tx brantect_github.TransactionInterface,
	fixCd, divKey string,
) (bool, error) {
	var exists bool
	err := tx.Get().(*sql.Tx).QueryRow(sqlExistsByFixCdAndDivKey, fixCd, divKey).Scan(&exists)
	if err != nil {
		log.Errorf("ExistsByFixCdAndDivKey error: %v (fixCd=%s, divKey=%s)", err, fixCd, divKey)
		return false, errors.New(2030010004032, "SQL error ExistsByFixCdAndDivKey", err)
	}
	return exists, nil
}

// ExistsByRegCdAndLangId checks if a registration code exists in tm_reg_mst for a given reg_cd.
func (t *TmRightsSystemRepositoryPGSQL) ExistsByRegCdAndLangId(
	tx brantect_github.TransactionInterface,
	regCd string,
) (bool, error) {
	stmt, err := tx.Get().(*sql.Tx).Prepare(sqlGetByRegCd)
	if err != nil {
		log.Errorf("ExistsByRegCdAndLangId error: %v (regCd=%s)", err, regCd)
		return false, errors.New(2030010004033, err.Error(), nil)
	}
	defer stmt.Close()
	row := stmt.QueryRow(regCd)
	var reg_cd, reg_nm string
	if err := row.Scan(&reg_cd, &reg_nm); err != nil {
		log.Errorf("ExistsByRegCdAndLangId error: %v (regCd=%s)", err, regCd)
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, errors.New(2030010004034, err.Error(), nil)
	}
	return true, nil
}

// ExistsByPrcCdAndRegCdAndLangId checks if a process code exists in tm_prc_mst for a given prc_cd and reg_cd.
func (r *TmRightsSystemRepositoryPGSQL) ExistsByPrcCdAndRegCdAndLangId(
	tx brantect_github.TransactionInterface,
	prcCd string,
	regCd string,
) (bool, error) {
	stmt, err := tx.Get().(*sql.Tx).Prepare(sqlGetByPrcCd)
	if err != nil {
		log.Errorf("ExistsByPrcCdAndRegCdAndLangId error: %v (prcCd=%s, regCd=%s)", err, prcCd, regCd)
		return false, errors.New(2030010004035, err.Error(), nil)
	}
	defer stmt.Close()
	row := stmt.QueryRow(regCd, prcCd)
	var prc_cd, prc_nm string
	if err := row.Scan(&prc_cd, &prc_nm); err != nil {
		log.Errorf("ExistsByPrcCdAndRegCdAndLangId error: %v (prcCd=%s, regCd=%s)", err, prcCd, regCd)
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, errors.New(2030010004036, err.Error(), nil)
	}
	return true, nil
}

// ExistsOrg checks if an organization exists in mst_org for a given client_cd, dept_cd, and org_flg.
func (t *TmRightsSystemRepositoryPGSQL) ExistsOrg(
	tx brantect_github.TransactionInterface,
	clientCd string,
	deptCd string,
	orgFlg string,
) (bool, error) {
	stmt, err := tx.Get().(*sql.Tx).Prepare(sqlExistsOrg)
	if err != nil {
		log.Errorf("ExistsOrg error: %v (clientCd=%s, deptCd=%s, orgFlg=%s)", err, clientCd, deptCd, orgFlg)
		return false, errors.New(2030010004037, err.Error(), nil)
	}
	defer stmt.Close()
	var exists bool
	err = stmt.QueryRow(clientCd, deptCd, orgFlg).Scan(&exists)
	if err != nil {
		log.Errorf("ExistsOrg error: %v (clientCd=%s, deptCd=%s, orgFlg=%s)", err, clientCd, deptCd, orgFlg)
		return false, errors.New(2030010004038, err.Error(), nil)
	}
	return exists, nil
}

// ExistsPerson checks if a person exists in mst_person for a given client_cd and person_cd.
func (t *TmRightsSystemRepositoryPGSQL) ExistsPerson(
	tx brantect_github.TransactionInterface,
	clientCd string,
	personCd string,
) (bool, error) {
	stmt, err := tx.Get().(*sql.Tx).Prepare(sqlExistsPerson)
	if err != nil {
		log.Errorf("ExistsPerson error: %v (clientCd=%s, personCd=%s)", err, clientCd, personCd)
		return false, errors.New(2030010004039, err.Error(), nil)
	}
	defer stmt.Close()
	var exists bool
	err = stmt.QueryRow(clientCd, personCd).Scan(&exists)
	if err != nil {
		log.Errorf("ExistsPerson error: %v (clientCd=%s, personCd=%s)", err, clientCd, personCd)
		return false, errors.New(2030010004040, err.Error(), nil)
	}
	return exists, nil
}

// ExistsBaseCstField checks if a custom field_id exists in tm_base_cst_field for a client and is not deleted.
func (t *TmRightsSystemRepositoryPGSQL) ExistsBaseCstField(
	tx brantect_github.TransactionInterface,
	clientCd string,
	fieldId int,
) (bool, error) {
	stmt, err := tx.Get().(*sql.Tx).Prepare(sqlExistsBaseCstField)
	if err != nil {
		log.Errorf("ExistsBaseCstField error: %v (clientCd=%s, fieldId=%d)", err, clientCd, fieldId)
		return false, errors.New(2030010004041, err.Error(), nil)
	}
	defer stmt.Close()
	var exists bool
	if err := stmt.QueryRow(clientCd, fieldId).Scan(&exists); err != nil {
		log.Errorf("ExistsBaseCstField error: %v (clientCd=%s, fieldId=%d)", err, clientCd, fieldId)
		return false, errors.New(2030010004042, err.Error(), nil)
	}
	return exists, nil
}

// UpdateTmBaseByClientCdAndSysRef cập nhật động các trường được truyền vào (field khác nil)
func (t *TmRightsSystemRepositoryPGSQL) UpdateTmBaseByClientCdAndSysRef(
	tx brantect_github.TransactionInterface,
	clientCd string, // VARCHAR(11)
	sysRef int, // INTEGER
	req *brantect.TmRightUpdateRequest,
) error {
	var (
		fields []string
		args   []interface{}
		idx    = 1
	)

	// Helper closure for adding fields
	addField := func(fieldName string, value interface{}) {
		fields = append(fields, fmt.Sprintf("%s=$%d", fieldName, idx))
		args = append(args, value)
		idx++
	}

	// Map struct fields to DB columns, đảm bảo đúng kiểu dữ liệu
	if req.OurRef != nil {
		addField("our_ref", sqlNullString(*req.OurRef)) // VARCHAR(30)
	}
	if req.TmRef != nil {
		if *req.TmRef == 0 {
			addField("tm_ref", nil) // INTEGER
		} else {
			addField("tm_ref", *req.TmRef) // INTEGER
		}
	}
	if req.AppType != nil {
		addField("app_type", sqlNullString(*req.AppType)) // CHAR(2)
	}
	if req.CountryCd != nil {
		addField("country_cd", sqlNullString(*req.CountryCd)) // CHAR(2)
	}
	if req.ClassType != nil {
		addField("class_type", sqlNullString(*req.ClassType)) // CHAR(2)
	}
	if req.TmClass != nil {
		addField("tm_class", sqlNullString(*req.TmClass)) // VARCHAR(10100)
	}
	if req.RegType != nil {
		addField("reg_type", sqlNullString(*req.RegType)) // CHAR(2)
	}
	if req.Registrant != nil {
		addField("registrant", sqlNullString(*req.Registrant)) // VARCHAR(1000)
	}
	if req.RegCountryCd != nil {
		addField("reg_country_cd", sqlNullString(*req.RegCountryCd)) // CHAR(2)
	}
	if req.RegAdd != nil {
		addField("reg_add", sqlNullString(*req.RegAdd)) // VARCHAR(1000)
	}
	if req.RegCd != nil {
		addField("reg_cd", sqlNullString(*req.RegCd)) // CHAR(4)
	}
	if req.PrcCd != nil {
		addField("prc_cd", sqlNullString(*req.PrcCd)) // CHAR(4)
	}
	if req.ChkDate != nil {
		addField("chk_date", sqlNullString(*req.ChkDate)) // CHAR(8)
	}
	if req.ChkFileNo != nil {
		addField("chk_fileno", sqlNullString(*req.ChkFileNo)) // VARCHAR(50)
	}
	if req.AppDate != nil {
		addField("app_date", sqlNullString(*req.AppDate)) // CHAR(8)
	}
	if req.AppNo != nil {
		addField("app_no", sqlNullString(*req.AppNo)) // VARCHAR(50)
	}
	if req.AppDateUpd != nil {
		addField("app_date_upd", sqlNullString(*req.AppDateUpd)) // CHAR(8)
	}
	if req.AppNoUpd != nil {
		addField("app_no_upd", sqlNullString(*req.AppNoUpd)) // VARCHAR(50)
	}
	if req.PubDate != nil {
		addField("pub_date", sqlNullString(*req.PubDate)) // CHAR(8)
	}
	if req.PubNo != nil {
		addField("pub_no", sqlNullString(*req.PubNo)) // VARCHAR(50)
	}
	if req.PubNoUpd != nil {
		addField("pub_no_upd", sqlNullString(*req.PubNoUpd)) // VARCHAR(50)
	}
	if req.PubDateUpd != nil {
		addField("pub_date_upd", sqlNullString(*req.PubDateUpd)) // CHAR(8)
	}
	if req.RegDate != nil {
		addField("reg_date", sqlNullString(*req.RegDate)) // CHAR(8)
	}
	if req.RegNo != nil {
		addField("reg_no", sqlNullString(*req.RegNo)) // VARCHAR(50)
	}
	if req.RegDateUpd != nil {
		addField("reg_date_upd", sqlNullString(*req.RegDateUpd)) // CHAR(8)
	}
	if req.RegNoUpd != nil {
		addField("reg_no_upd", sqlNullString(*req.RegNoUpd)) // VARCHAR(50)
	}
	if req.Expire != nil {
		addField("expire", sqlNullString(*req.Expire)) // CHAR(8)
	}
	if req.NextLegalProc != nil {
		addField("next_legal_proc", sqlNullString(*req.NextLegalProc)) // CHAR(2)
	}
	if req.NextLegalProcDate != nil {
		addField("next_legal_proc_date", sqlNullString(*req.NextLegalProcDate)) // VARCHAR(8)
	}
	if req.AgentJp != nil {
		addField("agent_jp", sqlNullString(*req.AgentJp)) // VARCHAR(12)
	}
	if req.AgentItn != nil {
		addField("agent_itn", sqlNullString(*req.AgentItn)) // VARCHAR(12)
	}
	if req.ReqDept != nil {
		addField("req_dept", sqlNullString(*req.ReqDept)) // CHAR(5)
	}
	if req.ReqPerson != nil {
		addField("req_person", sqlNullString(*req.ReqPerson)) // CHAR(5)
	}
	if req.ExpDept != nil {
		addField("exp_dept", sqlNullString(*req.ExpDept)) // CHAR(5)
	}
	if req.ExpPerson != nil {
		addField("exp_person", sqlNullString(*req.ExpPerson)) // CHAR(5)
	}
	if req.PrcDept != nil {
		addField("prc_dept", sqlNullString(*req.PrcDept)) // CHAR(5)
	}
	if req.PrcPerson != nil {
		addField("prc_person", sqlNullString(*req.PrcPerson)) // CHAR(5)
	}
	if req.PerFlg != nil {
		addField("per_flg", sqlNullString(*req.PerFlg)) // CHAR(1)
	}
	if req.Remarks != nil {
		addField("remarks", sqlNullString(*req.Remarks)) // TEXT
	}
	if req.IntRegNo != nil {
		addField("int_reg_no", sqlNullString(*req.IntRegNo)) // CHAR(10)
	}
	if req.IntRegDate != nil {
		addField("int_reg_date", sqlNullString(*req.IntRegDate)) // CHAR(8)
	}
	if req.IntRegSubDate != nil {
		addField("int_reg_sub_date", sqlNullString(*req.IntRegSubDate)) // CHAR(8)
	}
	if req.DomAttorneyClientCd != nil {
		addField("dom_attorney_client_cd", sqlNullString(*req.DomAttorneyClientCd)) // VARCHAR(11)
	}
	if req.IntAttorneyClientCd != nil {
		addField("int_attorney_client_cd", sqlNullString(*req.IntAttorneyClientCd)) // VARCHAR(11)
	}
	if req.UpdUser != nil {
		addField("upd_user", sqlNullString(*req.UpdUser)) // VARCHAR(50)
	}
	if req.RyomaItemNo != nil {
		addField("ryoma_item_no", sqlNullString(*req.RyomaItemNo)) // TEXT
	}
	if req.RyomaUserNo != nil {
		addField("ryoma_user_no", sqlNullString(*req.RyomaUserNo)) // TEXT
	}

	if len(fields) == 0 {
		log.Debugf("UpdateTmBaseByClientCdAndSysRef: No fields to update for clientCd=%v, sysRef=%v", clientCd, sysRef)
		return nil
	}

	// set upd_date time
	loc, _ := time.LoadLocation("Asia/Tokyo")
	addField("upd_date", time.Now().In(loc).Format("2006-01-02 15:04:05.000000")) // TIMESTAMP

	// Debug: log all fields and args
	log.Debugf("UpdateTmBaseByClientCdAndSysRef: fields=%v, args=%v, clientCd=%v (%T), sysRef=%v (%T)", fields, args, clientCd, clientCd, sysRef, sysRef)

	// WHERE: client_cd (VARCHAR) và sys_ref (INTEGER)
	args = append(args, clientCd) // Đảm bảo là string
	args = append(args, sysRef)   // Đảm bảo là int

	query := fmt.Sprintf(strings.ReplaceAll(sqlUpdateTmBaseByClientCdAndSysRef, "{fields}", strings.Join(fields, ", ")), idx, idx+1)
	log.Debugf("UpdateTmBaseByClientCdAndSysRef: query=%s", query)
	log.Debugf("UpdateTmBaseByClientCdAndSysRef: final query for debug: %s", substituteQueryArgs(query, args))

	stmt, err := tx.Get().(*sql.Tx).Prepare(query)
	if err != nil {
		log.Errorf("UpdateTmBaseByClientCdAndSysRef error: %v", err)
		return errors.New(2030010004043, err.Error(), nil)
	}
	defer stmt.Close()
	_, errExec := stmt.Exec(args...)
	if errExec != nil {
		log.Errorf("UpdateTmBaseByClientCdAndSysRef exec error: %v", errExec)
		return errors.New(2030010004044, errExec.Error(), nil)
	}
	return nil
}

// substituteQueryArgs is a helper for debugging: replaces $n in query with the actual value for easier copy-paste debug
func substituteQueryArgs(query string, args []interface{}) string {
	for i, arg := range args {
		placeholder := fmt.Sprintf("$%d", i+1)
		var val string
		switch v := arg.(type) {
		case string:
			val = "'" + v + "'"
		case int:
			val = fmt.Sprintf("%d", v)
		case int64:
			val = fmt.Sprintf("%d", v)
		case float64:
			val = fmt.Sprintf("%f", v)
		default:
			val = fmt.Sprintf("'%v'", v)
		}
		query = strings.ReplaceAll(query, placeholder, val)
	}
	return query
}

// UpdateBaseCstValue cập nhật custom field value
func (t *TmRightsSystemRepositoryPGSQL) UpdateBaseCstValue(
	tx brantect_github.TransactionInterface,
	field *brantect.TmBaseCstField,
	sysRef int,
	updUser *string,
) error {
	user := ""
	if updUser != nil {
		user = *updUser
	}
	stmt, err := tx.Get().(*sql.Tx).Prepare(sqlUpdateBaseCstValue)
	if err != nil {
		log.Errorf("UpdateBaseCstValue error: %v", err)
		return errors.New(2030010004045, err.Error(), nil)
	}
	defer stmt.Close()
	_, errExec := stmt.Exec(field.FieldValue, user, field.FieldId, sysRef)
	if errExec != nil {
		log.Errorf("UpdateBaseCstValue exec error: %v", errExec)
		return errors.New(2030010004046, errExec.Error(), nil)
	}
	return nil
}

// GetMaxDoneTaskDate returns the max date (yyyymmdd) from tm_task for a client_cd, sysRef, done_flg=1, delete_flg=0
func (t *TmRightsSystemRepositoryPGSQL) GetMaxDoneTaskDate(
	tx brantect_github.TransactionInterface,
	clientCd string,
	sysRef int,
) (string, error) {
	stmt, err := tx.Get().(*sql.Tx).Prepare(sqlGetMaxDoneTaskDate)
	if err != nil {
		log.Errorf("GetMaxDoneTaskDate error: %v (clientCd=%s, sysRef=%d)", err, clientCd, sysRef)
		return "", errors.New(2030010004047, err.Error(), nil)
	}
	defer stmt.Close()
	var maxDate sql.NullString
	err = stmt.QueryRow(clientCd, sysRef).Scan(&maxDate)
	if err != nil && err != sql.ErrNoRows {
		log.Errorf("GetMaxDoneTaskDate scan error: %v (clientCd=%s, sysRef=%d)", err, clientCd, sysRef)
		return "", errors.New(2030010004048, err.Error(), nil)
	}
	if maxDate.Valid {
		return maxDate.String, nil
	}
	return "", nil
}

// ExistsMstClient checks if a client exists in mst_client for a given client_cd.
func (t *TmRightsSystemRepositoryPGSQL) ExistsMstClient(
	tx brantect_github.TransactionInterface,
	clientCd string,
) (bool, error) {
	stmt, err := tx.Get().(*sql.Tx).Prepare(sqlExistsMstClient)
	if err != nil {
		log.Errorf("ExistsMstClient error: %v (clientCd=%s)", err, clientCd)
		return false, errors.New(2030010004049, err.Error(), nil)
	}
	defer stmt.Close()
	var exists bool
	err = stmt.QueryRow(clientCd).Scan(&exists)
	if err != nil {
		log.Errorf("ExistsMstClient error: %v (clientCd=%s)", err, clientCd)
		return false, errors.New(2030010004050, err.Error(), nil)
	}
	return exists, nil
}

// AttorneyAuthorizedToUpdateRights checks if a given client_cd is authorized to update rights for a specific tm_ref.
func (t *TmRightsSystemRepositoryPGSQL) AttorneyAuthorizedToUpdateRights(
	tx brantect_github.TransactionInterface,
	clientCd string,
	sysRef int,
	searchAttorney string,
) (bool, error) {
	log.Infof("AttorneyAuthorizedToUpdateRights called with clientCd: %s, sysRef: %d, searchAttorney: %s", clientCd, sysRef, searchAttorney)
	stmt, err := tx.Get().(*sql.Tx).Prepare(sqlAttorneyAuthorizedToUpdateRights)
	if err != nil {
		log.Errorf("AttorneyAuthorizedToUpdateRights error: %v (clientCd=%s, sysRef=%d, searchAttorney=%s)", err, clientCd, sysRef, searchAttorney)
		return false, errors.New(2030010004051, err.Error(), nil)
	}
	defer stmt.Close()
	var authorized bool
	err = stmt.QueryRow(clientCd, sysRef, searchAttorney).Scan(&authorized)
	if err != nil {
		log.Errorf("AttorneyAuthorizedToUpdateRights scan error: %v (clientCd=%s, sysRef=%d, searchAttorney=%s)", err, clientCd, sysRef, searchAttorney)
		return false, errors.New(2030010004052, err.Error(), nil)
	}
	return authorized, nil
}

func (t *TmRightsSystemRepositoryPGSQL) UpdateTmBaseSoftDelete(
	tx brantect_github.TransactionInterface,
	clientCd string,
	sysRef int,
	req *brantect.TmRightDeleteRequest,
) error {
	stmt, errPrepare := tx.Get().(*sql.Tx).Prepare(sqlUpdateTmBaseSoftDelete)
	if errPrepare != nil {
		log.Errorf("UpdateTmBaseSoftDelete error: %v", errPrepare)
		return errors.New(2030010004053, errPrepare.Error(), nil)
	}
	defer stmt.Close()

	_, errExec := stmt.Exec(req.DeleteReason, req.Remarks, req.UpdUser, clientCd, sysRef)
	if errExec != nil {
		log.Errorf("UpdateTmBaseSoftDelete error: %v", errExec)
		return errors.New(2030010004054, errExec.Error(), nil)
	}
	return nil
}
