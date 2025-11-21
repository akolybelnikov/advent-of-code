fn main() {
    println!("--- Part One ---");
    println!("Result: {}", part1("inputs/day{{DAY}}.txt"));

    println!("--- Part Two ---");
    println!("Result: {}", part2("inputs/day{{DAY}}.txt"));
}

fn part1(filename: &str) -> i32 {
    let input = std::fs::read_to_string(filename).expect("Failed to read input file");
    println!("{}", input);
    0
}

fn part2(filename: &str) -> i32 {
    let input = std::fs::read_to_string(filename).expect("Failed to read input file");
    println!("{}", input);
    0
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "";
        let expected = 0;
        let actual = part1(input);
        assert_eq!(actual, expected);
    }

    #[test]
    fn test_part2() {
        let input = "";
        let expected = 0;
        let actual = part2(input);
        assert_eq!(actual, expected);
    }
}
