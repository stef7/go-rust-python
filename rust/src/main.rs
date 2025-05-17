#[cfg(test)]
mod tests {
    #[test]
    fn hello_world() {
        use crate::hello_world;
        assert_eq!(hello_world(), "Hello, world!");
    }
}

fn hello_world() -> &'static str {
    return "Hello, world!";
}

fn main() {
    let string = hello_world();
    println!("{}", string);
}
