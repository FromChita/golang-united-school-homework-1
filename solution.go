package solution

import "github.com/kyokomi/emoji"
func GetMessage() string {
	msg := emoji.Sprint("Hello :world_map:!")
	return msg
}
