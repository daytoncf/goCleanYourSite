module github.com/daytoncf/goCleanYourSite

go 1.19

replace github.com/daytoncf/goCleanYourSite/css v0.0.0-unpublished => ./css

replace github.com/daytoncf/goCleanYourSite/pkg/lib v0.0.0-unpublished => ./pkg/lib

require (
	github.com/daytoncf/goCleanYourSite/css v0.0.0-unpublished
	github.com/daytoncf/goCleanYourSite/pkg/lib v0.0.0-unpublished
	github.com/deckarep/golang-set/v2 v2.1.0
	golang.org/x/net v0.4.0
)
