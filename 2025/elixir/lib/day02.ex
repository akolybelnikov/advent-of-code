defmodule Day02 do
  @moduledoc """
  Advent of Code - Day 02
  """

  @doc """
  Solves Part 1
  """
  def part1(input) do
    input
    |> String.trim()
    |> String.split(",")
    |> Enum.sum_by(&count_repeated_twice/1)
  end

  @doc """
  Solves Part 2
  """
  def part2(input) do
    input
    |> String.trim()
    |> String.split(",")
    |> Enum.sum_by(&count_repeated_sequence/1)
  end

  defp parse(str) do
    [start_s, end_s] = String.split(str, "-")
    start = String.to_integer(start_s)
    finish = String.to_integer(end_s)

    {start, finish}
  end

  defp count_repeated_twice(range_str) do
    {start, finish} = parse(range_str)

    count_in_range(start, finish, &has_identical_halves?/1)
  end

  defp count_repeated_sequence(range_str) do
    {start, finish} = parse(range_str)

    count_in_range(start, finish, &invalid_id?/1)
  end

  # Avoid creating full range, iterate and count
  defp count_in_range(start, finish, predicate) do
    do_count(start, finish, predicate, 0)
  end

  defp do_count(n, finish, _predicate, acc) when n > finish, do: acc

  defp do_count(n, finish, predicate, acc) do
    new_acc = if predicate.(n), do: acc + n, else: acc
    do_count(n + 1, finish, predicate, new_acc)
  end

  # Work with integers directly
  defp has_identical_halves?(num) do
    digits = Integer.digits(num)
    len = length(digits)

    rem(len, 2) == 0 and
      Enum.split(digits, div(len, 2)) |> then(fn {left, right} -> left == right end)
  end

  defp invalid_id?(id) do
    digits = Integer.digits(id)
    len = length(digits)

    # Cache divisors for common lengths
    divisors = divisors_for_length(len)

    Enum.any?(divisors, fn d ->
      pattern = Enum.take(digits, d)
      chunks = Enum.chunk_every(digits, d)
      Enum.all?(chunks, &(&1 == pattern))
    end)
  end

  # Most numbers will be 6-10 digits, cache those
  defp divisors_for_length(len) do
    case len do
      1 -> []
      2 -> [1]
      3 -> [1]
      4 -> [2, 1]
      5 -> [1]
      6 -> [3, 2, 1]
      8 -> [4, 2, 1]
      10 -> [5, 2, 1]
      _ -> Enum.filter(len..1//-1, &(rem(len, &1) == 0)) |> tl()
    end
  end

  @doc """
  Reads input file and runs both parts
  """
  def solve do
    input = read_input()

    IO.puts("--- Part One ---")
    IO.puts("Result: #{part1(input)}")

    IO.puts("--- Part Two ---")
    IO.puts("Result: #{part2(input)}")
  end

  defp read_input do
    File.read!("inputs/day02.txt")
  end
end
