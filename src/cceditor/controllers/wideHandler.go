package controllers

import (
	"github.com/astaxie/beego"
	"github.com/b3log/wide"
	"net/http"
	"github.com/b3log/wide/session"
	"github.com/b3log/wide/output"
	"github.com/b3log/wide/file"
	"github.com/b3log/wide/editor"
	"github.com/b3log/wide/notification"
)

type WideHandler struct {
	beego.Controller
}

var wideHandlers map[string]func(w http.ResponseWriter, r *http.Request)

func init(){
	wideHandlers = make(map[string]func(w http.ResponseWriter, r *http.Request))

	wideHandlers["IndexHandler"] = wide.HandlerGzWrapper(wide.IndexHandler)
	wideHandlers["LoginHandler"] = wide.HandlerGzWrapper(session.LoginHandler)


	wideHandlers["StartHandler"] = wide.HandlerGzWrapper(wide.StartHandler)
	wideHandlers["WSHandler"] = wide.HandlerGzWrapper(session.WSHandler)
	wideHandlers["SaveContentHandler"] = wide.HandlerGzWrapper(session.SaveContentHandler)
	wideHandlers["BuildHandler"] = wide.HandlerGzWrapper(output.BuildHandler)
	wideHandlers["RunHandler"] = wide.HandlerGzWrapper(output.RunHandler)
	wideHandlers["StopHandler"] = wide.HandlerGzWrapper(output.StopHandler)
	wideHandlers["GoTestHandler"] = wide.HandlerGzWrapper(output.GoTestHandler)
	wideHandlers["GoVetHandler"] = wide.HandlerGzWrapper(output.GoVetHandler)
	wideHandlers["GoGetHandler"] = wide.HandlerGzWrapper(output.GoGetHandler)
	wideHandlers["GoInstallHandler"] = wide.HandlerGzWrapper(output.GoInstallHandler)
	wideHandlers["WSHandlerOutput"] = wide.HandlerGzWrapper(output.WSHandler)
	wideHandlers["CrossCompilationHandler"] = wide.HandlerGzWrapper(output.CrossCompilationHandler)
	wideHandlers["GetFilesHandler"] = wide.HandlerGzWrapper(file.GetFilesHandler)
	wideHandlers["RefreshDirectoryHandler"] = wide.HandlerGzWrapper(file.RefreshDirectoryHandler)
	wideHandlers["GetFileHandler"] = wide.HandlerGzWrapper(file.GetFileHandler)
	wideHandlers["SaveFileHandler"] = wide.HandlerGzWrapper(file.SaveFileHandler)
	wideHandlers["NewFileHandler"] = wide.HandlerGzWrapper(file.NewFileHandler)
	wideHandlers["RemoveFileHandler"] = wide.HandlerGzWrapper(file.RemoveFileHandler)
	wideHandlers["RenameFileHandler"] = wide.HandlerGzWrapper(file.RenameFileHandler)
	wideHandlers["SearchTextHandler"] = wide.HandlerGzWrapper(file.SearchTextHandler)
	wideHandlers["FindHandler"] = wide.HandlerGzWrapper(file.FindHandler)
	wideHandlers["GetOutlineHandler"] = wide.HandlerGzWrapper(file.GetOutlineHandler)
	wideHandlers["CreateZipHandler"] = wide.HandlerGzWrapper(file.CreateZipHandler)
	wideHandlers["GetZipHandler"] = wide.HandlerGzWrapper(file.GetZipHandler)
	wideHandlers["UploadHandler"] = wide.HandlerGzWrapper(file.UploadHandler)
	wideHandlers["DecompressHandler"] = wide.HandlerGzWrapper(file.DecompressHandler)
	wideHandlers["WSHandlerEditor"] = wide.HandlerGzWrapper(editor.WSHandler)
	wideHandlers["GoFmtHandler"] = wide.HandlerGzWrapper(editor.GoFmtHandler)
	wideHandlers["AutocompleteHandler"] = wide.HandlerGzWrapper(editor.AutocompleteHandler)
	wideHandlers["GetExprInfoHandler"] = wide.HandlerGzWrapper(editor.GetExprInfoHandler)
	wideHandlers["FindDeclarationHandler"] = wide.HandlerGzWrapper(editor.FindDeclarationHandler)
	wideHandlers["FindUsagesHandler"] = wide.HandlerGzWrapper(editor.FindUsagesHandler)
	wideHandlers["WSHandlerNotify"] = wide.HandlerGzWrapper(notification.WSHandler)
	wideHandlers["LogoutHandler"] = wide.HandlerGzWrapper(session.LogoutHandler)
}


func (c *WideHandler) LogoutHandler() {
	handle("LogoutHandler",c)
}
func (c *WideHandler) WSHandlerNotify() {
	handle("WSHandlerNotify",c)
}
func (c *WideHandler) FindUsagesHandler() {
	handle("FindUsagesHandler",c)
}
func (c *WideHandler) FindDeclarationHandler() {
	handle("FindDeclarationHandler",c)
}
func (c *WideHandler) GetExprInfoHandler() {
	handle("GetExprInfoHandler",c)
}
func (c *WideHandler) AutocompleteHandler() {
	handle("AutocompleteHandler",c)
}
func (c *WideHandler) GoFmtHandler() {
	handle("GoFmtHandler",c)
}
func (c *WideHandler) WSHandlerEditor() {
	handle("WSHandlerEditor",c)
}
func (c *WideHandler) DecompressHandler() {
	handle("DecompressHandler",c)
}
func (c *WideHandler) UploadHandler() {
	handle("UploadHandler",c)
}

func (c *WideHandler) GetZipHandler() {
	handle("GetZipHandler",c)
}

func (c *WideHandler) CreateZipHandler() {
	handle("CreateZipHandler",c)
}

func (c *WideHandler) GetOutlineHandler() {
	handle("GetOutlineHandler",c)
}

func (c *WideHandler) FindHandler() {
	handle("FindHandler",c)
}

func (c *WideHandler) SearchTextHandler() {
	handle("SearchTextHandler",c)
}

func (c *WideHandler) RenameFileHandler() {
	handle("RenameFileHandler",c)
}

func (c *WideHandler) RemoveFileHandler() {
	handle("RemoveFileHandler",c)
}

func (c *WideHandler) NewFileHandler() {
	handle("NewFileHandler",c)
}

func (c *WideHandler) StartHandler() {
	handle("StartHandler",c)
}

func (c *WideHandler) IndexHandler() {
	handle("IndexHandler",c)
}

func (c *WideHandler) LoginHandler() {
	handle("LoginHandler",c)
}

func (c *WideHandler) WSHandler() {
	handle("WSHandler",c)
}

func (c *WideHandler) SaveContentHandler() {
	handle("SaveContentHandler",c)
}

func (c *WideHandler) BuildHandler() {
	handle("BuildHandler",c)
}

func (c *WideHandler) RunHandler() {
	handle("RunHandler",c)
}

func (c *WideHandler) StopHandler() {
	handle("StopHandler",c)
}

func (c *WideHandler) GoTestHandler() {
	handle("GoTestHandler",c)
}

func (c *WideHandler) GoVetHandler() {
	handle("GoVetHandler",c)
}

func (c *WideHandler) GoGetHandler() {
	handle("GoGetHandler",c)
}

func (c *WideHandler) GoInstallHandler() {
	handle("GoInstallHandler",c)
}


func (c *WideHandler) WSHandlerOutput() {
	handle("WSHandlerOutput",c)
}

func (c *WideHandler) CrossCompilationHandler() {
	handle("CrossCompilationHandler",c)
}

func (c *WideHandler) GetFilesHandler() {
	handle("GetFilesHandler",c)
}

func (c *WideHandler) RefreshDirectoryHandler() {
	handle("RefreshDirectoryHandler",c)
}

func (c *WideHandler) GetFileHandler() {
	handle("GetFileHandler",c)
}

func (c *WideHandler) SaveFileHandler() {
	handle("SaveFileHandler",c)
}

func handle(handlerName string, c *WideHandler){
	wideHandlers[handlerName](c.Ctx.ResponseWriter.ResponseWriter, c.Ctx.Request)
	c.TplName = "empty.html"
}