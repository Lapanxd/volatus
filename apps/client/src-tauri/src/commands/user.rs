use tauri::{AppHandle, Wry};
use tauri_plugin_store::StoreExt;
use crate::api::dtos::user::UserOutputDto;
use crate::{API_CLIENT, API_URL};

#[tauri::command]
pub async fn get_me(app_handle: AppHandle<Wry>) -> Result<(), String> {
    let store = app_handle.store("store.json").map_err(|e| e.to_string())?;

    let res = API_CLIENT.get(
        &format!("{}/users/me", *API_URL))
        .await
        .map_err(|e| e.to_string())?;

    if res.status().is_success() {
        let body: UserOutputDto = res.json().await.map_err(|e| e.to_string())?;
        store.set("user", serde_json::json!(body));
        Ok(())
    } else {
        Err("Cannot get me".to_string())
    }
}

#[tauri::command]
pub async fn get_user_by_id(id: u32) -> Result<UserOutputDto, String> {
    let res = API_CLIENT.get(
        &format!("{}/users/{}", *API_URL, id))
        .await
        .map_err(|e| e.to_string())?;

    if res.status().is_success() {
        let body: UserOutputDto = res.json().await.map_err(|e| e.to_string())?;
        Ok(body)
    } else {
        Err("Cannot get me".to_string())
    }
}