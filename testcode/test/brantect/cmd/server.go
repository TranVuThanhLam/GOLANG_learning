package cmd

import (
	brantect_local "brantect-api-tm/brantect"
	"brantect-api-tm/interfaces"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/BrightsHoldings/attorneydx-go/attorneydx"
	"github.com/BrightsHoldings/brantect-go/brantect"
	"github.com/BrightsHoldings/brostools-go/brostools"
	"github.com/BrightsHoldings/gmobrs-go-lib/log"
	"github.com/labstack/echo/v4"
)

const API_SERVER_DEFAULT_PORT string = "8080"

type ApiServer struct {
	authHandler brantect.BrantectAuthenticationHandlerInterface

	iamApiClient       brantect.IamApiClientInterface
	jwtApiClient       brantect.JwtApiClientInterface
	brantectRepository brantect.RepositoryInterface
	brantectDatabase   brantect.BrantectDatabaseInterface
	auditLogClient     brantect.AuditLogSystemApiClientInterface

	tmRightsUserHandler interfaces.TmRightsHandlerInterface
	tmRightsService     interfaces.TmRightsServiceInterface
	tmRightsRepository  brantect_local.TmRightsRepositoryInterface

	tmImagesUserHandler interfaces.TmImagesHandlerInterface
	tmImagesService     interfaces.TmImagesServiceInterface
	tmImagesRepository  brantect_local.TmImagesRepositoryInterface

	tmRightsSystemHandler    interfaces.TmRightsSystemHandlerInterface
	tmRightsSystemService    interfaces.TmRightsSystemServiceInterface
	tmRightsSystemRepository brantect_local.TmRightsSystemRepositoryInterface

	tmImagesSystemHandler    interfaces.TmImagesSystemHandlerInterface
	tmImagesSystemService    interfaces.TmImagesSystemServiceInterface
	tmImagesSystemRepository brantect_local.TmImagesSystemRepositoryInterface

	tmTasksSystemHandler    interfaces.TmTasksSystemHandlerInterface
	tmTasksSystemService    interfaces.TmTasksSystemServiceInterface
	tmTasksSystemRepository brantect_local.TmTasksSystemRepositoryInterface

	mstClientIdRepository brantect_local.MstClientIdRepositoryInterface
	mstCountryRepository  brantect.MstCountryRepositoryInterface
	tmMstRepository       brantect.TmMstRepositoryInterface
	mstOrgRepository      brantect.MstOrgRepositoryInterface
	mstPersonRepository   brantect.MstPersonRepositoryInterface
	tmTaskAtRepository    brantect.TmTaskAtRepositoryInterface

	tmAttorneyOrgTmUserHandler interfaces.TmAttorneyOrgTmHandlerInterface
	tmAttorneyOrgTmUserService interfaces.TmAttorneyOrgTmServiceInterface
	tmAttorneyOrgTmUserClient  attorneydx.ApiAttorneyOrgTmClientInterface
	brostoolsEapiClient        brostools.ClientEapiClientInterface

	attorneydxClientApiUri    string
	attorneydxClientApiSecret string
	brostoolsClientApiUri     string
	brostoolsClientApiSecret  string

	tmGroupSystemHandler    interfaces.TmGroupSystemHandlerInterface
	tmGroupSystemService    interfaces.TmGroupSystemServiceInterface
	tmGroupSystemRepository brantect_local.TmGroupSystemRepositoryInterface

	tmAreaSystemHandler    interfaces.TmAreaSystemHandlerInterface
	tmAreaSystemService    interfaces.TmAreaSystemServiceInterface
	tmAreaSystemRepository brantect_local.TmAreaSystemRepositoryInterface

	tmMstCstValueRepository brantect_local.TmMstCstValueRepositoryInterface

	tmClassSystemHandler    interfaces.TmClassSystemHandlerInterface
	tmClassSystemService    interfaces.TmClassSystemServiceInterface
	tmClassSystemRepository brantect_local.TmClassSystemRepositoryInterface

	tmCountrySystemHandler    interfaces.TmCountrySystemHandlerInterface
	tmCountrySystemService    interfaces.TmCountrySystemServiceInterface
	tmCountrySystemRepository brantect_local.TmCountrySystemRepositoryInterface

	tmDesignedOrgSystemHandler    interfaces.TmDesignedOrgSystemHandlerInterface
	tmDesignedOrgSystemService    interfaces.TmDesignedOrgSystemServiceInterface
	tmDesignedOrgSystemRepository brantect_local.TmDesignedOrgSystemRepositoryInterface

	tmUseStockSystemHandler    interfaces.TmUseStockSystemHandlerInterface
	tmUseStockSystemService    interfaces.TmUseStockSystemServiceInterface
	tmUseStocksystemRepository brantect_local.TmUseStockSystemRepositoryInterface

	tmPersonSystemHandler    interfaces.TmPersonSystemHandlerInterface
	tmPersonSystemService    interfaces.TmPersonSystemServiceInterface
	tmPersonSystemRepository brantect_local.TmPersonSystemRepositoryInterface

	tmRenewalPolicySystemHandler    interfaces.TmRenewalPolicySystemHandlerInterface
	tmRenewalPolicySystemService    interfaces.TmRenewalPolicySystemServiceInterface
	tmRenewalPolicySystemRepository brantect_local.TmRenewalPolicySystemRepositoryInterface

	tmApplicationCategorySystemHandler    interfaces.TmApplicationCategorySystemHandlerInterface
	tmApplicationCategorySystemService    interfaces.TmApplicationCategorySystemServiceInterface
	tmApplicationCategorySystemRepository brantect_local.TmApplicationCategorySystemRepositoryInterface

	tmKeywordSystemHandler    interfaces.TmKeywordSystemHandlerInterface
	tmKeywordSystemService    interfaces.TmKeywordSystemServiceInterface
	tmKeywordSystemRepository brantect_local.TmKeywordSystemRepositoryInterface

	tmRegistrantSystemHandler    interfaces.TmRegistrantSystemHandlerInterface
	tmRegistrantSystemService    interfaces.TmRegistrantSystemServiceInterface
	tmRegistrantSystemRepository brantect_local.TmRegistrantSystemRepositoryInterface

	tmRightsHolderAddressSystemHandler    interfaces.TmRightsHolderAddressSystemHandlerInterface
	tmRightsHolderAddressSystemService    interfaces.TmRightsHolderAddressSystemServiceInterface
	tmRightsHolderAddressSystemRepository brantect_local.TmRightsHolderAddressSystemRepositoryInterface

	tmOwnershipTypeSystemHandler    interfaces.TmOwnershipTypeSystemHandlerInterface
	tmOwnershipTypeSystemService    interfaces.TmOwnershipTypeSystemServiceInterface
	tmOwnershipTypeSystemRepository brantect_local.TmOwnershipTypeSystemRepositoryInterface

	tmClassCategorySystemHandler    interfaces.TmClassCategorySystemHandlerInterface
	tmClassCategorySystemService    interfaces.TmClassCategorySystemServiceInterface
	tmClassCategorySystemRepository brantect_local.TmClassCategorySystemRepositoryInterface

	tmRegStatusSystemHandler    interfaces.TmRegStatusSystemHandlerInterface
	tmRegStatusSystemService    interfaces.TmRegStatusSystemServiceInterface
	tmRegStatusSystemRepository brantect_local.TmRegStatusSystemRepositoryInterface

	tmCategorySystemHandler    interfaces.TmCategorySystemHandlerInterface
	tmCategorySystemService    interfaces.TmCategorySystemServiceInterface
	tmCategorySystemRepository brantect_local.TmCategorySystemRepositoryInterface

	tmEmbodimentSystemHandler    interfaces.TmEmbodimentSystemHandlerInterface
	tmEmbodimentSystemService    interfaces.TmEmbodimentSystemServiceInterface
	tmEmbodimentSystemRepository brantect_local.TmEmbodimentSystemRepositoryInterface

	tmLegalProcedureSystemHandler    interfaces.TmLegalProcedureSystemHandlerInterface
	tmLegalProcedureSystemService    interfaces.TmLegalProcedureSystemServiceInterface
	tmLegalProcedureSystemRepository brantect_local.TmLegalProcedureSystemRepositoryInterface

	tmMstClientIdSystemHandler    interfaces.TmMstClientIdSystemHandlerInterface
	tmMstClientIdSystemService    interfaces.TmMstClientIdSystemServiceInterface
	tmMstClientIdSystemRepository brantect_local.TmMstClientIdSystemRepositoryInterface

	tmActionSystemHandler    interfaces.TmActionSystemHandlerInterface
	tmActionSystemService    interfaces.TmActionSystemServiceInterface
	tmActionSystemRepository brantect_local.TmActionSystemRepositoryInterface

	tmProcessStatusSystemHandler    interfaces.TmProcessStatusSystemHandlerInterface
	tmProcessStatusSystemService    interfaces.TmProcessStatusSystemServiceInterface
	tmProcessStatusSystemRepository brantect_local.TmProcessStatusSystemRepositoryInterface

	tmCurrencySystemHandler    interfaces.TmCurrencySystemHandlerInterface
	tmCurrencySystemService    interfaces.TmCurrencySystemServiceInterface
	tmCurrencySystemRepository brantect_local.TmCurrencySystemRepositoryInterface

	tmMstManageNoRepository brantect_local.TmMstManageNoRepositoryInterface
	tmMstManageNoHandler    interfaces.TmMstManageNoSystemHandlerInterface
	tmMstManageNoService    interfaces.TmMstManageNoSystemServiceInterface

	iamApiUri    string
	iamApiSecret string

	brantectFileApiUri    string
	brantectFileApiSecret string
	clientFile            brantect.ClientFileApiInterface

	pgsqlInstance string
	pgsqlPass     string
	pgsqlUser     string
	pgsqlDbname   string

	echo           *echo.Echo
	groupSystemAPI *echo.Group
	groupUserAPI   *echo.Group
	groupSwagger   *echo.Group

	jwtApiUri      string
	jwtApiSecret   string
	selfSecret     string
	runtimeEnv     string
	apiRunType     string
	logLevel       string
	auditLogUrl    string
	auditLogSecret string
}

func (server *ApiServer) Run() {

	port := os.Getenv("PORT")
	if port == "" {
		port = API_SERVER_DEFAULT_PORT
	}

	server.Start(port)

	//wait for signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	log.Infof("signal waiting:")
	s := <-sigCh
	log.Infof("signal received:%s", s.String())
	server.Teardown() //TODO:DockerComposeでSIGNALが受信できないため、正しく検証ができない

}

func (server *ApiServer) Start(port string) {

	_, errLoadEnv := server.loadEnv()
	if len(errLoadEnv) != 0 {
		for _, e := range errLoadEnv {
			log.Errorf(e)
		}
		return
	}

	server.echo = echo.New()
	server.setMiddleware()
	server.dependenciesInjection()
	server.route()

	go func() { //go routine
		if err := server.echo.Start(":" + port); err != nil && err != http.ErrServerClosed {
			fmt.Println(err)
			server.echo.Logger.Fatal("shutting down the server")
		}
	}()
	log.Infof("server running with go routine")

}

func (server *ApiServer) Teardown() {

	log.Infof("server graceful shutdown")
}
