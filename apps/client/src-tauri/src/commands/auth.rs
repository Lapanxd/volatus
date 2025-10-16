use crate::api::dtos::auth::{LoginResponse, RegisterResponse};
use crate::{API_CLIENT, API_URL};

#[tauri::command]
pub async fn login(username: String, password: String) -> Result<LoginResponse, String> {
    let client = &API_CLIENT.client;

    let res = client.post(format!("{}/auth/login", *API_URL))
        .json(&serde_json::json!({ "username": username, "password": password }))
        .send()
        .await
        .map_err(|e| e.to_string())?;

    if res.status().is_success() {
        let body: LoginResponse = res.json().await.map_err(|e| e.to_string())?;

        API_CLIENT.set_token(body.token.clone());

        Ok(body)
    } else {
        Err("Login failed".into())
    }
}

#[tauri::command]
pub async fn register(username: String, password: String) -> Result<RegisterResponse, String> {
    let client = &API_CLIENT.client;

    let res = client.post(format!("{}/auth/register", *API_URL))
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