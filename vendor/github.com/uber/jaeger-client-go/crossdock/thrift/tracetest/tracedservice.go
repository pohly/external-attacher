// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package tracetest

import (
	"bytes"
	"fmt"
	"github.com/uber/jaeger-client-go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type TracedService interface {
	// Parameters:
	//  - Request
	StartTrace(request *StartTraceRequest) (r *TraceResponse, err error)
	// Parameters:
	//  - Request
	JoinTrace(request *JoinTraceRequest) (r *TraceResponse, err error)
}

type TracedServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewTracedServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *TracedServiceClient {
	return &TracedServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewTracedServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *TracedServiceClient {
	return &TracedServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Request
func (p *TracedServiceClient) StartTrace(request *StartTraceRequest) (r *TraceResponse, err error) {
	if err = p.sendStartTrace(request); err != nil {
		return
	}
	return p.recvStartTrace()
}

func (p *TracedServiceClient) sendStartTrace(request *StartTraceRequest) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("startTrace", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := TracedServiceStartTraceArgs{
		Request: request,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *TracedServiceClient) recvStartTrace() (value *TraceResponse, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "startTrace" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "startTrace failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "startTrace failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1 error
		error1, err = error0.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "startTrace failed: invalid message type")
		return
	}
	result := TracedServiceStartTraceResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Request
func (p *TracedServiceClient) JoinTrace(request *JoinTraceRequest) (r *TraceResponse, err error) {
	if err = p.sendJoinTrace(request); err != nil {
		return
	}
	return p.recvJoinTrace()
}

func (p *TracedServiceClient) sendJoinTrace(request *JoinTraceRequest) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("joinTrace", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := TracedServiceJoinTraceArgs{
		Request: request,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *TracedServiceClient) recvJoinTrace() (value *TraceResponse, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "joinTrace" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "joinTrace failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "joinTrace failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error3 error
		error3, err = error2.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error3
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "joinTrace failed: invalid message type")
		return
	}
	result := TracedServiceJoinTraceResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type TracedServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      TracedService
}

func (p *TracedServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *TracedServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *TracedServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewTracedServiceProcessor(handler TracedService) *TracedServiceProcessor {

	self4 := &TracedServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self4.processorMap["startTrace"] = &tracedServiceProcessorStartTrace{handler: handler}
	self4.processorMap["joinTrace"] = &tracedServiceProcessorJoinTrace{handler: handler}
	return self4
}

func (p *TracedServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x5 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x5.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x5

}

type tracedServiceProcessorStartTrace struct {
	handler TracedService
}

func (p *tracedServiceProcessorStartTrace) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := TracedServiceStartTraceArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("startTrace", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := TracedServiceStartTraceResult{}
	var retval *TraceResponse
	var err2 error
	if retval, err2 = p.handler.StartTrace(args.Request); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing startTrace: "+err2.Error())
		oprot.WriteMessageBegin("startTrace", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("startTrace", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tracedServiceProcessorJoinTrace struct {
	handler TracedService
}

func (p *tracedServiceProcessorJoinTrace) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := TracedServiceJoinTraceArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("joinTrace", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := TracedServiceJoinTraceResult{}
	var retval *TraceResponse
	var err2 error
	if retval, err2 = p.handler.JoinTrace(args.Request); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing joinTrace: "+err2.Error())
		oprot.WriteMessageBegin("joinTrace", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("joinTrace", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Request
type TracedServiceStartTraceArgs struct {
	Request *StartTraceRequest `thrift:"request,1" json:"request"`
}

func NewTracedServiceStartTraceArgs() *TracedServiceStartTraceArgs {
	return &TracedServiceStartTraceArgs{}
}

var TracedServiceStartTraceArgs_Request_DEFAULT *StartTraceRequest

func (p *TracedServiceStartTraceArgs) GetRequest() *StartTraceRequest {
	if !p.IsSetRequest() {
		return TracedServiceStartTraceArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *TracedServiceStartTraceArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *TracedServiceStartTraceArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TracedServiceStartTraceArgs) readField1(iprot thrift.TProtocol) error {
	p.Request = &StartTraceRequest{}
	if err := p.Request.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
	}
	return nil
}

func (p *TracedServiceStartTraceArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("startTrace_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TracedServiceStartTraceArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err)
	}
	if err := p.Request.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err)
	}
	return err
}

func (p *TracedServiceStartTraceArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TracedServiceStartTraceArgs(%+v)", *p)
}

// Attributes:
//  - Success
type TracedServiceStartTraceResult struct {
	Success *TraceResponse `thrift:"success,0" json:"success,omitempty"`
}

func NewTracedServiceStartTraceResult() *TracedServiceStartTraceResult {
	return &TracedServiceStartTraceResult{}
}

var TracedServiceStartTraceResult_Success_DEFAULT *TraceResponse

func (p *TracedServiceStartTraceResult) GetSuccess() *TraceResponse {
	if !p.IsSetSuccess() {
		return TracedServiceStartTraceResult_Success_DEFAULT
	}
	return p.Success
}
func (p *TracedServiceStartTraceResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *TracedServiceStartTraceResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TracedServiceStartTraceResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &TraceResponse{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *TracedServiceStartTraceResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("startTrace_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TracedServiceStartTraceResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *TracedServiceStartTraceResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TracedServiceStartTraceResult(%+v)", *p)
}

// Attributes:
//  - Request
type TracedServiceJoinTraceArgs struct {
	Request *JoinTraceRequest `thrift:"request,1" json:"request"`
}

func NewTracedServiceJoinTraceArgs() *TracedServiceJoinTraceArgs {
	return &TracedServiceJoinTraceArgs{}
}

var TracedServiceJoinTraceArgs_Request_DEFAULT *JoinTraceRequest

func (p *TracedServiceJoinTraceArgs) GetRequest() *JoinTraceRequest {
	if !p.IsSetRequest() {
		return TracedServiceJoinTraceArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *TracedServiceJoinTraceArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *TracedServiceJoinTraceArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TracedServiceJoinTraceArgs) readField1(iprot thrift.TProtocol) error {
	p.Request = &JoinTraceRequest{}
	if err := p.Request.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
	}
	return nil
}

func (p *TracedServiceJoinTraceArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("joinTrace_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TracedServiceJoinTraceArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err)
	}
	if err := p.Request.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err)
	}
	return err
}

func (p *TracedServiceJoinTraceArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TracedServiceJoinTraceArgs(%+v)", *p)
}

// Attributes:
//  - Success
type TracedServiceJoinTraceResult struct {
	Success *TraceResponse `thrift:"success,0" json:"success,omitempty"`
}

func NewTracedServiceJoinTraceResult() *TracedServiceJoinTraceResult {
	return &TracedServiceJoinTraceResult{}
}

var TracedServiceJoinTraceResult_Success_DEFAULT *TraceResponse

func (p *TracedServiceJoinTraceResult) GetSuccess() *TraceResponse {
	if !p.IsSetSuccess() {
		return TracedServiceJoinTraceResult_Success_DEFAULT
	}
	return p.Success
}
func (p *TracedServiceJoinTraceResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *TracedServiceJoinTraceResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TracedServiceJoinTraceResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &TraceResponse{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *TracedServiceJoinTraceResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("joinTrace_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TracedServiceJoinTraceResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *TracedServiceJoinTraceResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TracedServiceJoinTraceResult(%+v)", *p)
}
