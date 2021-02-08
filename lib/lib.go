package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type GetVersionRequest struct {
	Hostname string
}
type GetVersionReply struct {
	SomeReply  string
	OSVersion  int
	OSPlatform string
}

type GetHOstnameReply struct {
	Hostname string
}

type CC interface {
	GetV(context.Context, *GetVersionRequest, ...string) (*GetVersionReply, error)
	TheBigTest(string) bool
	TheBigTest2(string) bool
}

// fake implementation
type Fake struct {
}

func (x *Fake) GetV(ctx context.Context, gvr *GetVersionRequest, strs ...string) (*GetVersionReply, error) {
	fmt.Println("i'm in fake")
	return &GetVersionReply{
		SomeReply: "foo",
	}, nil
}

func (x *Fake) TheBigTest(theString string) bool {
	fmt.Println("i'm in fake")
	return true
}

func (x *Fake) TheBigTest2(theString string) bool {
	fmt.Println("i'm in fake")
	return true
}

// real implementation
type Cisco struct {
	HostName    string
	IOSVersion  int
	IOSPlatform string

	Fake
}

type CiscoNXAPIResponse struct {
	InsAPI struct {
		Type    string `json:"type"`
		Version string `json:"version"`
		Sid     string `json:"sid"`
		Outputs struct {
			Output struct {
				Body  interface{} `json:"body"`
				Input string      `json:"input"`
				Msg   string      `json:"msg"`
				Code  string      `json:"code"`
			} `json:"output"`
		} `json:"outputs"`
	} `json:"ins_api"`
}

type CiscoNXOShowSwitchnameOutput struct {
	Hostname string `json:"hostname"`
}

type ver struct {
	OsVersion string `json:"hostname"`
}

func (x *Cisco) doNxOSAPI(username string, password string, url string, o interface{}) error {
	req := http.Client()
	resp := req.do()
	json.Unmarshal(resp.Body, o)
}

func (x *Cisco) getVersion() error {
	fmt.Printf("SSh'ing to %s\n", x.HostName)
	// parse stuff
	x.IOSPlatform = "NxOS"
	x.IOSVersion = 111
	return nil
}

func (x *Cisco) GetV(ctx context.Context, gvr *GetVersionRequest, strs ...string) (reply *GetVersionReply, err error) {

	resp := &CiscoNXAPIResponse{}
	out := &CiscoNXOShowSwitchnameOutput{}
	resp.InsAPI.Outputs.Output.Body = out

	if err = x.doNxOSAPI("bob", "ff", "https....", resp); err != nil {
		return nil, fmt.Errorf("Something went wrong: %s", err.Error())
	}

	fmt.Println(out.Hostname)

	reply = &GetVersionReply{
		OSVersion:  x.IOSVersion,
		OSPlatform: x.IOSPlatform,
	}

	return reply, nil

}

func (x *Cisco) TheBigTest(theString string) bool {
	fmt.Println("i'm in Cisco")
	return true

}
func (x *Cisco) TheBigTest2(theString string) bool {
	fmt.Println("i'm in Cisco")
	return true
}

func (x *Cisco) Foo() {
	fmt.Println("foo")
}

func DoitWithCC(theCC CC) {
	rep := theCC.TheBigTest("aaa")
	fmt.Println("The response was %r", rep)

	c, ok := theCC.(*Cisco)
	if ok {
		c.Foo()
	}

	theCC.TheBigTest2("aa")
}
