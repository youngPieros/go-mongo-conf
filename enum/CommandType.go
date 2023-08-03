package enum

type CommandType string

const (
	Clean      CommandType = "clean"
	Delete     CommandType = "delete"
	Update     CommandType = "update"
	BadCommand CommandType = "bad-command"
)

func GetCommandTypeFrom(command string) CommandType {
	switch command {
	case string(Clean):
		return Clean
	case string(Delete):
		return Delete
	case string(Update):
		return Update
	default:
		return BadCommand
	}
}
