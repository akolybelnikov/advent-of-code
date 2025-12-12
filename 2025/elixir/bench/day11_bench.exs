input = File.read!("inputs/day11.txt")
input = File.read!("inputs/day11.txt")

Benchee.run(
  %{
    "part1" => fn -> Day11.part1(input) end,
    "part2" => fn -> Day11.part2(input) end
  },
  warmup: 2,
  time: 3,
  memory_time: 2,
   formatters: [{Benchee.Formatters.HTML, file: "bench/day11/bench.html"}]
)
