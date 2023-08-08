use std::path::PathBuf;

use rocket::serde::{json::Json, Deserialize};

#[derive(Deserialize)]
#[serde(crate = "rocket::serde")]
pub struct URL<'r> {
    url: &'r str,
}

#[derive(Deserialize)]
#[serde(crate = "rocket::serde")]
pub struct ReturnValue<'r> {
    message: &'r str,
}
#[get("/")]
pub fn index() -> &'static str {
    "Welcome to the unfurling api"
}

#[post("/api/unfurl", format = "json", data = "<url>")]
pub fn api_handler(url: Json<URL<'_>>) -> &'_ str {
    let user_url = url.url;
    user_url
}

#[get("/apis/<path..>")]
pub fn mult_path(path: PathBuf) {
    println!("{:?}", path);
}
