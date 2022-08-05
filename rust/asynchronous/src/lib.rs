#[cfg(test)]
mod tests {
    use core::time;
    use futures::{executor::block_on, Future};
    use std::thread;

    #[test]
    fn test_thread() {
        let t_1 = thread::spawn(|| thread::sleep(time::Duration::from_secs(5)));
        let t_2 = thread::spawn(|| thread::sleep(time::Duration::from_secs(1)));

        t_1.join().expect("thread one panic");
        t_2.join().expect("thread two panic");
    }

    #[test]
    fn test_async() {
        async fn hello_world() {
            println!("hello world");
        }

        let future = hello_world();
        block_on(future);
    }

    #[test]
    fn test_await() {
        struct Song;
        async fn learn_song() -> Song {
            println!("learn song");
            Song {}
        }
        async fn sing_song(_song: Song) {
            println!("sing_song");
        }
        async fn dance() {
            println!("dance");
        }

        async fn learn_and_sing() {
            let song = learn_song().await;
            sing_song(song).await;
        }

        async fn async_main() {
            let f1 = learn_and_sing();
            let f2 = dance();

            futures::join!(f1, f2);
        }

        block_on(async_main());
    }

    #[test]
    fn test_async_syntax() {
        async fn foo() -> u8 {
            5
        }
        fn bar() -> impl Future<Output = u8> {
            async {
                let x: u8 = foo().await;
                x + 5
            }
        }
    }

    #[test]
    fn test() {
        struct My {
            a: i32,
            b: i32,
            c: i32,
        }

        let y1 = My { a: 1, b: 1, c: 1 };
        let x = My { ..y1 };
    }
}
