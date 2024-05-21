package process

import "fmt"

// 因为UserMgr实例在服务端只有一个
// 在很多地方会使用所以写成全局变量

var userMgr *UserMgr

type UserMgr struct {
	onlineUser map[int]*UserProcess
	UserId     int
}

func InitMgr() {
	userMgr = &UserMgr{
		onlineUser: make(map[int]*UserProcess, 1024),
	}
}

func (this *UserMgr) Add(up *UserProcess) {
	this.onlineUser[up.UserId] = up
}

func (this *UserMgr) DelOnlineUser(userId int) {
	delete(this.onlineUser, userId)
}

func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUser
}

func (this *UserMgr) GetAllOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := this.onlineUser[userId]
	if !ok {
		err = fmt.Errorf("用户%d,不存在", userId)
		return
	}
	return
}
