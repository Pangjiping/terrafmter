package format

type Formatter interface {
	Format() error
	Cleanup()
}
