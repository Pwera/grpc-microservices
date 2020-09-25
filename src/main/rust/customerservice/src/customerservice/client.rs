use pb::customer_service_client::CustomerServiceClient;
use pb::CustomerRequest;

pub mod hello_world {
    tonic::include_proto!("pb");
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut client = CustomerServiceClient::connect("http://[::1]:50051").await?;

    let request = tonic::Request::new(CustomerRequest {
        value: "Tonic".into(),
        id: 45,
    });

    let response = client.my_service_unary(request).await?;

    // let request2 = tonic::Request::new(CustomerRequest {
    //     value: "Tonic".into(),
    //     id: 45,
    // });

    // client.my_service_server_streaming(request2).await?;

    // client.my_service_client_streaming(request3).await?;

    println!("RESPONSE={:?}", response);

    Ok(())
}
