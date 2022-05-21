package requests

type MemberDeleteRequest struct {
	MemberId int `form:"member_id" binding:"required"`
}
