CREATE TABLE events
(
    Id        UInt64,
    EventDate Date,
    Data0     String,
    Data1     String,
    Data2     String,
    Data3     String,
    Data4     String,
    Data5     String,
    Data6     String,
    Data7     String,
    Data8     String,
    Data9     String,
    Data10    String,
    Data11    String,
    Data12    String,
    Data13    String,
    Data14    String,
    Data15    String,
    Data16    String,
    Data17    String,
    Data18    String,
    Data19    String,
    Data20    String,
    Data21    String,
    Data22    String,
    Data23    String,
    Data24    String,
    Data25    String,
    Data26    String,
    Data27    String,
    Data28    String,
    Data29    String
) engine = MergeTree(EventDate, (Id, EventDate), 8192);
