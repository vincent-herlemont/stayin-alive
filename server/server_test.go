package server

import (
	"testing"
	"reflect"
)

const TEST_ADDR = ":8080"

func TestServerInfo(t *testing.T) {
	testServer := NewServer(&CfgWebServer{
		Addr: TEST_ADDR,
	})

	testServer.Start()
	output := &TOInfo{}
	err := testServer.GetClient().GetTO("/info",output)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(output,&TOInfo{Message:"alive !"}) {
		t.Error("Info are incorrect")
	}
	testServer.Stop()
}
//
//type TOTestInput struct {
//
//	string `json:"input"`
//}
//
//type TOTestOutput struct {
//	MessageOutput string `json:"output"`
//}
//
//
//func HandleTest(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	json.NewEncoder(w).Encode(TOInfo{Message:"alive !"})
//}
//
//func TestServerCustom(t *testing.T) {
//	testServer := NewServer(&CfgWebServer{
//		Addr: TEST_ADDR,
//	})
//	testServer.Init(func(r *mux.Router) {
//		r.HandleFunc("/send/{}",HandleTest)
//	})
//
//	testServer.Start()
//	testServer.Stop()
//}