namespace go errno

enum Err {
    Success            = 0,
    NoRoute            = 1,
    NoMethod           = 2,
    BadRequest         = 10000,
    ParamsErr          = 10001,
    AuthorizeFail      = 10002,
    TooManyRequest     = 10003,
    ServiceErr         = 20000,
    RecordNotFound     = 30000,
    RecordAlreadyExist = 30001,
    DirtyData          = 30003,
    RPCUserSrvErr      = 40000,
    UserSrvErr         = 40001,
    RPCBlobSrvErr      = 50000,
    BlobSrvErr         = 50001,
    RPCProfileSrvErr   = 60000,
    ProfileSrvErr      = 60001,
}