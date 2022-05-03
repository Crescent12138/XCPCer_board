package luogu

import (
	_ "XCPCer_board/config"
	_ "XCPCer_board/db/mysql"
	"XCPCer_board/model"
	"testing"
)

///////////////////方法函数//////////////////////

//检查函数，若不一致返回1
func isIntMsgDifferent(funcRet, ans map[string]int) bool {
	for k, v := range ans {
		if _, r := funcRet[k]; r == false || v != funcRet[k] {
			return true
		}
	}
	return false
}
func isStringMsgDifferent(funcRet, ans map[string]string) bool {
	for k, v := range ans {
		if _, r := funcRet[k]; r == false || v != funcRet[k] {
			return true
		}
	}
	return false
}

//判断并输出错误
func checkIntError(t *testing.T, uid string, tp string, all func(uid string) (map[string]int, error),
	acInt map[string]int) {
	if ret, err := all(uid); isIntMsgDifferent(ret, acInt) {
		if err != nil {
			t.Errorf("Error of %v in all msg: %v", tp, err)
		}
		t.Errorf("Error of %v in all msg\n ret= %v  \nbut the ans is %v", tp, ret, acInt)
	}
}
func checkStrError(t *testing.T, uid string, tp string, all func(uid string) (map[string]string, error),
	acInt map[string]string) {
	if ret, err := all(uid); isStringMsgDifferent(ret, acInt) {
		if err != nil {
			t.Errorf("Error of %v in all msg: %v", tp, err)
		}
		t.Errorf("Error of %v in str msg\n ret= %v  \nbut the ans is %v", tp, ret, acInt)
	}
}

//SQL方法函数
func checkSqlAll(t *testing.T, uid string) {
	if err := InsertSql(uid); err != nil {
		t.Errorf("luogu InserSql uid=%v error:%v", uid, err)
	}
	if _, err := QuerySql(uid); err != nil {
		t.Errorf("luogu QuerySql uid=%v error:%v", uid, err)
	}
	if err := UpdateSql(uid); err != nil {
		t.Errorf("luogu UpdateSql uid=%v error:%v", uid, err)
	}
	if err := DeleteSql(uid); err != nil {
		t.Errorf("luogu DeleteSql uid=%v error:%v", uid, err)
	}
}

///////////////////测试函数//////////////////////
//用户测试
func UserTest(t *testing.T) {
	//开始测试
	checkIntError(t, model.TestLuoGuIdLYF, "luoGu", ScrapeAll, map[string]int{
		basicProblem:      30,
		elevatedProblem:   19,
		hardProblem:       6,
		passProblemNumber: 55,
		ranting:           94572,
		simpleProblem:     0,
		unKnowProblem:     0,
	})

}
func SqlTest(t *testing.T) {
	checkSqlAll(t, model.TestLuoGuIdLYF)
}

func TestLg(t *testing.T) {
	UserTest(t)
	SqlTest(t)
}
