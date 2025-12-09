input = File.read!("inputs/day08.txt")

Benchee.run(
  %{
    "part1" => fn -> Day08.part1(input) end,
    "part2" => fn -> Day08.part2(input) end
  },
  warmup: 2,
  time: 3,
  memory_time: 2,
   formatters: [{Benchee.Formatters.HTML, file: "bench/day08/bench.html"}]
)
