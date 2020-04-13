package userDao

import (
	"RSs/plugins/db"
	"RSs/plugins/etcd"
	"RSs/types/users"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

func CreateToken(username string, userType int) (string, error) {
	tokenkey := fmt.Sprintf("%s%d", username, time.Now().Unix())
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(120)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["jti"] = tokenkey
	token.Claims = claims

	tokenString, err := token.SignedString([]byte("hello"))
	if err != nil {
		return "", err
	}
	err = writeToken(tokenString, username, userType)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func writeToken(tokenString, username string, userType int) error {
	userInfo := usersType.TokenValue{
		UserName: username,
		UserType: strconv.Itoa(userType),
	}
	data, err := json.Marshal(userInfo)
	if err != nil {
		return err
	}
	err = etcd.EtcdPutLease(etcd.ETCDROOT+"/"+tokenString, string(data), 999993600)
	if err != nil {
		return err
	}
	return nil
}
func DeleteToken(tokenkey string) error {
	_, err := etcd.EtcdDel(etcd.ETCDROOT + "/" + tokenkey)
	if err != nil {
		return err
	}
	return err
}
func UserExist(username string) (bool, error) {
	orm, err := db.GetEngine()
	if err != nil {
		return false, err
	}
	user := new(usersType.Users)
	user.Username = username
	return orm.Exist(user)
}
func CheckPassword(username, password string) (int, error) {
	orm, err := db.GetEngine()
	if err != nil {
		return -1, err
	}
	sql := fmt.Sprintf("select password,user_type from users where user_name = '%s'", username)
	result, err := orm.QueryString(sql)
	if err != nil {
		return -1, err
	}
	if len(result) < 1 {
		return -1, errors.New("errors")
	} else {
		if password == result[0]["password"] {
			userType, _ := strconv.Atoi(result[0]["user_type"])
			return userType, nil
		} else {
			return -1, errors.New("密码不正确")
		}
	}
}

func IncreaseUser(user *usersType.AddUserReq) error {
	orm, err := db.GetEngine()
	if err != nil {
		return err
	}
	age,_:=strconv.ParseInt(user.Age,10,64)
	_, err = orm.InsertOne(&usersType.Users{
		Username:    user.UserName,
		Password:    user.Password,
		Age:         int(age),
		Occupation:  user.Occupation,
		Gender:      user.Gender,
		CreatedTime: time.Now(),
		Status:      0,
		UserType:    0,
	})
	if err != nil {
		return err
	}
	return nil
}


func GetUserIdByUserName(userName string)(int,error){
	orm, err := db.GetEngine()
	if err != nil {
		return 0,err
	}
	user :=new(usersType.Users)
	_,err = orm.Where("user_name = ?",userName).Get(user)
	if err!=nil{
		return 0,err
	}
	return user.ID,nil
}