package main

import (
	//"os"
	"io"
	"net"
	"fmt"
	"syscall"
	"log"
	"strings"
	"unsafe"
	"time"
	"errors"
	"strconv"
	otdd "otdd-test-runner/thrift/gen-go/otdd"
	"github.com/apache/thrift/lib/go/thrift"
	//"bytes"
	//"encoding/base64"
)

// use socket directly https://gist.github.com/jbenet/5c191d698fe9ec58c49d
// get original destination https://github.com/ryanchapman/go-any-proxy

const SO_ORIGINAL_DST = 80
var logfileName = "otdd-local-proxy.log"
var rootDirectory = "/home/otdd-local-proxy/otdd-local-proxy/"
var gTestCaseRunning = false
var passthroughConnections []string
var listenIpPort = "127.0.0.1:8746"
var redirectedConnections map[string]net.Conn
var passthroughContentFile = rootDirectory+"passthrough_content_file"
var passthroughContentFileChan = make(chan string,10)

var redirectIpPort = ""

/*
type TestCase struct {
	testId string
	runId string
	port int
	passthroughConnections[] string
	inboundRequest [] byte
}

type TestResult struct {
	testId string
	runId string
	inboundRequest [] byte
	inboundRequestErr string
	inboundResponse [] byte
	inboundResponseErr string
}
*/

type TestRunner struct {
	username string
	tag string
	macAddr string
	listenPort int
	otddServerHost string
	otddServerPort int
	currentTestCase *otdd.TestCase
	thriftClient *otdd.OtddTestRunnerServiceClient
	transport thrift.TTransport
}

func NewTestRunner(username string,tag string,macAddr string,listenPort int,
	otddServerHost string,otddServerPort int) *TestRunner{
	testRunner := &TestRunner{
                username:     username,
                tag:   tag,
                macAddr:   macAddr,
		listenPort:  listenPort,
                otddServerHost: otddServerHost,
                otddServerPort: otddServerPort,
        }
	return testRunner
}

func (t *TestRunner) Start() error {
	
	//listen on port to receive out-bound requests redirected by iptables.
	go t.listen()
	for {
		time.Sleep(10*time.Second)
		//fetch a test from otdd server.
		test,err := t.fetchTest(); 
		if err !=nil {
			continue
		}
		//run the test
		result := t.runTest(test)
		
		//report the test's result
		t.reportTestResult(result)

	}
}

func (t *TestRunner) getOtddThriftClient() (*otdd.OtddTestRunnerServiceClient,error) {

	if t.transport.IsOpen() {
		return t.thriftClient,nil
	}

	// connecto to otdd server
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(fmt.Sprintf("%s:%d",t.otddServerHost, t.otddServerPort))
	if err != nil {
		log.Println(fmt.Sprintf("error resolving address of otdd server %s:%v err:%v", t.otddServerHost, t.otddServerPort, err))
		return nil,err
	}
	t.transport = transport
	
	useTransport,err := transportFactory.GetTransport(t.transport)
	t.thriftClient = otdd.NewOtddTestRunnerServiceClientFactory(useTransport, protocolFactory)
	if err := t.transport.Open(); err != nil {
		log.Println(fmt.Sprintf("Error connect to otdd server %s:%v err:%v", t.otddServerHost, t.otddServerPort,err))
		return nil,err
	}
	return t.thriftClient,nil
	
}

func (t *TestRunner) fetchTest() (*otdd.TestCase,error) {
	log.Println(fmt.Sprintf("fetch test from otdd server: %s:%d",t.otddServerHost,t.otddServerPort))
	//t.thriftClient.fetchTest()
	return nil,errors.New("no test fetched.")
}

func (t *TestRunner) runTest(test *otdd.TestCase) *otdd.TestResult_ {
	log.Println(fmt.Sprintf("start to run test. test id: %s",test.TestId))
	result := &otdd.TestResult_ {
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%v", test.Port));
	if err !=nil {
		result.InboundRequestErr = err.Error()
		return result
	}
	defer conn.Close()
	t.setTestStarted(test)
	defer t.setTestStoped()
	conn.Write(test.InboundRequest)
	tmp := make([]byte, 2048)
	bytesRead := 0
	for{
                n, err := conn.Read(tmp);
                if err != nil {
                       	if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
				if bytesRead > 0 {
                               		return result
                        	} else {
					continue
				}
			} else if err == io.EOF {
				if bytesRead > 0 {
                               		return result
				} else {
					result.InboundResponseErr = err.Error()
					return result
				}
			} else { 
				result.InboundResponseErr = err.Error() 
				return result
			}
                }
		bytesRead += n
		conn.SetReadDeadline(time.Now().Add(10 * time.Millisecond))
		result.InboundResponse = append(result.InboundResponse,tmp[:n]...)
	}
	return result
}

func (t *TestRunner) setTestStarted(test *otdd.TestCase) {
	log.Println(fmt.Sprintf("test started. test id: %s",test.TestId))
	t.currentTestCase = test
}

func (t *TestRunner) setTestStoped() {
	if t.currentTestCase != nil {
		log.Println(fmt.Sprintf("test stopped. test id: %s",t.currentTestCase.TestId))
		t.currentTestCase = nil
	}
}

func (t *TestRunner) isTestRunning() bool {
	return t.currentTestCase != nil
}

func (t *TestRunner) reportTestResult(result *otdd.TestResult_) error{
	log.Println(fmt.Sprintf("report test result, test id: %s",result.TestId))
	return nil
}

func (t *TestRunner) listen() {
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d",t.listenPort))
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	log.Println(fmt.Sprintf("accepting connection on %d",t.listenPort))
	for {
		leftConn, err := listener.Accept()
		if err != nil {
			log.Println("accept error: " ,err)
			continue
		}
		go t.connHandler(leftConn)
	}
}

func (t *TestRunner) connHandler(conn net.Conn) {
	log.Println("connection accepted, local:",conn.LocalAddr()," remote:",conn.RemoteAddr())
        defer conn.Close()
	var originalConn net.Conn
	if !t.isTestRunning() || t.needPassthrough(conn) {
		if dst, dport, err := t.getOriginalDestination(conn); err == nil {
			log.Println(fmt.Sprintf("connecting to original destination on %s:%d",dst,dport))
			originalConn, err = net.Dial("tcp", fmt.Sprintf("%s:%v", dst, dport)); 
			if err != nil{
				log.Println(fmt.Sprintf("failed to connect to original destination %s:%d, err:%v",dst,dport,err))
				return
			}
			defer originalConn.Close()
		}
	}
	tmp := make([]byte, 2048)
	var accumulatedBytes [] byte
	bytesRead := 0
	nothingEverReceived := true
	connectButSendNothingCnt := 0
        for{
		conn.SetReadDeadline(time.Now().Add(10 * time.Millisecond))
                n, err := conn.Read(tmp);
                if err != nil {
                        if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
				if nothingEverReceived {
					connectButSendNothingCnt ++ 
				}
				tmpBytesRead := bytesRead
				bytesRead = 0
				if t.isTestRunning() && !t.needPassthrough(conn) { // fetch outbound response from otdd server.
					//only if the is outboundRequest bytes or connect but send nothing at the begining that need to fetch a matching outboundResponse 
					if len(accumulatedBytes)>0 || connectButSendNothingCnt > 10 {
						nothingEverReceived = false
						connectButSendNothingCnt = 0
						outboundResponse, err := t.fetchOutboundRespFromOtdd(accumulatedBytes[:tmpBytesRead]); 
						if err != nil {
							return
						}
						conn.Write(outboundResponse)
					}
					accumulatedBytes = accumulatedBytes[:0]
				} else { // send to original destination to get response.  
					outboundResponse, err := t.fetchOutboundRespFromConn(accumulatedBytes[:tmpBytesRead],originalConn);
					if err != nil {
						if err == io.EOF {
							return
						}
					}
					if len(outboundResponse)>0 {
						conn.Write(outboundResponse[:])
					}
					accumulatedBytes = accumulatedBytes[:0]
				}
                        } else {
                        	return
			}
                }
		bytesRead += n
		accumulatedBytes = append(accumulatedBytes,tmp[:n]...)
	}
}

func (t *TestRunner) needPassthrough(conn net.Conn) bool {
	if t.currentTestCase == nil {
		return false
	}

	dst, dport, err := t.getOriginalDestination(conn) 
	if err != nil {
		return false
	}

        for _,ipPort := range t.currentTestCase.PassthroughConnections {
                if ipPort == "" {
                        continue
                }
                if !strings.Contains(ipPort,":") {
                        if strings.Contains(ipPort,".") {//ip
                                if strings.Contains(dst,ipPort) {
                                        return true
                                }
                        } else {//port
                                port, _ := strconv.Atoi(ipPort)
                                if dport==uint16(port) {
					return true
                                }
                        }
                } else {
                        ip:=strings.Split(ipPort,":")[0]
                        port, _ := strconv.Atoi(strings.Split(ipPort,":")[1])
                        if strings.EqualFold(ip,dst) && dport==uint16(port) {
				return true
                        }
                }
        }

	return false
}

func (t *TestRunner) fetchOutboundRespFromOtdd(outbountRequests [] byte) ([] byte, error) {
	return nil,nil
}

func (t *TestRunner) fetchOutboundRespFromConn(outbountRequests [] byte, conn net.Conn) ([] byte, error) {

	if conn == nil {
		return nil, io.EOF
	}
	if len(outbountRequests) > 0 {
		log.Println(fmt.Sprintf("fetchOutboundRespFromConn:\n %s",string(outbountRequests[:])))
		conn.Write(outbountRequests[:])
	}

	tmp := make([]byte, 2048)
	var accumulatedBytes [] byte
	bytesRead := 0
        for{
		conn.SetReadDeadline(time.Now().Add(10 * time.Millisecond))
                n, err := conn.Read(tmp)
                if err != nil {
			return accumulatedBytes[:bytesRead],err
                }
		log.Println(fmt.Sprintf("read %s:",string(tmp[:n])))
		bytesRead += n
		accumulatedBytes = append(accumulatedBytes,tmp[:n]...)
	}
	return accumulatedBytes[:bytesRead],nil
}

func (t *TestRunner) getOriginalDestination(leftConn net.Conn) (string, uint16, error) {
        tcpConn := leftConn.(*net.TCPConn)

        /*
        //connection => file, will make a copy
        tcpConnFile, err := tcpConn.File()
        if err != nil {
                panic(err)
        } else {
                defer tcpConnFile.Close()
        }
        addr, err :=  syscall.GetsockoptIPv6Mreq(int(tcpConnFile.Fd()), syscall.IPPROTO_IP, SO_ORIGINAL_DST)
        */

        //https://zhuanlan.zhihu.com/p/22617140
        fdAddr := *(**int)(unsafe.Pointer(tcpConn))
        sysfd := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(fdAddr)) + 16))
        addr, err :=  syscall.GetsockoptIPv6Mreq(sysfd, syscall.IPPROTO_IP, SO_ORIGINAL_DST)
        if err != nil {
                //panic(err)
                return "", 0,err
        }

        dst := itod(uint(addr.Multiaddr[4])) + "." +
                itod(uint(addr.Multiaddr[5])) + "." +
                itod(uint(addr.Multiaddr[6])) + "." +
                itod(uint(addr.Multiaddr[7]))
        dport := uint16(addr.Multiaddr[2]) << 8 + uint16(addr.Multiaddr[3])
        return dst, dport, nil
}
// from pkg/net/parse.go
// Convert i to decimal string.
func itod(i uint) string {
        if i == 0 {
                return "0"
        }

        // Assemble decimal in reverse order.
        var b [32]byte
        bp := len(b)
        for ; i > 0; i /= 10 {
                bp--
                b[bp] = byte(i%10) + '0'
        }

        return string(b[bp:])
}

