defmodule Day09Test do
  use ExUnit.Case
  doctest Day09

  @sample_input File.read!("test/inputs/day09_sample.txt")
  @input File.read!("inputs/day09.txt")

  describe "part 1" do
    test "sample input" do
      expected = 0
      actual = Day09.part1(@sample_input)
      assert actual == expected
    end

    test "real input" do
      expected = 0
      actual = Day09.part1(@input)
      assert actual == expected
    end
  end

  describe "part 2" do
    test "sample input" do
      expected = 0
      actual = Day09.part2(@sample_input)
      assert actual == expected
    end

    test "real input" do
      expected = 0
      actual = Day09.part2(@input)
      assert actual == expected
    end
  end
end
