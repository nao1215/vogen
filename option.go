package vogen

// Option is a parameter to be specified when creating a new Vogen struct.
type Option func(*Vogen) error

// WithFilePath specifies the file path to be generated.
func WithFilePath(filePath string) Option {
	return func(vo *Vogen) error {
		if filePath == "" {
			return ErrInvalidFilePath
		}
		vo.filePath = filePath
		return nil
	}
}

// WithPackageName specifies the package name to be generated.
func WithPackageName(packageName string) Option {
	return func(vo *Vogen) error {
		if packageName == "" {
			return ErrInvalidPackageName
		}
		vo.packageName = packageName
		return nil
	}
}
