---
title: Usage
expires_at: never
tags: [diego-release,bytefmt]
---

# Usage


```go
bytefmt.ByteSize(100.5*bytefmt.MEGABYTE) // returns "100.5M"
bytefmt.ByteSize(uint64(1024)) // returns "1K"
```

