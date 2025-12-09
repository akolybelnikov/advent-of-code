defmodule Day09 do
  @moduledoc """
  Advent of Code - Day 09
  """

  @doc """
  Solves Part 1
  """
  def part1(input) do
    IO.puts(input)
    0
  end

  @doc """
  Solves Part 2
  """
  def part2(input) do
    IO.puts(input)
    0
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
    File.read!("inputs/day09.txt")
  end
end
