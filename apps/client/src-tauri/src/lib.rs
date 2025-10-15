use dotenv::dotenv;
use std::env;
use once_cell::sync::Lazy;
use reqwest::Client;
use serde::{Deserialize, Serialize};

pub static API_URL: Lazy<String> = Lazy::new(|| {
    env::var("API_URL").expect("API_URL must be set in .env")
});

pub static HTTP_CLIENT: Lazy<Client> = Lazy::new(|| {
    Client::builder()
        .timeout(std::time::Duration::from_secs(10))
        .build()
        .expect("Failed to build HTTP client")
});

#[derive(Serialize, Deserialize)]
struct LoginResponse {
    token: String,
}

#[derive(Serialize, Deserialize)]
struct RegisterResponse {
    id: u32,
    username: String,
}

#[tauri::command]
async fn login(username: String, password: String) -> Result<LoginResponse, String> {
    let res = HTTP_CLIENT.post(format!("{}/auth/login", *API_URL))
        .json(&serde_json::json!({ "username": username, "password": password }))
        .send()
        .await
        .map_err(|e| e.to_string())?;

    if res.status().is_success() {
        let body: LoginResponse = res.json().await.map_err(|e| e.to_string())?;
        Ok(body)
    } else {
        Err("Login failed".into())
    }
}

#[tauri::command]
async fn register(username: String, password: String) -> Result<RegisterResponse, String> {
    let res = HTTP_CLIENT.post(format!("{}/auth/register", *API_URL))
        .json(&serde_json::json!({ "username": username, "password": password }))
        .send()
        .await
        .map_err(|e| e.to_string())?;

    if res.status().is_success() {
        let body: RegisterResponse = res.json().await.map_err(|e| e.to_string())?;
        Ok(body)
    } else {
        Err("Register failed".into())
    }
}

#[cfg_attr(mobile, tauri::mobile_entry_point)]
pub fn run() {
    dotenv().ok();
    tauri::Builder::default()
        .plugin(tauri_plugin_opener::init())
        .invoke_handler(tauri::generate_handler![login, register])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
