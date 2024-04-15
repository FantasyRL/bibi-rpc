namespace go video

include "user.thrift"
include"base.thrift"

struct Video{
    1:i64 id,
    2:string title,
    3:user.User author,
    4:i64 uid,
    5:string play_url,
    6:string cover_url,
    7:i64 like_count,
    8:i64 comment_count,
    9:i64 is_like,
    10:string publish_time,
}

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
    3:optional list<Video> video_list,
}

struct SearchVideoRequest{
    1:string param,
    2:i64 page_num,
}

struct SearchVideoResponse{
    1:base.BaseResp base,
    2:optional i64 count,
    3:optional list<Video> video_list,
}

struct HotVideoRequest{
    1:i64 user_id,
}

struct HotVideoResponse{
    1:base.BaseResp base,
    2:optional list<Video> video_list,
}

service VideoHandler{
    PutVideoResponse PutVideo(1:PutVideoRequest req),
    ListUserVideoResponse ListVideo(1:ListUserVideoRequest req),
    SearchVideoResponse SearchVideo(1:SearchVideoRequest req),
    HotVideoResponse HotVideo(1:HotVideoRequest req),
}