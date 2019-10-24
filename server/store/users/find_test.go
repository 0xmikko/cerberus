/*
* Cerberus Wallet App
* Designed for CoinList ChainLink Hackathon
*
* https://github.com/MikaelLazarev/cerberus
*
* Copyright (c) 2019, Mikhail Lazarev
 */

package users

import (
	"context"
	"fmt"
	"github.com/MikaelLazarev/cerberus/server/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"testing"
	"time"
)

func Test_store_FindByID(t *testing.T) {

	// clear collection
	filter := bson.D{}
	opts := options.Delete()
	s.Col.DeleteMany(context.TODO(), filter, opts)

	user1 := &core.User{
		ID:            "",
		Username:      "",
		Email:         "",
		Password:      nil,
		LastLogin:     time.Time{},
		EmailVerified: false,
		Active:        false,
	}
	s.Col.InsertOne(context.TODO(), user1)

	user2 := &core.User{
		ID:              "505",
		FirstName:       "Ringo",
		LastName:        "Star",
		Type:            "",
		AvatarURL:       "",
		WrikeToken:      nil,
		Host:            "",
		LastLogin:       time.Time{},
		ActiveMeetingID: "",
		AllowSpeech:     false,
		MembersID:       []*string{&user1.ID},
		Members:         []*core.User{},
	}

	s.Col.InsertOne(context.TODO(), user2)

	user2.Members = append(user2.Members, user1)

	type args struct {
		ID string
	}
	tests := []struct {
		name    string
		args    args
		want    *core.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "#1 .Inserting simple user",
			args: args{
				ID: user1.ID,
			},
			want: user1,
		},
		{
			name: "#2 .Inserting user with members",
			args: args{
				ID: user2.ID,
			},
			want: user2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.FindByIDWithMembers(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByIDWithMembers() error = %#v, wantErr %#v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByIDWithMembers() got = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func BenchmarkStore_FindByID(b *testing.B) {
	// clear collection
	filter := bson.D{}
	opts := options.Delete()
	s.Col.DeleteMany(context.TODO(), filter, opts)

	user1 := &core.User{
		ID:              "123",
		FirstName:       "John",
		LastName:        "Lennon",
		Type:            "User",
		AvatarURL:       "gif",
		Host:            "www.wrike.com",
		ActiveMeetingID: "",
		AllowSpeech:     false,
		Members:         []*core.User{},
	}
	s.Col.InsertOne(context.TODO(), user1)

	user2 := &core.User{
		ID:              "505",
		FirstName:       "Ringo",
		LastName:        "Star",
		Type:            "",
		AvatarURL:       "",
		WrikeToken:      nil,
		Host:            "",
		LastLogin:       time.Time{},
		ActiveMeetingID: "",
		AllowSpeech:     false,
		MembersID:       []*string{&user1.ID},
		Members:         []*core.User{},
	}

	s.Col.InsertOne(context.TODO(), user2)

	for i := 0; i < b.N; i++ {
		s.FindByID("505")
	}
}

func BenchmarkStore_FindByIDWithMembers(b *testing.B) {
	// clear collection
	filter := bson.D{}
	opts := options.Delete()
	s.Col.DeleteMany(context.TODO(), filter, opts)

	user1 := &core.User{
		ID:              "123",
		FirstName:       "John",
		LastName:        "Lennon",
		Type:            "User",
		AvatarURL:       "gif",
		Host:            "www.wrike.com",
		ActiveMeetingID: "",
		AllowSpeech:     false,
		Members:         []*core.User{},
	}
	s.Col.InsertOne(context.TODO(), user1)

	user2 := &core.User{
		ID:              "505",
		FirstName:       "Ringo",
		LastName:        "Star",
		Type:            "",
		AvatarURL:       "",
		WrikeToken:      nil,
		Host:            "",
		LastLogin:       time.Time{},
		ActiveMeetingID: "",
		AllowSpeech:     false,
		MembersID:       []*string{&user1.ID},
		Members:         []*core.User{},
	}

	s.Col.InsertOne(context.TODO(), user2)

	for i := 0; i < b.N; i++ {
		_, err := s.FindByIDWithMembers("505")
		if err != nil {
			fmt.Println(err)
		}
	}
}
