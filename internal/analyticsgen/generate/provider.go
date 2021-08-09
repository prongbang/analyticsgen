package generate

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewIOSUseCase,
	NewFlutterUseCase,
	NewAndroidUseCase,
)
