defmodule Day07Test do
  use ExUnit.Case
  doctest Day07

  @sample_input File.read!("test/inputs/day07_sample.txt")
  @input File.read!("inputs/day07.txt")

  describe "part 1" do
    test "sample input" do
      expected = 21
      actual = Day07.part1(@sample_input)
      assert actual == expected
    end

    test "real input" do
      expected = 1626
      actual = Day07.part1(@input)
      assert actual == expected
    end
  end

  describe "part 2" do
    test "sample input" do
      expected = 40
      actual = Day07.part2(@sample_input)
      assert actual == expected
    end

    test "real input" do
      expected = 48_989_920_237_096
      actual = Day07.part2(@input)
      assert actual == expected
    end
  end
end
