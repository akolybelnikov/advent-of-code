defmodule Day09Test do
  use ExUnit.Case
  doctest Day09

  @sample_input File.read!("test/inputs/day09_sample.txt")
  @input File.read!("inputs/day09.txt")

  describe "part 1" do
    test "sample input" do
      expected = 50
      actual = Day09.part1(@sample_input)
      assert actual == expected
    end

    test "real input" do
      expected = 4777967538
      actual = Day09.part1(@input)
      assert actual == expected
    end
  end

  describe "part 2" do
    test "sample input" do
      expected = 24
      actual = Day09.part2(@sample_input)
      assert actual == expected
    end

    test "real input" do
      expected = 1439894345
      actual = Day09.part2(@input)
      assert actual == expected
    end
  end

  describe "area" do
    test "{2,5}, {11,1} -> 50" do
      assert Day09.area({2, 5}, {11, 1}) == 50
    end

    test "{7,3}, {2,3} -> 6" do
      assert Day09.area({7, 3}, {2, 3}) == 6
    end

    test "{7,1}, {11,7} -> 35" do
      assert Day09.area({7, 1}, {11, 7}) == 35
    end

    test "{2,5}, {9,7} -> 24" do
      assert Day09.area({2, 5}, {9, 7}) == 24
    end
  end
end
