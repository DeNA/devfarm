package platforms

type DeviceFinder func(EitherDevice) (bool, error)
