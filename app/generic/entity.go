package generic

type IEntity interface {
	GetId() uint

	Validate() error
}
