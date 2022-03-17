package repos

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"

	"rainbow-umbrella/internal/consts"
	"rainbow-umbrella/internal/interfaces"
	"rainbow-umbrella/internal/objects/dao"
	"rainbow-umbrella/internal/utils"
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
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			if driverErr.Number == 1452 {
				return fmt.Errorf("[friendshipRepo.InsertOne][1]: %w: %v", utils.AppErrorAlreadyExist, err.Error())
			}
		}

		return fmt.Errorf("[friendshipRepo.InsertOne][2]: %w", err)
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
			&user.FriendshipID,
			&requestingUserID, &targetingUserID, &status,
			&user.ID, &user.Login,
			&user.FirstName, &user.LastName, &user.Birthday, &user.Gender, &user.City,
			&user.CreatedAt,
		)

		if status == consts.FriendshipStatusAccept {
			friendList.Friends = append(friendList.Friends, *user)
		} else if requestingUserID == userID {
			friendList.WaitingForResponse = append(friendList.WaitingForResponse, *user)
		} else if targetingUserID == userID {
			friendList.Requested = append(friendList.Requested, *user)
		}

		if err != nil {
			return nil, fmt.Errorf("[friendshipRepo.FriendList][2]: %+v", err)
		}
	}

	return friendList, nil
}

func (r friendshipRepo) UpdateStatus(id uint64, status string) error {
	q := buildUpdateOneFriendshipQuery(id, status)

	if _, err := r.dbClient.Exec(q.Query, q.Args...); err != nil {
		return fmt.Errorf("[friendshipRepo.UpdateStatus][1]: %+v", err)
	}

	return nil
}
