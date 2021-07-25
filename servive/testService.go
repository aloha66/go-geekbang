package servive

import (
	"homework2/dao"
	"net/http"
)

func TestService(writer http.ResponseWriter, request *http.Request)  {
	if err := getSomething();err != nil {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("出错了"))
		return
	}
	writer.WriteHeader(http.StatusOK)

}

func getSomething() error  {
	err := dao.QuerySomeThing()
	return err
}