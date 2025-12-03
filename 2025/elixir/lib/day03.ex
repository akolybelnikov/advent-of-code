defmodule Day03 do
  @moduledoc """
  Advent of Code - Day 03
  """

  @doc """
  Solves Part 1
  """
  def part1(input) do
    input
    |> String.trim()
    |> String.split("\n")
    |> Enum.sum_by(&find_max_joltage/1)
  end

  @doc """
  Solves Part 2
  """
  def part2(input) do
    IO.puts(input)
    0
  end

  def find_max_joltage(battery) do
    String.trim(battery)
    |> String.to_charlist()
    |> Enum.chunk_every(2, 1, :discard)
    |> Enum.reduce([?0, ?0], &by_max_sequence/2)
    |> List.to_integer()
  end

  def by_max_sequence(next, cur) do
    if to_num(next) > to_num(cur) do
      next
    else
      largest_last(cur, next)
    end
  end

  def to_num([a, b]) do
    10 * (a - ?0) + (b - ?0)
  end

  def largest_last([a, b], [_, d]) do
    if b > d do
      [a, b]
    else
      [a, d]
    end
  end

  @doc """
  Reads input file and runs both parts
  811111111111119
  """
  def solve do
    input = read_input()

    IO.puts("--- Part One ---")
    IO.puts("Result: #{part1(input)}")

    IO.puts("--- Part Two ---")
    IO.puts("Result: #{part2(input)}")
  end

  defp read_input do
    File.read!("inputs/day03.txt")
  end
end
