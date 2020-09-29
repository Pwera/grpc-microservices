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
```
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

## Server streaming
In gRPC server streaming calls are defined using keyword "stream"

```
service GreetService {
     rpc Greet(GreetRequest) returns (stream GreetResponse){};
}
```

## Client streaming
In gRPC client streaming calls are defined using keyword "stream".
The client will send many message to the server and will receive one response form the server, at any time.

``` 
service GreetService {
     rpc ManyGreetRequest(stream GreetRequest) returns (GreetResponse){};
}
```

## Bi Directional streaming

``` 
service GreetService {
     rpc ManyGreetRequest(stream GreetRequest) returns (stream GreetResponse){};
}
```


## gRPC Error codes
http://avi.im/grpc-errors/
If an application needs to return extra information on top of an error code, it can use the metadata context.
| Code   |       Code 	Numeric Value      |  Description |
|----------|:-------------:|------:|
| OK |  0 | The zero value is the code used for successful operations. It means that no error occured. |
| Cancelled |    1   | This code is an acknowledgement of a client cancellation. If a client chooses to cancel an outstanding RPC, this is the error code that the server should record. When used correctly, a client will never actually see this code: the RPC has been cancelled, so any error code will not be delivered.   |
| Unknown | 2 |    cccccc |
| Invalid Argument | 2 |    cccccc |
| Unknown | 3 |    This indicates that the request is malformed. When validating a request, if an argument is invalid and could never be valid, then this is the right code. If the argument is invalid due to the current state of the system, then “Failed Precondition” is the better choice. |
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


## gRPC Deadlines
https://grpc.io/blog/deadlines
Deadlines allow gRPC clients to specify how long they are willing to wait for an RPC to complete before the RPC is terminated with the DEADLINE_EXEEDED.
The server should check if the deadline has exeeded and cancel the work it is doing.

## SSL Encryption in gRPC
https://grpc.io/docs/guides/auth/
In production gRPC calls should be running with encryption enabled.
gRPC can use both Encryption (1-way verification) and Authentication (2-wayverification)

## gRPC Reflection & CLI
We may want reflection for two reasons:
- Having servers "expose" which endpoints are available
- Allowing command line interfaces to talk to our server without have a preliminary .proto


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
