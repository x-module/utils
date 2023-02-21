/**
 * Created by Goland.
 * @file   account.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2022/4/8 19:32
 * @desc   account.go
 */

package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-xmodule/utils/global"
	"github.com/go-xmodule/utils/nakama/common"
	"github.com/go-xmodule/utils/utils"
	"github.com/go-xmodule/utils/utils/request"
	"github.com/go-xmodule/utils/utils/xlog"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

type Account struct {
	common.NakamaApi
}

type AccountInfo struct {
	Account     AccountData `json:"account"`
	DisableTime any         `json:"disable_time"`
}
type AccountData struct {
	User        User   `json:"user"`
	Wallet      string `json:"wallet"`
	Email       string `json:"email"`
	Devices     []any  `json:"devices"`
	CustomID    string `json:"custom_id"`
	VerifyTime  any    `json:"verify_time"`
	DisableTime string `json:"disable_time"`
}
type Accounts struct {
	Users      []User `json:"users"`
	TotalCount int    `json:"total_count"`
	NextCursor string `json:"next_cursor"`
}

type BanPlayer struct {
	ID          string     `json:"id"`
	Username    string     `json:"username"`
	DisplayName string     `json:"display_name"`
	AvatarURL   string     `json:"avatar_url"`
	LangTag     string     `json:"lang_tag"`
	Metadata    string     `json:"metadata"`
	EdgeCount   int        `json:"edge_count"`
	CreateTime  CreateTime `json:"create_time"`
	UpdateTime  UpdateTime `json:"update_time"`
	SteamID     string     `json:"steam_id,omitempty"`
}
type CreateTime struct {
	Seconds int `json:"seconds"`
}
type UpdateTime struct {
	Seconds int `json:"seconds"`
}

type User struct {
	ID                    string    `json:"id"`
	Username              string    `json:"username"`
	DisplayName           string    `json:"display_name"`
	AvatarURL             string    `json:"avatar_url"`
	LangTag               string    `json:"lang_tag"`
	Location              string    `json:"location"`
	Timezone              string    `json:"timezone"`
	Metadata              string    `json:"metadata"`
	FacebookID            string    `json:"facebook_id"`
	GoogleID              string    `json:"google_id"`
	GamecenterID          string    `json:"gamecenter_id"`
	SteamID               string    `json:"steam_id"`
	Online                bool      `json:"online"`
	EdgeCount             int       `json:"edge_count"`
	CreateTime            time.Time `json:"create_time"`
	UpdateTime            time.Time `json:"update_time"`
	FacebookInstantGameID string    `json:"facebook_instant_game_id"`
	AppleID               string    `json:"apple_id"`
}

type Encoder struct{}
type Params struct {
	Updates   any     `json:"updates"`
	CloneFrom any     `json:"cloneFrom"`
	Encoder   Encoder `json:"encoder"`
	Map       any     `json:"map"`
}
type NormalizedNames struct{}
type LazyUpdate struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Op    string `json:"op"`
}
type LazyInit struct {
	NormalizedNames NormalizedNames   `json:"normalizedNames"`
	LazyUpdate      any               `json:"lazyUpdate"`
	Headers         map[string]string `json:"headers"`
}
type Headers struct {
	NormalizedNames NormalizedNames   `json:"normalizedNames"`
	LazyUpdate      []LazyUpdate      `json:"lazyUpdate"`
	Headers         map[string]string `json:"headers"`
	LazyInit        LazyInit          `json:"lazyInit"`
}
type FriendResponse struct {
	Friends []Friends `json:"friends"`
	Cursor  string    `json:"cursor"`
}
type Friends struct {
	State      int       `json:"state"`
	UpdateTime time.Time `json:"update_time"`
	User       User      `json:"user,omitempty"`
}
type Payload struct {
	Params  Params  `json:"params"`
	Headers Headers `json:"headers"`
}

func NewAccount(token string) *Account {
	account := new(Account)
	account.Token = token
	return account
}

// GetAccountList 获取用户列表
func (a *Account) GetAccountList(apiUrl string, filter string, cursor string, mode string) (Accounts, error) {
	apiUrl = apiUrl + "?a=a"
	if filter != "" {
		filter = url.QueryEscape(filter)
		apiUrl = fmt.Sprintf("%s&filter=%s", apiUrl, filter)
	}
	if cursor != "" {
		apiUrl = fmt.Sprintf("%s&cursor=%s", apiUrl, cursor)
	}
	xlog.Logger.Info("当前运行模式为:", mode)
	response, err := request.NewRequest().Debug(mode == gin.DebugMode).SetHeaders(a.GetNakamaHeader(a.Token)).SetTimeout(10).Get(apiUrl)
	if utils.HasErr(err, global.GetAccountListErr) {
		xlog.Logger.Error("request api[accounts-list] error:", err)
		return Accounts{}, err
	}
	defer response.Close()
	if !utils.Success(response.StatusCode()) {
		content, _ := response.Content()
		xlog.Logger.Error("request api[accounts-list] error,result:", content)
		return Accounts{}, errors.New("request nakama server error")
	}
	var accounts Accounts
	err = response.Json(&accounts)
	if utils.HasErr(err, global.ParseJsonDataErr) {
		return Accounts{}, err
	}
	return accounts, nil
}

// GetAccountBanList 获取用用列表
func (a *Account) GetAccountBanList(apiUrl string, UserID string, UserName string, Offset int, Limit int, mode string) ([]BanPlayer, error) {
	xlog.Logger.Info("当前运行模式为:", mode)
	params := map[string]any{
		"user_id":   UserID,
		"user_name": UserName,
		"offset":    Offset,
		"limit":     Limit,
	}
	response, err := new(request.Request).Debug(mode == gin.DebugMode).SetHeaders(map[string]string{
		"Accept": "application/json",
	}).SetTimeout(20).Post(apiUrl, params)
	if utils.HasErr(err, global.GetAccountBanListErr) {
		return nil, err
	}
	defer response.Close()

	if !utils.Success(response.StatusCode()) {
		content, _ := response.Content()
		xlog.Logger.Error("request api[accounts-ban-list] response code error，code:", response.StatusCode(), ",result:", content)
		return nil, errors.New("request nakama server error")
	}
	// c, _ := response.Content()
	// utils.JsonDisplay(c)
	var accounts []BanPlayer
	err = response.Json(&accounts)
	if utils.HasErr(err, global.ParseJsonDataErr) {
		return nil, err
	}
	return accounts, nil
}

// GetAccountDetail 获取用户详情
func (a *Account) GetAccountDetail(id string, url string, mode string) (AccountInfo, error) {
	xlog.Logger.Info("当前运行模式为:", mode)
	response, err := new(request.Request).Debug(mode == gin.DebugMode).SetHeaders(a.GetNakamaHeader(a.Token)).SetTimeout(10).Get(url)
	if utils.HasErr(err, global.GetAccountDetailErr) {
		return AccountInfo{}, err
	}
	defer response.Close()
	if !utils.Success(response.StatusCode()) {
		xlog.Logger.Error("request api[accounts-detail] error,code:", response.StatusCode())
		return AccountInfo{}, errors.New("request nakama server error")
	}
	var accountInfo AccountInfo
	err = response.Json(&accountInfo)
	if utils.HasErr(err, global.ParseJsonDataErr) {
		return AccountInfo{}, err
	}
	return accountInfo, nil
}

func (a *Account) UpdateAccount(id string, params []byte, url string, mode string) (string, error) {
	type Payload struct {
		Username    string `json:"username"`
		DisplayName string `json:"display_name"`
		AvatarURL   string `json:"avatar_url"`
		Location    string `json:"location"`
		Timezone    string `json:"timezone"`
		Metadata    string `json:"metadata"`
	}
	var data Payload
	_ = json.Unmarshal(params, &data)
	xlog.Logger.Info("当前运行模式为:", mode)
	response, err := new(request.Request).Debug(mode == gin.DebugMode).SetHeaders(a.GetNakamaHeader(a.Token)).SetTimeout(10).Post(url, data)
	if utils.HasErr(err, global.EditeAccountErr) {
		return "", err
	}
	defer response.Close()
	type ErrorResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Details []any  `json:"details"`
	}
	if !utils.Success(response.StatusCode()) {
		res, _ := response.Content()
		var errorResp ErrorResponse
		_ = json.Unmarshal([]byte(res), &errorResp)
		xlog.Logger.Error("request api[accounts-detail] error,code:", res)
		return errorResp.Message, errors.New(errorResp.Message)
	}
	return "success", nil
}

// Unlink account unlink
func (a *Account) Unlink(url string, mode string) error {
	data := Payload{
		Params: Params{},
		Headers: Headers{
			NormalizedNames: NormalizedNames{},
			LazyUpdate: []LazyUpdate{
				{
					Name:  "Authorization",
					Value: "Bearer ",
					Op:    "s",
				},
			},
			Headers:  map[string]string{},
			LazyInit: LazyInit{},
		},
	}
	xlog.Logger.Info("当前运行模式为:", mode)
	response, err := new(request.Request).Debug(mode == gin.DebugMode).SetHeaders(a.GetNakamaHeader(a.Token)).Json().SetTimeout(10).Post(url, data)
	if utils.HasErr(err, global.AccountUnlinkErr) {
		return err
	}
	defer response.Close()
	type ErrorResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Details []any  `json:"details"`
	}
	if !utils.Success(response.StatusCode()) {
		res, _ := response.Content()
		var errorResp ErrorResponse
		_ = json.Unmarshal([]byte(res), &errorResp)
		xlog.Logger.Error("request api[accounts-detail] error,code:", res)
		return errors.New(errorResp.Message)
	}
	return nil
}

// ChangeAccount 修改邮箱密码
func (a *Account) ChangeAccount(email string, password string, url string, mode string) error {
	type Payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	data := Payload{
		Email:    email,
		Password: password,
	}

	xlog.Logger.Info("当前运行模式为:", mode)
	response, err := new(request.Request).Debug(mode == gin.DebugMode).SetHeaders(a.GetNakamaHeader(a.Token)).Json().SetTimeout(10).Post(url, data)
	if utils.HasErr(err, global.EditeAccountErr) {
		return err
	}
	defer response.Close()
	type ErrorResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Details []any  `json:"details"`
	}
	if !utils.Success(response.StatusCode()) {
		res, _ := response.Content()
		var errorResp ErrorResponse
		_ = json.Unmarshal([]byte(res), &errorResp)
		xlog.Logger.Error("request api[accounts-change-account] error,code:", res)
		return errors.New(errorResp.Message)
	}
	return nil
}

// GetFriends 获取账户朋友
func (a *Account) GetFriends(url string, mode string) (FriendResponse, error) {
	xlog.Logger.Info("当前运行模式为:", mode)
	response, err := new(request.Request).Debug(mode == gin.DebugMode).SetHeaders(a.GetNakamaHeader(a.Token)).SetTimeout(10).Get(url)
	if utils.HasErr(err, global.GetAccountFriendErr) {
		return FriendResponse{}, err
	}
	defer response.Close()
	if !utils.Success(response.StatusCode()) {
		errorMsg, _ := response.Content()
		xlog.Logger.Error("request api[accounts-friend-get] error:", errorMsg)
		return FriendResponse{}, errors.New(errorMsg)
	}
	var friendResponse FriendResponse
	err = response.Json(&friendResponse)
	if utils.HasErr(err, global.ParseJsonDataErr) {
		return FriendResponse{}, err
	}
	return friendResponse, nil
}

// DeleteFriend 删除好友
func (a *Account) DeleteFriend(url string, mode string) error {
	xlog.Logger.Info("当前运行模式为:", mode)
	response, err := new(request.Request).Debug(mode == gin.DebugMode).SetHeaders(a.GetNakamaHeader(a.Token)).SetTimeout(10).Delete(url)
	if utils.HasErr(err, global.DeleteAccountFriendErr) {
		return err
	}
	defer response.Close()
	if !utils.Success(response.StatusCode()) {
		errorMsg, _ := response.Content()
		xlog.Logger.Error("request api[accounts-friend-delete] error:", errorMsg)
		return errors.New(errorMsg)
	}
	return nil
}

// DeleteAccount 删除账户
func (a *Account) DeleteAccount(url string, mode string) error {
	xlog.Logger.Info("当前运行模式为:", mode)
	response, err := new(request.Request).Debug(mode == gin.DebugMode).SetHeaders(a.GetNakamaHeader(a.Token)).SetTimeout(10).Delete(url)
	if utils.HasErr(err, global.DeleteAccountErr) {
		return err
	}
	defer response.Close()
	if !utils.Success(response.StatusCode()) {
		errorMsg, _ := response.Content()
		xlog.Logger.Error("request api[accounts-delete] error:", errorMsg)
		return errors.New(errorMsg)
	}
	return nil
}

func (a *Account) Enable(url string, mode string) error {
	data := Payload{
		Params: Params{},
		Headers: Headers{
			NormalizedNames: NormalizedNames{},
			LazyUpdate: []LazyUpdate{
				{
					Name:  "Authorization",
					Value: "Bearer ",
					Op:    "s",
				},
			},
			Headers:  map[string]string{},
			LazyInit: LazyInit{},
		},
	}
	xlog.Logger.Info("当前运行模式为:", mode)
	response, err := new(request.Request).Debug(mode == gin.DebugMode).SetHeaders(a.GetNakamaHeader(a.Token)).Json().SetTimeout(10).Post(url, data)
	if utils.HasErr(err, global.AccountEnableErr) {
		return err
	}
	defer response.Close()
	type ErrorResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Details []any  `json:"details"`
	}
	if !utils.Success(response.StatusCode()) {
		res, _ := response.Content()
		var errorResp ErrorResponse
		_ = json.Unmarshal([]byte(res), &errorResp)
		xlog.Logger.Error("request api[accounts-enable] error,code:", res)
		return errors.New(errorResp.Message)
	}
	return nil
}
func (a *Account) Disable(url string, mode string) error {
	data := Payload{
		Params: Params{},
		Headers: Headers{
			NormalizedNames: NormalizedNames{},
			LazyUpdate: []LazyUpdate{
				{
					Name:  "Authorization",
					Value: "Bearer ",
					Op:    "s",
				},
			},
			Headers:  map[string]string{},
			LazyInit: LazyInit{},
		},
	}
	xlog.Logger.Info("当前运行模式为:", mode)
	response, err := new(request.Request).Debug(mode == gin.DebugMode).SetHeaders(a.GetNakamaHeader(a.Token)).Json().SetTimeout(10).Post(url, data)
	if utils.HasErr(err, global.AccountDisableErr) {
		return err
	}
	defer response.Close()
	type ErrorResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Details []any  `json:"details"`
	}
	if !utils.Success(response.StatusCode()) {
		res, _ := response.Content()
		var errorResp ErrorResponse
		_ = json.Unmarshal([]byte(res), &errorResp)
		xlog.Logger.Error("request api[accounts-disable] error,code:", res)
		return errors.New(errorResp.Message)
	}
	return nil
}
