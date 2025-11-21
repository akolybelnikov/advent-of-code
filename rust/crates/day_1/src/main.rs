fn main() {
    let puzzle_input = include_str!("input.txt");
    let max_fuel = calculate_fuel_1(puzzle_input);
    println!("Max fuel part 1: {}", max_fuel);

    let top_three_fuel = calculate_top_three_fuel(puzzle_input);
    println!("Top three fuel part 2: {}", top_three_fuel);
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
    // check the last fuel before the end of the file
    if total_fuel > max_fuel {
        max_fuel = total_fuel;
    }

    max_fuel
}

fn calculate_top_three_fuel(puzzle_input: &str) -> i32 {
    let mut top_fuel_sum = 0;
    let mut fuel = 0;
    let mut fuels = Vec::new();
    for line in puzzle_input.lines() {
        let mass = line.parse::<i32>();
        if let Ok(mass) = mass {
            fuel += mass;
        } else {
            fuels.push(fuel);
            fuel = 0;
        }
    }
    // push the last fuel before the end of the file
    fuels.push(fuel);

    fuels.sort();
    fuels.reverse();

    for i in 0..3 {
        top_fuel_sum += fuels[i];
    }
    top_fuel_sum
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_calculate_fuel_1() {
        let puzzle_input = include_str!("test_input.txt");
        assert_eq!(calculate_fuel_1(puzzle_input), 24000);
    }

    #[test]
    fn test_calculate_top_three_fuel() {
        let puzzle_input = include_str!("test_input.txt");
        assert_eq!(calculate_top_three_fuel(puzzle_input), 45000);
    }
}
