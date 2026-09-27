package configuration

type CliCommand struct {
	Name        string
	Description string
	Command     func(*Config) error
}

type LocationData struct{
	Offset uint16
	Page_limit uint16

}

type Config struct {
	AllCliCommands map[string]CliCommand
	LocationRequestState LocationData
}

var (
	CommandCallers []string // ordered callers
	ExitSignal chan struct{}
	SharedState Config
)

