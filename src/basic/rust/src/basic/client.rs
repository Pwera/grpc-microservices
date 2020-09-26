mod pb;

use env::args;
use futures::stream;
use pb::basic_service_client::BasicServiceClient;
use pb::BasicRequest;
use std::thread::sleep;
use std::{env, time};
use time::Duration;
use tonic::Request;

// "http://[::1]:50051"
#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let args: Vec<String> = args().collect();
    let mut client = BasicServiceClient::connect(args[1].to_string()).await?;

    client
        .my_service_unary(BasicRequest {
            value: "Unary Request".into(),
            id: 1,
        })
        .await?;

    client
        .my_service_server_streaming(BasicRequest {
            value: "Server Streaming Request".into(),
            id: 2,
        })
        .await?;

    client
        .my_service_client_streaming(Request::new(stream::iter(vec![BasicRequest {
            value: "Client Streaming Request".into(),
            id: 3,
        }])))
        .await?;

    client
        .my_service_bi_di_streaming(Request::new(stream::iter(vec![BasicRequest {
            value: "Bi Di Streaming Request".into(),
            id: 3,
        }])))
        .await?;

    sleep(Duration::from_millis(4000));

    Ok(())
}
