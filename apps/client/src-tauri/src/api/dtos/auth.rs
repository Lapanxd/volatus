use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
pub struct LoginOutputDto {
    pub token: String,
}

#[derive(Serialize, Deserialize)]
pub struct RegisterOutputDto {
    pub id: u32,
    pub username: String,
}
