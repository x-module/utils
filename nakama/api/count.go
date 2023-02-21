/**
 * Created by Goland.
 * @file   count.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2022/4/14 18:18
 * @desc   count.go
 */

package api

import (
	"errors"
	"github.com/go-xmodule/utils/global"
	"github.com/go-xmodule/utils/nakama/common"
	"github.com/go-xmodule/utils/utils"
	"github.com/go-xmodule/utils/utils/request"
	"github.com/go-xmodule/utils/utils/xlog"
	"time"
)

type Count struct {
	common.NakamaApi
}
type CountResponse struct {
	Nodes     []Node    `json:"nodes"`
	Timestamp time.Time `json:"timestamp"`
}
type Node struct {
	Name           string  `json:"name"`
	Health         int     `json:"health"`
	SessionCount   int     `json:"session_count"`
	PresenceCount  int     `json:"presence_count"`
	MatchCount     int     `json:"match_count"`
	GoroutineCount int     `json:"goroutine_count"`
	AvgLatencyMs   float64 `json:"avg_latency_ms"`
	AvgRateSec     float64 `json:"avg_rate_sec"`
	AvgInputKbs    float64 `json:"avg_input_kbs"`
	AvgOutputKbs   float64 `json:"avg_output_kbs"`
}

func NewCount(token string) *Count {
	count := new(Count)
	count.Token = token
	return count
}

func (a *Count) GetGameServerInfo(apiUrl string, mode string) (CountResponse, error) {
	xlog.Logger.Info("当前运行模式为:", mode)
	response, err := request.NewRequest().Debug(mode == xlog.DebugMode).SetHeaders(a.GetNakamaHeader(a.Token)).SetTimeout(10).Get(apiUrl)
	if utils.HasErr(err, global.GetGameDataErr) {
		return CountResponse{}, err
	}
	defer response.Close()
	if !utils.Success(response.StatusCode()) {
		content, _ := response.Content()
		xlog.Logger.Error("request api[count] error,result:", content)
		return CountResponse{}, errors.New("request nakama server error")
	}
	var countResponse CountResponse
	res, err := response.JsonReturn(&countResponse)
	if utils.HasErr(err, global.ParseJsonDataErr) {
		xlog.Logger.Error("parse nakama count json data error", err, " response:", res)
		return CountResponse{}, err
	}
	return countResponse, nil
}
