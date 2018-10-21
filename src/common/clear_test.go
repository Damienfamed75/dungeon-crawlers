package common

import (
	"runtime"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInitClear(t *testing.T) {
	Convey("When initClear is called", t, func() {
		initClear()
		Convey("clear should not equal nil", func() {
			So(clear, ShouldNotBeNil)
			So(clear["linux"], ShouldNotBeNil)
			So(clear["windows"], ShouldNotBeNil)
		})
	})
}

func TestClear(t *testing.T) {
	Convey("When Clear is called with real OS", t, func() {
		Convey("Then it should not panic", func() {
			So(Clear, ShouldNotPanic)
		})
	})

	Convey("When Clear is called with nil clear map", t, func() {
		clear = nil
		Convey("Then it should not panic", func() {
			So(Clear, ShouldNotPanic)
		})
	})

	defer func() { goos = func() string { return runtime.GOOS } }()

	Convey("When Clear is called with forced Windows", t, func() {
		goos = func() string { return "windows" }
		Convey("Then it should not panic", func() {
			So(Clear, ShouldNotPanic)
		})
	})

	Convey("When Clear is called with Linux", t, func() {
		goos = func() string { return "linux" }
		Convey("Then it should not panic", func() {
			So(Clear, ShouldNotPanic)
		})
	})

	Convey("When Clear is called with unsupported OS", t, func() {
		goos = func() string { return "test" }
		Convey("Then it should panic", func() {
			So(Clear, ShouldPanic)
		})
	})

}
