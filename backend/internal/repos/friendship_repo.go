package repos

import (
	"database/sql"
	"fmt"

	"rainbow-umbrella/internal/consts"
	"rainbow-umbrella/internal/interfaces"
	"rainbow-umbrella/internal/objects/dao"
)

type friendshipRepo struct {
	dbClient *sql.DB
}

func NewFriendshipRepo(dbClient *sql.DB) interfaces.IFriendshipRepo {
	return &friendshipRepo{dbClient: dbClient}
}

func (r friendshipRepo) InsertOne(friendship *dao.Friendship) error {
	q := buildInsertOneFriendshipQuery(friendship)

	if _, err := r.dbClient.Exec(q.Query, q.Args...); err != nil {
		return fmt.Errorf("[friendshipRepo.InsertOne][1]: %+v", err)
	}

	return nil
}

func (r friendshipRepo) FriendList(userID uint64) (*dao.FriendList, error) {
	q := buildFriendListQuery(userID)

	rows, err := r.dbClient.Query(q.Query, q.Args...)
	if err != nil {
		return nil, fmt.Errorf("[friendshipRepo.FriendList][1]: %+v", err)
	}

	friendList := new(dao.FriendList)

	for rows.Next() {
		var (
			requestingUserID uint64
			targetingUserID  uint64
			status           string
		)
		user := new(dao.User)

		err = rows.Scan(
			&requestingUserID, &targetingUserID, &status,
			&user.ID, &user.Login,
			&user.FirstName, &user.LastName, &user.Birthday, &user.Gender, &user.City,
			&user.CreatedAt,
		)

		if status == consts.FriendshipStatusAccept {
			friendList.Friends = append(friendList.Friends, *user)
		}

		if requestingUserID == userID {
			friendList.WaitingForResponse = append(friendList.WaitingForResponse, *user)
		}

		if targetingUserID == userID {
			friendList.Requested = append(friendList.Requested, *user)
		}

		if err != nil {
			return nil, fmt.Errorf("[friendshipRepo.FriendList][2]: %+v", err)
		}
	}

	return friendList, nil
}
