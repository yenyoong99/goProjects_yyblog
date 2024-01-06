package user

// 存放 需要入口的数据结构(PO)

// User 用户创建成功后返回一个User对象
// CreatedAt 为啥没用time.Time, int64(TimeStamp), 统一标准化, 避免时区你的程序产生影响
// 在需要对时间进行展示的时候，由前端根据具体展示那个时区的时间
type User struct {
	// 用户Id
	Id int
	// 创建时间, 时间戳 10位, 秒
	CreatedAt int64
	// 更新时间, 时间戳 10位, 秒
	UpdatedAt int64

	// 用户参数
	*CreateUserRequest
}
