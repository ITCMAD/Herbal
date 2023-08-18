namespace go user

include "../base/common.thrift"
include "../base/user.thrift"

struct RegisterReq{
1:string username
2:string password
3:string confirm_password
4:string role
}

struct RegisterResp{
1:common.BaseResponse base_resp
}

struct LoginReq{
1:string username
2:string password
}

struct LoginResp{
1:common.BaseResponse base_resp
2:string login_time
3:string access_token
4:string refresh_token
}

struct ChangePasswordReq{
1:string old_password
2:string new_password
}

struct ChangePasswordResp{
1:common.BaseResponse base_resp
}

service UserService{
RegisterResp Register(1: RegisterReq req)
LoginResp Login(1: LoginReq req)
ChangePasswordResp ChangePassword(1: ChangePasswordReq req)
}