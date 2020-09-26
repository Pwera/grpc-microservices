#[derive(Clone, PartialEq, ::prost::Message)]
pub struct BasicRequest {
    #[prost(string, tag = "1")]
    pub value: std::string::String,
    #[prost(int32, tag = "2")]
    pub id: i32,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct BasicResponse {
    #[prost(string, tag = "1")]
    pub status: std::string::String,
    #[prost(int32, tag = "2")]
    pub id: i32,
}
#[doc = r" Generated client implementations."]
pub mod basic_service_client {
    #![allow(unused_variables, dead_code, missing_docs)]
    use tonic::codegen::*;
    pub struct BasicServiceClient<T> {
        inner: tonic::client::Grpc<T>,
    }
    impl BasicServiceClient<tonic::transport::Channel> {
        #[doc = r" Attempt to create a new client by connecting to a given endpoint."]
        pub async fn connect<D>(dst: D) -> Result<Self, tonic::transport::Error>
        where
            D: std::convert::TryInto<tonic::transport::Endpoint>,
            D::Error: Into<StdError>,
        {
            let conn = tonic::transport::Endpoint::new(dst)?.connect().await?;
            Ok(Self::new(conn))
        }
    }
    impl<T> BasicServiceClient<T>
    where
        T: tonic::client::GrpcService<tonic::body::BoxBody>,
        T::ResponseBody: Body + HttpBody + Send + 'static,
        T::Error: Into<StdError>,
        <T::ResponseBody as HttpBody>::Error: Into<StdError> + Send,
    {
        pub fn new(inner: T) -> Self {
            let inner = tonic::client::Grpc::new(inner);
            Self { inner }
        }
        pub fn with_interceptor(inner: T, interceptor: impl Into<tonic::Interceptor>) -> Self {
            let inner = tonic::client::Grpc::with_interceptor(inner, interceptor);
            Self { inner }
        }
        #[doc = " Unary call"]
        pub async fn my_service_unary(
            &mut self,
            request: impl tonic::IntoRequest<super::BasicRequest>,
        ) -> Result<tonic::Response<super::BasicResponse>, tonic::Status> {
            self.inner.ready().await.map_err(|e| {
                tonic::Status::new(
                    tonic::Code::Unknown,
                    format!("Service was not ready: {}", e.into()),
                )
            })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static("/pb.BasicService/MyServiceUnary");
            self.inner.unary(request.into_request(), path, codec).await
        }
        #[doc = " Server Streaming"]
        pub async fn my_service_server_streaming(
            &mut self,
            request: impl tonic::IntoRequest<super::BasicRequest>,
        ) -> Result<tonic::Response<tonic::codec::Streaming<super::BasicResponse>>, tonic::Status>
        {
            self.inner.ready().await.map_err(|e| {
                tonic::Status::new(
                    tonic::Code::Unknown,
                    format!("Service was not ready: {}", e.into()),
                )
            })?;
            let codec = tonic::codec::ProstCodec::default();
            let path =
                http::uri::PathAndQuery::from_static("/pb.BasicService/MyServiceServerStreaming");
            self.inner
                .server_streaming(request.into_request(), path, codec)
                .await
        }
        #[doc = " Client Streaming"]
        pub async fn my_service_client_streaming(
            &mut self,
            request: impl tonic::IntoStreamingRequest<Message = super::BasicRequest>,
        ) -> Result<tonic::Response<super::BasicResponse>, tonic::Status> {
            self.inner.ready().await.map_err(|e| {
                tonic::Status::new(
                    tonic::Code::Unknown,
                    format!("Service was not ready: {}", e.into()),
                )
            })?;
            let codec = tonic::codec::ProstCodec::default();
            let path =
                http::uri::PathAndQuery::from_static("/pb.BasicService/MyServiceClientStreaming");
            self.inner
                .client_streaming(request.into_streaming_request(), path, codec)
                .await
        }
        #[doc = " BiDi Streaming"]
        pub async fn my_service_bi_di_streaming(
            &mut self,
            request: impl tonic::IntoStreamingRequest<Message = super::BasicRequest>,
        ) -> Result<tonic::Response<tonic::codec::Streaming<super::BasicResponse>>, tonic::Status>
        {
            self.inner.ready().await.map_err(|e| {
                tonic::Status::new(
                    tonic::Code::Unknown,
                    format!("Service was not ready: {}", e.into()),
                )
            })?;
            let codec = tonic::codec::ProstCodec::default();
            let path =
                http::uri::PathAndQuery::from_static("/pb.BasicService/MyServiceBiDiStreaming");
            self.inner
                .streaming(request.into_streaming_request(), path, codec)
                .await
        }
    }
    impl<T: Clone> Clone for BasicServiceClient<T> {
        fn clone(&self) -> Self {
            Self {
                inner: self.inner.clone(),
            }
        }
    }
    impl<T> std::fmt::Debug for BasicServiceClient<T> {
        fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
            write!(f, "BasicServiceClient {{ ... }}")
        }
    }
}
#[doc = r" Generated server implementations."]
pub mod basic_service_server {
    #![allow(unused_variables, dead_code, missing_docs)]
    use tonic::codegen::*;
    #[doc = "Generated trait containing gRPC methods that should be implemented for use with BasicServiceServer."]
    #[async_trait]
    pub trait BasicService: Send + Sync + 'static {
        #[doc = " Unary call"]
        async fn my_service_unary(
            &self,
            request: tonic::Request<super::BasicRequest>,
        ) -> Result<tonic::Response<super::BasicResponse>, tonic::Status>;
        #[doc = "Server streaming response type for the MyServiceServerStreaming method."]
        type MyServiceServerStreamingStream: Stream<Item = Result<super::BasicResponse, tonic::Status>>
            + Send
            + Sync
            + 'static;
        #[doc = " Server Streaming"]
        async fn my_service_server_streaming(
            &self,
            request: tonic::Request<super::BasicRequest>,
        ) -> Result<tonic::Response<Self::MyServiceServerStreamingStream>, tonic::Status>;
        #[doc = " Client Streaming"]
        async fn my_service_client_streaming(
            &self,
            request: tonic::Request<tonic::Streaming<super::BasicRequest>>,
        ) -> Result<tonic::Response<super::BasicResponse>, tonic::Status>;
        #[doc = "Server streaming response type for the MyServiceBiDiStreaming method."]
        type MyServiceBiDiStreamingStream: Stream<Item = Result<super::BasicResponse, tonic::Status>>
            + Send
            + Sync
            + 'static;
        #[doc = " BiDi Streaming"]
        async fn my_service_bi_di_streaming(
            &self,
            request: tonic::Request<tonic::Streaming<super::BasicRequest>>,
        ) -> Result<tonic::Response<Self::MyServiceBiDiStreamingStream>, tonic::Status>;
    }
    #[derive(Debug)]
    pub struct BasicServiceServer<T: BasicService> {
        inner: _Inner<T>,
    }
    struct _Inner<T>(Arc<T>, Option<tonic::Interceptor>);
    impl<T: BasicService> BasicServiceServer<T> {
        pub fn new(inner: T) -> Self {
            let inner = Arc::new(inner);
            let inner = _Inner(inner, None);
            Self { inner }
        }
        pub fn with_interceptor(inner: T, interceptor: impl Into<tonic::Interceptor>) -> Self {
            let inner = Arc::new(inner);
            let inner = _Inner(inner, Some(interceptor.into()));
            Self { inner }
        }
    }
    impl<T, B> Service<http::Request<B>> for BasicServiceServer<T>
    where
        T: BasicService,
        B: HttpBody + Send + Sync + 'static,
        B::Error: Into<StdError> + Send + 'static,
    {
        type Response = http::Response<tonic::body::BoxBody>;
        type Error = Never;
        type Future = BoxFuture<Self::Response, Self::Error>;
        fn poll_ready(&mut self, _cx: &mut Context<'_>) -> Poll<Result<(), Self::Error>> {
            Poll::Ready(Ok(()))
        }
        fn call(&mut self, req: http::Request<B>) -> Self::Future {
            let inner = self.inner.clone();
            match req.uri().path() {
                "/pb.BasicService/MyServiceUnary" => {
                    #[allow(non_camel_case_types)]
                    struct MyServiceUnarySvc<T: BasicService>(pub Arc<T>);
                    impl<T: BasicService> tonic::server::UnaryService<super::BasicRequest> for MyServiceUnarySvc<T> {
                        type Response = super::BasicResponse;
                        type Future = BoxFuture<tonic::Response<Self::Response>, tonic::Status>;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::BasicRequest>,
                        ) -> Self::Future {
                            let inner = self.0.clone();
                            let fut = async move { (*inner).my_service_unary(request).await };
                            Box::pin(fut)
                        }
                    }
                    let inner = self.inner.clone();
                    let fut = async move {
                        let interceptor = inner.1.clone();
                        let inner = inner.0;
                        let method = MyServiceUnarySvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = if let Some(interceptor) = interceptor {
                            tonic::server::Grpc::with_interceptor(codec, interceptor)
                        } else {
                            tonic::server::Grpc::new(codec)
                        };
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/pb.BasicService/MyServiceServerStreaming" => {
                    #[allow(non_camel_case_types)]
                    struct MyServiceServerStreamingSvc<T: BasicService>(pub Arc<T>);
                    impl<T: BasicService> tonic::server::ServerStreamingService<super::BasicRequest>
                        for MyServiceServerStreamingSvc<T>
                    {
                        type Response = super::BasicResponse;
                        type ResponseStream = T::MyServiceServerStreamingStream;
                        type Future =
                            BoxFuture<tonic::Response<Self::ResponseStream>, tonic::Status>;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::BasicRequest>,
                        ) -> Self::Future {
                            let inner = self.0.clone();
                            let fut =
                                async move { (*inner).my_service_server_streaming(request).await };
                            Box::pin(fut)
                        }
                    }
                    let inner = self.inner.clone();
                    let fut = async move {
                        let interceptor = inner.1;
                        let inner = inner.0;
                        let method = MyServiceServerStreamingSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = if let Some(interceptor) = interceptor {
                            tonic::server::Grpc::with_interceptor(codec, interceptor)
                        } else {
                            tonic::server::Grpc::new(codec)
                        };
                        let res = grpc.server_streaming(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/pb.BasicService/MyServiceClientStreaming" => {
                    #[allow(non_camel_case_types)]
                    struct MyServiceClientStreamingSvc<T: BasicService>(pub Arc<T>);
                    impl<T: BasicService> tonic::server::ClientStreamingService<super::BasicRequest>
                        for MyServiceClientStreamingSvc<T>
                    {
                        type Response = super::BasicResponse;
                        type Future = BoxFuture<tonic::Response<Self::Response>, tonic::Status>;
                        fn call(
                            &mut self,
                            request: tonic::Request<tonic::Streaming<super::BasicRequest>>,
                        ) -> Self::Future {
                            let inner = self.0.clone();
                            let fut =
                                async move { (*inner).my_service_client_streaming(request).await };
                            Box::pin(fut)
                        }
                    }
                    let inner = self.inner.clone();
                    let fut = async move {
                        let interceptor = inner.1;
                        let inner = inner.0;
                        let method = MyServiceClientStreamingSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = if let Some(interceptor) = interceptor {
                            tonic::server::Grpc::with_interceptor(codec, interceptor)
                        } else {
                            tonic::server::Grpc::new(codec)
                        };
                        let res = grpc.client_streaming(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/pb.BasicService/MyServiceBiDiStreaming" => {
                    #[allow(non_camel_case_types)]
                    struct MyServiceBiDiStreamingSvc<T: BasicService>(pub Arc<T>);
                    impl<T: BasicService> tonic::server::StreamingService<super::BasicRequest>
                        for MyServiceBiDiStreamingSvc<T>
                    {
                        type Response = super::BasicResponse;
                        type ResponseStream = T::MyServiceBiDiStreamingStream;
                        type Future =
                            BoxFuture<tonic::Response<Self::ResponseStream>, tonic::Status>;
                        fn call(
                            &mut self,
                            request: tonic::Request<tonic::Streaming<super::BasicRequest>>,
                        ) -> Self::Future {
                            let inner = self.0.clone();
                            let fut =
                                async move { (*inner).my_service_bi_di_streaming(request).await };
                            Box::pin(fut)
                        }
                    }
                    let inner = self.inner.clone();
                    let fut = async move {
                        let interceptor = inner.1;
                        let inner = inner.0;
                        let method = MyServiceBiDiStreamingSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = if let Some(interceptor) = interceptor {
                            tonic::server::Grpc::with_interceptor(codec, interceptor)
                        } else {
                            tonic::server::Grpc::new(codec)
                        };
                        let res = grpc.streaming(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                _ => Box::pin(async move {
                    Ok(http::Response::builder()
                        .status(200)
                        .header("grpc-status", "12")
                        .body(tonic::body::BoxBody::empty())
                        .unwrap())
                }),
            }
        }
    }
    impl<T: BasicService> Clone for BasicServiceServer<T> {
        fn clone(&self) -> Self {
            let inner = self.inner.clone();
            Self { inner }
        }
    }
    impl<T: BasicService> Clone for _Inner<T> {
        fn clone(&self) -> Self {
            Self(self.0.clone(), self.1.clone())
        }
    }
    impl<T: std::fmt::Debug> std::fmt::Debug for _Inner<T> {
        fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
            write!(f, "{:?}", self.0)
        }
    }
    impl<T: BasicService> tonic::transport::NamedService for BasicServiceServer<T> {
        const NAME: &'static str = "pb.BasicService";
    }
}
