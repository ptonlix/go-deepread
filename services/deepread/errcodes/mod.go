package errcodes

// ErrCode 错误码类型
//
// 全局错误码文档: https://developer.work.weixin.qq.com/document/path/90313
// 文档爬取时间: 2024-01-24 16:57:05 +0800
//
// NOTE: 关于错误码的名字为何如此无聊:
//
// 官方没有给出每个错误码对应的标识符，数量太多了
// 我也懒得帮他们想，反正有文档，就先这样吧
type ErrCode = int64

// ErrCodeServiceUnavailable 系统繁忙
//
// 排查方法: 服务器暂不可用，建议稍候重试。建议重试次数不超过3次。
const ErrCodeServiceUnavailable ErrCode = -1

// ErrCodeSuccess 请求成功
//
// 排查方法: 接口调用成功
const ErrCodeSuccess ErrCode = 0
