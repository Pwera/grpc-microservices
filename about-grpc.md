# gRPC

Is a free and open-source framework developed by Google, part of CNCF.
Allows to define request and response for RPC (Remote Procedure Calls) and handles all the rest for you.
On top of it, it's modern, fastm efficient, build on top of HTTP/2, low latency, support streaming, language independent.
It's used by :
- Google
- Netflix
- Square
- CoreOS
- CoackroachDB

# Samples:
 1)  Basic [https://github.com/Pwera/gRPC-Notes/tree/master/src/basic] 
 Golang & Rust & Java clients and servers.
 2) TBA
#  

## HTTP/2
gRPC levarages HTPP/2 as a backbone for communications, eg https://imagekit.io/demo/http2-vs-http1.
HTTP/2 is the newer standard for internet communications that address common pitfall of HTTP/1.1 on modern web pages
HTTP/1.1 open a new TCP connection to a server at each request
It does not compress headers (which are plaintext- heavy size)
It only works with Request/Response mechanism (no server push)

HTTP/2 was released in 2015, supports multiplexing
The client & server can push messages in parallel over the same TCP connection, this greatly reduces latency.

HTTP/2 supports server push, servers can push streams for one request from the client. It allso supports header compression.

HTTP 1.1 also has a severe limitation of how a single connection can be used for multiple requests: all requests must be sent back in the order the corresponding requests were received. So clients that use pipelining will see head-of-line blocking delays, but the later responses may have been computed quickly, and must wait for earlier responses to be computed and transmitted before they can be sent. And the other alternative, using a connection for only one request at a time and then using a pool of connections to issue parallel requests, consumes more resources in both clients and servers as well as potentially in between proxies and load balancers. HTTP/2 and Protocol Buffers do not have these problems. 


## Types of API in gRPC
- Unary
- Server Streaming
- Client Streaming
- BI Directional Streaming

## Scalability in gRPC
- gRPC servers are asynchronous by default
- This means they do not blosk threads on request.
- Therefore each gRPC server can serve milion of requests in parallel.
- gRPC clients can be asynchronous or synchronous
- The client decides which model works best
- The gRPC clients can perform client side load balancing. Allows to scale horizontally.

## Security in gRPC
Each language will provide an API to load gRPC with the required certificates and provide encryption capability out  of the box
Additionally using Interceptors, we can also provide authentication

## Unary 
Unary is what a traditional API looks like (http rest)
``` protobuf
service GreetService {
     rpc Greet(GreetRequest) returns (GreetResponse){};
}
```

## Architecture of gRPC is layered:
- he lowest layer is the transport: gRPC uses HTTP/2 as its transport protocol. HTTP/2 provides the same basic semantics as HTTP 1.1 (the version with which nearly all developers are familiar), but aims to be more efficient and more secure. 
- The next layer is the channel. This is a thin abstraction over the transport. The channel defines calling conventions and implements the mapping of an RPC onto the underlying transport. At this layer, a gRPC call consists of a client-provided service name and method name, optional request metadata (key-value pairs), and zero or more request messages. A call is completed when the server provides optional response header metadata, zero or more response messages, and response trailer metadata. The trailer metadata indicates the final disposition of the call: whether it was a success or a failure. At this layer, there is no knowledge of interface constraints, data types, or message encoding. A message is just a sequence of zero or more bytes. A call may have any number of request and response messages.
- The last layer is the stub. The stub layer is where interface constraints and data types are defined. Does a method accept exactly one request message or a stream of request messages? What kind of data is in each response message and how is it encoded? The answers to these questions are provided by the stub. The stub marries the IDL-defined interfaces to a channel. The stub code is generated from the IDL. The channel layer provides the ABI that these generated stubs use.


Another key component of gRPC is a technology called Protocol Buffers. Protocol Buffers, or “protobufs” for short, are an IDL for describing services, methods, and messages. A compiler turns IDL into generated code for a wide variety of programming languages, along with runtime libraries for each of those supported programming languages. You’ll learn much more about Protocol Buffers in the next chapter.

It is important to note that Protocol Buffers have a role only in the last layer in the list above: the stub. The lower layers of gRPC, the channel and the transport, are IDL-agnostic. This makes it possible to use any IDL with gRPC (though the core gRPC libraries only provide tools for using protobufs). You can even find unofficial open source implementations of the stub layer using other formats and IDLs, such as flatbuffers and messagepack.

Not only does gRPC support streaming, but it also supports full-duplex bidirectional streams. Bidirectional means that the client can use a stream to upload an arbitrary amount of request data and the server can use a stream to send back an arbitrary amount of response data, all in the same RPC. The novel part is the “full-duplex” part. Most request-response protocols, including HTTP 1.1 are “half-duplex.” They support bidirectional communication (HTTP 1.1 even supports bidirectional streaming), but the two directions cannot be used at the same time. A request must first be fully uploaded before the server begins responding; only after the client is done transmitting can the server then reply with its full response. gRPC is built on HTTP/2, which explicitly supports full-duplex streams, which means that the client can upload request data at the same time the server is sending back response data. This is very powerful and eliminates the need for things like web sockets, which is an extension of HTTP 1.1, to allow full-duplex communication over an HTTP 1.1 connection. Thanks to streaming, applications can build very sophisticated conversational protocols on top of gRPC.

Forget the days of manually wiring up server handlers based on URI paths and then manually marshalling paths, query string parameters, requests, and response bodies. Similarly, forget the days of manually creating HTTP request objects, with all of the same overhead on the server side.

Metadata is broken up into two categories: headers, which are sent before the response message data, and trailers, which are sent at the end (along with the gRPC status code for the operation).

## Server streaming
In gRPC server streaming calls are defined using keyword "stream"

``` protobuf
service GreetService {
     rpc Greet(GreetRequest) returns (stream GreetResponse){};
}
``` 

## Client streaming
In gRPC client streaming calls are defined using keyword "stream".
The client will send many message to the server and will receive one response form the server, at any time.

``` protobuf
service GreetService {
     rpc ManyGreetRequest(stream GreetRequest) returns (GreetResponse){};
}
```

## Bi Directional streaming

``` protobuf
service GreetService {
     rpc ManyGreetRequest(stream GreetRequest) returns (stream GreetResponse){};
}
```

## Default values for new fields
Any new fields added to protobuf messages will be given a default valuethat’s specific to the type. For example, a new int32 field will have a val-ue of 0, a string will have "", and a message field will be nil (or the lan-guage equivalent).For a full list of the default values for each field type, see the protobufdocs         at         https://developers.google.com/protocol-buffers/docs/proto3#default.This is the standard behavior for proto3. Working against the protocolcan lead to it’s own set of challenges and is generally not recommended.There  are  times,  however,  when  you  might  need  to  check  for  presenceand not just the default value.If it’s necessary, you can box primitive types so that you can check forpresence.  This  is  accomplished  by  using  wrappers.proto  from  the  Well-Known Types. There’s a wrapper for each of the defined primitive values.For example:

```protobuf
syntax = "proto3";
import "google/protobuf/wrapper.proto";
package practical_grpc.v1;
message SomeMessage {
//allows us to check for `nil` instead of just the default value  
google.protobuf.StringValue value = 1;  }
```

## gRPC Error codes
http://avi.im/grpc-errors/
If an application needs to return extra information on top of an error code, it can use the metadata context.
When  a  gRPC  server  responds,  it  includes  the  status  of  the  response  inthe  headers  portion  of  the  HTTP  response.  If  nothing  went  wrong  withthe request, the :status field is set on the response with the HTTP codeof 200. The gRPC status is encoded into the grpc-status field with the numeric  identifier.
Here  are  the  headers  that  will  be  set  on  the  response  for  the  originalcaller:
| Header   |       Usage
|----------|:-------------:|
| :status |  The HTTP status code |
| content-type |  The HTTP content type |
| grpc-status |  The gRPC status code |
| grpc-message |  A simple message included with the respons |
| grpc-status-details-bin |  An base64 encoded google.rpc.Status mes-sage |


| Code   |       Code 	Numeric Value      |  Description |
|----------|:-------------:|------:|
| OK |  0 | The zero value is the code used for successful operations. It means that no error occured. |
| Cancelled |    1   | This code is an acknowledgement of a client cancellation. If a client chooses to cancel an outstanding RPC, this is the error code that the server should record. When used correctly, a client will never actually see this code: the RPC has been cancelled, so any error code will not be delivered.   |
| Unknown | 2 |    This is the default error code. If the server can provideno information about the nature of the error, then thecause is unknown. Note that “Internal” is often a bet-ter choice when an unexpected error occurs |
| Invalid Argument | 3 |    This indicates that the request is malformed. When validating a request, if an argument is invalid and could never be valid, then this is the right code. If the argument is invalid due to the current state of the system, then “Failed Precondition” is the better choice. |
| Deadline Exceeded | 4 |    This code means that the client’s requested deadline has not been met. Like the “Cancelled” code, this error code should not usually be observed by clients: after their deadline has passed, the RPC will be automatically cancelled, and the client will not see any error response. It is instead more useful on the server for classifying failures. |
| Not Found | 5 |    This indicates that a request field refers to a resource that does not exist or instructs the server to query for information that does not exist. |
| Already Exists | 6 |    This indicates that a request has requested creation of a resource that already exists. This is a specialization of the “Failed Precondition” code. |
| Permission Denied | 7 |    This code means that the caller is not allowed to perform the operation. This should be due to authorization controls. If the caller cannot perform the operation due to a quota, use “Resource Exhausted” instead. If the caller is not allowed, but it might be if it were authenticated, use “Unauthenticated” instead. |
| Resource Exhausted | 8 |    The server has exhausted some resource and cannot complete the operation. This could be due to quota limitations for the caller, but it could also indicate some other resource, like disk space, has been exhausted. |
| Failed Precondition | 9 |    This code means the operation cannot be completed because some part of the system is not in the correct state. Validation that must query for a resource’s existing state is likely validating a precondition. If the state being checked is an optimistic concurrency checks, as part of a multi-operation sequence, “Aborted” is probably more appropriate. The caller typically must fix the system state before retrying. |
| Aborted | 10 |    This code means that the operation was aborted, such as due to an optimistic concurrency failure. This often means that retrying a multi-operation sequence may possibly correct the issue, but retrying just the one aborted step will not. |
| Out of Range | 11 |    This is a specialization of “Failed Precondition”. It means the precondition that failed was a bound check. It should be used in iteration operations to indicate that the end of the iteration has been reached. |
| Unimplemented | 12 |    The request operation is not implemented. This may be used by server’s that only expose part of a service, and not its full interface. |
| Internal | 13 |    This code means an internal error has occurred in the server. |
| Unavailable | 14 |    This code means that either the server or the requested operation is temporarily unavailable. The client may retry, with backoff, to see if it becomes available. |
| Data Loss | 15 |    This is a specialization of “Internal” and further conveys that some sort of unrecoverable loss or corruption of data has occurred. |
| Unauthenticated | 16 |    This code indicates the request operation requires authentication, but the request has either not included authentication credentials or has included invalid credentials. |

## Interceptors
There are scenarios where you may want the ability to inject in metadatabefore sending a request to a gRPC server. Another possible requirementis that you need to reject a request to a server based on some prerequi-site  such  as  authentication.  In  both  of  these  scenarios  you  might  use  apattern   called   “middleware”.   A   very   common   feature   of   applicationframeworks like Ruby on Rails is the ability to add and modify a middle-ware  stack.  This  includes  checking  sessions,  CSRF  protection,  or  evenjust compressing the response body before sending it back.
n  the  gRPC  world  this  functionality  is  called  “interceptors”.  An  inter-ceptor allows a client or server to wrap RPC requests and responses. This ncludes  streaming  and  unary  RPC  calls  as  well.  An  interceptor  has  theability  to  manipulate  almost  everything  about  a  procedure  call,  but  ismore  commonly  used  for  wrapping  around  logic  for  things  such  as  au-thentication, metrics, distributed tracing, etc.


```golang
func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handlergrpc.UnaryHandler) (out interface{}, err error) {    out, err = handler(ctx, req)    logrus.WithField("method", info.FullMethod).Info("handled rpc")return out, err}
```

```golang
func ErrorsInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (out interface{}, err error) {    out, err = handler(ctx, req)switch tErr := err.(type) {case NotFoundErr:return out, grpc.Errorf(codes.NotFound, tErr.Error())    }return out, err}
```
Interceptors are not only used for server side applications, they can alsobe added to gRPC client stubs. When an interceptor is added to a stub, itis  enacted  on  every  call  to  a  service  endpoint.  This  introduces  a  lot  ofpossibilities  for  client  side  communication,  such  as  retries,  distributedtracing  headers,  and  payload  normalization.  With  client  retries,  for  ex-ample,  we  can  add  a  timing  interceptor  that  times  how  long  each  RPCcall  takes  and  reports  it.

## Authentication
gRPC  provides  primitives  for  setting  up  TLS/SSL  between  clients  andservers.  It  also  provides  the  ability  to  setup  mutual  TLS  between  clientsand  servers!  However,  gRPC  doesn’t  provide  the  facility  of  authenticat-ing users/services easily. It is up to the developer to implement their ownauthentication  to  verify  services  have  provided  the  correct  credentialson RPC’s.

## gRPC Deadlines
https://grpc.io/blog/deadlines
In  gRPC,  clients  and  servers  can  specify  the  amount  of  time  a  requestmay take for both unary and streaming calls. The term used for this func-tionality  is  “timeout.”  The  implementation  of  the  timeout  is  languagespecific, but for the most part operates on the same concept.Timeouts  are  important  to  consider  if  you  are  working  with  an  appli-cation  that  may  take  a  long  time  to  process  a  call.  If  you  don’t  specify  atimeout  on  a  call,  it  can  potentially  operate  forever  (lets  say  your  appli-cation  accidentally  goes  into  an  infinite  loop),  causing  compute  resour-ces to eventually block.To  overcome  this,  you  specify  a  timeout  on  your  RPC  calls  for  howlong you are willing to wait for it to complete. If a timeout is not specifiedon a gRPC request, the call is timeout and is interpreted as infinite.When  a  client  specifies  a  timeout,  it  is  added  to  the  request  headersat the key grpc-timeout with a numeric value and unit. For example tospecify  a  timeout  of  1  minute,  the  header  on  the  request  will  be  grpc-timeout = 1M.
When a request has a designated timeout attached the server shouldhonor  the  timeout.  Client’s  may  also  cancel  the  call  if  they’ve  deter-mined  the  timeout  has  been  reached.  The  Ruby  client,  for  example,  willcancel the call client side when a timeout has been specified for a singleRPC even if the server is still processing it.If a client does not provide a timeout for the call that it is making, theserver may still provide it’s own deadline on the request. Server’s can in-dicate they have stopped processing the request due to a timeout by re-turning  with  a  gRPC  status  code  of  DeadlineExceeded  (Code  4).  How-ever,  a  server  may  still  have  completed  the  request  successfully.  Fromthe gRPC documentation on deadline exceeded codes:“DeadlineExceeded  means  operation  expired  before  completion.  Foroperations  that  change  the  state  of  the  system,  this  error  may  be  re-turned  even  if  the  operation  has  completed  successfully.  For  example,  asuccessful response from a server could have been delayed long enoughfor the deadline to expire.”Servers  and  Clients  are  independently  responsible  for  determining  ifan RPC is successful. If a client has set a timeout on an RPC call, the serv-er  may  have  completed  the  response  and  sent  it  back,  but  if  the  clientreceives   the   final   result   after   the   timeout,   it   will   error   with   DEAD-LINE_EXCEEDED.

## SSL Encryption in gRPC
https://grpc.io/docs/guides/auth/
In production gRPC calls should be running with encryption enabled.
gRPC can use both Encryption (1-way verification) and Authentication (2-wayverification)

## gRPC Reflection & CLI
We may want reflection for two reasons:
- Having servers "expose" which endpoints are available
- Allowing command line interfaces to talk to our server without have a preliminary .proto

## Load balancing
| Client load balancing   |       Proxy load balancing
|----------|:-------------:|
| The  main  advantage  of  client  load  balancing  is  obviously  its  perfor-mance.  No  middle  agents  between  client  and  servers  makes  it  optimalfor low latency constraints.As  you  can  imagine,  however,  client  load  balancing  adds  more  com-plexity  to  the  architecture.  The  client  is  much  more  sophisticated,  sinceit needs to apply strategies to equilibrate the traffic (Round Robin or oth-ers).The  client  will  need  to  keep  track  on  the  health  of  each  backend,  inorder  to  redirect  traffic  elsewhere  if  a  backend  malfunction  is  detected.This would be the thick client approach.Another  alternative  would  be  to  use  the  Lookaside  Load  Balancer,which  is  kind  of  a  Zookeeper/Consul/Eureka  who  communicates  the  cli-ent  that  is  the  best  backend  server  to  communicate  with. |  The  main  advantage  of  proxy  load  balancing  is  the  simplicity  of  the  cli-ent. It will only need a single endpoint to create a connection with. All ofthe  workload  problems,  the  security  issues  and  the  awareness  of  thehealth of every backend server, will be completely transparent to the cli-ent.Another  advantage  is  that  the  TLS  certificates  must  be  managed  di-rectly by the proxy. Backend servers don’t need to have secured connec-tions, which makes the architecture much less complex.In  return,  the  latency  will  be  increased  and  the  throughput  may  belimited by the capabilities of the proxy.There  are  two  kind  of  proxies  in  function  of  the  OSI  level  that  theywork on: 1) Transport Level: this option is dumber, but easier to implement.This is done on the TCP level, where the proxy just checks that thesocket is open, in order to know if the backend is up and running.This kind of proxy is very performant since there is no payloadtreatment. Client data is just copied to the backend connection.   2) Application Level: when the proxy needs to be a little bit smarterin terms of workload decisioning, this is the best option, in detri-ment of the added latency. On the application level, the HTTP/2protocol is parsed in order to inspect each request and make deci-sions on the fly. |
 
## CLI tools
Here  are  some  useful  tools  that  you  can  use  to  interact  with  a  running gRPC server on your local machine:
- grpcnode - CLI tool for quickly making servers and client, dynami-cally, in JavaScript
- grpcc - REPL gRPC command-line client
- grpc_cli - gRPC CLI tool
- Evans - Expressive universal gRPC (CLI) client
- grpcurl - Like cURL, but for gRPC: Command-line tool for interact-ing with gRPC servers
- danby - A grpc proxy for the browser
- docker-protoc - Dockerized protoc, grpc-gateway, and grpc_clicommands bundled with Google API libraries.
- prototool Useful “Swiss Army Knife” for processing proto files


## Evans CLI
``` bash
wget https://github.com/ktr0731/evans/releases/download/0.9.0/evans_linux_amd64.tar.gz
tar -xzf evans_linux_amd64.tar.gz
evans -p 50051 -r
show services
call Unary
```



## grpcurl

``` bash
##Installation
wget https://github.com/fullstorydev/grpcurl/releases/download/v1.7.0/grpcurl_1.7.0_linux_x86_64.tar.gz
tar -xzf grpcurl_1.7.0_linux_x86_64.tar.gz
## List available services, through reflection
grpcurl --plaintext 0.0.0.0:50051 list
grpcurl --plaintext 0.0.0.0:50051 describe
## Prepare dummy request body
grpcurl --plaintext -msg-template 0.0.0.0:50051 describe todo.GreetRequest
## Call gRPC request
grpcurl --plaintext  -d '{
  "greet": {
    "first": "foo",
    "second": "bar"
  }
}' 0.0.0.0:50051 todo.GreetService/Unary
## List available services, through .proto
grpcurl --import-path ../proto/customer -proto Customer.proto list
grpcurl --import-path ../proto/customer -proto Customer.proto list CustomerService
```


## HTTP
HTTP  (Hypertext  Transfer  Protocol)  is  the  most  expanded  way  to  maketwo  machines  communicate.  It  is  a  protocol  for  requesting  files  over  anetwork,  where  the  payload  (both  request  and  response)  is  transferredas plain text (ASCII).


Later on, HTTP/1.1 added some performance optimizations:

Keep  Alive  connections:  This  prevents  TCP  connections  from  beingclosed  every  time  that  a  response  is  sent.  In  the  headers,  you  can  set  atimeout  to  tell  how  long  an  idle  TCP  connection  may  be  kept  open,  aswell as the maximum number of requests that may be sent in one singleTCP connection.

Byte serving of Range requests: This allows you to request for a partof the hosted file that you want to download. For instance, if you need toshow a single page from a big .pdf file, you can ask only for the neededpage.

Request  pipelining:  This  allows  you  to  send  multiple  HTTP  requestswithin  the  same  TCP  connection,  without  being  forced  to  wait  for  theprecedent responses. That was one of the main handicaps of HTTP/1.0.Taking advantage of the Keep-Alive functionality, HTTP/1.1 makes itpossible  for  the  request  shooting  without  having  to  wait  for  the  re-sponse.

HTTP/2HTTP/2  was  born  to  accompany  the  new  needs  of  the  Internet’s  expo-nential  growth.  The  main  goals  are  optimizing  bandwidth,  lowering  la-tency,  and  allowing  a  higher  throughput.  The  semantics  and  use  casesshould not be affected.The  basic  improvement  of  HTTP/2  is  the  fact  that  it  isn’t  a  plain  textprotocol anymore, but it is binary. That of course decreases the payloadsent  all  over  the  network,  and  furthermore  it  will  allow  other  optimiza-tions in the protocol.Verbs  and  headers  stay  the  same,  and  the  only  thing  that  changes  isthe way data is encoded or serialized.HTTP/2  introduces  multiple  notions  and  they  will  be  detailed  in  thefollowing subsections.

HTTP2 verbs:
- Frames: frame in HTTP/2 is the smallest unit data representation. It contains:•
   - A frame Header that normally only contains the identifier of thestream that it belongs to.
   - The data content, which can have different formats depending onthe kind of data that it contains. Most of the frames may contain aPadding block at the end of the frame, and it is used to obfuscatethe  length  of  the  frame  for  security  reasons. 
- Messages: message is a sequence of frames representing a request or a response.
- Streams:  stream  is  a  bidirectional  flow  of  bytes  that  in  a  given  connection  maycarry one or more messages. Each stream is identified by an integer thatwill be written on the header of any frame.Note  that  in  a  single  TCP  connection,  several  streams  may  be  activeconcurrently at the same time.Streams   have   lifecycles   and   they   are   represented   by   transitionsamong states. These are all the possible states:
  - Idle: Initial state for any opened stream.
  - Open: State where both peers may send or receive frames at anymoment.
  - Reserved: One of the peers, having the stream in Idle state, sent aPUSH_PROMISE frame in order for the server to begin to push mes-sages to the client.
  - Half Closed: One of the peers finished to send frames.
  - Closed: Both peers agreed to terminate the connection.
- Multiplexing:  HTTP/1.x doesn’t allow you to do multiple parallelrequests  in  the  same  TCP  connection.  At  most,  with  HTTP/1.1  you’ll  beable to do multiple requests, but the responses will need to be receivedin the same order, inducing Header-of-Line blocking.HTTP/2  and  its  binary  convention  allows  you  to  multiplex  requestsand responses in the same connection so you can take full advantage ofthe network’s resources.
- Flow control: HTTP/2  multiplexed  frames  unleash  the  potential  to  take  full  advantageof  the  network  resources,  but  without  a  flow  control  you  lose  any  senseof traffic congestion. Any peer sending data needs to know the ingestioncapabilities  of  the  receiver,  otherwise  the  frames  can  be  lost.  If  the  re-ceiver  is  busy  doing  other  stuff,  it  needs  to  be  able  to  communicate  tothe sender to slow down the cadence.TCP  protocol  already  considers  flow  control  by  communicating  thereceive window of each peer. This is the equivalent buffer size of eachend point to hold incoming data. Each TCP ACK signal contains an upda-ted receive window in function of the receiver availability.The problem with TCP flow control, is that it doesn’t have enough Ap-plication Level granularity. Multiple streams in the same TCP connectionprevent  the  optimizing  of  the  receiving  flow  for  each  pair  of  stream-applications.So  in  addition  to  TCP  flow  control,  HTTP/2  uses  the  WINDOW_UPDATEframe   type   (type=0x8)   to   advertise   the   stream-application   receivewindow size.
- Server push: Server  push  is  the  capability  of  HTTP/2  to  initiate  the  transfer  of  dataframes  from  the  server  side,  instead  of  having  to  wait  for  the  request  ofthe  client.  This  is  obviously  a  big  win,  for  instance  when  browsers  needto  download  multiple  resources  from  a  webpage,  where  HTTP/1.x  isneeded to open one TCP connection for every one of them.PUSH_PROMISE frames are used for this purpose. They are sent by theserver  to  the  client,  and  the  client  is  free  to  accept  or  not,  and  even  de-cide  a  bunch  of  settings  via  the  SETTINGS  frame  on  how  the  download-ing of the resources will be done. The goal is that the client may have fullcontrol of the communication.
- HeadersIn  HTTP/1.x  headers  are  systematically  sent  in  any  request  or  response,in  plain  text.  This  is  a  big  waste  of  bandwidth,  that  HTTP/2  improves  byapplying HPACK compression on the HEADER frames.As  said  before,  the  purpose  of  HTTP/2  is  to  improve  performance  byreinventing the encoding, but keeping the syntax. This is to preserve theheaders as they are, and it’s just the way they will be transferred.HPACK compression uses three methods:
  - Static Dictionary:61 commonly used headers are predefinedwith default values, so that nothing is defined on the first request,and nothing is sent over the network.
  - Dynamic Dictionary: if new headers are sent over the transmis-sion, they are added into the dictionary in order to prevent send-ing them in further messages.
  - Huffman Encoding: an algorithm for the compression of ASCIIcharacters in the HTTP Headers context. Each character is associ-ated to a bit code. The more frequent a character will be, the short-er its bit code will be. You normally obtain an average of 37%smaller compressed texts.

