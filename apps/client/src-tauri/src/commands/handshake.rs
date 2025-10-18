use crate::api::dtos::handshake::PendingHandshakeOutputDto;
use crate::{API_CLIENT, API_URL};

#[tauri::command]
pub async fn handshake_init(to_user_id: u32, sdp_offer: String) -> Result<(), String> {
    let res = API_CLIENT
        .post(
            &format!("{}/handshake/init", *API_URL),
            serde_json::json!({ "to_user_id": to_user_id, "sdp_offer": sdp_offer }),
        )
        .await
        .map_err(|e| e.to_string())?;

    if res.status().is_success() {
        Ok(())
    } else {
        Err("Handshake init failed".into())
    }
}

#[tauri::command]
pub async fn handshake_response(
    session_id: String,
    accepted: bool,
    sdp_answer: String,
) -> Result<(), String> {
    let res = API_CLIENT
        .post(
            &format!("{}/handshake/response", *API_URL),
            serde_json::json!({ "session_id": session_id, "accepted": accepted, "sdp_answer": sdp_answer }),
        )
        .await
        .map_err(|e| e.to_string())?;

    if res.status().is_success() {
        // start p2p communication
        Ok(())
    } else {
        Err("Handshake response failed".into())
    }
}

#[tauri::command]
pub async fn get_pending_handshakes() -> Result<Vec<PendingHandshakeOutputDto>, String> {
    let res = API_CLIENT
        .get(&format!("{}/handshake/pending", *API_URL))
        .await
        .map_err(|e| e.to_string())?;

    if res.status().is_success() {
        let body: Vec<PendingHandshakeOutputDto> = res.json().await.map_err(|e| e.to_string())?;
        Ok(body)
    } else {
        Err("Handshake response failed".into())
    }
}
