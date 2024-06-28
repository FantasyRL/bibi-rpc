namespace go user

include "base.thrift"



struct RegisterRequest {
    1: string username,
    2: string email,
    3: string password,
}

struct RegisterResponse {
    1: base.BaseResp base,
    2: optional i64 user_id,
}

//struct OTP2FAReq{
//    1:i64 uid,
//}

//struct OTP2FAResp{
//    1:BaseResp base,
//}

struct Switch2FARequest{
    1:i64 user_id,
    2:i64 action_type,
    3:optional string totp,
}

struct Switch2FAResponse{
    1:base.BaseResp base,
}

struct LoginRequest {
    1: string username,
    2: string password,
    3: optional string otp,
}

struct LoginResponse {
    1: base.BaseResp base,
    2: optional base.User user,
}

struct InfoRequest {
    1:i64 user_id,
}

struct InfoResponse {
    1: base.BaseResp base,
    2: optional base.User user,
}

struct AvatarRequest{
    1:i64 user_id,
    2:binary avatar_file,
}
struct AvatarResponse{
    1: base.BaseResp base,
    2: optional base.User user,
}

//rpc
struct GetUsersRequest{
    1:list<i64> user_id_list,
}

struct GetUsersResponse{
    1:base.BaseResp base,
    2:list<base.User> user_list,
}

struct SearchAvatarRequest{
    1: required i64 dim, // 向量维数
    2: required list<double> vector,
    3: required i64 page_num,
}

struct SearchAvatarResponse{
    1: required base.BaseResp base,
    2: optional list<string> avatar,
}


service UserHandler {
    RegisterResponse Register(1: RegisterRequest req)(api.post="/bibi/user/register/"),
    LoginResponse Login(1: LoginRequest req)(api.post="/bibi/user/login/"),
    InfoResponse Info(1: InfoRequest req)(api.get="/bibi/user/"),
    AvatarResponse Avatar(1:AvatarRequest req)(api.put="/bibi/user/avatar/upload"),
//    OTP2FAResp OTP2FA(1:OTP2FAReq req)(api.get="/bibi/user/2fa"),
    Switch2FAResponse Switch2FA(1:Switch2FARequest req)(api.post="/bibi/user/switch2fa"),
    SearchAvatarResponse SearchAvatar(1:SearchAvatarRequest req)(api.post="/bibi/user/avatar/search"),
    GetUsersResponse GetUserList(1:GetUsersRequest req),
}