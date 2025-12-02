inputs = %{
  "sample" => File.read!("test/inputs/day02_sample.txt"),
  "real" => File.read!("inputs/day02.txt")
}

Benchee.run(
  %{
    "part1" => fn input -> Day02.part1(input) end,
    "part2" => fn input -> Day02.part2(input) end
  },
  inputs: inputs,
  time: 3,
  memory_time: 2,
  formatters: [{Benchee.Formatters.HTML, file: "bench/day02_bench.html"}]
)
