input = File.read!("inputs/day09.txt")

Benchee.run(
  %{
    "part1" => fn -> Day09.part1(input) end,
    "part2" => fn -> Day09.part2(input) end
  },
  warmup: 2,
  time: 3,
  memory_time: 2,
   formatters: [{Benchee.Formatters.HTML, file: "bench/day09/bench.html"}]
)
