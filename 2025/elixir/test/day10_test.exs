defmodule Day10Test do
  use ExUnit.Case
  doctest Day10

  @sample_input File.read!("test/inputs/day10_sample.txt")
  @input File.read!("inputs/day10.txt")

  describe "part 1" do
    test "sample input" do
      expected = 7
      actual = Day10.part1(@sample_input)
      assert actual == expected
    end

    test "real input" do
      expected = 527
      actual = Day10.part1(@input)
      assert actual == expected
    end
  end

  describe "part 2" do
    test "sample input" do
      expected =33
      actual = Day10.part2(@sample_input)
      assert actual == expected
    end

    test "real input" do
      expected = 0
      actual = Day10.part2(@input)
      assert actual == expected
    end
  end
end
