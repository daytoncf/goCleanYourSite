module github.com/daytoncf/goCleanSS

go 1.19

require (
	github.com/daytoncf/goCleanSS/css v0.0.0-unpublished
	golang.org/x/net v0.2.0
)

require github.com/daytoncf/goCleanSS/pkg/lib v0.0.0-unpublished

replace github.com/daytoncf/goCleanSS/css v0.0.0-unpublished => ./css

replace github.com/daytoncf/goCleanSS/pkg/lib v0.0.0-unpublished => ./pkg/lib
