package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Comment_reply struct {
	Id          int
	Content     string          `orm:"size(255)"`
	User        *User           `orm:"rel(fk)"`
	Comment     *Course_comment `orm:"rel(fk);column(comment_id)"`
	Reply       *Comment_reply  `orm:"null;rel(fk);column(reply_id)"`
	Create_time time.Time       `orm:"auto_now_add;type(datetime)"`
	Update_time time.Time       `orm:"auto_now_add;type(datetime)"`
	Delete_time time.Time       `orm:"type(datetime)"`
}

func GetCommentReplyByCommentId(commentId int, offset int, limit int) []Comment_reply {
	o := orm.NewOrm()
	var replies []Comment_reply
	o.QueryTable("comment_reply").Filter("comment_id", commentId).Offset(offset).Limit(limit).RelatedSel(2).All(&replies)
	return replies
}
func AddCommentReply(commentId int, content string, replyId int, userId int) (int64, error) {
	o := orm.NewOrm()
	var commentReply Comment_reply
	commentReply.Content = content
	commentReply.User = &User{
		Id: userId,
	}
	commentReply.Comment = &Course_comment{
		Id: commentId,
	}
	if replyId != 0 {
		commentReply.Reply = &Comment_reply{
			Id: replyId,
		}
	}
	id, err := o.Insert(&commentReply)
	if err == nil {
		return id, nil
	}
	return 0, err
}
func CommentReplyExist(replyId int, commentId int) bool {
	o := orm.NewOrm()
	return o.QueryTable("comment_reply").Filter("comment_id", commentId).Filter("reply_id", replyId).Exist()
}
