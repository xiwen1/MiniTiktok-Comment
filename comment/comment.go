package service

import (
	"MiniTiktok-Comment/proto/auth"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"time"

	"MiniTiktok-Comment/proto/comment"
	"MiniTiktok-Comment/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type CommentActionServer struct {
	comment.UnimplementedCommentActionServer
}

var (
	connStrUser = "postgres://postgres:RpB27iLmDV4z7ZU5tpkn0UPLQWTQx1zFGaUJixDZQhPght7WWLzfZ8PLhZjavGUZ@srv.paracraft.club:31294/nicognaw?sslmode=disable"
	connStr     = "postgres://root:zkw030813@127.0.0.1:5432/root?sslmode=disable"
	pool        *sql.DB
	poolUser    *sql.DB
	consulUser  = "https://10.233.70.76:14514/minitiktok-user-grpc"
	consulAuth  = "https://10.233.70.76:14514/minitiktok-auth-grpc"
	connUser    *grpc.ClientConn
	connAuth    *grpc.ClientConn
	clientAuth  auth.AuthServiceClient
	clientUser  user.UserServiceClient
)

func InitComment() error {
	var err error
	pool, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("unable to use data source name", err)
		return nil
	}
	poolUser, err = sql.Open("postgres", connStrUser)
	if err != nil {
		fmt.Println("unable to connect to user")
		return nil
	}
	connUser, err = grpc.Dial(consulUser, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err.Error())
	}
	connAuth, err = grpc.Dial(consulAuth, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err.Error())
	}
	//clientAuth = auth.NewAuthServiceClient(connAuth)
	//clientUser = user.NewUserServiceClient(connUser)
	return nil
}

func CloseComment() error {
	if pool == nil {
		return nil
	}
	err := pool.Close()
	if err != nil {
		return err
	}
	pool = nil
	err = connUser.Close()
	if err != nil {
		return err
	}
	return nil
}

func (CommentActionServer) CommentAction(ctx context.Context, request *comment.CommentActionRequest) (response *comment.CommentActionResponse, err error) {
	fmt.Println("CommentAction: running")
	defer func() {
		if err := recover(); err != nil {
			log.Println("error:", err)
		}
	}()
	token := request.Token
	clientAuth := auth.NewAuthServiceClient(connAuth)
	clientUser := user.NewUserServiceClient(connUser)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	authResp, err := clientAuth.Auth(ctx, &auth.AuthRequest{Token: token})
	if authResp.StatusCode != auth.AuthResponse_SUCCESS {
		response.StatusCode = comment.CommentActionResponse_FAIL
		response.StatusMsg = "auth fail"
		return
	}

	userId := authResp.UserId
	actionType := request.ActionType

	userResp, err := clientUser.GetInfo(ctx, &user.UserInfoRequest{TargetId: userId, SelfId: userId})
	if userResp.StatusCode == user.UserInfoResponse_TARGET_NOT_FOUND {
		response.StatusCode = comment.CommentActionResponse_FAIL
		response.StatusMsg = "userinfo target not found"
		return
	}
	if userResp.StatusCode == user.UserInfoResponse_UNSPECIFIED {
		response.StatusCode = comment.CommentActionResponse_FAIL
		response.StatusMsg = "userid unspecified"
		return
	}

	u := comment.User{
		Id:            userId,
		FollowCount:   userResp.FollowCount,
		FollowerCount: userResp.FollowerCount,
		Name:          userResp.Username,
		IsFollow:      userResp.IsFollow,
	}

	if actionType == comment.CommentActionRequest_PUBLISH {
		c := Comment{}
		c.User = userId
		c.Video_id = request.VideoId
		c.Content = request.CommentText
		c.CreateDate = time.Now().Format("01-02")
		Id, err := c.insertComment()
		if err != nil {
			response.StatusCode = comment.CommentActionResponse_FAIL
			response.StatusMsg = "unable to insert into database"
			log.Fatal(err.Error())
		}
		cc := comment.Comment{Id: Id, Content: c.Content, CreateDate: c.CreateDate, User: &u}
		response.Comment = &cc
		response.StatusMsg = "success"
		response.StatusCode = comment.CommentActionResponse_SUCCESS
	} else {
		err := Delete(request.CommentId)
		if err != nil {
			response.StatusCode = comment.CommentActionResponse_FAIL
			response.StatusMsg = "unable to delete from database"
			log.Fatal(err.Error())
		}
		response.StatusMsg = "success"
		response.StatusCode = comment.CommentActionResponse_SUCCESS
	}
	return
}

func (CommentActionServer) CommentList(ctx context.Context, request *comment.CommentListRequest) (response *comment.CommentListResponse, err error) {
	log.Println("CommentList running")
	defer func() {
		if err := recover(); err != nil {
			log.Println("error:", err)
		}
	}()
	token := request.Token
	clientAuth := auth.NewAuthServiceClient(connAuth)
	authResp, err := clientAuth.Auth(ctx, &auth.AuthRequest{Token: token})
	if authResp.StatusCode != auth.AuthResponse_SUCCESS {
		response.StatusCode = comment.CommentListResponse_FAIL
		response.StatusMsg = "CommentList: unable to auth"
		return
	}
	userId := authResp.UserId
	video_id := request.VideoId
	c, err := SearchComment(uint32(video_id), userId)
	if err != nil {
		response.StatusCode = comment.CommentListResponse_FAIL
		response.StatusMsg = "unable to search for comment"
		log.Fatal(err.Error())
		return
	}
	response.StatusMsg = "success"
	response.StatusCode = comment.CommentListResponse_SUCCESS
	response.CommentList = *c

	return
}
