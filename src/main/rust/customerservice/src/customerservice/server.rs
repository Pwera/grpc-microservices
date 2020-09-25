use tonic::{transport::Server, Request, Response, Status, Streaming};

use pb::customer_service_server::{CustomerService, CustomerServiceServer};
use pb::{CustomerResponse, CustomerRequest};
use futures::Stream;
// use std::net::SocketAddr;
use std::pin::Pin;
// use tokio::sync::mpsc;

pub mod pb {
    tonic::include_proto!("pb");
}

type CustomerResult<T> = Result<Response<T>, Status>;
type ResponseStream = Pin<Box<dyn Stream<Item=Result<CustomerResponse, Status>> + Send + Sync>>;

#[derive(Default)]
pub struct MyGreeter {}

#[tonic::async_trait]
impl CustomerService for MyGreeter {
    async fn my_service_unary(
        &self,
        request: Request<CustomerRequest>,
    ) -> Result<Response<CustomerResponse>, Status> {
        println!("Got a request from {:?}", request.remote_addr());

        let reply = pb::CustomerResponse {
            status: format!("Hello {}!", "c"),
            id: 34,
        };
        Ok(Response::new(reply))
    }

    type MyServiceServerStreamingStream = ResponseStream;
    async fn my_service_server_streaming(
        &self,
        _: Request<CustomerRequest>,
    ) -> CustomerResult<Self::MyServiceServerStreamingStream> {
        Err(Status::unimplemented("not implemented"))
    }

    //MyServiceClientStreaming
    // type MyServiceClientStreamingStream = ResponseStream;

    async fn my_service_client_streaming(
        &self,
        _: Request<Streaming<CustomerRequest>>,
    ) -> CustomerResult<CustomerResponse> {
        Err(Status::unimplemented("not implemented"))
    }

    // MyServiceBiDiStreaming
    type MyServiceBiDiStreamingStream = ResponseStream;

    async fn my_service_bi_di_streaming(
        &self,
        _: Request<Streaming<CustomerRequest>>,
    ) -> CustomerResult<Self::MyServiceBiDiStreamingStream> {
        Err(Status::unimplemented("not implemented"))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "[::1]:50051".parse().unwrap();
    let greeter = MyGreeter::default();

    println!("GreeterServer listening on {}", addr);

    Server::builder()
        .add_service(CustomerServiceServer::new(greeter))
        .serve(addr)
        .await?;

    Ok(())
}
