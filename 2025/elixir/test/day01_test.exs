defmodule Day01Test do
  use ExUnit.Case
  doctest Day01

  @sample_input ""

  test "part 1" do
    expected = 0
    actual = Day01.part1(@sample_input)
    assert actual == expected
  end

  test "part 2" do
    expected = 0
    actual = Day01.part2(@sample_input)
    assert actual == expected
  end
end
