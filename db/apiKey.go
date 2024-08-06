package db

import (
	"net/http"

	"github.com/dgraph-io/badger"
	"github.com/google/uuid"
	"github.com/s-ir/merpel/pbs"
	"google.golang.org/protobuf/proto"
)

func CreateAuthMockupRequest(req *http.Request) (*pbs.User, error) {
	user := &pbs.User{
		Id: uuid.NewString(),
	}
	apiKey, err := GenerateAPIKey(user)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	return user, err
}

func GenerateAPIKey(user *pbs.User) (string, error) {
	id := uuid.NewString()

	marshalledUser, err := proto.Marshal(user)
	if err != nil {
		return "", err
	}

	db.Update(func(txn *badger.Txn) error {
		return txn.SetEntry(badger.NewEntry([]byte(id), marshalledUser))
	})
	return id, nil

}
func GetUser(id uuid.UUID) (*pbs.User, error) {
	var user pbs.User
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(id.String()))
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			return proto.Unmarshal(val, &user)
		})
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}
