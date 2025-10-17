use crate::api::dtos::user::UserOutputDto;
use crate::{API_CLIENT, API_URL};

#[tauri::command]
pub async fn get_me() -> Result<UserOutputDto, String> {
    let res = API_CLIENT.get(
        &format!("{}/users/me", *API_URL))
        .await
        .map_err(|e| e.to_string())?;

    if res.status().is_success() {
        let body: UserOutputDto = res.json().await.map_err(|e| e.to_string())?;
        Ok(body)
    } else {
        Err("Cannot get me".to_string())
    }
}