package GeoIP

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFetch(t *testing.T) {
	Convey("Given the IP `thisshouldbreak`", t, func() {
		var ip = "thisshouldbreak"
		res, err := Fetch(ip)

		Convey("error should *not* be nil", func() {
			So(err, ShouldNotBeNil)
		})

		Convey("country code should be empty", func() {
			So(res.CountryCode, ShouldBeBlank)
		})
	})

	Convey("Given the IP `8.8.8.8`", t, func() {
		var ip = "8.8.8.8"
		res, err := Fetch(ip)

		Convey("error should be nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("country code should be `US`", func() {
			So(res.CountryCode, ShouldEqual, "US")
		})

		Convey("the region should be `California`", func() {
			So(res.RegionName, ShouldEqual, "California")
		})

	})

	Convey("Given the IP `213.133.107.227`", t, func() {
		var ip = "213.133.107.227"
		res, err := Fetch(ip)

		Convey("error should be nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("country code should be `DE`", func() {
			So(res.CountryCode, ShouldEqual, "DE")
		})
	})
}
