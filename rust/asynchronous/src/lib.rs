#[cfg(test)]
mod tests {
    use core::time;
    use std::thread;
    use super::*;

    #[test]
    fn test_thread() {
        let t_1 = thread::spawn(|| thread::sleep(time::Duration::from_secs(5)));
        let t_2 = thread::spawn(|| thread::sleep(time::Duration::from_secs(1)));

        t_1.join().expect("thread one panic");
        t_2.join().expect("thread two panic");
    }
    
    #[test]
    fn test_async() {
        async fn async_example() {

        }
    }

}