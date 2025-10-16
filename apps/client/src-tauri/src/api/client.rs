use std::sync::RwLock;
use reqwest::{header, Client};

pub struct ApiClient {
    pub(crate) client: Client,
    token: RwLock<Option<String>>,
}

impl ApiClient {
    pub fn new() -> Self {
        ApiClient {
            client: Client::builder()
                .timeout(std::time::Duration::from_secs(10))
                .build()
                .unwrap(),
            token: RwLock::new(None),
        }
    }

    pub fn set_token(&self, token: String) {
        *self.token.write().unwrap() = Some(token);
    }

    pub async fn get(&self, url: &str) -> Result<reqwest::Response, reqwest::Error> {
        let mut req = self.client.get(url);

        if let Some(token) = self.token.read().unwrap().clone() {
            req = req.header(header::AUTHORIZATION, format!("Bearer {}", token));
        }

        req.send().await
    }

    pub async fn post(&self, url: &str, body: serde_json::Value) -> Result<reqwest::Response, reqwest::Error> {
        let mut req = self.client.post(url).json(&body);

        if let Some(token) = self.token.read().unwrap().clone() {
            req = req.header(header::AUTHORIZATION, format!("Bearer {}", token));
        }

        req.send().await
    }
}
