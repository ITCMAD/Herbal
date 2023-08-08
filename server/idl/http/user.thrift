namespace go api

include "../base/common.thrift"
include "../base/user.thrift"

struct BaseResponse {
    1: i64 status_code,
    2: string status_msg,
}

struct RegisterReq{
1: string username
2: string password
}

struct RegisterResp{
1: BaseResponse base
}