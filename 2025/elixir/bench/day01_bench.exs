inputs = %{
  "sample" => File.read!("test/inputs/day01_sample.txt"),
  "real" => File.read!("inputs/day01.txt")
}

Benchee.run(
  %{
    "part1" => fn input -> Day01.part1(input) end,
    "part2" => fn input -> Day01.part2(input) end
  },
  inputs: inputs,
  time: 3,
  memory_time: 2,
  formatters: [{Benchee.Formatters.HTML, file: "bench/bench.html"}]
)
