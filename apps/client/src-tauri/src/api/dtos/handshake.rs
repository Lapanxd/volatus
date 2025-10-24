use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
pub struct PendingHandshakeOutputDto {
    pub session_id: String,
    pub from_user_id: u32,
    pub sdp_offer: String,
}