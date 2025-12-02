defmodule Day02Math do
  @moduledoc """
  Math-based implementation for Day 02, separate from existing code.

  Avoids string/list allocations by using integer arithmetic (`div`/`rem`) and
  tail-recursive iteration over ranges without materializing lists.
  """

  @doc """
  Solves Part 1 using math-based checks (identical halves).
  """
  def part1(input) do
    input
    |> String.trim()
    |> String.split(",")
    |> Enum.sum_by(&sum_identical_halves/1)
  end

  @doc """
  Solves Part 2 using math-based repeated pattern detection.
  """
  def part2(input) do
    input
    |> String.trim()
    |> String.split(",")
    |> Enum.sum_by(&sum_repeated_pattern/1)
  end

  defp sum_identical_halves(range_str) do
    {start, finish} = parse_range_to_tuple(range_str)
    count_sum_in_range(start, finish, &identical_halves?/1)
  end

  defp sum_repeated_pattern(range_str) do
    {start, finish} = parse_range_to_tuple(range_str)
    count_sum_in_range(start, finish, &repeated_pattern?/1)
  end

  # Iterate without building intermediate lists; add matching numbers to sum
  defp count_sum_in_range(start, finish, pred) do
    do_sum(start, finish, pred, 0)
  end

  defp do_sum(n, finish, _pred, acc) when n > finish, do: acc
  defp do_sum(n, finish, pred, acc) do
    acc = if pred.(n), do: acc + n, else: acc
    do_sum(n + 1, finish, pred, acc)
  end

  # Parse "a-b" into {a, b}
  defp parse_range_to_tuple(range_str) do
    [s, e] = String.split(range_str, "-")
    {String.to_integer(s), String.to_integer(e)}
  end

  # Compute number of base-10 digits for positive integers
  defp digit_count(n) when n > 0 do
    do_digit_count(n, 0)
  end
  defp do_digit_count(n, acc) when n > 0 do
    if n < 10, do: acc + 1, else: do_digit_count(div(n, 10), acc + 1)
  end

  # Precompute 10^k up to 20 digits (enough for inputs here)
  @pow10 Enum.reduce(0..20, %{0 => 1}, fn i, acc -> Map.put(acc, i, :math.pow(10, i) |> round) end)
  defp pow10(k), do: Map.fetch!(@pow10, k)

  # Check if number consists of two identical halves (len even, left == right)
  defp identical_halves?(n) when n > 0 do
    len = digit_count(n)
    if rem(len, 2) != 0 do
      false
    else
      k = div(len, 2)
      base = pow10(k)
      left = div(n, base)
      right = rem(n, base)
      left == right
    end
  end

  # Check if n is t repeats of k-digit pattern for some k|len, k < len
  defp repeated_pattern?(n) when n > 0 do
    len = digit_count(n)
    divisors(len)
    |> Enum.any?(fn k -> k < len and repeated_blocks?(n, k, div(len, k)) end)
  end

  # Compare t blocks of size k via div/rem without allocations
  defp repeated_blocks?(n, k, t) do
    base = pow10(k)
    first = rem(n, base)
    do_blocks_equal?(n, k, base, first, 1, t)
  end

  defp do_blocks_equal?(_n, _k, _base, _first, i, t) when i >= t, do: true
  defp do_blocks_equal?(n, k, base, first, i, t) do
    shift = pow10(i * k)
    block_i = rem(div(n, shift), base)
    if block_i == first do
      do_blocks_equal?(n, k, base, first, i + 1, t)
    else
      false
    end
  end

  # No need to extract exponent; we pass k directly.

  # Divisors of n (excluding n itself), in descending order
  defp divisors(n) do
    cond do
      n <= 1 -> []
      true ->
        (n - 1)..1//-1
        |> Enum.filter(&(rem(n, &1) == 0))
    end
  end

  @doc """
  Reads input file and runs both parts for this math-based variant.
  """
  def solve do
    input = File.read!("inputs/day02.txt")

    IO.puts("--- Math Part One ---")
    IO.puts("Result: #{part1(input)}")

    IO.puts("--- Math Part Two ---")
    IO.puts("Result: #{part2(input)}")
  end
end
