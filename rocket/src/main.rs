#[macro_use] extern crate rocket;

#[get("/ping")]
async fn ping() -> &'static str {
    "pong"
}

#[get("/hello/<name>")]
async fn hello(name: &str) -> String {
    format!("my name is{}", name)
}

#[launch]
fn rocket() -> _ {
    let figment = rocket::Config::figment()
        .merge(("port", 8080));
    
    rocket::custom(figment)
        .mount("/", routes![ping, hello])
}
