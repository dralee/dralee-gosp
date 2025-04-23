/*
登录会话
2015.4.23 by dralee
*/
package basic

type Session struct {
	userIds []uint32
	users   map[uint32]User
}

var session *Session = &Session{
	userIds: make([]uint32, 0),
	users:   make(map[uint32]User),
}

func Context() *Session {
	return session
}

func (s *Session) Login(user User) {
	userId := user.GetId()
	u := user.GetUser()
	s.userIds = append(s.userIds, userId)
	s.users[userId] = u
}

func (s *Session) Logout(userId uint32) {
	for i, id := range s.userIds {
		if id == userId {
			s.userIds = append(s.userIds[:i], s.userIds[i+1:]...)
			delete(s.users, userId)
			break
		}
	}
}
