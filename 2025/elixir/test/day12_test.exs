defmodule Day12Test do
  use ExUnit.Case
  doctest Day12

  @sample_input File.read!("test/inputs/day12_sample.txt")
  @input File.read!("inputs/day12.txt")

  describe "part 1" do
    test "sample input" do
      expected = 2
      actual = Day12.part1(@sample_input)
      assert actual == expected
    end

    test "real input" do
      expected = 557
      actual = Day12.part1(@input)
      assert actual == expected
    end
  end

  describe "part 2" do
    test "sample input" do
      expected = 0
      actual = Day12.part2(@sample_input)
      assert actual == expected
    end

    test "real input" do
      expected = 0
      actual = Day12.part2(@input)
      assert actual == expected
    end
  end
end
