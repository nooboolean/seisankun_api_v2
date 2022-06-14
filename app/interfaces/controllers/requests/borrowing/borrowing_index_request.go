package requests

type BorrowingIndexRequest struct {
	MemberId int `form:"member_id" binding:"required"`
}
