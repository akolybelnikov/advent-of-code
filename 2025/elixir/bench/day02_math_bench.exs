Code.compile_file("lib/day02.ex")
Code.compile_file("lib/day02_math.ex")

input_real = File.read!("inputs/day02.txt")

Benchee.run(%{
  "orig part1" => fn -> Day02.part1(input_real) end,
  "math part1" => fn -> Day02Math.part1(input_real) end,
  "orig part2" => fn -> Day02.part2(input_real) end,
  "math part2" => fn -> Day02Math.part2(input_real) end
},
  time: 3,
  warmup: 2,
  memory_time: 2,
  formatters: [
    {Benchee.Formatters.HTML, file: "bench/day02_math_bench.html"},
    Benchee.Formatters.Console
  ]
)
