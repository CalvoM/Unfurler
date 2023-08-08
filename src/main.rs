#[macro_use]
extern crate rocket;

mod api_routes;
#[launch]
fn rocket() -> _ {
    rocket::build().mount(
        "/",
        routes![
            api_routes::index,
            api_routes::api_handler,
            api_routes::mult_path
        ],
    )
}
