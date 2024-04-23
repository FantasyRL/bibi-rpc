namespace go video

include"base.thrift"

struct PutVideoRequest{
    1:i64 user_id,
    2:binary video_file,
    3:string title,
    4:binary cover,
}

struct PutVideoResponse{
    1:base.BaseResp base,
}

struct ListUserVideoRequest{
    1:i64 user_id,
    2:i64 page_num,
}

struct ListUserVideoResponse{
    1:base.BaseResp base,
    2:optional i64 count,
    3:optional list<base.Video> video_list,
}

struct SearchVideoRequest{
    1:string param,
    2:i64 page_num,
}

struct SearchVideoResponse{
    1:base.BaseResp base,
    2:optional i64 count,
    3:optional list<base.Video> video_list,
}

struct HotVideoRequest{
    1:i64 user_id,
}

struct HotVideoResponse{
    1:base.BaseResp base,
    2:optional list<base.Video> video_list,
}

service VideoHandler{
    PutVideoResponse PutVideo(1:PutVideoRequest req),
    ListUserVideoResponse ListVideo(1:ListUserVideoRequest req),
    SearchVideoResponse SearchVideo(1:SearchVideoRequest req),
    HotVideoResponse HotVideo(1:HotVideoRequest req),
}