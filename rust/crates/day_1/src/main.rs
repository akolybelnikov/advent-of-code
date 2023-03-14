fn main() {
    let puzzle_input = include_str!("test_input.txt");
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
            total_fuel_with_fuel = 0;
        }
    }
}
