use tauri::{AppHandle, Wry};
use tauri_plugin_store::StoreExt;
use crate::api::dtos::auth::{LoginOutputDto, RegisterOutputDto};
use crate::{API_CLIENT, API_URL};

#[tauri::command]
pub async fn login(app_handle: AppHandle<Wry>, username: String, password: String) -> Result<(), String> {
    let store = app_handle.store("store.json").map_err(|e| e.to_string())?;

    let res = API_CLIENT.post(
        &format!("{}/auth/login", *API_URL),
        serde_json::json!({ "username": username, "password": password }))
        .await
        .map_err(|e| e.to_string())?;

    if res.status().is_success() {
        let body: LoginOutputDto = res.json().await.map_err(|e| e.to_string())?;

        API_CLIENT.set_token(body.token.clone());
        store.set("auth_token", serde_json::json!(body.token.clone()));
        Ok(())
    } else {
        Err("Login failed".into())
    }
}

#[tauri::command]
pub async fn register(username: String, password: String) -> Result<RegisterOutputDto, String> {
    let res = API_CLIENT.post(
        &format!("{}/auth/register", *API_URL),
        serde_json::json!({ "username": username, "password": password }))
        .await
        .map_err(|e| e.to_string())?;

    if res.status().is_success() {
        let body: RegisterOutputDto = res.json().await.map_err(|e| e.to_string())?;
        Ok(body)
    } else {
        Err("Register failed".into())
    }
}

#[tauri::command]
pub async fn logout(app_handle: AppHandle<Wry>) -> Result<(), String> {
    let store = app_handle.store("store.json").map_err(|e| e.to_string())?;
    store.clear();
    store.save().map_err(|e| e.to_string())?;
    Ok(())
}
