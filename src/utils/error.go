package utils

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/23 11:28
  @describe :
*/

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}
