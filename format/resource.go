package format

type Resource struct {
	AccessKey string
	SecretKey string
	Region    string
	Name      string
	Fields    map[string]Field
}

func NewResource(r string) (Formatter, error) {
	var (
		accKey string
		secKey string
		region string
		err    error
	)

	if accKey, err = getAccessKey(); err != nil {
		return nil, err
	}
	if secKey, err = getSecretKey(); err != nil {
		return nil, err
	}
	if region, err = getRegion(); err != nil {
		return nil, err
	}
	return &Resource{
		Name:      r,
		AccessKey: accKey,
		SecretKey: secKey,
		Region:    region,
		Fields:    make(map[string]Field),
	}, nil
}

func (r *Resource) Format() error {
	return nil
}

func (r *Resource) Cleanup() {

}
