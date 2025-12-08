defmodule Day05Test do
  use ExUnit.Case
  doctest Day05

  @sample_input File.read!("test/inputs/day05_sample.txt")
  @input File.read!("inputs/day05.txt")

  describe "part 1" do
    test "sample input" do
      expected = 3
      actual = Day05.part1(@sample_input)
      assert actual == expected
    end

    test "real input" do
      expected = 726
      actual = Day05.part1(@input)
      assert actual == expected
    end
  end

  describe "part 2" do
    test "sample input" do
      expected = 14
      actual = Day05.part2(@sample_input)
      assert actual == expected
    end

    test "real input" do
      expected = 354226555270043
      actual = Day05.part2(@input)
      assert actual == expected
    end
  end
end
