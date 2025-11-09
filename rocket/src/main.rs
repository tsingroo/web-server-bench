#[macro_use] extern crate rocket;

#[get("/ping")]
fn ping() -> &'static str {
    "pong"
}

#[get("/hello/<name>")]
fn hello(name: &str) -> String {
    format!("my name is{}", name)
}

#[launch]
fn rocket() -> _ {
    rocket::build()
        .mount("/", routes![ping, hello])
}
