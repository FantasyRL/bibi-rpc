namespace go base
//公共数据结构
struct BaseResp {
    1: i64 code,
    2: string msg,
}

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

struct Comment {
    1: i64 id,
    2: i64 video_id,
    3: optional i64 parent_id,
    4: User user,
    5: string content,
    6: string publish_time,
}

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

struct Message{
    1:i64 id,
    2:i64 target_id,
    3:i64 from_id,
    4:string content,
    5:string create_time,
}
