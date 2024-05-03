namespace go chat
include"base.thrift"

struct MessageSaveRequest{
    1:i64 target_id,
    2:i64 user_id,
    3:string content,
    4:bool is_online,
}

struct MessageSaveResponse{
    1:base.BaseResp base,
}

struct MessageRecordRequest{
    1:i64 target_id,
    2:string from_time,
    3:string to_time,
    4:i64 action_type,//todo:群聊
    5:i64 page_num,
    6:i64 user_id,
}

struct MessageRecordResponse{
    1:base.BaseResp base,
    2:i64 message_count,
    3:list<base.Message> record,
}

//rpc
struct IsNotReadMessageRequest{
    1:i64 user_id,
}

struct IsNotReadMessageResponse{
    1:base.BaseResp base,
    2:i64 count,
    3:list<base.Message> message_list,
}

service ChatHandler{
    MessageSaveResponse MessageSave(1: MessageSaveRequest req),
    MessageRecordResponse MessageRecord(1: MessageRecordRequest req),
    IsNotReadMessageResponse IsNotReadMessage(1:IsNotReadMessageRequest req),
}