package format

type DataSource struct {
	AccessKey string
	SecretKey string
	Region    string
	Name      string
	Fields    map[string]Field
}

func NewDataSource(d string) (Formatter, error) {
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
	return &DataSource{
		Name:      d,
		AccessKey: accKey,
		SecretKey: secKey,
		Region:    region,
		Fields:    make(map[string]Field),
	}, nil
}

func (d *DataSource) Format() error {
	return nil
}

func (d *DataSource) Cleanup() {

}
