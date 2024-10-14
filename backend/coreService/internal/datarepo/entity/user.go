package entity

import (
	v1 "github.com/Gishinkou/kker-kratos/backend/coreService/api/v1"
	"github.com/Gishinkou/kker-kratos/backend/coreService/internal/infrastructure/persistence/model"
)

type User struct {
	ID              int64
	AccountID       int64
	Name            string
	Avatar          string
	BackgroundImage string
	Signature       string
	Mobile          string
	Email           string
}

// ↓这里是proto文件中的User结构体
// type User struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	Id              int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                                 // 用户id
// 	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                              // 用户名称
// 	Avatar          string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`                                          // 用户头像Url
// 	BackgroundImage string `protobuf:"bytes,4,opt,name=background_image,json=backgroundImage,proto3" json:"background_image,omitempty"` // 用户个人页顶部大图
// 	Signature       string `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`                                    // 个人简介
// 	Mobile          string `protobuf:"bytes,6,opt,name=mobile,proto3" json:"mobile,omitempty"`                                          // 手机号
// 	Email           string `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`                                            // 邮箱
// 	FollowCount     int64  `protobuf:"varint,8,opt,name=follow_count,json=followCount,proto3" json:"follow_count,omitempty"`            // 关注总数
// 	FollowerCount   int64  `protobuf:"varint,9,opt,name=follower_count,json=followerCount,proto3" json:"follower_count,omitempty"`      // 粉丝总数
// 	TotalFavorited  int64  `protobuf:"varint,10,opt,name=total_favorited,json=totalFavorited,proto3" json:"total_favorited,omitempty"`  // 获赞数量
// 	WorkCount       int64  `protobuf:"varint,11,opt,name=work_count,json=workCount,proto3" json:"work_count,omitempty"`                 // 作品数量
// 	FavoriteCount   int64  `protobuf:"varint,12,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count,omitempty"`     // 点赞数量
// }

// ↓这里是将User结构体转换为UserResp结构体(protoc生成)
func (u *User) ToUserResp() *v1.User {
	return &v1.User{
		Id:              u.ID,
		Name:            u.Name,
		Avatar:          u.Avatar,
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
		Mobile:          u.Mobile,
		Email:           u.Email,
	}
}

// ↓这里是将User结构体(独立定义)转换为UserModel结构体(gorm-gen生成)
func (u *User) ToUserModel() *model.User {
	return &model.User{
		ID:              u.ID,
		AccountID:       u.AccountID,
		Mobile:          u.Mobile,
		Email:           u.Email,
		Name:            u.Name,
		Avatar:          u.Avatar,
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
	}
}

// ↓这里是将UserModel结构体(gorm-gen生成)转换为User结构体(独立定义)
func FromUserModel(user *model.User) *User {
	return &User{
		ID:              user.ID,
		AccountID:       user.AccountID,
		Mobile:          user.Mobile,
		Email:           user.Email,
		Name:            user.Name,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
	}
}
