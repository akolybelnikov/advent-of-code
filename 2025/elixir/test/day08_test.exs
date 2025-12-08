defmodule Day08Test do
  use ExUnit.Case
  doctest Day08

  @sample_input File.read!("test/inputs/day08_sample.txt")
  @input File.read!("inputs/day08.txt")

  describe "part 1" do
    test "sample input" do
      expected = 40
      actual = Day08.part1(@sample_input, 10)
      assert actual == expected
    end

    test "real input" do
      expected = 57564
      actual = Day08.part1(@input)
      assert actual == expected
    end
  end

  describe "part 2" do
    test "sample input" do
      expected = 25272
      actual = Day08.part2(@sample_input)
      assert actual == expected
    end

    test "real input" do
      expected = 133296744
      actual = Day08.part2(@input)
      assert actual == expected
    end
  end
end
