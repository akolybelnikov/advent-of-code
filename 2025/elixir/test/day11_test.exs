defmodule Day11Test do
  use ExUnit.Case
  doctest Day11

  @sample_input File.read!("test/inputs/day11_sample.txt")
  @sample_input_2 File.read!("test/inputs/day11_sample_2.txt")
  @input File.read!("inputs/day11.txt")

  describe "part 1" do
    test "sample input" do
      expected = 5
      actual = Day11.part1(@sample_input)
      assert actual == expected
    end

    test "real input" do
      expected = 749
      actual = Day11.part1(@input)
      assert actual == expected
    end
  end

  describe "part 2" do
    test "sample input" do
      expected = 2
      actual = Day11.part2(@sample_input_2)
      assert actual == expected
    end

    test "real input" do
      expected = 420257875695750
      actual = Day11.part2(@input)
      assert actual == expected
    end
  end
end
