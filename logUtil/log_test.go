package logUtil

import (
	"fmt"
	"testing"
)

func TestLog(t *testing.T) {
	InitLog()
	DebugLog("LoginDebug", "debug %dtimes", 7)
	InfoLog("LogInfoTest", "log_test", "TestLog", "testtesttsert")
	ErrLog("errTest", fmt.Errorf("Err Test"))

	//	time.Sleep(time.Second * 20)
	FailLog("%s,failTestLog", "TestLog")

}
