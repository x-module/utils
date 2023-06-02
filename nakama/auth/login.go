/**
 * Created by Goland.
 * @file   login.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2022/4/11 11:37
 * @desc   login.go
 */

package auth

import (
	"crypto"
	"errors"
	"fmt"
	"github.com/x-module/utils/global"
	"github.com/x-module/utils/nakama/common"
	"github.com/x-module/utils/utils"
	"github.com/x-module/utils/utils/request"
	"github.com/x-module/utils/utils/xlog"
	"time"
)

// LoginToken 身份验证token
type LoginToken struct {
	Token string   `json:"token"`
	Uname string   `json:"uname"`
	Email string   `json:"email"`
	Role  UserRole `json:"role"`
}

type ConsoleTokenClaims struct {
	Username  string   `json:"usn,omitempty"`
	Email     string   `json:"ema,omitempty"`
	Role      UserRole `json:"rol,omitempty"`
	ExpiresAt int64    `json:"exp,omitempty"`
	Cookie    string   `json:"cki,omitempty"`
}

// InvalidToken 无效token
const InvalidToken = 2

// EffectiveToken 有效token
const EffectiveToken = 1

// ExpireToken 过期token
const ExpireToken = 3

type UserRole int32

type Auth struct {
	common.NakamaApi
	userName string
	password string
	url      string
	model    string
	signKey  string
}

func NewAuth(userName string, password string, url string, signKey string, model string) *Auth {
	auth := new(Auth)
	auth.userName = userName
	auth.password = password
	auth.url = url
	auth.model = model
	auth.signKey = signKey
	return auth
}

// Valid 校验
func (stc *ConsoleTokenClaims) Valid() error {
	// Verify expiry.
	if stc.ExpiresAt <= time.Now().UTC().Unix() {
		vErr := new(jwt.ValidationError)
		xlog.Logger.Warning("Token is expired")
		vErr.Inner = errors.New("Token is expired")
		vErr.Errors |= jwt.ValidationErrorExpired
		return vErr
	}
	return nil
}

// 解析token
func (a *Auth) parseConsoleToken(hmacSecretByte []byte, tokenString string) (username, email string, role UserRole, exp int64, ok bool) {
	token, err := jwt.ParseWithClaims(tokenString, &ConsoleTokenClaims{}, func(token *jwt.Token) (any, error) {
		if s, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || s.Hash != crypto.SHA256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecretByte, nil
	})
	if utils.HasErr(err, global.GetTokenErr) {
		return
	}
	claims, ok := token.Claims.(*ConsoleTokenClaims)
	if !ok || !token.Valid {
		return
	}
	return claims.Username, claims.Email, claims.Role, claims.ExpiresAt, true
}

// token 检测
func (a *Auth) testToken(loginToken LoginToken) (int, error) {
	token, err := jwt.Parse(loginToken.Token, func(token *jwt.Token) (any, error) {
		if s, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || s.Hash != crypto.SHA256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.signKey), nil
	})
	if utils.HasErr(err, global.GetTokenErr) {
		xlog.Logger.Error("parse token error:", err, "  config:", a.signKey, ", token:", loginToken.Token)
		return InvalidToken, err
	}
	uname, email, role, exp, ok := a.parseConsoleToken([]byte(a.signKey), loginToken.Token)
	xlog.Logger.Debug("parse_console_token:", " uname:", uname, "  email:", email, "  role:", role, "  exp:", exp, "  ok:", ok)
	if !ok || !token.Valid {
		// The token or its claims are invalid.
		xlog.Logger.Error("console login  token or its claims are invalid")
		return InvalidToken, err
	}
	if exp <= time.Now().UTC().Unix() {
		// Token expired.
		xlog.Logger.Error("console login  token login expired")
		return ExpireToken, err
	}
	return EffectiveToken, nil
}

// GetToken 获取身份token
func (a *Auth) GetToken(loginToken LoginToken) (LoginToken, error) {
	if loginToken.Token == "" {
		token, err := a.login()
		if utils.HasErr(err, global.AccountLoginErr) {
			return LoginToken{}, err
		} else {
			return token, err
		}
	} else {
		_, err := a.testToken(loginToken)
		if utils.HasErr(err, global.AccountTokenExpressErr) {
			// if checkResult == ExpireToken { // token过期
			return a.GetToken(LoginToken{})
			// }
			// return LoginToken{}, err
		} else {
			return loginToken, err
		}
	}
}

// 登录操作
func (a *Auth) login() (LoginToken, error) {
	data := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: a.userName,
		Password: a.password,
	}

	xlog.Logger.Info("当前运行模式为:", a.model)
	response, err := request.NewRequest().Debug(a.model == xlog.DebugMode).Json().SetTimeout(10).Post(a.url, data)
	if utils.HasErr(err, global.AccountLoginErr) {
		return LoginToken{}, err
	}
	defer response.Close()
	if !utils.Success(response.StatusCode()) {
		xlog.Logger.Error("request nakama server error", response)
		return LoginToken{}, errors.New("request nakama server error")
	}
	var loginToken LoginToken
	err = response.Json(&loginToken)
	if utils.HasErr(err, global.ParseJsonDataErr) {
		return LoginToken{}, err
	}
	xlog.Logger.Info("success login nakama console. token info: ", loginToken)

	// uname, email, role, exp, ok := a.parseConsoleToken([]byte(a.Config.NakamaConfig.Account.SignKey), loginToken.Token)
	// xlog.Logger.Debug("parseConsoleToken:", " uname:", uname, "  email:", email, "  role:", role, "  exp:", exp, "  ok:", ok)
	// loginToken.Uname = uname
	// loginToken.Email = email
	// loginToken.Role = role
	return loginToken, nil
}
