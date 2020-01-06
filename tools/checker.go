package tools

func CheckInfo(id string, pwd string) bool {
	userInfo := ReadYaml()

	if id == userInfo.ID {
		if pwd == userInfo.PASSWORD {
			return true
		} else if pwd != userInfo.PASSWORD {
			return false
		}
	} else {
		return false
	}

	return false
}
