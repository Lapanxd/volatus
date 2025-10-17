use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
pub struct UserOutputDto {
    pub id: u32,
    pub username: String,
}