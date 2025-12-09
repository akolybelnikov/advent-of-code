input = File.read!("inputs/day05.txt")

Benchee.run(
  %{
    "part1" => fn -> Day05.part1(input) end,
    "part2" => fn -> Day05.part2(input) end
  },
  warmup: 2,
  time: 3,
  memory_time: 2,
   formatters: [{Benchee.Formatters.HTML, file: "bench/day05/bench.html"}]
)
