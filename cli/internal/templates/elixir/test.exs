defmodule Day{{DAY}}Test do
  use ExUnit.Case
  doctest Day{{DAY}}

  @sample_input File.read!("test/inputs/day{{DAY}}_sample.txt")
  @input File.read!("inputs/day{{DAY}}.txt")

  test "part 1 sample" do
    expected = 0
    actual = Day{{DAY}}.part1(@sample_input)
    assert actual == expected
  end

  test "part 1" do
    expected = 0
    actual = Day{{DAY}}.part1(@input)
    assert actual == expected
  end

  test "part 2 sample" do
    expected = 0
    actual = Day{{DAY}}.part2(@sample_input)
    assert actual == expected
  end

  test "part 2" do
    expected = 0
    actual = Day{{DAY}}.part2(@input)
    assert actual == expected
  end
end
