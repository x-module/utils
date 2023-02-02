/**
 * Created by Goland.
 * @file   const.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2022/4/11 14:45
 * @desc   const.go
 */

package common

// AuthenticateApiUrl 登录地址
const AuthenticateApiUrl = "/v2/console/authenticate"

// CountApi 统计
const CountApi = "/v2/console/status"

// AccountListApiUrl 接口-账户列表地址
const AccountListApiUrl = "/v2/console/account"

// AccountBanListApiUrl 接口-账户禁用列表地址
const AccountBanListApiUrl = "/v2/rpc/player/ban"

// AccountDetailApiUrl 接口-账户详情地址
const AccountDetailApiUrl = "/v2/console/account"

// UpdateAccountApiUrl 接口-更新账户信息
const UpdateAccountApiUrl = "/v2/console/account"

// LeaderboardApi 排行榜
const LeaderboardApi = "/v2/console/leaderboard"

// MatchApi 比赛
const MatchApi = "/v2/console/match"

// AccountStatistical 账号总数统计
const AccountStatistical = "/v2/rpc/account_statistical"

// AccountCount 账号总数
const AccountCount = "/v2/rpc/account_count"

// WalletGet 用户钱包活的
const WalletGet = "/v2/rpc/wallet/get"

// AddGoods 添加商品
const AddGoods = "/v2/rpc/add_goods"

// UpdateGoods 更新商品
const UpdateGoods = "/v2/rpc/update_goods"

type NakamaApi struct {
	Token string
}
type NakamaRpc struct {
}

// GetNakamaHeader 获取请求header
func (n *NakamaRpc) GetNakamaHeader() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
}

// GetNakamaHeader 获取请求header
func (n *NakamaApi) GetNakamaHeader(token string) map[string]string {
	Authorization := "Bearer " + token
	return map[string]string{
		"Authority":       "contractors.us-east1.nakamacloud.io",
		"Accept":          "application/json, text/plain, */*",
		"Accept-Language": "zh-CN,zh;q=0.9",
		"Cache-Control":   "no-cache",
		"Authorization":   Authorization,
		// "Cookie":"ajs_anonymous_id=9f7473db-f5a4-47ba-977b-740577facd86; _hp2_ses_props.203618332=%7B%22ts%22%3A1649854522120%2C%22d%22%3A%22contractors.us-east1.nakamacloud.io%22%2C%22h%22%3A%22%2F%22%2C%22g%22%3A%22%23%2Fleaderboards%22%7D; _hp2_id.203618332=%7B%22userId%22%3A%225710885299828656%22%2C%22pageviewId%22%3A%225167391082143752%22%2C%22sessionId%22%3A%228516329491793770%22%2C%22identity%22%3Anull%2C%22trackerVersion%22%3A%224.0%22%7D",
		"Pragma":             "no-cache",
		"Referer":            "https://contractors.us-east1.nakamacloud.io/",
		"Sec-Ch-Ua":          "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"100\", \"Google Chrome\";v=\"100\"",
		"Sec-Ch-Ua-Mobile":   "?0",
		"Sec-Ch-Ua-Platform": "\"macOS\"",
		"Sec-Fetch-Dest":     "empty",
		"Sec-Fetch-Mode":     "cors",
		"Sec-Fetch-Site":     "same-origin",
		"User-Agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.75 Safari/537.36",
	}

}
