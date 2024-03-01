package sredis

import (
	"context"
	"encoding/json"
	"fmt"
	conf "go-exam/user-service/config"
	pbu "go-exam/user-service/genproto/user"
	pkgdb "go-exam/user-service/pkg/db"
	pgf "go-exam/user-service/storage/postgres"
	se "go-exam/user-service/storage/sendEmail"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisRepo struct {
	rConn *redis.Client
}

func NewRedisRepo(rds *redis.Client) *redisRepo {
	return &redisRepo{
		rConn: rds,
	}
}

func (r *redisRepo) Sign(user *pbu.UserDetail) (*pbu.ResponseMessage, error) {
	byteData, err := json.Marshal(&user)
	if err != nil {
		// log.Fatal(err)
		fmt.Println("here 1 eror")
		return nil, err
	}
	code := se.SendEmail(user.Email)

	//
	err = r.rConn.Set(context.Background(), code, byteData, time.Minute*3).Err()
	if err != nil {
		// log.Fatal(err)
		fmt.Println("here 2 eror")
		return nil, err
	}

	reply := pbu.ResponseMessage{
		Content: "Get your code from your email!",
	}

	return &reply, err
}

func (r *redisRepo) Verification(req *pbu.VerificationUserRequest) (*pbu.User, error) {

	respUser, err := r.rConn.Get(context.Background(), req.Code).Result()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var newUser pbu.UserDetail
	if err := json.Unmarshal([]byte(respUser), &newUser); err != nil {
		log.Fatal(err)
		return nil, err
	}

	// fmt.Print(newUser.LastName, "helo")

	// Write Get data to postgres
	db, err := pkgdb.ConnectToDB(conf.Load())
	if err != nil {
		return nil, err
	}
	user, err := pgf.SubVerification(db, &newUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *redisRepo) Get(id string) (*pbu.UserDetail, error) {
	// get
	respUser, err := r.rConn.Get(context.Background(), id).Result()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var newUser pbu.UserDetail
	if err := json.Unmarshal([]byte(respUser), &newUser); err != nil {
		log.Fatal(err)
		return nil, err
	} else {
		fmt.Println(newUser)
	}

	// fmt.Print(newUser.FirstName,"\n")
	// fmt.Print(newUser.LastName, "hello")

	return &newUser, err
}
