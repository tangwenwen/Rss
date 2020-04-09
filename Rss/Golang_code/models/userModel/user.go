package userModel

import (
	"RSs/dao/userDao"
	"RSs/plugins/etcd"
	"RSs/types/users"
	"encoding/json"
	"errors"
)

func Login(username,password string)(usersType.LoginResp,error){

	var userType =0
	var token = ""
	has,err:=userDao.UserExist(username)
	if err!=nil{
		return usersType.LoginResp{}, err
	}
	if has{
		userType,err=userDao.CheckPassword(username,password)
		if err!=nil{
			return usersType.LoginResp{}, err
		}else{
			token,err=userDao.CreateToken(username,userType)
		}
	}else{
		return usersType.LoginResp{}, errors.New("用户不存在")
	}
	return usersType.LoginResp{Token:token,UserType:userType},nil
}


func Logout(token string)error{
	err := userDao.DeleteToken(token)
	if err!=nil{
		return err
	}
	return nil
}

func GetInfo(token string) (*usersType.GetInfoResp,error){
	userInfo := new(usersType.GetInfoResp)
	data,status,err:=etcd.EtcdGet(etcd.ETCDROOT+"/"+token)
	if err!=nil{
		return userInfo,err
	}
	if status!=1{
		json.Unmarshal([]byte(data),userInfo)
	}
	return userInfo,nil
}