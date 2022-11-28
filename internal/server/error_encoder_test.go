package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
)

// 描述: TODO
// 作者: hgy
// 创建日期: 2022/11/28

func TestFromError(t *testing.T) {
	err := errors.New(http.StatusUnauthorized, "JWT_PARSE_ERROR", "no authorization")
	fmt.Println(err)
	httpError := FromError(err)
	marshal, _ := json.Marshal(httpError)
	fmt.Println(string(marshal))
}
