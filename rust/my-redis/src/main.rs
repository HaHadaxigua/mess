use mini_redis::{client, Result};

#[tokio::main]
async fn main() -> Result<()> {
    let mut client = client::connect("127.0.0.1:6379").await?;

    client.set("hello", "world".into()).await?;

    let result = client.get("hello").await?;

    println!("got value from the server; result={:?}", result);

    Ok(())
}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    async fn test_await() {
        async fn say_world() {
            println!("world")
        }

        async fn try_await() {
            let op = say_world();
            println!("hello");
            op.await;
        }

        try_await().await
    }
}
