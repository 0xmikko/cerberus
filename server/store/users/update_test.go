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
	"github.com/MikaelLazarev/cerberus/server/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"testing"
)

func Test_store_Update(t *testing.T) {

	// clear collection
	filter := bson.D{}
	opts := options.Delete()
	s.Col.DeleteMany(context.TODO(), filter, opts)

	user1 := &core.User{
		ID: "123",
	}
	s.Col.InsertOne(context.TODO(), user1)

	user2 := &core.User{
		ID: "123",
	}

	s.UpdateUserFromMembersInfo(user2)

	userWant := *user1
	userWant.FirstName = user2.FirstName
	userWant.LastName = user2.LastName
	userWant.Type = user2.Type
	userWant.AvatarURL = user2.AvatarURL

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
			name: "#1 .Partial update",
			args: args{
				ID: user1.ID,
			},
			want: &userWant,
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
