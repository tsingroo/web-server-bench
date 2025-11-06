use actix_web::{web, App, HttpResponse, HttpServer, Responder};

async fn ping() -> impl Responder {
    HttpResponse::Ok().body("pong")
}

async fn hello(name: web::Path<String>) -> impl Responder {
    HttpResponse::Ok().body(format!("my name is {}", name))
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .route("/ping", web::get().to(ping))
            .route("/hello/{name}", web::get().to(hello))
    })
    .bind("127.0.0.1:8080")?
    .run()
    .await
}