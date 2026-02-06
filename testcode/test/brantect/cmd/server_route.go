package cmd

func (server *ApiServer) route() {

	// System API
	server.groupSystemAPI = server.echo.Group("/api/tm2")
	server.groupSystemAPI.Use(server.authHandler.Secret())

	server.groupSystemAPI.GET("/rights", server.tmRightsSystemHandler.Search())
	server.groupSystemAPI.POST("/rights", server.tmRightsSystemHandler.CreateRights())
	server.groupSystemAPI.GET("/rights/:client_cd/:sys_ref", server.tmRightsSystemHandler.GetByClientCdAndSysRef())
	server.groupSystemAPI.GET("/rights/:client_cd/:sys_ref/tasks", server.tmRightsSystemHandler.GetListOfTasksForRightsByClientCdAndSysRef())x
	server.groupSystemAPI.PUT("/rights/:client_cd/:sys_ref", server.tmRightsSystemHandler.UpdateRights())x
	server.groupSystemAPI.DELETE("/rights/:client_cd/:sys_ref", server.tmRightsSystemHandler.DeleteRights())x

	server.groupSystemAPI.POST("/tasks", server.tmTasksSystemHandler.CreateTmTask())x
	server.groupSystemAPI.PUT("/tasks/:client_cd/:tm_task_seq", server.tmTasksSystemHandler.UpdateTmTask())x
	server.groupSystemAPI.DELETE("/tasks/files/:client_cd/:tm_task_seq/:tm_at_seq/task_at", server.tmTasksSystemHandler.DeleteTmTaskAt())x
	server.groupSystemAPI.DELETE("/tasks/:client_cd/:tm_task_seq", server.tmTasksSystemHandler.DeleteTmTask()) x
	server.groupSystemAPI.GET("/tasks", server.tmTasksSystemHandler.Search()) x
	server.groupSystemAPI.GET("/tasks/:client_cd/:tm_task_seq", server.tmTasksSystemHandler.GetByClientCdAndTmTaskSeq())x
	server.groupSystemAPI.POST("/tasks/files/:client_cd/task_at", server.tmTasksSystemHandler.ProcessTmTaskAt())x

	server.groupSystemAPI.GET("/images", server.tmImagesSystemHandler.Search())
	server.groupSystemAPI.POST("/images", server.tmImagesSystemHandler.CreateImage())x
	server.groupSystemAPI.GET("/img_dsp/:client_cd/:file_name", server.tmImagesSystemHandler.DisplayImageByFileName())
	server.groupSystemAPI.GET("/images/:client_cd/:tm_ref/rights", server.tmImagesSystemHandler.GetListRighsOfImagesByClientCdAndTmRef())
	server.groupSystemAPI.GET("/images/:client_cd/:tm_ref", server.tmImagesSystemHandler.GetByTmRefAndClientCd())
	server.groupSystemAPI.PUT("/images/:client_cd/:tm_ref", server.tmImagesSystemHandler.UpdateByTmRefAndClientCd())x
	server.groupSystemAPI.DELETE("/images/:client_cd/:tm_ref", server.tmImagesSystemHandler.DeleteByTmRefAndClientCd())
	//================================== User API
	server.groupUserAPI = server.echo.Group("/uapi")
	server.groupUserAPI.Use(server.authHandler.VerifyJwt())

	server.groupUserAPI.GET("/tm2/images", server.tmImagesUserHandler.Search())
	server.groupUserAPI.GET("/tm2/images/:tm_ref", server.tmImagesUserHandler.GetByTmRefAndClientCd())
	server.groupUserAPI.GET("/tm2/img_dsp/:file_name", server.tmImagesUserHandler.DisplayImageByFileName())

	server.groupUserAPI.GET("/tm2/rights", server.tmRightsUserHandler.Search())
	server.groupUserAPI.GET("/tm2/rights/search", server.tmRightsUserHandler.SearchAdvanced())
	server.groupUserAPI.GET("/tm2/rights/:sys_ref", server.tmRightsUserHandler.GetByClientCdAndSysRef())
	server.groupUserAPI.GET("/tm2/images/:tm_ref/rights", server.tmRightsUserHandler.GetListRighsOfImagesByClientCdAndTmRef())

	// User API for get AttorneyDx info
	server.groupUserAPI.GET("/attorney_org/tm/dom_attorney_candidate", server.tmAttorneyOrgTmUserHandler.GetDomAttorneyCandidate())
	server.groupUserAPI.GET("/attorney_org/tm/int_attorney_candidate", server.tmAttorneyOrgTmUserHandler.GetIntAttorneyCandidate())
}
