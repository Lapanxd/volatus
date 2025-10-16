mod api;
mod commands;

use dotenv::dotenv;
use std::env;
use once_cell::sync::Lazy;
use crate::api::client::ApiClient;
use crate::commands::auth::login;
use crate::commands::auth::register;

pub static API_URL: Lazy<String> = Lazy::new(|| {
    env::var("API_URL").expect("API_URL must be set in .env")
});

pub static API_CLIENT: Lazy<ApiClient> = Lazy::new(|| ApiClient::new());

#[cfg_attr(mobile, tauri::mobile_entry_point)]
pub fn run() {
    dotenv().ok();
    tauri::Builder::default()
        .plugin(tauri_plugin_opener::init())
        .invoke_handler(tauri::generate_handler![login, register])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
