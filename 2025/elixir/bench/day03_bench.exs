input = File.read!("inputs/day03.txt")

Benchee.run(
  %{
    "part1 (stack k=2)" => fn -> Day03.part1(input) end,
    "part2 (stack k=12)" => fn -> Day03.part2(input) end
  },
  warmup: 2,
  time: 3,
  memory_time: 2,
    formatters: [{Benchee.Formatters.HTML, file: "bench/day03/bench.html"}]
)
