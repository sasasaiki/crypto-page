package sample

//GetFullName フルネームを返すサンプル
func GetFullName(vars map[string]string) string {
	f := vars["firstName"]
	l := vars["lastName"]
	return f + l
}
