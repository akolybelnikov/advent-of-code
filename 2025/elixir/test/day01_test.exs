defmodule Day01Test do
  use ExUnit.Case
  doctest Day01

  @sample_input File.read!("test/inputs/day01_sample.txt")
  @input File.read!("inputs/day01.txt")

  test "part 1 sample" do
    expected = 3
    actual = Day01.part1(@sample_input)
    assert actual == expected
  end

  test "part 1 real input" do
    expected = 982
    actual = Day01.part1(@input)
    assert actual == expected
  end

  describe "part 2" do
    test "sample input" do
      expected = 6
      actual = Day01.part2(@sample_input)
      assert actual == expected
    end

    test "real" do
      expected = 6106
      actual = Day01.part2(@input)
      assert actual == expected
    end
  end

  describe "rotate/2 - Left (Lxx)" do
    test "no wrap: pos 50, L10 -> pos 40, times 0" do
      assert Day01.rotate("L10", 50) == {40, 0}
    end

    test "land on 0: pos 5, L5 -> pos 0, times 1" do
      assert Day01.rotate("L5", 5) == {0, 1}
    end

    test "wrap once (pass 0 once): pos 5, L6 -> pos 99, times 1" do
      assert Day01.rotate("L6", 5) == {99, 1}
    end

    test "wrap twice: pos 5, L106 -> pos 99, times 2" do
      assert Day01.rotate("L106", 5) == {99, 2}
    end

    test "wrap once and land on 0: pos 5, L105 -> pos 0, times 2" do
      assert Day01.rotate("L105", 5) == {0, 2}
    end

    test "start on 0 and don't wrap: pos 0, L5 -> pos 95, times 0" do
      assert Day01.rotate("L5", 0) == {95, 0}
    end

    test "start on 0 and land on 0: pos 0 L500 -> pos 0, times 5" do
      assert Day01.rotate("L500", 0) == {0, 5}
    end
  end

  describe "rotate/2 - Right (Rxx)" do
    test "no wrap: pos 50, R10 -> pos 60, times 0" do
      assert Day01.rotate("R10", 50) == {60, 0}
    end

    test "land on 0: pos 50, R50 -> pos 0, times 1" do
      assert Day01.rotate("R50", 50) == {0, 1}
    end

    test "wrap once (pass 0 once): pos 99, R2 -> pos 1, times 1" do
      assert Day01.rotate("R2", 99) == {1, 1}
    end

    test "wrap once (pass 0 once): pos 80, R25 -> pos 5, times 1" do
      assert Day01.rotate("R25", 80) == {5, 1}
    end

    test "wrap once and land on 0: pos 80, R120 -> pos 0, times 2" do
      assert Day01.rotate("R120", 80) == {0, 2}
    end

    test "wrap twice: pos 80, R121 -> pos 1, times 2" do
      assert Day01.rotate("R121", 80) == {1, 2}
    end

    test "start on 0 and don't wrap: pos 0, R5 -> pos 5, times 0" do
      assert Day01.rotate("R5", 0) == {5, 0}
    end

    test "start on 0 and end on 0: pos 0, R100 -> pos 0, times 1" do
      assert Day01.rotate("R100", 0) == {0, 1}
    end
  end
end
