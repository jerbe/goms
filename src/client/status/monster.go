package status

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/11/2 22:20
  @describe :
*/

var (
	// 武器攻击
	MonsterStatusWatk = &MonsterStatus{status: 0x10}

	MonsterStatusWdef = &MonsterStatus{status: 0x20}
)

type MonsterStatus struct {
	status int
	first  bool
}
