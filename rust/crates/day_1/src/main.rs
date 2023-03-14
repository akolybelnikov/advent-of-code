fn main() {
    let puzzle_input = include_str!("input.txt");
    let max_fuel = calculate_fuel_1(puzzle_input);
    println!("Max fuel part 1: {}", max_fuel);
}

fn calculate_fuel_1(puzzle_input: &str) -> i32 {
    let mut total_fuel = 0;
    let mut max_fuel = 0;
    for line in puzzle_input.lines() {
        let mass = line.parse::<i32>();
        if let Ok(mass) = mass {
            total_fuel += mass;
        } else {
            if total_fuel > max_fuel {
                max_fuel = total_fuel;
            }
            total_fuel = 0;
        }
    }
    max_fuel
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_calculate_fuel_1() {
        let puzzle_input = include_str!("test_input.txt");
        assert_eq!(calculate_fuel_1(puzzle_input), 24000);
    }
}
