# agora_rtc_sdk_cgo

Agora RTC SDK for GO

base on `Agora-RTC-x86_64-linux-gnu-v4.2.31-20240327_111832-296989`

base on [agora_rtc_sdk_c](https://github.com/Lensual/agora_rtc_sdk_c)

## How to use

1. Use `git submodule`

```bash
# On your project
git submodule add github.com/Lensual/agora_rtc_sdk_cgo
git submodule update --init --recursive
cd agora_rtc_sdk_cgo
cmake .
make
```

go.mod

```go.mod
require github.com/Lensual/agora_rtc_sdk_cgo v0.0.0  //TODO tagged version
replace github.com/Lensual/agora_rtc_sdk_cgo => ../agora_rtm_sdk_cgo
```

2. Or use `go mod`

```bash
go get -u github.com/Lensual/agora_rtc_sdk_cgo
```

NOTE: `go mod` is not support `git submodule`

```bash
# Provide `Lensual/agora_rtc_sdk_c` & `Agora RTC SDK` by yourself.

# Set ENV
CGO_CFLAGS=-I${AGORA_RTC_SDK_C_INCLUDE_PATH}
CGO_LDFLAGS=-L${AGORA_RTC_SDK_C_LIB_PATH} -L${AGORA_RTC_SDK_LIB_PATH}
LD_LIBRARY_PATH=${LD_LIBRARY_PATH}:${AGORA_RTC_SDK_C_LIB_PATH}:${AGORA_RTC_SDK_LIB_PATH}
```
