/**
 * Created by Goland.
 * @file   leaderboard.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2022/4/13 17:56
 * @desc   leaderboard.go
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

	"github.com/gin-gonic/gin"
)

type Leaderboard struct {
	common.NakamaApi
}
type LeaderboardList struct {
	Leaderboards []LeaderboardInfo `json:"leaderboards"`
}
type LeaderboardInfo struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Category      int    `json:"category"`
	SortOrder     int    `json:"sort_order"`
	Size          int    `json:"size"`
	MaxSize       int    `json:"max_size"`
	MaxNumScore   int    `json:"max_num_score"`
	Operator      int    `json:"operator"`
	EndActive     int    `json:"end_active"`
	ResetSchedule string `json:"reset_schedule"`
	Metadata      string `json:"metadata"`
	CreateTime    any    `json:"create_time"`
	StartTime     any    `json:"start_time"`
	EndTime       any    `json:"end_time"`
	Duration      int    `json:"duration"`
	StartActive   int    `json:"start_active"`
	JoinRequired  bool   `json:"join_required"`
	Authoritative bool   `json:"authoritative"`
	Tournament    bool   `json:"tournament"`
}

type LeaderboardRecord struct {
	Records      []Records `json:"records"`
	OwnerRecords []any     `json:"owner_records"`
	NextCursor   string    `json:"next_cursor"`
	PrevCursor   string    `json:"prev_cursor"`
}

type Records struct {
	LeaderboardID string    `json:"leaderboard_id"`
	OwnerID       string    `json:"owner_id"`
	Username      string    `json:"username"`
	Score         string    `json:"score"`
	Subscore      string    `json:"subscore"`
	NumScore      int       `json:"num_score"`
	Metadata      string    `json:"metadata"`
	CreateTime    time.Time `json:"create_time"`
	UpdateTime    time.Time `json:"update_time"`
	ExpiryTime    any       `json:"expiry_time"`
	Rank          string    `json:"rank"`
	MaxNumScore   int       `json:"max_num_score"`
}

func NewLeaderboard(token string) *Leaderboard {
	leaderboard := new(Leaderboard)
	leaderboard.Token = token
	return leaderboard
}

// GetLeaderboardList 获取排行榜列表
func (a *Leaderboard) GetLeaderboardList(url string, mode string) (LeaderboardList, error) {
	xlog.Logger.Info("当前运行模式为:", mode)
	response, err := request.NewRequest().Debug(mode == gin.DebugMode).SetHeaders(a.GetNakamaHeader(a.Token)).SetTimeout(10).Get(url)
	if utils.HasErr(err, global.GetLeaderboardListErr) {
		return LeaderboardList{}, err
	}
	defer response.Close()
	if !utils.Success(response.StatusCode()) {
		content, _ := response.Content()
		xlog.Logger.Error("request api[leaderboard-list] error,result:", content)
		return LeaderboardList{}, errors.New("request nakama server error")
	}
	var leaderboardList LeaderboardList
	err = response.Json(&leaderboardList)
	if utils.HasErr(err, global.ParseJsonDataErr) {
		return LeaderboardList{}, err
	}
	return leaderboardList, nil
}

// DeleteLeaderboard 删除排行榜
func (a *Leaderboard) DeleteLeaderboard(url string, mode string) error {
	xlog.Logger.Info("当前运行模式为:", mode)
	response, err := new(request.Request).Debug(mode == gin.DebugMode).SetHeaders(a.GetNakamaHeader(a.Token)).SetTimeout(10).Delete(url)
	if utils.HasErr(err, global.DeleteLeaderboardErr) {
		return err
	}
	defer response.Close()
	if !utils.Success(response.StatusCode()) {
		errorMsg, _ := response.Content()
		xlog.Logger.Error("request api[leaderboard-delete] error:", errorMsg)
		return errors.New(errorMsg)
	}
	return nil
}

// GetLeaderboardDetail 获取排行榜详情
func (a *Leaderboard) GetLeaderboardDetail(url string, mode string) (LeaderboardInfo, error) {
	xlog.Logger.Info("当前运行模式为:", mode)
	response, err := new(request.Request).Debug(mode == gin.DebugMode).SetHeaders(a.GetNakamaHeader(a.Token)).SetTimeout(10).Get(url)
	if utils.HasErr(err, global.GetLeaderboardDetailErr) {
		return LeaderboardInfo{}, err
	}
	defer response.Close()
	if !utils.Success(response.StatusCode()) {
		xlog.Logger.Error("request api[leaderboard-detail] error,code:", response.StatusCode())
		return LeaderboardInfo{}, errors.New(global.GetLeaderboardDetailErr.String())
	}
	var leaderboardInfo LeaderboardInfo
	err = response.Json(&leaderboardInfo)
	if utils.HasErr(err, global.ParseJsonDataErr) {
		return LeaderboardInfo{}, err
	}
	return leaderboardInfo, nil
}

// GetLeaderboardRecord 获取排行榜记录
func (a *Leaderboard) GetLeaderboardRecord(url string, mode string) (LeaderboardRecord, error) {
	xlog.Logger.Info("当前运行模式为:", mode)
	response, err := new(request.Request).Debug(mode == gin.DebugMode).SetHeaders(a.GetNakamaHeader(a.Token)).SetTimeout(10).Get(url)
	if utils.HasErr(err, global.GetAccountListErr) {
		return LeaderboardRecord{}, err
	}
	defer response.Close()
	if !utils.Success(response.StatusCode()) {
		content, _ := response.Content()
		xlog.Logger.Error("request api[accounts-list] error,result:", content)
		return LeaderboardRecord{}, errors.New(global.GetAccountListErr.String())
	}
	var leaderboardRecord LeaderboardRecord
	err = response.Json(&leaderboardRecord)
	if utils.HasErr(err, global.ParseJsonDataErr) {
		return LeaderboardRecord{}, err
	}
	return leaderboardRecord, nil
}
