package process

import "fmt"

//userMgr实例很多地方都会使用
//而且,服务器端有且只有一个
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

func (userMgr *UserMgr) AddOnlineUser(up *UserProcess) {
	userMgr.onlineUsers[up.UserId] = up
}

func (userMgr *UserMgr) DelOnlineUser(userId int) {
	delete(userMgr.onlineUsers, userId)
}

func (userMgr *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return userMgr.onlineUsers
}

//根据id返回对应的值
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {

	//如何从map取出一个值，带检测方式
	up, ok := this.onlineUsers[userId]
	if !ok { //说明，你要查找的这个用户，当前不在线。
		err = fmt.Errorf("用户%d 不存在", userId)
		return
	}
	return
}
