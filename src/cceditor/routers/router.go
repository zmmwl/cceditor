package routers

import (
	"cceditor/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Router("/cc", &controllers.Cccontroller{}, "*:Index")
	beego.Router("/", &controllers.WideHandler{}, "*:IndexHandler")


	beego.Router("/start", &controllers.WideHandler{}, "*:StartHandler")
	beego.Router("/session/ws", &controllers.WideHandler{}, "*:WSHandler")
	beego.Router("/session/save", &controllers.WideHandler{}, "*:SaveContentHandler")
	beego.Router("/build", &controllers.WideHandler{}, "*:BuildHandler")
	beego.Router("/run", &controllers.WideHandler{}, "*:RunHandler")
	beego.Router("/stop", &controllers.WideHandler{}, "*:StopHandler")
	beego.Router("/go/test", &controllers.WideHandler{}, "*:GoTestHandler")
	beego.Router("/go/vet", &controllers.WideHandler{}, "*:GoVetHandler")
	beego.Router("/go/get", &controllers.WideHandler{}, "*:GoGetHandler")
	beego.Router("/go/install", &controllers.WideHandler{}, "*:GoInstallHandler")
	beego.Router("/output/ws", &controllers.WideHandler{}, "*:WSHandlerOutput")
	beego.Router("/cross", &controllers.WideHandler{}, "*:CrossCompilationHandler")
	beego.Router("/files", &controllers.WideHandler{}, "*:GetFilesHandler")
	beego.Router("/file/refresh", &controllers.WideHandler{}, "*:RefreshDirectoryHandler")
	beego.Router("/file", &controllers.WideHandler{}, "*:GetFileHandler")
	beego.Router("/file/save", &controllers.WideHandler{}, "*:SaveFileHandler")
	beego.Router("/file/new", &controllers.WideHandler{}, "*:NewFileHandler")
	beego.Router("/file/remove", &controllers.WideHandler{}, "*:RemoveFileHandler")
	beego.Router("/file/rename", &controllers.WideHandler{}, "*:RenameFileHandler")
	beego.Router("/file/search/text", &controllers.WideHandler{}, "*:SearchTextHandler")
	beego.Router("/file/find/name", &controllers.WideHandler{}, "*:FindHandler")
	beego.Router("/outline", &controllers.WideHandler{}, "*:GetOutlineHandler")
	beego.Router("/file/zip/new", &controllers.WideHandler{}, "*:CreateZipHandler")
	beego.Router("/file/zip", &controllers.WideHandler{}, "*:GetZipHandler")
	beego.Router("/file/upload", &controllers.WideHandler{}, "*:UploadHandler")
	beego.Router("/file/decompress", &controllers.WideHandler{}, "*:DecompressHandler")
	beego.Router("/editor/ws", &controllers.WideHandler{}, "*:WSHandlerEditor")
	beego.Router("/go/fmt", &controllers.WideHandler{}, "*:GoFmtHandler")
	beego.Router("/autocomplete", &controllers.WideHandler{}, "*:AutocompleteHandler")
	beego.Router("/exprinfo", &controllers.WideHandler{}, "*:GetExprInfoHandler")
	beego.Router("/find/decl", &controllers.WideHandler{}, "*:FindDeclarationHandler")
	beego.Router("/find/usages", &controllers.WideHandler{}, "*:FindUsagesHandler")
	beego.Router("/notification/ws", &controllers.WideHandler{}, "*:WSHandlerNotify")
	beego.Router("/login", &controllers.WideHandler{}, "*:LoginHandler")
	beego.Router("/logout", &controllers.WideHandler{}, "*:LogoutHandler")

}
