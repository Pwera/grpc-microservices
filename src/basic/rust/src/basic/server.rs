mod pb;

use futures::{Stream, StreamExt};
use pb::basic_service_server::{BasicService, BasicServiceServer};
use pb::{BasicRequest, BasicResponse};
use std::env;
use std::error::Error;
use std::pin::Pin;
use tonic::{transport::Server, Request, Response, Status, Streaming};

type BasicResult<T> = Result<Response<T>, Status>;
type ResponseStream = Pin<Box<dyn Stream<Item=Result<BasicResponse, Status>> + Send + Sync>>;

#[derive(Default)]
pub struct BasicServer {}

#[tonic::async_trait]
impl BasicService for BasicServer {
    async fn my_service_unary(
        &self,
        request: Request<BasicRequest>,
    ) -> Result<Response<BasicResponse>, Status> {
        println!(
            "Got a request from {:?} through Unary request",
            request.remote_addr()
        );

        let reply = pb::BasicResponse {
            status: format!("Unary BasicResponse"),
            id: request.get_ref().id,
        };
        Ok(Response::new(reply))
    }

    type MyServiceServerStreamingStream = ResponseStream;
    async fn my_service_server_streaming(
        &self,
        request: Request<BasicRequest>,
    ) -> BasicResult<Self::MyServiceServerStreamingStream> {
        println!(
            "Got a request from {:?} through Server streaming request",
            request.remote_addr()
        );
        let responses = vec!["!", "@", "#", "$"]
            .iter()
            .map(|x| -> Result<pb::BasicResponse, E> {
                Ok(pb::BasicResponse {
                    status: format!("Server Streaming BasicResponse"),
                    id: request.get_ref().id,
                })
            })
            .collect::<Vec<Result<pb::BasicResponse, E>>>();

        let output = futures::stream::iter(responses);
        Ok(Response::new(
            Box::pin(output) as Self::MyServiceServerStreamingStream
        ))
    }

    async fn my_service_client_streaming(
        &self,
        request: Request<Streaming<BasicRequest>>,
    ) -> BasicResult<BasicResponse> {
        println!(
            "Got a request from {:?} through Client streaming request",
            request.remote_addr()
        );

        let mut stream = request.into_inner();
        while let Some(re) = stream.next().await {
            let re = re?;
            println!("Received {}", re.value);
        }

        let reply = pb::BasicResponse {
            status: format!("Client streaming BasicResponse {}!", "c"),
            id: 34,
        };
        Result::Ok(Response::new(reply))
    }

    type MyServiceBiDiStreamingStream = ResponseStream;

    async fn my_service_bi_di_streaming(
        &self,
        request: Request<Streaming<BasicRequest>>,
    ) -> Result<Response<Self::MyServiceBiDiStreamingStream>, Status> {
        println!(
            "Got a request from {:?} through Bidirectional streaming request",
            request.remote_addr()
        );

        let mut stream = request.into_inner();

        let output = async_stream::try_stream! {
            while let Some(re) = stream.next().await {
                let re = re?;
                println!("Received {}", re.value);

                let reply = pb::BasicResponse {
                    status: format!("Hello {}!", "c"),
                    id: 34,
                };
                yield  reply.clone();
            }
        };

        Ok(Response::new(
            Box::pin(output) as Self::MyServiceBiDiStreamingStream
        ))
    }
}

// Run with argument:  "127.0.0.1:50051"
#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    let args: Vec<String> = env::args().collect();
    let addr = (&args[1] as &str).parse().unwrap();
    let basicserver = BasicServer::default();

    println!("Server listening on {}", addr);

    Server::builder()
        .add_service(BasicServiceServer::new(basicserver))
        .serve(addr)
        .await?;

    Ok(())
}
