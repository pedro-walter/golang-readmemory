THIS IS A WORK IN PROGRESS, IT'S STILL BUGGY!

This code hangs up after running for a few minutes when monitoring ZSnes 1.42, maybe other processes too, need to find a way to fix it

# Usage

- `go get https://github.com/pedro-walter/golang-readmemory/`
- Import it:
```go
import (
	...
	"github.com/pedro-walter/golang-readmemory/readmemory"
)
```
- Initialize the handle, passing in the process name:
```go
handle := readmemory.NewHandle("zsnesw.exe")
```
- Use the read functions to grab data from the process, example based on Top Gear ROM and ZSnes 1.42:
```go
// this is where the ROM data is located on the process memory
var BASE_ADDRESS int64 = 0x2EECD0
p1_position := handle.ReadMemoryAtByte2(BASE_ADDRESS + 0x46)
p1_lap := handle.ReadMemoryAtByte1(BASE_ADDRESS + 0x1E76)
```

# Next milestones

- Add other read memory function that return other data types:
  - Little Endian vs Big Endian
  - Signed integers
  - Floats
  - Other memory sizes (usually 1, 2 or 4 bytes)

Based on a gist made by [@VityaSchel](https://gist.github.com/VityaSchel): https://gist.github.com/VityaSchel/32801695707420c2a22816a34a5b6cb5