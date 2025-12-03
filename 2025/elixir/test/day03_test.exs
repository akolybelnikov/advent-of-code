defmodule Day03Test do
  use ExUnit.Case
  doctest Day03

  @sample_input File.read!("test/inputs/day03_sample.txt")
  @input File.read!("inputs/day03.txt")

  test "part 1 sample" do
    expected = 357
    actual = Day03.part1(@sample_input)
    assert actual == expected
  end

  test "part 1" do
     expected = 17034
    actual = Day03.part1(@input)
    assert actual == expected
  end

  test "part 2 sample" do
    # Sum of picked-only 12-digit prefixes per line
    # 987654321111 + 811111111119 + 434234234278 + 888911112111
    expected = 3_121_910_778_619
    actual = Day03.part2(@sample_input)
    assert actual == expected
  end

  test "part 2" do
    expected = 168_798_209_663_590
    actual = Day03.part2(@input)
    assert actual == expected
  end

  describe "find_max_joltage/1" do
    test "line 1: 987654321111111" do
      assert Day03.find_max_joltage("987654321111111") == 98
    end

    test "line 2: 811111111111119" do
      assert Day03.find_max_joltage("811111111111119") == 89
    end

    test "line 3: 234234234234278" do
      assert Day03.find_max_joltage("234234234234278") == 78
    end

    test "line 4: 818181911112111" do
      assert Day03.find_max_joltage("818181911112111") == 92
    end
  end

  describe "to_num/1" do
    test "converts '98' to 98" do
      assert Day03.to_num(~c"98") == 98
    end

    test "converts '12' to 12" do
      assert Day03.to_num(~c"12") == 12
    end

    test "converts '00' to 0" do
      assert Day03.to_num(~c"00") == 0
    end

    test "converts '99' to 99" do
      assert Day03.to_num(~c"99") == 99
    end

    test "converts '01' to 1" do
      assert Day03.to_num(~c"01") == 1
    end
  end

  describe "largest_last/2" do
    test "returns [3,8] when comparing [3,7] and [9,8]" do
      assert Day03.largest_last(~c"37", ~c"98") == ~c"38"
    end

    test "returns [9,7] when comparing [9,5] and [1,7]" do
      assert Day03.largest_last(~c"95", ~c"17") == ~c"97"
    end

    test "returns [8,9] when comparing [8,9] and [8,1]" do
      assert Day03.largest_last(~c"89", ~c"81") == ~c"89"
    end

    test "returns same when both equal" do
      assert Day03.largest_last(~c"99", ~c"99") == ~c"99"
    end

    test "returns [1,9] when comparing [1,2] and [1,9]" do
      assert Day03.largest_last(~c"12", ~c"19") == ~c"19"
    end
  end

  describe "find_max_joltage_12/1 (picked-only prefix)" do
    test "line 1: 987654321111111 -> 987654321111" do
      result = Day03.find_max_joltage_12("987654321111111")
      assert List.to_integer(result) == 987_654_321_111
    end

    test "line 2: 811111111111119 -> 811111111119" do
      result = Day03.find_max_joltage_12("811111111111119")
      assert List.to_integer(result) == 811_111_111_119
    end

    test "line 3: 234234234234278 -> 434234234234" do
      result = Day03.find_max_joltage_12("234234234234278")
      assert List.to_integer(result) == 434_234_234_278
    end

    test "line 4: 818181911112111 -> 888911112111" do
      result = Day03.find_max_joltage_12("818181911112111")
      assert List.to_integer(result) == 888_911_112_111
    end
  end
end
