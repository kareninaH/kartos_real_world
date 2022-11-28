package server

import (
	stdhttp "net/http"

	myerror "real_world/pkg/error"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// 描述: 将内部错误转为http error
// 作者: hgy
// 创建日期: 2022/11/28

func FromError(err error) *myerror.HttpError {
	if err == nil {
		return nil
	}
	if se := new(myerror.HttpError); errors.As(err, &se) {
		return se
	}
	return myerror.NewHttpError(stdhttp.StatusInternalServerError, "error", "FromError()")
}

func errorEncoder(w stdhttp.ResponseWriter, r *stdhttp.Request, err error) {
	se := FromError(err)
	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(stdhttp.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/"+codec.Name())
	w.WriteHeader(se.Code)
	_, _ = w.Write(body)
}
