package users

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Check if user with this name exists
func (s *store) IsExists(Username string) (bool, error) {
	// Creating filter which find / create an item with the same meeting.ID
	filter := bson.D{{
		"username", Username,
	}}

	res, err := s.Col.Find(context.TODO(), filter, options.Find())
	if err != nil {
		return false, err
	}

	return res.Next(context.TODO()), nil

}
