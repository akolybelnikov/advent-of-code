defmodule Day02Test do
  use ExUnit.Case
  doctest Day02

  @sample_input File.read!("test/inputs/day02_sample.txt")
  @input File.read!("inputs/day02.txt")

  describe "part 1" do
    test "sample input" do
      expected = 1227775554
      actual = Day02.part1(@sample_input)
      assert actual == expected
    end

    test "real input" do
      expected = 38310256125
      actual = Day02.part1(@input)
      assert actual == expected
    end
  end

  describe "part 2" do
    test "sample input" do
      expected = 4174379265
      actual = Day02.part2(@sample_input)
      assert actual == expected
    end

    test "real input" do
      expected = 58961152806
      actual = Day02.part2(@input)
      assert actual == expected
    end
  end
end
