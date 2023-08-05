#[macro_use]
extern crate rocket;

use rocket::serde::{json::Json, Deserialize};

#[derive(Deserialize)]
#[serde(crate = "rocket::serde")]
struct URL<'r> {
    url: &'r str,
}
#[get("/")]
fn index() -> &'static str {
    "Welcome to the unfurling api"
}

#[post("/api/unfurl", format = "json", data = "<url>")]
fn api_handler(url: Json<URL<'_>>) -> &'_ str {
    let user_url = url.url;
    user_url
}

#[launch]
fn rocket() -> _ {
    rocket::build().mount("/", routes![index, api_handler])
}
