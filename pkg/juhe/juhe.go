package juhe

import "github.com/pkg/errors"

const (
	// _baseURL 聚合https://www.juhe.cn/docs请求基础URL
	_baseURL = "http://web.juhe.cn:8080"
)

type CommonResp struct {
	ErrorCode int    `json:"error_code"`
	Reason    string `json:"reason"`
}

func (c CommonResp) Error() error {
	if c.ErrorCode != 0 {
		return errors.New(c.Reason)
	}
	return nil
}
