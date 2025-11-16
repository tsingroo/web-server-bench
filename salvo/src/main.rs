use salvo::prelude::*;

#[handler]
async fn ping() -> &'static str {
    "pong"
}

#[handler]
async fn hello(req: &mut Request) -> String {
    let name = req.param::<String>("name").unwrap_or_default();
    format!("my name is {}", name)
}

#[tokio::main]
async fn main() {
    let router = Router::new()
        .push(Router::with_path("ping").get(ping))
        .push(Router::with_path("hello/<name>").get(hello));

    let acceptor = TcpListener::new("0.0.0.0:8080").bind().await;
    Server::new(acceptor).serve(router).await;
}
