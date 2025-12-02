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
    |> Enum.reduce(0, fn range, sum ->
      sum + process(range)
    end)
    |> IO.puts()

    0
  end

  @doc """
  Solves Part 2
  """
  def part2(input) do
    IO.puts(input)
    0
  end

  def process(range) do
    range
    |> String.split("-")
    |> then(fn [s, e] -> String.to_integer(s)..String.to_integer(e) end)
    |> Enum.reduce(0, fn num, sum ->
      s = Integer.to_string(num)
      len = String.length(s)

      cond do
        rem(len, 2) == 0 ->
          {l, r} = String.split_at(s, div(String.length(s), 2))

          if l == r do
            sum + num
          else
            sum
          end

        true ->
          sum
      end
    end)
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
