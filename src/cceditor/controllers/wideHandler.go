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
			"fmt"
	"strings"
	"io"
)

type WideHandler struct {
	beego.Controller
}

var wideHandlers map[string]func(w http.ResponseWriter, r *http.Request)

func init(){
	wideHandlers = make(map[string]func(w http.ResponseWriter, r *http.Request))

	wideHandlers["IndexHandler"] = wide.HandlerGzWrapper(wide.IndexHandler)
	wideHandlers["LoginHandler"] = wide.HandlerWrapper(session.LoginHandler)


	wideHandlers["StartHandler"] = wide.HandlerWrapper(wide.StartHandler)
	wideHandlers["WSHandler"] = wide.HandlerWrapper(session.WSHandler)
	wideHandlers["SaveContentHandler"] = wide.HandlerWrapper(session.SaveContentHandler)
	wideHandlers["BuildHandler"] = wide.HandlerWrapper(output.BuildHandler)
	wideHandlers["RunHandler"] = wide.HandlerWrapper(output.RunHandler)
	wideHandlers["StopHandler"] = wide.HandlerWrapper(output.StopHandler)
	wideHandlers["GoTestHandler"] = wide.HandlerWrapper(output.GoTestHandler)
	wideHandlers["GoVetHandler"] = wide.HandlerWrapper(output.GoVetHandler)
	wideHandlers["GoGetHandler"] = wide.HandlerWrapper(output.GoGetHandler)
	wideHandlers["GoInstallHandler"] = wide.HandlerWrapper(output.GoInstallHandler)
	wideHandlers["WSHandlerOutput"] = wide.HandlerWrapper(output.WSHandler)
	wideHandlers["CrossCompilationHandler"] = wide.HandlerWrapper(output.CrossCompilationHandler)
	wideHandlers["GetFilesHandler"] = wide.HandlerWrapper(file.GetFilesHandler)
	wideHandlers["RefreshDirectoryHandler"] = wide.HandlerWrapper(file.RefreshDirectoryHandler)
	wideHandlers["GetFileHandler"] = wide.HandlerWrapper(file.GetFileHandler)
	wideHandlers["SaveFileHandler"] = wide.HandlerWrapper(file.SaveFileHandler)
	wideHandlers["NewFileHandler"] = wide.HandlerWrapper(file.NewFileHandler)
	wideHandlers["RemoveFileHandler"] = wide.HandlerWrapper(file.RemoveFileHandler)
	wideHandlers["RenameFileHandler"] = wide.HandlerWrapper(file.RenameFileHandler)
	wideHandlers["SearchTextHandler"] = wide.HandlerWrapper(file.SearchTextHandler)
	wideHandlers["FindHandler"] = wide.HandlerWrapper(file.FindHandler)
	wideHandlers["GetOutlineHandler"] = wide.HandlerWrapper(file.GetOutlineHandler)
	wideHandlers["CreateZipHandler"] = wide.HandlerWrapper(file.CreateZipHandler)
	wideHandlers["GetZipHandler"] = wide.HandlerWrapper(file.GetZipHandler)
	wideHandlers["UploadHandler"] = wide.HandlerWrapper(file.UploadHandler)
	wideHandlers["DecompressHandler"] = wide.HandlerWrapper(file.DecompressHandler)
	wideHandlers["WSHandlerEditor"] = wide.HandlerWrapper(editor.WSHandler)
	wideHandlers["GoFmtHandler"] = wide.HandlerWrapper(editor.GoFmtHandler)
	wideHandlers["AutocompleteHandler"] = wide.HandlerWrapper(editor.AutocompleteHandler)
	wideHandlers["GetExprInfoHandler"] = wide.HandlerWrapper(editor.GetExprInfoHandler)
	wideHandlers["FindDeclarationHandler"] = wide.HandlerWrapper(editor.FindDeclarationHandler)
	wideHandlers["FindUsagesHandler"] = wide.HandlerWrapper(editor.FindUsagesHandler)
	wideHandlers["WSHandlerNotify"] = wide.HandlerWrapper(notification.WSHandler)
	wideHandlers["LogoutHandler"] = wide.HandlerWrapper(session.LogoutHandler)
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

type myReadCloser struct{
	io.Reader
	io.Closer
}
func (mrc *myReadCloser)Close() error{
	fmt.Println("do nothing")
	return nil;
}

type myCloser struct{
}

func (my *myCloser) Close() error{
	return nil
}


func handle(handlerName string, c *WideHandler){
	switch handlerName{
	case "WSHandlerOutput","WSHandler","WSHandlerNotify","WSHandlerEditor": //do nothing
	default:
		sReader := strings.NewReader(string(c.Ctx.Input.RequestBody))
		mr := &myReadCloser{sReader,&myCloser{}};
		c.Ctx.Request.Body = mr;
	}

	wideHandlers[handlerName](c.Ctx.ResponseWriter, c.Ctx.Request)

	c.TplName = "empty.html"
}