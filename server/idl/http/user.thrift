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
3: string confirm_password
4: string role
}

struct RegisterResp{
1: BaseResponse base
}

struct LoginReq{
1: string username
2: string password
}

struct LoginResp{
1: BaseResponse base
2: string login_time
3: string access_token
4: string refresh_token
}

struct ChangePasswordReq{
1: string new_password
2: string old_password
}

struct ChangePasswordResp{
1: BaseResponse base
}

service UserService {
    common.NilResponse Register(1: RegisterReq req) (api.post = "/user/register")
    common.NilResponse Login(1: LoginReq req) (api.post = "/user/login")
    common.NilResponse ChangePassword(1: ChangePasswordReq  req) (api.post = "/user/password")
}