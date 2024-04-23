namespace go api

struct BaseResp {
    1: i64 code
    2: string msg
}


//user
struct User {
    1: i64 id,
    2: string name,
    3: string email,
    4: i64 follow_count,
    5: i64 follower_count,
    6: bool is_follow,
    7: string avatar,
    8: i64 video_count,
}

struct RegisterRequest {
    1: required string username,
    2: required string email,
    3: required string password,
}

struct RegisterResponse {
    1: BaseResp base,
    2: optional i64 user_id,
}

//struct OTP2FAReq{
//    1:i64 uid,
//}

//struct OTP2FAResp{
//    1:BaseResp base,
//}

struct Switch2FARequest{
    1:required i64 action_type,
    2:optional string totp,
}

struct Switch2FAResponse{
    1:BaseResp base,
}

struct LoginRequest {
    1: required string username,
    2: required string password,
    3: optional string otp,
}

struct LoginResponse {
    1: BaseResp base,
    2: optional User user,
    3: optional string access_token,
    4: optional string refresh_token,
}

struct InfoRequest {
    1:required i64 user_id,
}

struct GetAccessTokenRequest{

}

struct GetAccessTokenResponse{
    1:BaseResp base,
    2:optional string access_token,
}


struct InfoResponse {
    1: BaseResp base,
    2: optional User user,
}

struct AvatarRequest{
    1:required binary avatar_file,
}
struct AvatarResponse{
    1: BaseResp base,
    2: optional User user,
}
service UserHandler {
    RegisterResponse Register(1: RegisterRequest req)(api.post="/bibi/user/register/"),
    LoginResponse Login(1: LoginRequest req)(api.post="/bibi/user/login/"),
    InfoResponse Info(1: InfoRequest req)(api.get="/bibi/user/info"),
    AvatarResponse Avatar(1:AvatarRequest req)(api.put="/bibi/user/avatar/upload"),
//    OTP2FAResp OTP2FA(1:OTP2FAReq req)(api.get="/bibi/user/2fa"),
    Switch2FAResponse Switch2FA(1:Switch2FARequest req)(api.post="/bibi/user/switch2fa"),
    GetAccessTokenResponse GetAccessToken(1:GetAccessTokenRequest req)(api.get="/bibi/access_token/get")
}

//video
struct Video{
    1:i64 id,
    2:string title,
    3:User author,
    4:i64 uid,
    5:string play_url,
    6:string cover_url,
    7:i64 like_count,
    8:i64 comment_count,
    9:i64 is_like,
    10:string publish_time,
}

struct PutVideoRequest{
    1:required binary video_file,
    2:required string title,
    3:required binary cover,
}

struct PutVideoResponse{
    1:BaseResp base,
}

struct ListUserVideoRequest{
    1:required i64 page_num,
}

struct ListUserVideoResponse{
    1:BaseResp base,
    2:optional i64 count,
    3:optional list<Video> video_list,
}

struct SearchVideoRequest{
    1:required string param,
    2:required i64 page_num,
}

struct SearchVideoResponse{
    1:BaseResp base,
    2:optional i64 count,
    3:optional list<Video> video_list,
}

struct HotVideoRequest{
}

struct HotVideoResponse{
    1:BaseResp base,
    2:optional list<Video> video_list,
}

service VideoHandler{
    PutVideoResponse PutVideo(1:PutVideoRequest req)(api.post="/bibi/video/upload"),
    ListUserVideoResponse ListVideo(1:ListUserVideoRequest req)(api.get="/bibi/video/published"),
    SearchVideoResponse SearchVideo(1:SearchVideoRequest req)(api.post="/bibi/video/search"),
    HotVideoResponse HotVideo(1:HotVideoRequest req)(api.get="/bibi/video/hot"),
}

//interaction
struct Comment {
    1: i64 id,
    2: i64 video_id,
    3: optional i64 parent_id,
    4: User user,
    5: string content,
    6: string publish_time,
}

struct LikeActionRequest{
    1:optional i64 video_id,
    2:optional i64 comment_id,
    3:required i64 action_type,
}

struct LikeActionResponse{
    1:BaseResp base,
}

struct LikeListRequest{
    1:required i64 page_num,
}

struct LikeListResponse{
    1:BaseResp base,
    2:optional i64 video_count,
    3:optional list<Video> video_list,
}

struct CommentCreateRequest{
    1:required i64 video_id,
    2:optional i64 parent_id,
    3:required string content,
}

struct CommentCreateResponse{
    1:BaseResp base,
}

struct CommentDeleteRequest{
    1:required i64 video_id,
    2:required i64 comment_id,
}

struct CommentDeleteResponse{
    1:BaseResp base,
}

struct CommentListRequest{
    1:required i64 video_id,
    2:required i64 page_num,
}

struct CommentListResponse{
    1:BaseResp base,
    2:optional i64 comment_count,
    3:optional list<Comment> comment_list,
}

service InteractionHandler{
    LikeActionResponse LikeAction(1:LikeActionRequest req)(api.post="/bibi/interaction/like/action"),
    LikeListResponse LikeList(1:LikeListRequest req)(api.get="/bibi/interaction/like/list"),
    CommentCreateResponse CommentCreate(1:CommentCreateRequest req)(api.post="/bibi/interaction/comment/create"),
    CommentDeleteResponse CommentDelete(1:CommentDeleteRequest req)(api.post="/bibi/interaction/comment/delete"),
    CommentListResponse CommentList(1:CommentListRequest req)(api.post="/bibi/interaction/comment/list"),
}