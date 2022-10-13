namespace go learn.fx.item

struct FilterVisibleItemsReq {
    1: required i64 UserID
    2: required list<i64> ItemIDs
}

struct FilterVisibleItemsResp {
    1: required list<i64> VisibleItemIDs
}

struct Item {
    1: required i64 ID
    2: required string Name
    3: required string Desc
    4: required list<i64> VisibleUsers
}

struct CreateItemReq {
    1: required Item Item
}

struct CreateItemResp {
    1: required i64 ID
}

service LearnFxService {
    CreateItemResp CreateItem(1: CreateItemReq req)
    FilterVisibleItemsResp FilterVisibleItems(1: FilterVisibleItemsReq req)
}

