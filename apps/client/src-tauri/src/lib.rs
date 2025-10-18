mod api;
mod commands;

use dotenv::dotenv;
use once_cell::sync::Lazy;
use std::env;
use tauri_plugin_store::StoreExt;

use crate::api::client::ApiClient;
use crate::commands::auth::{login, logout, register};
use crate::commands::handshake::{get_pending_handshakes, handshake_init, handshake_response};
use crate::commands::user::get_me;

pub static API_URL: Lazy<String> =
    Lazy::new(|| env::var("API_URL").expect("API_URL must be set in .env"));

pub static API_CLIENT: Lazy<ApiClient> = Lazy::new(|| ApiClient::new());

#[cfg_attr(mobile, tauri::mobile_entry_point)]
pub fn run() {
    dotenv().ok();
    tauri::Builder::default()
        .plugin(tauri_plugin_store::Builder::new().build())
        .plugin(tauri_plugin_opener::init())
        .setup(|app| {
            let store = app.store("store.json")?;

            if let Some(token) = store
                .get("auth_token")
                .and_then(|v| v.as_str().map(|s| s.to_string()))
            {
                API_CLIENT.set_token(token);
            }

            Ok(())
        })
        .invoke_handler(tauri::generate_handler![
            login,
            register,
            logout,
            get_me,
            handshake_init,
            handshake_response,
            get_pending_handshakes,
        ])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
