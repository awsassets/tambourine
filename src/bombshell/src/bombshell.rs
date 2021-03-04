use tonic::{transport::Server, Request, Response, Status};

use bombshell::io_server::{Io, IoServer};
use bombshell::{Record, IoResponse, Error};

pub mod bombshell {
    tonic::include_proto!("bombshell"); // The string specified here must match the proto package name
}

#[derive(Debug, Default)]
pub struct IOInstance {}

#[tonic::async_trait]
impl Io for IOInstance {

    // ------- Set() -------
    async fn set(
        &self,
        request: Request<Record>, // <-- Input Record
    ) -> Result<Response<IoResponse>, Status> { // --> Output IOResponse
        println!("Got a request: {:?}", request);
        let ioresponse = bombshell::IoResponse {
            ok: true,
            code: 17,
            record: Some(bombshell::Record{
                name: format!("name {}", request.into_inner().name).into(),
                value: format!("value {}", request.into_inner().value).into(),
            }),
            error: Some(bombshell::Error {
                msg: format!("msg {}", request.into_inner().name).into(),
            }),
        };
        Ok(Response::new(ioresponse))
    }

    // ------- Get() -------
    async fn get(
        &self,
        request: Request<Record>, // <-- Input Record
    ) -> Result<Response<IoResponse>, Status> { // --> Output IOResponse
        println!("Got a request: {:?}", request);
        let ioresponse = bombshell::IoResponse {
            ok: true,
            code: 17,
            record: bombshell::Record{
                name: format!("name {}", request.into_inner().name).into(),
                value: format!("value {}", request.into_inner().value).into(),
            },
            error: bombshell::Error {
                msg: format!("msg {}", request.into_inner().name).into(),
            },
        };
        Ok(Response::new(ioresponse))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "[::1]:1313".parse()?;
    let io = IOInstance::default();

    Server::builder()
        .add_service(IoServer::new(io))
        .serve(addr)
        .await?;

    Ok(())
}