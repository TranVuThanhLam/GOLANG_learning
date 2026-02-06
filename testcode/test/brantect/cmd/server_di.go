package cmd

import (
	"brantect-api-tm/handler"
	"brantect-api-tm/infrastructure"
	"brantect-api-tm/service"

	attorneydx_client "github.com/BrightsHoldings/attorneydx-go/client"
	brantect_client "github.com/BrightsHoldings/brantect-go/client"
	brantect_database "github.com/BrightsHoldings/brantect-go/database"
	brantect_handler "github.com/BrightsHoldings/brantect-go/handler"
	brantect_infra "github.com/BrightsHoldings/brantect-go/infrastructure"
	brostools_client "github.com/BrightsHoldings/brostools-go/client"
)

func (server *ApiServer) dependenciesInjection() {
	server.jwtApiClient = brantect_client.NewJwtApiClient(
		server.jwtApiUri,
		server.jwtApiSecret,
	)

	server.iamApiClient = brantect_client.NewIamApiClient(
		server.iamApiUri,
		server.iamApiSecret,
	)

	//Setting handler
	server.authHandler = brantect_handler.NewBrantectAuthenticationHandler(
		server.selfSecret,
		server.jwtApiClient,
		server.iamApiClient,
	)

	server.brantectDatabase = brantect_database.NewBrantectDatabase(
		brantect_database.DATABASE_DRIVER_PGSQL,
		server.pgsqlInstance,
		server.pgsqlPass,
		server.pgsqlUser,
		server.pgsqlDbname,
		server.runtimeEnv,
		server.apiRunType,
	)

	server.brantectRepository = brantect_infra.NewBrantectRepositoryPGSQL(server.brantectDatabase)

	server.tmRightsRepository = infrastructure.NewTmRightsRepositoryPGSQL(server.brantectDatabase)
	server.tmRightsService = service.NewTmRightsService(server.tmRightsRepository, server.brantectRepository)

	// audit log client
	server.auditLogClient = brantect_client.NewAuditLogSystemApiClient(server.auditLogUrl, server.auditLogSecret)

	server.tmRightsUserHandler = handler.NewTmRightsUserHandler(
		server.iamApiClient,
		server.authHandler,
		server.tmRightsService,
		server.auditLogClient,
	)

	server.mstClientIdRepository = infrastructure.NewMstClientIdRepositoryPGSQL(server.brantectDatabase)
	server.mstCountryRepository = brantect_infra.NewMstCountryRepositoryPGSQL(server.brantectDatabase)
	server.tmMstRepository = brantect_infra.NewTmMstRepositoryPGSQL(server.brantectDatabase)
	server.mstOrgRepository = brantect_infra.NewMstOrgRepositoryPGSQL(server.brantectDatabase)
	server.mstPersonRepository = brantect_infra.NewMstPersonRepositoryPGSQL(server.brantectDatabase)
	server.tmTaskAtRepository = brantect_infra.NewTmTaskAtRepositoryPGSQL(server.brantectDatabase)
//////////////////////////////////////////////////////////////
	// tm rights system api
	server.tmRightsSystemRepository = infrastructure.NewTmRightsSystemRepositoryPGSQL(server.brantectDatabase)
	server.tmTasksSystemRepository = infrastructure.NewTmTasksSystemRepositoryPGSQL(server.brantectDatabase)
	server.tmRightsSystemService = service.NewTmRightsSystemService(
		server.tmRightsSystemRepository,
		server.brantectRepository,
		server.tmTasksSystemRepository,
		server.mstClientIdRepository,
		server.mstCountryRepository,
		server.tmMstRepository,
		server.mstOrgRepository,
		server.mstPersonRepository,
	)

	server.tmRightsSystemHandler = handler.NewTmRightsSystemHandler(
		server.tmRightsSystemService,
	)
	// end tm rights system api
//////////////////////////////////////////////////////////////
	// tm tasks system api
	server.tmTasksSystemService = service.NewTmTasksSystemService(
		server.tmTasksSystemRepository,
		server.brantectRepository,
		server.tmRightsSystemRepository,
		server.tmMstRepository,
		server.mstOrgRepository,
		server.mstPersonRepository,
		server.mstCountryRepository,
		server.tmTaskAtRepository,
		server.brantectFileApiUri,
		server.brantectFileApiSecret,
	)

	server.tmTasksSystemHandler = handler.NewTmTasksSystemHandler(
		server.tmTasksSystemService,
	)
	// end tm tasks system api

	server.tmImagesRepository = infrastructure.NewTmImagesRepositoryPGSQL(server.brantectDatabase)
	server.tmImagesService = service.NewTmImagesService(server.tmImagesRepository, server.brantectRepository)

	server.clientFile = brantect_client.NewBrantectFileApi(
		server.brantectFileApiUri,
		server.brantectFileApiSecret,
	)
	server.tmImagesUserHandler = handler.NewTmImagesUserHandler(
		server.iamApiClient,
		server.authHandler,
		server.tmImagesService,
		server.clientFile,
		server.auditLogClient,
	)

	server.tmImagesSystemRepository = infrastructure.NewTmImagesSystemRepositoryPGSQL(server.brantectDatabase)
	server.tmImagesSystemService = service.NewTmImagesSystemService(server.tmImagesSystemRepository, server.brantectRepository)
	server.tmImagesSystemHandler = handler.NewTmImagesSystemHandler(
		server.tmImagesSystemService,
		server.clientFile,
	)

	server.tmAttorneyOrgTmUserClient = attorneydx_client.NewAttorneyOrgTmApi(
		server.attorneydxClientApiUri,
		server.attorneydxClientApiSecret,
	)

	server.brostoolsEapiClient = brostools_client.NewClientEapiClient(
		server.brostoolsClientApiUri,
		server.brostoolsClientApiSecret,
	)

	server.tmAttorneyOrgTmUserService = service.NewTmAttorneyOrgTmUserService(
		server.tmAttorneyOrgTmUserClient,
		server.brostoolsEapiClient,
	)

	server.tmAttorneyOrgTmUserHandler = handler.NewTmAttorneyOrgTmUserHandler(
		server.authHandler,
		server.tmAttorneyOrgTmUserService,
	)

	server.tmAreaSystemRepository = infrastructure.NewTmAreaSystemRepositoryPGSQL(server.brantectDatabase)
	server.tmAreaSystemService = service.NewTmAreaSystemService(server.tmAreaSystemRepository, server.brantectRepository)
	server.tmAreaSystemHandler = handler.NewTmAreaSystemHandler(
		server.tmAreaSystemService,
	)

	server.tmClassSystemRepository = infrastructure.NewTmClassSystemRepositoryPGSQL(server.brantectDatabase)
	server.tmClassSystemService = service.NewTmClassSystemService(server.tmClassSystemRepository, server.brantectRepository)
	server.tmClassSystemHandler = handler.NewTmClassSystemHandler(server.tmClassSystemService)

	server.tmCountrySystemRepository = infrastructure.NewTmCountrySystemRepositoryPGSQL(server.brantectDatabase)
	server.tmCountrySystemService = service.NewTmCountrySystemService(server.tmCountrySystemRepository, server.brantectRepository)
	server.tmCountrySystemHandler = handler.NewTmCountrySystemHandler(
		server.tmCountrySystemService,
	)

	server.tmDesignedOrgSystemRepository = infrastructure.NewTmDesignedOrgSystemRepositoryPGSQL(server.brantectDatabase)
	server.tmDesignedOrgSystemService = service.NewTmDesignedOrgSystemService(server.tmDesignedOrgSystemRepository, server.brantectRepository)
	server.tmDesignedOrgSystemHandler = handler.NewTmDesignedOrgSystemHandler(
		server.tmDesignedOrgSystemService,
	)

	server.tmUseStocksystemRepository = infrastructure.NewTmUseStockSystemRepositoryPGSQL(server.brantectDatabase)
	server.tmUseStockSystemService = service.NewTmUseStockSystemService(server.tmUseStocksystemRepository, server.brantectRepository)
	server.tmUseStockSystemHandler = handler.NewTmUseStockSystemHandler(server.tmUseStockSystemService)

	server.tmRenewalPolicySystemRepository = infrastructure.NewTmRenewalPolicySystemRepositoryPGSQL(server.brantectDatabase)
	server.tmRenewalPolicySystemService = service.NewTmRenewalPolicySystemService(server.tmRenewalPolicySystemRepository, server.brantectRepository)
	server.tmRenewalPolicySystemHandler = handler.NewTmRenewalPolicySystemHandler(server.tmRenewalPolicySystemService)

	server.tmKeywordSystemRepository = infrastructure.NewTmKeywordSystemRepositoryPGSQL(server.brantectDatabase)
	server.tmKeywordSystemService = service.NewTmKeywordSystemService(server.tmKeywordSystemRepository, server.brantectRepository)
	server.tmKeywordSystemHandler = handler.NewTmKeywordSystemHandler(server.tmKeywordSystemService)

	server.tmRegistrantSystemRepository = infrastructure.NewTmRegistrantSystemRepositoryPGSQL(server.brantectDatabase)
	server.tmRegistrantSystemService = service.NewTmRegistrantSystemService(server.tmRegistrantSystemRepository, server.brantectRepository)
	server.tmRegistrantSystemHandler = handler.NewTmRegistrantSystemHandler(server.tmRegistrantSystemService)

	server.tmPersonSystemRepository = infrastructure.NewTmPersonSystemRepositoryPGSQL(server.brantectDatabase)
	server.tmPersonSystemService = service.NewTmPersonSystemService(server.tmPersonSystemRepository, server.brantectRepository)
	server.tmPersonSystemHandler = handler.NewTmPersonSystemHandler(
		server.tmPersonSystemService,
	)

	server.tmRightsHolderAddressSystemRepository = infrastructure.NewTmRightsHolderAddressSystemRepositoryPGSQL(server.brantectDatabase)
	server.tmRightsHolderAddressSystemService = service.NewTmRightsHolderAddressSystemService(server.tmRightsHolderAddressSystemRepository, server.brantectRepository)
	server.tmRightsHolderAddressSystemHandler = handler.NewTmRightsHolderAddressSystemHandler(server.tmRightsHolderAddressSystemService)

	server.tmOwnershipTypeSystemRepository = infrastructure.NewTmOwnershipTypeSystemRepositoryPGSQL(server.brantectDatabase)
	server.tmOwnershipTypeSystemService = service.NewTmOwnershipTypeSystemService(server.tmOwnershipTypeSystemRepository, server.brantectRepository)
	server.tmOwnershipTypeSystemHandler = handler.NewTmOwnershipTypeSystemHandler(server.tmOwnershipTypeSystemService)

	server.tmClassCategorySystemRepository = infrastructure.NewTmClassCategorySystemRepositoryPGSQL(server.brantectDatabase)
	server.tmClassCategorySystemService = service.NewTmClassCategorySystemService(server.tmClassCategorySystemRepository, server.brantectRepository)
	server.tmClassCategorySystemHandler = handler.NewTmClassCategorySystemHandler(server.tmClassCategorySystemService)

	server.tmRegStatusSystemRepository = infrastructure.NewTmRegStatusSystemRepositoryPGSQL(server.brantectDatabase)
	server.tmRegStatusSystemService = service.NewTmRegStatusSystemService(server.tmRegStatusSystemRepository, server.brantectRepository)
	server.tmRegStatusSystemHandler = handler.NewTmRegStatusSystemHandler(server.tmRegStatusSystemService)

	server.tmCategorySystemRepository = infrastructure.NewTmCategorySystemRepositoryPGSQL(server.brantectDatabase)
	server.tmCategorySystemService = service.NewTmCategorySystemService(server.tmCategorySystemRepository, server.brantectRepository)
	server.tmCategorySystemHandler = handler.NewTmCategorySystemHandler(
		server.tmCategorySystemService,
	)

	server.tmEmbodimentSystemRepository = infrastructure.NewTmEmbodimentSystemRepositoryPGSQL(server.brantectDatabase)
	server.tmEmbodimentSystemService = service.NewTmEmbodimentSystemService(server.tmEmbodimentSystemRepository, server.brantectRepository)
	server.tmEmbodimentSystemHandler = handler.NewTmEmbodimentSystemHandler(
		server.tmEmbodimentSystemService,
	)

	server.tmApplicationCategorySystemRepository = infrastructure.NewTmApplicationCategoryRepositoryPGSQL(server.brantectDatabase)
	server.tmApplicationCategorySystemService = service.NewTmApplicationCategorySystemService(server.tmApplicationCategorySystemRepository, server.brantectRepository)
	server.tmApplicationCategorySystemHandler = handler.NewTmApplicationCategorySystemHandler(
		server.tmApplicationCategorySystemService,
	)

	server.tmLegalProcedureSystemRepository = infrastructure.NewTmLegalProcedureRepositoryPGSQL(server.brantectDatabase)
	server.tmLegalProcedureSystemService = service.NewTmLegalProcedureSystemService(server.tmLegalProcedureSystemRepository, server.brantectRepository)
	server.tmLegalProcedureSystemHandler = handler.NewTmLegalProcedureSystemHandler(server.tmLegalProcedureSystemService)

	server.tmMstClientIdSystemRepository = infrastructure.NewTmMstClientIdSystemRepositoryPGSQL(server.brantectDatabase)
	server.tmMstClientIdSystemService = service.NewTmMstClientIdSystemService(server.tmMstClientIdSystemRepository, server.brantectRepository)
	server.tmMstClientIdSystemHandler = handler.NewTmMstClientIdSystemHandler(server.tmMstClientIdSystemService)

	server.tmActionSystemRepository = infrastructure.NewTmActionSystemRepositoryPGSQL(server.brantectDatabase)
	server.tmActionSystemService = service.NewTmActionSystemService(server.tmActionSystemRepository, server.brantectRepository)
	server.tmActionSystemHandler = handler.NewTmActionSystemHandler(server.tmActionSystemService)

	server.tmProcessStatusSystemRepository = infrastructure.NewTmProcessStatusSystemRepositoryPGSQL(server.brantectDatabase)
	server.tmProcessStatusSystemService = service.NewTmProcessStatusSystemService(server.tmProcessStatusSystemRepository, server.brantectRepository)
	server.tmProcessStatusSystemHandler = handler.NewTmProcessStatusSystemHandler(server.tmProcessStatusSystemService)

	server.tmGroupSystemRepository = infrastructure.NewTmGroupSystemRepositoryPGSQL(server.brantectDatabase)
	server.tmGroupSystemService = service.NewTmGroupSystemService(server.tmGroupSystemRepository, server.brantectRepository)
	server.tmGroupSystemHandler = handler.NewTmGroupSystemHandler(server.tmGroupSystemService)

	server.tmCurrencySystemRepository = infrastructure.NewTmCurrencySystemRepositoryPGSQL(server.brantectDatabase)
	server.tmCurrencySystemService = service.NewTmCurrencySystemService(server.tmCurrencySystemRepository, server.brantectRepository)
	server.tmCurrencySystemHandler = handler.NewTmCurrencySystemHandler(server.tmCurrencySystemService)

	server.tmMstManageNoRepository = infrastructure.NewTmMstManageNoRepositoryPGSQL(server.brantectDatabase)
	server.tmMstManageNoService = service.NewTmMstManageNoSystemService(server.tmMstManageNoRepository, server.brantectRepository)
	server.tmMstManageNoHandler = handler.NewTmMstManageNoSystemHandler(server.tmMstManageNoService)
}
