// Code generated by "stringer -type ErrCode -linecomment"; DO NOT EDIT.

package global

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Success-200]
	_ = x[StartServerErr-101000]
	_ = x[SystemErr-101001]
	_ = x[SystemInitFail-101002]
	_ = x[listenConfigErr-101003]
	_ = x[ParamsError-101004]
	_ = x[ParamsErr-101005]
	_ = x[ConnectMysqlErr-101006]
	_ = x[RequestOvertimeErr-101007]
	_ = x[SignErr-101008]
	_ = x[NoSignParamsErr-101009]
	_ = x[GetNoticeConfigErr-101010]
	_ = x[GetGameConfigErr-101011]
	_ = x[GetChannelConfigErr-101012]
	_ = x[GetLogConfigErr-101013]
	_ = x[GetApiConfigErr-101014]
	_ = x[GetDbConfigErr-101015]
	_ = x[GetGRPCConfigErr-101016]
	_ = x[GetSystemConfigErr-101017]
	_ = x[GetNacosConfigErr-101018]
	_ = x[RedisPushErr-101019]
	_ = x[RedisPublishErr-101020]
	_ = x[NeTRequestErr-101021]
	_ = x[RPCRequestErr-101022]
	_ = x[DataSaveErr-101023]
	_ = x[DataAddErr-101024]
	_ = x[DataGetErr-101025]
	_ = x[GetNakamaConfigErr-101026]
	_ = x[PublishDataErr-101027]
	_ = x[DbErr-101028]
	_ = x[DataDeleteErr-101029]
	_ = x[NoTokenErr-101030]
	_ = x[TokenErr-101031]
	_ = x[GetTokenErr-101032]
	_ = x[GetLeaderboardListErr-101033]
	_ = x[GetLeaderboardDetailErr-101034]
	_ = x[ParseJsonDataErr-101035]
	_ = x[GetAccountListErr-101036]
	_ = x[DeleteAccountErr-101037]
	_ = x[EditeAccountErr-101038]
	_ = x[GetAccountDetailErr-101039]
	_ = x[GetAccountBanListErr-101040]
	_ = x[DeleteLeaderboardErr-101041]
	_ = x[AccountUnlinkErr-101042]
	_ = x[GetAccountFriendErr-101043]
	_ = x[DeleteAccountFriendErr-101044]
	_ = x[AccountEnableErr-101045]
	_ = x[AccountDisableErr-101046]
	_ = x[GetMatchDataErr-101047]
	_ = x[GetMatchStateErr-101048]
	_ = x[AccountLoginErr-101049]
	_ = x[AccountTokenExpressErr-101050]
	_ = x[GetGameDataErr-101051]
	_ = x[ExecuteAfterDeleteFunErr-101052]
	_ = x[ExecuteAfterEditFunErr-101053]
	_ = x[ExecuteBeforeEditFunErr-101054]
	_ = x[ExecuteBeforeAddFunErr-101055]
	_ = x[LinkMysqlErr-101056]
	_ = x[GetSingleDataErr-101057]
	_ = x[CreateUploadFileDirErr-101058]
	_ = x[SystemError-101059]
	_ = x[ParamsEmptyError-101060]
	_ = x[ParamsFormatError-101061]
	_ = x[RepeatRequestError-101062]
	_ = x[InitSessionRedisErr-101063]
	_ = x[InitMysqlErr-101064]
	_ = x[InitRedisErr-101065]
	_ = x[GetSystemNoticeConfigErr-101066]
	_ = x[RegisterServerErr-101067]
	_ = x[GetServerErr-101068]
	_ = x[GetConfigErr-101069]
	_ = x[ListenConfigErr-101070]
	_ = x[GetNamingClientErr-101071]
	_ = x[GetConfigClientErr-101072]
	_ = x[GetInstanceErr-101073]
	_ = x[RunModeErr-101074]
	_ = x[SubscribeServerErr-101075]
	_ = x[UnknownServerErr-101076]
	_ = x[RPCLinkErr-101077]
	_ = x[SubscribeDataErr-101078]
	_ = x[NoRecordErr-101079]
	_ = x[PublishErr-101080]
	_ = x[TransDataTypeErr-101081]
}

const (
	_ErrCode_name_0 = "Success"
	_ErrCode_name_1 = "启动服务异常系统异常系统初始化失败配置文件监控失败参数异常，请检查参数异常，请检查连接数据库异常请求发起时间超时参数签名异常参数签名时间戳或签名为异常获取系统通知配置异常获取游戏配置异常获取发布频道配置异常获取日志配置异常获取Api配置异常获取数据库配置异常获取GRPC配置异常获取系统配置异常获取Nacos配置异常Redis push 数据异常Redis 发布消息异常网络请求异常RPC请求异常DB数据编辑异常DB数据添加异常DB数据获取异常获取Nakama配置异常数据发布异常数据库异常DB数据删除异常无Token认证信息Token认证信息无效获取Token信息异常获取Nakama排行榜数据列表异常获取Nakama排行榜数据详情异常解析Nakama json数据异常获取Nakama账户列表异常删除Nakama账户列表异常编辑Nakama账户列表异常获取Nakama账户详情异常获取Nakama禁用账户列表异常删除Nakama排行榜数据异常删除Nakama账户好友关联异常获取Nakama账户好友异常删除Nakama账户好友异常启用Nakama账户异常禁用Nakama账户异常获取Nakama比赛数据异常获取Nakama比赛状态数据异常Nakama账户登录异常Nakama Token过期异常获取Nakama数据异常执行删除后方法异常执行编辑后方法异常执行编辑前方法异常执行添加前方法异常连接数据库异常获取单条数据异常创建上传目录异常系统异常，请稍后重试参数不可空，请检查参数格式错误，请检查重复请求初始化sessionRedis连接异常初始化系统-连接管理后台数据库异常。初始化系统-连接Redis数据库异常获取系统通知配置文件异常服务注册异常获取服务异常获取配置异常监听配置异常获取服务实例异常获取配置实例异常获取服务实例异常运行模式异常服务监听异常未知服务RPC连接异常定义数据异常数据查询为空！发布消息异常数据类型转换异常"
)

var (
	_ErrCode_index_1 = [...]uint16{0, 18, 30, 51, 75, 99, 123, 144, 168, 186, 225, 255, 279, 309, 333, 354, 381, 403, 427, 450, 473, 497, 515, 530, 550, 570, 590, 614, 632, 647, 667, 687, 710, 733, 772, 811, 840, 870, 900, 930, 960, 996, 1029, 1065, 1095, 1125, 1149, 1173, 1203, 1239, 1263, 1287, 1311, 1338, 1365, 1392, 1419, 1440, 1464, 1488, 1518, 1545, 1575, 1587, 1620, 1672, 1714, 1750, 1768, 1786, 1804, 1822, 1846, 1870, 1894, 1912, 1930, 1942, 1957, 1975, 1996, 2014, 2038}
)

func (i ErrCode) String() string {
	switch {
	case i == 200:
		return _ErrCode_name_0
	case 101000 <= i && i <= 101081:
		i -= 101000
		return _ErrCode_name_1[_ErrCode_index_1[i]:_ErrCode_index_1[i+1]]
	default:
		return "ErrCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
