defmodule Day01 do
  @moduledoc """
  Advent of Code - Day 01
  """

  @doc """
  Solves Part 1
  """
  def part1(input) do
    String.split(input)
    |> Enum.reduce({50, 0}, fn instruction, {pos, cur} ->
      {new_pos, _} = rotate(instruction, pos)
      new_cur = if new_pos == 0, do: cur + 1, else: cur
      {new_pos, new_cur}
    end)
    |> elem(1)
  end

  @doc """
  Solves Part 2
  """
  def part2(input) do
    String.split(input)
    |> Enum.reduce({50, 0}, fn instruction, {pos, cur} ->
      {new_pos, times} = rotate(instruction, pos)
      {new_pos, cur + times}
    end)
    |> elem(1)
  end

  @doc """
  Rotates position based on L or R instruction, wrapping around at 0 and 99
  """
  @spec rotate(String.t(), integer()) :: {integer(), integer()}
  def rotate("L" <> clicks, pos) do
    clicks = String.to_integer(clicks)

    times =
      cond do
        clicks < pos -> 0
        pos != 0 && 100 > clicks && clicks >= pos -> 1
        pos == 0 && 100 > clicks -> 0
        pos == 0 && 100 == clicks -> 1
        true -> div(clicks + 99, 100)
      end

    {Integer.mod(pos - clicks, 100), times}
  end

  def rotate("R" <> clicks, pos) do
    clicks = String.to_integer(clicks)

    times =
      cond do
        clicks < 100 - pos -> 0
        true -> div(clicks + 99, 100)
      end

    {Integer.mod(pos + clicks, 100), times}
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
    File.read!("inputs/day01.txt")
  end
end
