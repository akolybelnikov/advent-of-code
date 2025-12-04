input = File.read!("inputs/day04.txt")

Benchee.run(
  %{
    "part1" => fn -> Day04.part1(input) end,
    "part2" => fn -> Day04.part2(input) end
  },
  warmup: 2,
  time: 3,
  memory_time: 2,
   formatters: [{Benchee.Formatters.HTML, file: "bench/day04/bench.html"}]
)
