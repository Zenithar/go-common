package api

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPaginator(t *testing.T) {
	Convey("Given a paginator (perPage: 50, page: 1)", t, func() {
		paginator := NewPaginator(1, 50)

		Convey("When there are 20 elements", func() {
			paginator.SetTotal(20)

			Convey("Then there is no pagination enabled", func() {
				So(paginator.Count(), ShouldEqual, 20)
				So(paginator.HasNext(), ShouldBeFalse)
				So(paginator.HasPrev(), ShouldBeFalse)
				So(paginator.HasOtherPages(), ShouldBeFalse)
				So(paginator.NumPages(), ShouldEqual, 1)
			})
		})

		Convey("When there are 52 elements", func() {
			paginator.SetTotal(52)

			Convey("Then pagination should be enabled", func() {
				So(paginator.Count(), ShouldEqual, 52)
				So(paginator.HasNext(), ShouldBeTrue)
				So(paginator.HasPrev(), ShouldBeFalse)
				So(paginator.HasOtherPages(), ShouldBeTrue)
				So(paginator.NumPages(), ShouldEqual, 2)
			})

		})
	})
}

func TestAnotherPaginator(t *testing.T) {
	Convey("Given another paginator (perPage: 50, page: 2)", t, func() {
		paginator := NewPaginator(2, 50)

		Convey("When there are 20 elements", func() {
			paginator.SetTotal(20)

			Convey("Then there is no pagination enabled", func() {
				So(paginator.Count(), ShouldEqual, 20)
				So(paginator.HasNext(), ShouldBeFalse)
				So(paginator.HasPrev(), ShouldBeTrue)
				So(paginator.HasOtherPages(), ShouldBeTrue)
				So(paginator.NumPages(), ShouldEqual, 1)
			})
		})

		Convey("When there are 52 elements", func() {
			paginator.SetTotal(52)

			Convey("Then pagination should be enabled", func() {
				So(paginator.Count(), ShouldEqual, 52)
				So(paginator.HasNext(), ShouldBeFalse)
				So(paginator.HasPrev(), ShouldBeTrue)
				So(paginator.HasOtherPages(), ShouldBeTrue)
				So(paginator.NumPages(), ShouldEqual, 2)
			})

		})
	})
}
